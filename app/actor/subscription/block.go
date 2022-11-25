package subscription

import (
	"okp4/nemeton-leaderboard/app/event"
	"okp4/nemeton-leaderboard/app/message"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/rs/zerolog/log"
)

type Block struct{}

func NewBlock() *Block {
	return &Block{}
}

func (b *Block) Receive(ctx actor.Context) {
	switch e := ctx.Message().(type) {
	case *actor.Started:
		log.Info().Msg("Start looking for NewBlock event")
	case *message.NewEventMessage:
		b.receiveNewBlock(e.Event)
	case *actor.Stopping:
		log.Info().Msg("Stop looking new block event")
	}
}

func (b *Block) receiveNewBlock(e event.Event) {
	log.Info().Msg("📦 Receive block event")
}
