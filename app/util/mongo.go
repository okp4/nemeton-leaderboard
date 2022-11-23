package util

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenMongoDatabase(ctx context.Context, mongoURI, dbName string) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	return client.Database(dbName), nil
}
