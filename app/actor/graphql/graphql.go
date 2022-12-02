package graphql

import (
	"context"

	"okp4/nemeton-leaderboard/app/keybase"

	"okp4/nemeton-leaderboard/app/nemeton"
	"okp4/nemeton-leaderboard/graphql"
	"okp4/nemeton-leaderboard/graphql/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/asynkron/protoactor-go/actor"
)

func NewGraphQLServer(ctx context.Context, actorCTX actor.Context, mongoURI, db string, eventStore *actor.PID) (*handler.Server, error) {
	store, err := nemeton.NewStore(ctx, mongoURI, db)
	if err != nil {
		return nil, err
	}

	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: graphql.NewResolver(actorCTX, store, keybase.NewClient(), eventStore)},
		),
	), nil
}
