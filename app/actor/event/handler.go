package event

import (
	"okp4/nemeton-leaderboard/app/messages"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type Handler struct{}

func NewEventHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *messages.NewEvent[messages.ReceiveNewBlock]:
		handler.handleNewBlockEvent(msg.Event)
	default:
		log.Warn().Msg("❌ Could not handle event")
	}
}

func (handler *Handler) handleNewBlockEvent(event *messages.ReceiveNewBlock) {
	log.Info().Str("event-name", event.Name()).Msg("📦 Receive new block event")
	// TODO: handle new block event
}
