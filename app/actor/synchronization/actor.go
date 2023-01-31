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
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/rs/zerolog/log"
)

const ownerOffset = "block-synchronization"

type Actor struct {
	context      context.Context
	grpcClient   *actor.PID
	eventStore   *actor.PID
	offsetStore  *offset.Store
	currentBlock int64
	txDecoder    types.TxDecoder
}

func NewActor(eventStore, grpcClient *actor.PID, mongoURI, dbName string) (*Actor, error) {
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

	txDecoder := simapp.MakeTestEncodingConfig().TxConfig.TxDecoder()

	return &Actor{
		context:      ctx,
		grpcClient:   grpcClient,
		eventStore:   eventStore,
		offsetStore:  store,
		currentBlock: currentBlock,
		txDecoder:    txDecoder,
	}, nil
}

func (a *Actor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		log.Info().Msg("🔁 Start block syncing")

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
		log.Err(err).Msg("❌ Failed publish block event")
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
	filteredMsgs(&blockEvent, a.txDecoder, block.Data.Txs)

	blockData, err := blockEvent.Marshal()
	if err != nil {
		return err
	}

	ctx.Send(a.eventStore, &message.PublishEventMessage{Event: event.NewEvent(NewBlockEventType, blockData)})

	if a.offsetStore.Save(a.context, block.Header.Height) != nil {
		return err
	}

	return nil
}

func filteredMsgs(block *NewBlockEvent, decoder types.TxDecoder, txs [][]byte) {
	msgVotes := make([]v1.MsgVote, 0)
	for _, tx := range txs {
		txDecoded, err := decoder(tx)
		if err != nil {
			log.Err(err).Msg("💱 Failed decode transaction")
			continue
		}

		for _, msg := range txDecoded.GetMsgs() {
			switch voteMsg := msg.(type) {
			case *v1.MsgVote:
				msgVotes = append(msgVotes, *voteMsg)
			default:
				log.Info().Interface("msg", msg).Msg("Skip message from transaction")
			}
		}
	}
	block.MsgVotes = msgVotes
}
