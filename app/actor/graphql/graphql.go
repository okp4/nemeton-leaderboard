package graphql

import (
	"context"

	"okp4/nemeton-leaderboard/app/nemeton"
	"okp4/nemeton-leaderboard/graphql"
	"okp4/nemeton-leaderboard/graphql/generated"

	"github.com/99designs/gqlgen/graphql/handler"
)

func NewGraphQLServer(ctx context.Context, mongoURI, db string) (*handler.Server, error) {
	store, err := nemeton.NewStore(ctx, mongoURI, db)
	if err != nil {
		return nil, err
	}

	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: graphql.NewResolver(store)},
		),
	), nil
}
