package nemeton

import (
	"context"
	"errors"
	"fmt"

	"okp4/nemeton-leaderboard/app/util"

	"github.com/cosmos/cosmos-sdk/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	phasesCollectionName     = "phases"
	validatorsCollectionName = "validators"
)

type Store struct {
	db     *mongo.Database
	phases []*Phase
}

func NewStore(ctx context.Context, mongoURI, dbName string) (*Store, error) {
	db, err := util.OpenMongoDatabase(ctx, mongoURI, dbName)
	if err != nil {
		return nil, err
	}

	store := &Store{
		db: db,
	}

	if err := store.init(ctx); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) init(ctx context.Context) error {
	phases := s.db.Collection(phasesCollectionName)
	count, err := phases.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		for _, phase := range bootstrapPhases() {
			_, err := phases.InsertOne(ctx, phase)
			if err != nil {
				return err
			}
		}
	}

	c, err := phases.Find(ctx, bson.M{})
	if err != nil {
		return err
	}

	for c.Next(ctx) {
		var phase Phase
		if err := c.Decode(&phase); err != nil {
			return err
		}
		s.phases = append(s.phases, &phase)
	}

	if err := s.ensureIndexes(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Store) ensureIndexes(ctx context.Context) error {
	_, err := s.db.Collection(validatorsCollectionName).
		Indexes().
		CreateMany(
			ctx,
			[]mongo.IndexModel{
				{Keys: bson.M{"points": 1}},
				{Keys: bson.M{"moniker": 1}},
				{Keys: bson.M{"valoper": 1}, Options: options.Index().SetUnique(true)},
				{Keys: bson.M{"delegator": 1}, Options: options.Index().SetUnique(true)},
				{Keys: bson.M{"twitter": 1}, Options: options.Index().SetSparse(true).SetUnique(true)},
				{Keys: bson.M{"discord": 1}, Options: options.Index().SetUnique(true)},
			},
		)
	return err
}

func (s *Store) GetPhase(number int) *Phase {
	for _, phase := range s.phases {
		if phase.Number == number {
			return phase
		}
	}
	return nil
}

func (s *Store) GetCurrentPhase() *Phase {
	for _, phase := range s.phases {
		if phase.InProgress() {
			return phase
		}
	}
	return nil
}

func (s *Store) GetAllPhases() []*Phase {
	return s.phases
}

func (s *Store) GetUnstartedPhases() []*Phase {
	return s.GetPhases(func(p Phase) bool {
		return !p.Started()
	})
}

func (s *Store) GetFinishedPhases() []*Phase {
	return s.GetPhases(func(p Phase) bool {
		return p.Finished()
	})
}

func (s *Store) GetPhases(criteriaFn func(p Phase) bool) []*Phase {
	var filtered []*Phase
	for _, phase := range s.phases {
		if criteriaFn(*phase) {
			filtered = append(filtered, phase)
		}
	}
	return filtered
}

func (s *Store) GetValidatorByCursor(ctx context.Context, c Cursor) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"_id": c.objectID})
}

func (s *Store) GetValidatorByValoper(ctx context.Context, addr types.ValAddress) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"valoper": addr.String()})
}

func (s *Store) GetValidatorByDelegator(ctx context.Context, addr types.AccAddress) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"delegator": addr.String()})
}

func (s *Store) GetValidatorByDiscord(ctx context.Context, discord string) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"discord": discord})
}

func (s *Store) GetValidatorByTwitter(ctx context.Context, twitter string) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"twitter": twitter})
}

func (s *Store) GetValidatorBy(ctx context.Context, filter bson.M) (*Validator, error) {
	res := s.db.Collection(validatorsCollectionName).FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	var val Validator
	return &val, res.Decode(&val)
}

