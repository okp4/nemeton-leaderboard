package event

import (
	"context"
	"time"

	"okp4/nemeton-leaderboard/app/message"

	"okp4/nemeton-leaderboard/app/event"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type Actor struct {
	mongoURI string
	dbName   string
	store    *event.Store
}

func NewPublisherActor(mongoURI, dbName string) *Actor {
	return &Actor{
		mongoURI: mongoURI,
		dbName:   dbName,
	}
}

func (a *Actor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		a.handleStart()
	case *message.PublishEventMessage:
		a.handlePublishEvent(msg)
	}
}

func (a *Actor) handleStart() {
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)
	defer cancelFn()

	store, err := event.NewStore(ctx, a.mongoURI, a.dbName)
	if err != nil {
		log.Fatal().Err(err).Str("uri", a.mongoURI).Str("db", a.dbName).Msg("❌ Couldn't create event store")
	}
	a.store = store
}

func (a *Actor) handlePublishEvent(msg *message.PublishEventMessage) {
	if err := a.store.Store(context.Background(), msg.Event); err != nil {
		log.Fatal().Err(err).Str("type", msg.Event.Type()).Msg("❌ Couldn't publish event")
	}
	log.Info().Str("type", msg.Event.Type()).Msg("💌 Event published")
}
