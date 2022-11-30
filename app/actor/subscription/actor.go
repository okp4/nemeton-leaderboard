package subscription

import (
	"context"

	"okp4/nemeton-leaderboard/app/actor/synchronization"
	"okp4/nemeton-leaderboard/app/event"
	"okp4/nemeton-leaderboard/app/message"
	"okp4/nemeton-leaderboard/app/nemeton"
	"okp4/nemeton-leaderboard/app/offset"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ownerOffset = "subscription"

type Actor struct {
	store       *nemeton.Store
	ctx         context.Context
	eventPID    *actor.PID
	offsetStore *offset.Store
}

func NewSubscriber(mongoURI, dbName string, eventPID *actor.PID) (*Actor, error) {
	ctx := context.Background()
	store, err := nemeton.NewStore(ctx, mongoURI, dbName)
	if err != nil {
		return nil, err
	}

	offsetStore, err := offset.NewStore(ctx, mongoURI, dbName, ownerOffset)
	if err != nil {
		return nil, err
	}

	return &Actor{
		store:       store,
		ctx:         ctx,
		eventPID:    eventPID,
		offsetStore: offsetStore,
	}, nil
}

func (a *Actor) Receive(ctx actor.Context) {
	switch e := ctx.Message().(type) {
	case *actor.Started:
		var from *primitive.ObjectID
		value, _ := a.offsetStore.Get(a.ctx)
		switch v := value.(type) {
		case primitive.ObjectID:
			from = &v
		default:
			from = nil
		}

		log.Info().Msg("🕵️ Start looking for new event")
		ctx.Send(a.eventPID, &message.SubscribeEventMessage{
			PID:  ctx.Self(),
			From: from,
		})
	case *message.NewEventMessage:
		a.receiveNewEvent(e.Event)
	case *actor.Stopping:
		log.Info().Msg("✋ Stop looking new event")
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
	if err := a.offsetStore.Save(a.ctx, e.ID); err != nil {
		log.Panic().Err(err).Msg("❌ Failed save offset state.")
	}
}

func (a *Actor) handleNewBlockEvent(data map[string]interface{}) {
	log.Info().Interface("event", data).Msg("Handle NewBlock event")

	e, err := synchronization.Unmarshall(data)
	if err != nil {
		log.Panic().Err(err).Msg("❌ Failed unmarshall event to NewBlockEvent")
		return
	}

	consensusAddr := make([]string, len(e.Signatures))
	for i, signature := range e.Signatures {
		var addr types.ConsAddress = signature.GetValidatorAddress()
		consensusAddr[i] = addr.String()
	}

	if err := a.store.UpdateValidatorUptime(a.ctx, consensusAddr, e.Height); err != nil {
		log.Panic().Err(err).Msg("🤕 Failed update validator uptime.")
	}
}
