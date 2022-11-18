package event

import (
	"okp4/nemeton-leaderboard/app/messages"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type EventHandler struct{}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (handler *EventHandler) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *messages.NewEvent[ReceiveNewBlock]:
		handler.handleNewBlockEvent(msg.Event)
	}
}

func (handler *EventHandler) handleNewBlockEvent(event *ReceiveNewBlock) {
	log.Info().Interface("event", &event).Msg("📦 Receive new block event")
}
