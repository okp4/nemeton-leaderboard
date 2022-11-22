package system

import (
	"time"

	"okp4/nemeton-leaderboard/app/actor/event"
	"okp4/nemeton-leaderboard/app/actor/graphql"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type App struct {
	ctx  *actor.RootContext
	init *actor.PID
}

func Bootstrap(listenAddr, mongoURI, dbName string) *App {
	initProps := actor.PropsFromFunc(func(ctx actor.Context) {
		if _, ok := ctx.Message().(*actor.Started); ok {
			boot(ctx, listenAddr, mongoURI, dbName)
		}
	}).Configure(actor.WithSupervisor(actor.NewAllForOneStrategy(
		3, time.Second, func(reason interface{}) actor.Directive {
			return actor.EscalateDirective
		},
	)))

	ctx := actor.NewActorSystem().Root.
		WithGuardian(actor.NewOneForOneStrategy(
			1, time.Second, func(reason interface{}) actor.Directive {
				return actor.EscalateDirective
			},
		))
	initPID, err := ctx.SpawnNamed(initProps, "init")
	if err != nil {
		log.Panic().Err(err).Msg("❌ Could not create init actor")
	}

	return &App{
		ctx:  ctx,
		init: initPID,
	}
}

func (app *App) Stop() error {
	return app.ctx.StopFuture(app.init).Wait()
}

func boot(ctx actor.Context, listenAddr, mongoURI, dbName string) {
	graphqlProps := actor.PropsFromProducer(func() actor.Actor {
		return graphql.NewActor(listenAddr)
	})
	if _, err := ctx.SpawnNamed(graphqlProps, "graphql"); err != nil {
		log.Panic().Err(err).Str("actor", "graphql").Msg("❌ Could not create actor")
	}

	eventStoreProps := actor.PropsFromProducer(func() actor.Actor {
		return event.NewEventStoreActor(mongoURI, dbName)
	})
	if _, err := ctx.SpawnNamed(eventStoreProps, "event-store"); err != nil {
		log.Panic().Err(err).Str("actor", "event-store").Msg("❌ Could not create actor")
	}
}
