package graphql

import (
	"context"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type Actor struct {
	addr       string
	mongoURI   string
	dbName     string
	eventStore *actor.PID
	grpcClient *actor.PID
	bearer     *string
	srv        *server
}

func NewActor(httpAddr, mongoURI, dbName string, eventStore, grpcClient *actor.PID, bearer *string) *Actor {
	return &Actor{
		addr:       httpAddr,
		mongoURI:   mongoURI,
		dbName:     dbName,
		eventStore: eventStore,
		bearer:     bearer,
	}
}

func (a *Actor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		a.handleStart(ctx)
	case *actor.Restarting, *actor.Stopping:
		a.handleStop()
	}
}

func (a *Actor) handleStart(ctx actor.Context) {
	graphqlServer, err := NewGraphQLServer(context.Background(), ctx, a.mongoURI, a.dbName, a.eventStore, a.bearer)
	if err != nil {
		log.Fatal().Err(err).Str("db", a.dbName).Msg("❌ Couldn't create graphql server")
	}
	a.srv = makeHTTPServer(
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
