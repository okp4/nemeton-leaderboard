package block

import (
	"fmt"
	"time"

	"okp4/nemeton-leaderboard/app/messages"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/rs/zerolog/log"
)

type Actor struct {
	grpcClientProps *actor.Props
	grpcClient      *actor.PID
	currentBlock    int64
}

func NewActor(grpcClientProps *actor.Props, blockHeight int64) *Actor {
	return &Actor{
		grpcClientProps: grpcClientProps,
		grpcClient:      nil,
		currentBlock:    blockHeight,
	}
}

func (a *Actor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		a.grpcClient = ctx.Spawn(a.grpcClientProps)
		a.startSynchronization(ctx)
	}
}

func (a *Actor) startSynchronization(ctx actor.Context) {
	go func() {
		for range time.Tick(8 * time.Second) {
			result, err := ctx.RequestFuture(a.grpcClient, &messages.GetBlock{Height: a.currentBlock}, 5*time.Second).Result()
			if err != nil {
				log.Err(err).Msg("⚠️ Failed request current block.")
				continue
			}

			var block *tmservice.Block
			switch resp := result.(type) {
			case *messages.GetBlockResponse:
				block = resp.Block
			default:
				log.Panic().Err(fmt.Errorf("wrong response message")).Msg("❌ Could not get block.")
			}

			log.Info().Int64("blockHeight", block.Header.Height).Msg("Successful request block")

			// TODO: Send to event handler the new block received

			a.currentBlock += 1
		}
	}()
}
