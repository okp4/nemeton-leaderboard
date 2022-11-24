package nemeton

import (
	"context"

	"okp4/nemeton-leaderboard/app/util"

	"github.com/cosmos/cosmos-sdk/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	return nil
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

func (s *Store) GetValidatorByID(ctx context.Context, id primitive.ObjectID) (*Validator, error) {
	return s.GetValidatorBy(ctx, bson.M{"_id": id})
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
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var val Validator
	return &val, res.Decode(&val)
}