func (s *Store) GetValidatorRank(ctx context.Context, cursor Cursor) (int, error) {
	count, err := s.db.Collection(validatorsCollectionName).
		CountDocuments(
			ctx,
			bson.M{
				"$or": bson.A{
					bson.M{
						"points": bson.M{"$gt": cursor.points},
					},
					bson.M{
						"points": cursor.points,
						"_id":    bson.M{"$lt": cursor.objectID},
					},
				},
			},
		)

	return int(count) + 1, err
}

func (s *Store) CountValidators(ctx context.Context) (int64, error) {
	return s.db.Collection(validatorsCollectionName).CountDocuments(ctx, bson.M{})
}

func (s *Store) GetBoard(ctx context.Context, search *string, limit int, after *Cursor) ([]*Validator, bool, error) {
	c, err := s.db.Collection(validatorsCollectionName).Find(
		ctx,
		makeBoardFilter(search, after),
		options.Find().
			SetSort(
				bson.D{
					bson.E{Key: "points", Value: -1},
					bson.E{Key: "_id", Value: 1},
				},
			).
			SetLimit(int64(limit+1)),
	)
	if err != nil {
		return nil, false, err
	}
	defer func() {
		_ = c.Close(ctx)
	}()

	validators := make([]*Validator, 0, limit)
	count := 0
	for count < limit && c.Next(ctx) {
		count++
		var validator Validator
		if err := c.Decode(&validator); err != nil {
			return nil, false, err
		}
		validators = append(validators, &validator)
	}

	return validators, c.Next(ctx), nil
}

func makeBoardFilter(search *string, after *Cursor) bson.M {
	filters := bson.A{}
	if after != nil {
		filters = append(filters, bson.M{
			"$or": bson.A{
				bson.M{
					"points": bson.M{"$lt": after.points},
				},
				bson.M{
					"points": after.points,
					"_id":    bson.M{"$gt": after.objectID},
				},
			},
		})
	}
	if search != nil {
		filters = append(filters, bson.M{
			"$or": bson.A{
				bson.M{"moniker": bson.M{"$regex": fmt.Sprintf(".*%s.*", *search), "$options": "i"}},
				bson.M{"valoper": bson.M{"$regex": fmt.Sprintf(".*%s.*", *search), "$options": "i"}},
				bson.M{"delegator": bson.M{"$regex": fmt.Sprintf(".*%s.*", *search), "$options": "i"}},
			},
		})
	}

	var filter bson.M
	if len(filters) > 0 {
		filter = bson.M{
			"$and": filters,
		}
	}
	return filter
}

func (s *Store) UpdateValidatorUptime(ctx context.Context, validatorsUp []string, height int64) error {
	model := []mongo.WriteModel{
		mongo.NewUpdateManyModel().
			SetFilter(
				bson.M{
					"$and": bson.A{
						bson.M{"valoper": bson.M{"$nin": validatorsUp}},
						bson.M{"missedBlocks.to": bson.M{"$eq": height}},
					},
				}).
			SetUpdate(
				bson.M{"$inc": bson.M{"missedBlocks.$[down].to": 1}},
			).
			SetArrayFilters(options.ArrayFilters{
				Filters: bson.A{bson.M{"down.to": bson.M{"$eq": height}}},
			}),
		mongo.NewUpdateManyModel().
			SetFilter(
				bson.M{
					"$and": bson.A{
						bson.M{"valoper": bson.M{"$nin": validatorsUp}},
						bson.M{"$or": bson.A{
							bson.M{"missedBlocks": bson.M{"$size": 0}},
							bson.M{"missedBlocks.to": bson.M{"$not": bson.M{"$gte": height}}},
						}},
					},
				}).
			SetUpdate(
				bson.M{
					"$push": bson.M{"missedBlocks": bson.M{"from": height, "to": height + 1}},
				},
			),
	}
	opts := options.BulkWrite().SetOrdered(true)
	_, err := s.db.Collection(validatorsCollectionName).BulkWrite(ctx, model, opts)
	return err
}
