package store

import (
	"context"
	"time"

	"okp4/nemeton-leaderboard/app/event"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type Actor struct {
	mongoURI   string
	dbName     string
	eventStore *event.Store
}

func NewActor(mongoURI, dbName string) *Actor {
	return &Actor{
		mongoURI: mongoURI,
		dbName:   dbName,
	}
}

func (a *Actor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		a.handleStart()
	}
}

func (a *Actor) handleStart() {
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second)
	defer cancelFn()

	store, err := event.NewStore(ctx, a.mongoURI, a.dbName)
	if err != nil {
		log.Fatal().Err(err).Str("uri", a.mongoURI).Str("db", a.dbName).Msg("❌ Couldn't create event store")
	}
	a.eventStore = store
}
