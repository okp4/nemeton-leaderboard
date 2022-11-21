package offset

import (
	"context"

	"okp4/nemeton-leaderboard/app/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "offsets"

type Store struct {
	owner string
	db    *mongo.Database
}

func NewStore(ctx context.Context, mongoURI, dbName, owner string) (*Store, error) {
	db, err := util.OpenMongoDatabase(ctx, mongoURI, dbName)
	if err != nil {
		return nil, err
	}

	return &Store{
		owner: owner,
		db:    db,
	}, nil
}

func (s *Store) Save(ctx context.Context, offset interface{}) error {
	_, err := s.db.Collection(collectionName).
		UpdateOne(
			ctx,
			bson.M{
				"owner": s.owner,
			},
			bson.M{
				"$set": bson.M{
					"owner":  s.owner,
					"offset": offset,
				},
			},
		)
	return err
}

func (s *Store) Get(ctx context.Context) (interface{}, error) {
	res := s.db.Collection(collectionName).
		FindOne(
			ctx,
			bson.M{
				"owner": s.owner,
			},
		)
	if err := res.Err(); err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := res.Decode(&data); err != nil {
		return nil, err
	}

	return data["offset"], nil
}