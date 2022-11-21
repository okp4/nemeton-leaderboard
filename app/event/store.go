package event

import (
	"context"

	"okp4/nemeton-leaderboard/app/util"

	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "events"

type Store struct {
	db *mongo.Database
}

func NewStore(ctx context.Context, mongoURI, dbName string) (*Store, error) {
	db, err := util.OpenMongoDatabase(ctx, mongoURI, dbName)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) Store(ctx context.Context, evt Event) error {
	_, err := s.db.Collection(collectionName).InsertOne(ctx, evt)
	return err
}
