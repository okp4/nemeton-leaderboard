package synchronization

import (
	"context"
	"fmt"
	"time"

	"okp4/nemeton-leaderboard/app/event"
	"okp4/nemeton-leaderboard/app/message"
	"okp4/nemeton-leaderboard/app/offset"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/scheduler"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/rs/zerolog/log"
)

const ownerOffset = "block-synchronization"

type Actor struct {
	context         context.Context
	grpcClientProps *actor.Props
	grpcClient      *actor.PID
	eventStore      *actor.PID
	offsetStore     *offset.Store
	currentBlock    int64
}

func NewActor(grpcClientProps *actor.Props, eventStore *actor.PID, mongoURI, dbName string) (*Actor, error) {
	ctx := context.Background()
	store, err := offset.NewStore(ctx, mongoURI, dbName, ownerOffset)
	if err != nil {
		return nil, err
	}

	storeValue, _ := store.Get(ctx)
	var currentBlock int64
	switch resp := storeValue.(type) {
	case int64:
		currentBlock = resp
	default:
		currentBlock = 1
	}

	return &Actor{
		context:         ctx,
		grpcClientProps: grpcClientProps,
		grpcClient:      nil,
		eventStore:      eventStore,
		offsetStore:     store,
		currentBlock:    currentBlock,
	}, nil
}

func (a *Actor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		log.Info().Msg("🔁 Start block syncing")

		a.grpcClient = ctx.Spawn(a.grpcClientProps)

		err := a.catchUpSyncBlocks(ctx)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Could not catch up to latest block sync")
		}

		scheduler.NewTimerScheduler(ctx).SendRepeatedly(0, 5*time.Second, ctx.Self(), &message.SyncBlock{})
	case *message.SyncBlock:
		a.syncBlock(ctx)
	case *actor.Restarting, *actor.Stopping:
		log.Info().Msg("🛑 Stop block syncing")
		if err := a.offsetStore.Close(context.Background()); err != nil {
			log.Err(err).Msg("❌ Couldn't properly close offset store")
		}
	}
}

func (a *Actor) syncBlock(ctx actor.Context) {
	block, err := a.getBlock(ctx, a.currentBlock+1)
	if err != nil {
		log.Err(err).Msg("❌ Could not get block.")
		return
	}

	err = a.publishEvent(ctx, block)
	if err != nil {
		log.Err(err).Msg("❌ Failed to marshall event to map interface")
		return
	}

	a.currentBlock = block.Header.Height
}

func (a *Actor) catchUpSyncBlocks(ctx actor.Context) error {
	// Check first the latest latestBlock height before sync
	result, err := ctx.RequestFuture(a.grpcClient, &message.GetLatestBlock{}, 5*time.Second).Result()
	if err != nil {
		return err
	}

	var latestBlock *tmservice.Block
	switch resp := result.(type) {
	case *message.GetBlockResponse:
		latestBlock = resp.Block
	default:
		return fmt.Errorf("wrong response message")
	}

	if a.currentBlock+1 >= latestBlock.Header.Height {
		return nil
	}

	log.Info().
		Int64("currentBlock", a.currentBlock).
		Int64("latestBlock", latestBlock.Header.Height).
		Msg("Need to catch up to latest block.")

	for i := a.currentBlock + 1; i <= latestBlock.Header.Height; i++ {
		block, err := a.getBlock(ctx, i)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Could not get block for sync.")
		}

		err = a.publishEvent(ctx, block)
		if err != nil {
			log.Panic().Err(err).Msg("❌ Failed publish block event on sync")
		}

		log.Info().Int64("blockHeight", block.Header.Height).Msg("Successful request block on sync")
	}

	a.currentBlock = latestBlock.Header.Height
	return nil
}

func (a *Actor) getBlock(ctx actor.Context, height int64) (*tmservice.Block, error) {
	result, err := ctx.RequestFuture(a.grpcClient, &message.GetBlock{Height: height}, 5*time.Second).Result()
	if err != nil {
		log.Err(err).Msg("⚠️ Failed request current block.")
		return nil, err
	}

	var block *tmservice.Block
	switch resp := result.(type) {
	case *message.GetBlockResponse:
		block = resp.Block
	default:
		return nil, fmt.Errorf("wrong response message")
	}

	return block, nil
}

func (a *Actor) publishEvent(ctx actor.Context, block *tmservice.Block) error {
	blockEvent := NewBlockEvent{
		Height:     block.Header.Height,
		Time:       block.Header.Time,
		Signatures: block.LastCommit.Signatures,
	}

	blockData, err := blockEvent.Marshall()
	if err != nil {
		return err
	}

	ctx.Send(a.eventStore, &message.PublishEventMessage{Event: event.NewEvent(NewBlockEventType, blockData)})

	if a.offsetStore.Save(a.context, block.Header.Height) != nil {
		return err
	}

	return nil
}
