package graphql

import (
	"okp4/nemeton-leaderboard/app/keybase"
	"okp4/nemeton-leaderboard/app/nemeton"

	"github.com/asynkron/protoactor-go/actor"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type ContextKey int32

const CTXBearerKey = ContextKey(0)

type Resolver struct {
	actorCTX      actor.Context
	store         *nemeton.Store
	keybaseClient *keybase.Client
	eventStore    *actor.PID
	grpcClient    *actor.PID
}

func NewResolver(
	ctx actor.Context,
	store *nemeton.Store,
	keybaseClient *keybase.Client,
	eventStore, grpcClient *actor.PID,
) *Resolver {
	return &Resolver{
		actorCTX:      ctx,
		store:         store,
		keybaseClient: keybaseClient,
		eventStore:    eventStore,
		grpcClient:    grpcClient,
	}
}
