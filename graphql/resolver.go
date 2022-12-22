package graphql

import (
	"fmt"
	"time"

	"okp4/nemeton-leaderboard/app/keybase"
	"okp4/nemeton-leaderboard/app/message"
	"okp4/nemeton-leaderboard/app/nemeton"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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

func (r *Resolver) FetchValidator(valoper types.ValAddress) (*stakingtypes.Validator, error) {
	res, err := r.actorCTX.RequestFuture(
		r.grpcClient,
		&message.GetValidator{
			Valoper: valoper,
		},
		5*time.Second,
	).Result()
	if err != nil {
		return nil, err
	}

	val, ok := res.(*message.GetValidatorResponse)
	if !ok {
		return nil, fmt.Errorf("cannot read grpc validator response")
	}

	if val.Validator == nil {
		return nil, fmt.Errorf("could not find validator")
	}

	return val.Validator, nil
}
