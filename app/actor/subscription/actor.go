package subscription

import (
	"context"

	"okp4/nemeton-leaderboard/app/actor/synchronization"
	"okp4/nemeton-leaderboard/app/event"
	"okp4/nemeton-leaderboard/app/message"
	"okp4/nemeton-leaderboard/app/nemeton"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
)

type Actor struct {
	store *nemeton.Store
	ctx   context.Context
}

func NewBlock(mongoURI, dbName string) (*Actor, error) {
	ctx := context.Background()
	store, err := nemeton.NewStore(ctx, mongoURI, dbName)
	if err != nil {
		return nil, err
	}
	return &Actor{
		store: store,
		ctx:   ctx,
	}, nil
}

func (a *Actor) Receive(ctx actor.Context) {
	switch e := ctx.Message().(type) {
	case *actor.Started:
		log.Info().Msg("Start looking for NewBlock event")
	case *message.NewEventMessage:
		a.receiveNewEvent(e.Event)
	case *actor.Stopping:
		log.Info().Msg("Stop looking new block event")
	}
}

func (a *Actor) receiveNewEvent(e event.Event) {
	log.Info().Str("type", e.Type).Msg("📦 Receive event")
	switch e.Type {
	case synchronization.NewBlockEventType:
		a.handleNewBlockEvent(e.Data)
	default:
		log.Warn().Msg("⚠️ No event handler for this event.")
	}
}

func (a Actor) handleNewBlockEvent(data map[string]interface{}) {
	e, err := synchronization.Unmarshall(data)
	if err != nil {
		return
	}

	var valopers []types.ValAddress
	for _, signature := range e.Signatures {
		valopers = append(valopers, signature.GetValidatorAddress())
	}

	if err := a.store.UpdateValidatorUptime(a.ctx, valopers, e.Height); err != nil {
		log.Panic().Err(err).Msg("🤕 Failed update validator uptime.")
	}

	log.Info().Interface("event", data).Msg("Handle NewBlock event")
}
