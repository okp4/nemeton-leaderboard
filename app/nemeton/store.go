package nemeton

import (
	"context"

	"okp4/nemeton-leaderboard/app/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "phases"

type Store struct {
	db     *mongo.Database
	phases []Phase
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
	phases := s.db.Collection(collectionName)
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
		s.phases = append(s.phases, phase)
	}

	return nil
}

func (s *Store) GetPhase(number int) *Phase {
	for i, phase := range s.phases {
		if phase.Number == number {
			return &s.phases[i]
		}
	}
	return nil
}