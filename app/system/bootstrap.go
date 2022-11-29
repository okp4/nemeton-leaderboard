package system

import (
	"okp4/nemeton-leaderboard/app/actor/cosmos"
	"okp4/nemeton-leaderboard/app/actor/event"
	"okp4/nemeton-leaderboard/app/actor/graphql"
	"okp4/nemeton-leaderboard/app/actor/subscription"
	"okp4/nemeton-leaderboard/app/actor/synchronization"
	"okp4/nemeton-leaderboard/app/actor/tweet"

	"okp4/nemeton-leaderboard/app/message"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/credentials"
)

type App struct {
	ctx  *actor.RootContext
	init *actor.PID
}

func Bootstrap(listenAddr, mongoURI, dbName, grpcAddr, twitterToken, hashtag string, tls credentials.TransportCredentials) *App {
	initProps := actor.PropsFromFunc(func(ctx actor.Context) {
		if _, ok := ctx.Message().(*actor.Started); ok {
			boot(ctx, listenAddr, mongoURI, dbName, grpcAddr, twitterToken, hashtag, tls)
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

func boot(ctx actor.Context, listenAddr, mongoURI, dbName, grpcAddr, twitterToken, hashtag string,
	tls credentials.TransportCredentials,
) {
	grpcClientProps := actor.PropsFromProducer(func() actor.Actor {
		grpcClient, err := cosmos.NewGrpcClient(grpcAddr, tls)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Could not create grpc client")
		}

		return grpcClient
	})

	eventStoreProps := actor.PropsFromProducer(func() actor.Actor {
		return event.NewEventStoreActor(mongoURI, dbName)
	})
	eventStorePID, err := ctx.SpawnNamed(eventStoreProps, "event-store")
	if err != nil {
		log.Panic().Err(err).Str("actor", "event-store").Msg("❌ Could not create actor")
	}

	startSubscriber(ctx, eventStorePID, mongoURI, dbName)

	blockSync := actor.PropsFromProducer(func() actor.Actor {
		sync, err := synchronization.NewActor(grpcClientProps, eventStorePID, mongoURI, dbName)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Could not start block synchronisation actor")
		}
		return sync
	})
	if _, err := ctx.SpawnNamed(blockSync, "blockSync"); err != nil {
		log.Panic().Err(err).Msg("❌ Could not create block sync actor")
	}

	tweetProps := actor.PropsFromProducer(func() actor.Actor {
		actor, err := tweet.NewSearchActor(eventStorePID, mongoURI, dbName, twitterToken, hashtag)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Could not start tweet actor")
		}
		return actor
	})
	if _, err := ctx.SpawnNamed(tweetProps, "tweet"); err != nil {
		log.Panic().Err(err).Msg("❌ Could not create tweet sync actor")
	}

	graphqlProps := actor.PropsFromProducer(func() actor.Actor {
		return graphql.NewActor(listenAddr, mongoURI, dbName)
	})
	if _, err := ctx.SpawnNamed(graphqlProps, "graphql"); err != nil {
		log.Panic().Err(err).Str("actor", "graphql").Msg("❌ Could not create actor")
	}
}

func startSubscriber(ctx actor.Context, eventPID *actor.PID, mongoURI, dbName string) {
	blockSubscriberProps := actor.PropsFromProducer(func() actor.Actor {
		s, err := subscription.NewBlock(mongoURI, dbName)
		if err != nil {
			log.Panic().Err(err).Msg("❌ failed instantiate event subscription actor")
		}
		return s
	})
	blockSubscriberPID, err := ctx.SpawnNamed(blockSubscriberProps, "blockSubscriber")
	if err != nil {
		log.Panic().Err(err).Str("actor", "graphql").Msg("❌ Could not create actor")
	}
	ctx.Send(eventPID, &message.SubscribeEventMessage{
		PID:  blockSubscriberPID,
		From: nil,
	})
}
