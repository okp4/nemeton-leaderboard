package graphql

import (
	"context"
	"fmt"

	"okp4/nemeton-leaderboard/app/keybase"

	"okp4/nemeton-leaderboard/app/nemeton"
	"okp4/nemeton-leaderboard/graphql"
	"okp4/nemeton-leaderboard/graphql/generated"

	graphql2 "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/asynkron/protoactor-go/actor"
)

var ErrUnauthorized = fmt.Errorf("unauthorized")

const bearerCTXKey = "bearer"

func NewGraphQLServer(
	ctx context.Context,
	actorCTX actor.Context,
	mongoURI, db string,
	eventStore *actor.PID,
	bearer *string,
) (*handler.Server, error) {
	store, err := nemeton.NewStore(ctx, mongoURI, db)
	if err != nil {
		return nil, err
	}

	cfg := generated.Config{Resolvers: graphql.NewResolver(actorCTX, store, keybase.NewClient(), eventStore)}
	cfg.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql2.Resolver) (interface{}, error) {
		if err := Authorize(ctx, bearer); err != nil {
			return nil, err
		}
		return next(ctx)
	}
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(
			cfg,
		),
	), nil
}

func Authorize(ctx context.Context, expectedBearer *string) error {
	if expectedBearer == nil {
		return nil
	}

	if b := GetBearerFrom(ctx); b == nil || *expectedBearer != *b {
		return ErrUnauthorized
	}
	return nil
}

func GetBearerFrom(ctx context.Context) *string {
	if bearer, ok := ctx.Value(bearerCTXKey).(string); ok {
		return &bearer
	}
	return nil
}
