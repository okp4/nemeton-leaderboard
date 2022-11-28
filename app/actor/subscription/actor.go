package subscription

import (
	"okp4/nemeton-leaderboard/app/actor/synchronization"
	"okp4/nemeton-leaderboard/app/event"
	"okp4/nemeton-leaderboard/app/message"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type Actor struct{}

func NewBlock() *Actor {
	return &Actor{}
}

func (a *Actor) Receive(ctx actor.Context) {
	switch e := ctx.Message().(type) {
	case *actor.Started:
		log.Info().Msg("Start looking for NewBlock event")
	case *message.NewEventMessage:
		a.receiveNewEvent(e.Event)
	case *actor.Stopping:
		log.Info().Msg("Stop looking new block event")
	}
}

func (a *Actor) receiveNewEvent(e event.Event) {
	log.Info().Str("type", e.Type).Msg("📦 Receive event")
	switch e.Type {
	case synchronization.NewBlockEventType:
		a.handleNewBLockEvent(e.Data)
	default:
		log.Warn().Msg("⚠️ No event handler for this event.")
	}
}

func (a Actor) handleNewBLockEvent(data map[string]interface{}) {
	_, err := synchronization.Unmarshall(data)
	if err != nil {
		return
	}

	log.Info().Interface("event", data).Msg("Handle NewBlock event")
}
