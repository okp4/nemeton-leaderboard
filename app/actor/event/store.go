package event

import (
	"context"

	"okp4/nemeton-leaderboard/app/message"

	"okp4/nemeton-leaderboard/app/event"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type StoreActor struct {
	mongoURI string
	dbName   string
	store    *event.Store
}

func NewEventStoreActor(mongoURI, dbName string) *StoreActor {
	return &StoreActor{
		mongoURI: mongoURI,
		dbName:   dbName,
	}
}

func (a *StoreActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		a.handleStart()
	case *actor.Stopping:
		log.Info().Msg("\U0001F9EF Stopping Event store...")
	case *message.PublishEventMessage:
		a.handlePublishEvent(msg)
	case *message.SubscribeEventMessage:
		a.handleSubscribeEvent(ctx, msg)
	}
}

func (a *StoreActor) handleStart() {
	store, err := event.NewStore(context.Background(), a.mongoURI, a.dbName)
	if err != nil {
		log.Fatal().Err(err).Str("db", a.dbName).Msg("❌ Couldn't create event store")
	}
	a.store = store
	log.Info().Msg("🚌 Event store started")
}

func (a *StoreActor) handlePublishEvent(msg *message.PublishEventMessage) {
	if err := a.store.Store(context.Background(), msg.Event); err != nil {
		log.Fatal().Err(err).Str("type", msg.Event.Type).Msg("❌ Couldn't publish event")
	}
	log.Info().Str("type", msg.Event.Type).Msg("💌 Event published")
}

func (a *StoreActor) handleSubscribeEvent(ctx actor.Context, msg *message.SubscribeEventMessage) {
	stream, err := a.store.StreamFrom(context.Background(), msg.From)
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Couldn't create stream")
	}

	streamProps := actor.PropsFromProducer(func() actor.Actor {
		return NewStreamHandlerActor(stream, msg.PID)
	})

	ctx.Spawn(streamProps)
	log.Info().Str("to", msg.PID.Address).Msg("®️ Event subscriber registered")
}
