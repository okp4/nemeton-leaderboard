package system

import (
	"okp4/nemeton-leaderboard/app/actor/cosmos"
	"okp4/nemeton-leaderboard/app/actor/event"
	"okp4/nemeton-leaderboard/app/actor/graphql"
	"okp4/nemeton-leaderboard/app/actor/synchronization"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/credentials"
)

type App struct {
	ctx  *actor.RootContext
	init *actor.PID
}

func Bootstrap(listenAddr, grpcAddr string, tls credentials.TransportCredentials) *App {
	initProps := actor.PropsFromFunc(func(ctx actor.Context) {
		if _, ok := ctx.Message().(*actor.Started); ok {
			boot(ctx, listenAddr, grpcAddr, tls)
		}
	})

	ctx := actor.NewActorSystem().Root
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

func boot(ctx actor.Context, listenAddr, grpcAddr string, tls credentials.TransportCredentials) {
	grpcClientProps := actor.PropsFromProducer(func() actor.Actor {
		grpcClient, err := cosmos.NewGrpcClient(grpcAddr, tls)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Could not create grpc client")
		}

		return grpcClient
	})

	eventProps := actor.PropsFromProducer(func() actor.Actor {
		return event.NewEventHandler()
	})

	eventPID, err := ctx.SpawnNamed(eventProps, "event")
	if err != nil {
		log.Panic().Err(err).Msg("❌Could not create event actor")
	}

	blockSync := actor.PropsFromProducer(func() actor.Actor {
		return synchronization.NewActor(grpcClientProps, eventPID, 16757)
	})

	graphqlProps := actor.PropsFromProducer(func() actor.Actor {
		return graphql.NewActor(listenAddr)
	})

	if _, err := ctx.SpawnNamed(graphqlProps, "graphql"); err != nil {
		log.Panic().Err(err).Msg("❌Could not create graphql actor")
	}

	if _, err := ctx.SpawnNamed(blockSync, "blockSync"); err != nil {
		log.Panic().Err(err).Msg("❌Could not create block sync actor")
	}
}
