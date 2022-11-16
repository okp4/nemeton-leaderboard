package system

import (
    "okp4/nemeton-leaderboard/app/actor/graphql"
    "okp4/nemeton-leaderboard/app/message"

    "github.com/asynkron/protoactor-go/actor"
    "github.com/rs/zerolog/log"
)

type App struct {
    ctx  *actor.RootContext
    init *actor.PID
}

func Bootstrap(listenAddr string) *App {
    initProps := actor.PropsFromFunc(func(ctx actor.Context) {
        switch ctx.Message().(type) {
        case *actor.Started:
            boot(ctx, listenAddr)
        }
    })

    ctx := actor.NewActorSystem().Root
    initPID, err := ctx.SpawnNamed(initProps, "init")
    if err != nil {
        log.Panic().Err(err).Msg("Could not create init actor")
    }

    return &App{
        ctx:  ctx,
        init: initPID,
    }
}

func (app *App) Stop() error {
    return app.ctx.StopFuture(app.init).Wait()
}

func boot(ctx actor.Context, listenAddr string) {
    graphqlProps := actor.PropsFromProducer(func() actor.Actor {
        return graphql.NewActor()
    })

    graphqlPID, err := ctx.SpawnNamed(graphqlProps, "graphql")
    if err != nil {
        log.Panic().Err(err).Msg("Could not create graphql actor")
    }

    ctx.Send(graphqlPID, &message.Start{
        ListenAddr: listenAddr,
    })
}
