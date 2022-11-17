package store

import "github.com/asynkron/protoactor-go/actor"

type Actor struct {
	mongoURI string
	dbName   string
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
}
