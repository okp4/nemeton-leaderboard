package system

import (
	"okp4/nemeton-leaderboard/app/actor/cosmos"
	"okp4/nemeton-leaderboard/app/actor/event"
	"okp4/nemeton-leaderboard/app/actor/graphql"
	"okp4/nemeton-leaderboard/app/actor/subscription"
	"okp4/nemeton-leaderboard/app/actor/synchronization"
	"okp4/nemeton-leaderboard/app/actor/tweet"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/credentials"
)

type App struct {
	ctx  *actor.RootContext
	init *actor.PID
}

func Bootstrap(
	listenAddr, mongoURI, dbName, grpcAddr, twitterToken, twitterAccount string,
	tls credentials.TransportCredentials,
	accessToken *string, noBlockSync bool,
) *App {
	initProps := actor.PropsFromFunc(func(ctx actor.Context) {
		if _, ok := ctx.Message().(*actor.Started); ok {
			boot(ctx, listenAddr, mongoURI, dbName, grpcAddr, twitterToken, twitterAccount, tls, accessToken, noBlockSync)
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

func boot(ctx actor.Context, listenAddr, mongoURI, dbName, grpcAddr, twitterToken, twitterAccount string,
	tls credentials.TransportCredentials,
	accessToken *string, noBlockSync bool,
) {
	grpcClientProps := actor.PropsFromProducer(func() actor.Actor {
		grpcClient, err := cosmos.NewGrpcClient(grpcAddr, tls)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Could not create grpc client")
		}

		return grpcClient
	})
	grpcClientPID, err := ctx.SpawnNamed(grpcClientProps, "grpc-client")
	if err != nil {
		log.Panic().Err(err).Msg("❌ Could not create grpc client actor")
	}

	eventStoreProps := actor.PropsFromProducer(func() actor.Actor {
		return event.NewEventStoreActor(mongoURI, dbName)
	})
	eventStorePID, err := ctx.SpawnNamed(eventStoreProps, "event-store")
	if err != nil {
		log.Panic().Err(err).Str("actor", "event-store").Msg("❌ Could not create actor")
	}

	startSubscriber(ctx, eventStorePID, mongoURI, dbName)

	if !noBlockSync {
		blockSync := actor.PropsFromProducer(func() actor.Actor {
			sync, err := synchronization.NewActor(eventStorePID, grpcClientPID, mongoURI, dbName)
			if err != nil {
				log.Panic().Err(err).Msg("❌ Could not start block synchronisation actor")
			}
			return sync
		})
		if _, err := ctx.SpawnNamed(blockSync, "blockSync"); err != nil {
			log.Panic().Err(err).Msg("❌ Could not create block sync actor")
		}
	}

	tweetProps := actor.PropsFromProducer(func() actor.Actor {
		actor, err := tweet.NewSearchActor(eventStorePID, mongoURI, dbName, twitterToken, twitterAccount)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Could not start tweet actor")
		}
		return actor
	})
	if _, err := ctx.SpawnNamed(tweetProps, "tweet"); err != nil {
		log.Panic().Err(err).Msg("❌ Could not create tweet sync actor")
	}

	graphqlProps := actor.PropsFromProducer(func() actor.Actor {
		return graphql.NewActor(listenAddr, mongoURI, dbName, eventStorePID, grpcClientPID, accessToken)
	})
	if _, err := ctx.SpawnNamed(graphqlProps, "graphql"); err != nil {
		log.Panic().Err(err).Str("actor", "graphql").Msg("❌ Could not create actor")
	}
}

func startSubscriber(ctx actor.Context, eventPID *actor.PID, mongoURI, dbName string) {
	subscriberProps := actor.PropsFromProducer(func() actor.Actor {
		s, err := subscription.NewSubscriber(mongoURI, dbName, eventPID)
		if err != nil {
			log.Panic().Err(err).Msg("❌ failed instantiate event subscriber actor")
		}
		return s
	})
	_, err := ctx.SpawnNamed(subscriberProps, "subscriber")
	if err != nil {
		log.Panic().Err(err).Str("actor", "subscriber").Msg("❌ Could not create actor")
	}
}
