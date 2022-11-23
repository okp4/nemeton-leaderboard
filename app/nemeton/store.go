package nemeton

import (
	"context"

	"okp4/nemeton-leaderboard/app/util"

	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "phases"

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
