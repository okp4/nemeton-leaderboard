package nemeton

import (
	"context"
	"errors"
	"fmt"
	"time"

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

var omitMissedBlocks = bson.M{"missedBlocks": 0}

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

func (s *Store) Close(ctx context.Context) error {
	return s.db.Client().Disconnect(ctx)
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
				{Keys: bson.M{"valcons": 1}, Options: options.Index().SetUnique(true)},
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

// GetCurrentPhase returns the phase which is in progress, if any.
// WARNING: Do not use this when processing and event, the event sourcing shall use only the event's context: we need
// the phase in progress at the time of the event, not now.
func (s *Store) GetCurrentPhase() *Phase {
	return s.GetCurrentPhaseAt(time.Now())
}

func (s *Store) GetCurrentPhaseAt(at time.Time) *Phase {
	for _, phase := range s.phases {
		if phase.InProgressAt(at) {
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
	return s.GetValidatorBy(ctx, bson.M{"valoper": addr})
}

func (s *Store) GetValidatorByDelegator(ctx context.Context, addr types.AccAddress) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"delegator": addr})
}

func (s *Store) GetValidatorByDiscord(ctx context.Context, discord string) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"discord": discord})
}

func (s *Store) GetValidatorByTwitter(ctx context.Context, twitter string) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"twitter": twitter})
}

func (s *Store) GetValidatorBy(ctx context.Context, filter bson.M) (*Validator, error) {
	res := s.db.Collection(validatorsCollectionName).
		FindOne(ctx, filter, options.FindOne().SetProjection(omitMissedBlocks))
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
		options.Find().SetProjection(omitMissedBlocks),
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

func (s *Store) CreateValidator(
	ctx context.Context,
	createdAt time.Time,
	discord, country string,
	twitter *string,
	genTX map[string]interface{},
) error {
	msgCreateVal, err := ParseGenTX(genTX)
	if err != nil {
		return err
	}

	validator, err := MakeValidator(msgCreateVal, discord, country, twitter)
	if err != nil {
		return err
	}

	points := uint64(0)
	var tasks map[int]map[string]TaskState
	if p := s.GetCurrentPhaseAt(createdAt); p != nil {
		for _, task := range p.Tasks {
			if task.Type == taskTypeGentx && task.InProgressAt(createdAt) {
				points = *task.Rewards
				tasks = map[int]map[string]TaskState{
					p.Number: {
						task.ID: {
							Completed:    true,
							EarnedPoints: *task.Rewards,
						},
					},
				}
				break
			}
		}
	}

	validator.Points = points
	validator.Tasks = tasks

	_, err = s.db.Collection(validatorsCollectionName).InsertOne(ctx, validator)
	return err
}

func (s *Store) UpdateValidatorUptime(ctx context.Context, consensusAddrs []types.ConsAddress, height int64) error {
	model := []mongo.WriteModel{
		mongo.NewUpdateManyModel().
			SetFilter(
				bson.M{
					"$and": bson.A{
						bson.M{"valcons": bson.M{"$nin": consensusAddrs}},
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
						bson.M{"valcons": bson.M{"$nin": consensusAddrs}},
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

func (s *Store) CompleteTweetTask(ctx context.Context, when time.Time, username string, phase *Phase, task Task) error {
	filter := bson.M{"twitter": username}
	if !phase.InProgressAt(when) || !task.InProgressAt(when) {
		return fmt.Errorf("could not complete task since task or phase is not in progress at %s", when.Format(time.RFC3339))
	}

	return s.ensureTaskCompleted(ctx, filter, phase.Number, task.ID, *task.Rewards)
}

func (s *Store) CompleteNodeSetupTask(ctx context.Context, when time.Time, vals []types.ConsAddress) error {
	phase := s.GetCurrentPhaseAt(when)
	if phase != nil {
		for _, task := range phase.Tasks {
			if task.Type == taskTypeNodeSetup && task.InProgressAt(when) {
				return s.ensureTaskCompleted(ctx, bson.M{"valcons": bson.M{"$in": vals}}, phase.Number, task.ID, *task.Rewards)
			}
		}
	}
	return nil
}

func (s *Store) ensureTaskCompleted(ctx context.Context, filter bson.M, phase int, task string, rewards uint64) error {
	_, err := s.db.Collection(validatorsCollectionName).UpdateMany(ctx,
		bson.M{
			"$and": bson.A{
				filter,
				bson.M{fmt.Sprintf("tasks.%d.%s.completed", phase, task): bson.M{"$ne": true}},
			},
		},
		bson.M{
			"$set": bson.M{
				fmt.Sprintf("tasks.%d.%s.completed", phase, task): true,
				fmt.Sprintf("tasks.%d.%s.points", phase, task):    rewards,
			},
			"$inc": bson.M{
				"points": rewards,
			},
		})
	return err
}

func (s *Store) UpdatePhaseBlocks(ctx context.Context, blockTime time.Time, height int64) error {
	_, err := s.db.Collection(phasesCollectionName).UpdateOne(ctx, bson.M{
		"startDate": bson.M{"$lte": blockTime},
		"endDate":   bson.M{"$gt": blockTime},
	}, bson.A{
		bson.M{
			"$set": bson.M{
				"blocks": bson.M{
					"$ifNull": bson.A{
						"$blocks",
						bson.M{"from": height, "to": height + 1},
						"$blocks",
					},
				},
			},
		},
		bson.M{"$set": bson.M{"blocks.to": height + 1}},
	})
	return err
}

func (s *Store) GetPhaseBlocks(ctx context.Context, number int) (*BlockRange, error) {
	var phase struct {
		Blocks *BlockRange `bson:"blocks"`
	}
	if err := s.db.Collection(phasesCollectionName).FindOne(ctx,
		bson.M{"_id": number},
		options.FindOne().SetProjection(bson.M{"blocks": 1}),
	).Decode(&phase); err != nil {
		return nil, err
	}

	return phase.Blocks, nil
}

func (s *Store) GetPreviousPhaseByBlock(ctx context.Context, height int64) (*Phase, error) {
	res := s.db.Collection(phasesCollectionName).
		FindOne(ctx, bson.M{"blocks.to": height - 1}, options.FindOne())
	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	var phase Phase
	return &phase, res.Decode(&phase)
}
