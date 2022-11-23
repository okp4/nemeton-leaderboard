package graphql

import (
	"context"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type Actor struct {
	addr     string
	mongoURI string
	dbName   string
	srv      *server
}

func NewActor(httpAddr, mongoURI, dbName string) *Actor {
	return &Actor{
		addr:     httpAddr,
		mongoURI: mongoURI,
		dbName:   dbName,
	}
}

func (a *Actor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		a.handleStart()
	case *actor.Stopping:
		a.handleStop()
	}
}

func (a *Actor) handleStart() {
	graphqlServer, err := NewGraphQLServer(context.Background(), a.mongoURI, a.dbName)
	if err != nil {
		log.Fatal().Err(err).Str("uri", a.mongoURI).Str("db", a.dbName).Msg("❌ Couldn't create graphql server")
	}
	makeHTTPServer(
		a.addr,
		makeRouter(graphqlServer),
	)
	a.srv.start()
	log.Info().Str("addr", a.addr).Msg("🔥 GraphQL server started")
}

func (a *Actor) handleStop() {
	ctx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFn()

	log.Info().Msg("\U0001F9EF Stopping GraphQL server...")
	if err := a.srv.stop(ctx); err != nil {
		log.Err(err).Msg("❌ Couldn't stop GraphQL server")
	}
}
