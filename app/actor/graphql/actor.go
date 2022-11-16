package graphql

import (
	"context"
	"okp4/nemeton-leaderboard/app/message"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type Actor struct {
	srv *server
}

func NewActor() *Actor {
	return &Actor{}
}

func (a *Actor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *message.Start:
		a.handleStart(msg)
	case *actor.Stopping:
		a.handleStop()
	}
}

func (a *Actor) handleStart(msg *message.Start) {
	if a.srv != nil {
		log.Warn().Msg("GraphQL server already started.")
		return
	}
	a.srv = makeHTTPServer(
		msg.ListenAddr,
		makeRouter(),
	)
	a.srv.start()
	log.Info().Str("addr", msg.ListenAddr).Msg("GraphQL server started")
}

func (a *Actor) handleStop() {
	ctx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFn()

	log.Info().Msg("Stopping GraphQL server...")
	if err := a.srv.stop(ctx); err != nil {
		log.Err(err).Msg("Couldn't stop GraphQL server")
	}
}
