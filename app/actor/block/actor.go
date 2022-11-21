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
		log.Info().Msg("🔁 Start block syncing")
		a.grpcClient = ctx.Spawn(a.grpcClientProps)
		a.startSynchronization(ctx)
	case *actor.Stopping:
		log.Info().Msg("🛑 Stop block syncing")
	}
}

func (a *Actor) startSynchronization(ctx actor.Context) {
	err := a.catchUpSyncBlocks(ctx)
	if err != nil {
		log.Err(err).Msg("❌ Could not catch up to latest block sync")
		return
	}

	go func() {
		for range time.Tick(8 * time.Second) {
			block, err := a.getBlock(ctx, a.currentBlock)
			if err != nil {
				log.Err(err).Msg("❌ Could not get block.")
				continue
			}

			// TODO: Send to event handler the new block received
			log.Info().Int64("blockHeight", block.Header.Height).Msg("Successful request block")
			a.currentBlock++
		}
	}()
}

func (a *Actor) catchUpSyncBlocks(ctx actor.Context) error {
	// Check first the latest latestBlock height before sync
	result, err := ctx.RequestFuture(a.grpcClient, &messages.GetLatestBlock{}, 5*time.Second).Result()
	if err != nil {
		return err
	}

	var latestBlock *tmservice.Block
	switch resp := result.(type) {
	case *messages.GetBlockResponse:
		latestBlock = resp.Block
	default:
		return fmt.Errorf("wrong response message")
	}

	if a.currentBlock >= latestBlock.Header.Height {
		return nil
	}

	log.Info().
		Int64("currentBlock", a.currentBlock).
		Int64("latestBlock", latestBlock.Header.Height).
		Msg("Need to catch up to latest block.")

	for i := a.currentBlock; i <= latestBlock.Header.Height; i++ {
		block, err := a.getBlock(ctx, i)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Could not get block for sync.")
			continue
		}
		// TODO: Send to event handler the new latestBlock received
		log.Info().Int64("blockHeight", block.Header.Height).Msg("Successful request block on sync")
	}

	a.currentBlock = latestBlock.Header.Height + 1
	return nil
}

func (a *Actor) getBlock(ctx actor.Context, height int64) (*tmservice.Block, error) {
	result, err := ctx.RequestFuture(a.grpcClient, &messages.GetBlock{Height: height}, 5*time.Second).Result()
	if err != nil {
		log.Err(err).Msg("⚠️ Failed request current block.")
		return nil, err
	}

	var block *tmservice.Block
	switch resp := result.(type) {
	case *messages.GetBlockResponse:
		block = resp.Block
	default:
		return nil, fmt.Errorf("wrong response message")
	}

	return block, nil
}
