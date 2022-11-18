package system

import (
	"okp4/nemeton-leaderboard/app/actor/cosmos"
	"okp4/nemeton-leaderboard/app/actor/graphql"

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
	_ = actor.PropsFromProducer(func() actor.Actor {
		grpcClient, err := cosmos.NewGrpcClient(grpcAddr, tls)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Could not create grpc client")
		}

		return grpcClient
	})

	graphqlProps := actor.PropsFromProducer(func() actor.Actor {
		return graphql.NewActor(listenAddr)
	})

	if _, err := ctx.SpawnNamed(graphqlProps, "graphql"); err != nil {
		log.Panic().Err(err).Msg("❌Could not create graphql actor")
	}
}
