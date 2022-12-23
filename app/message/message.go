package message

import (
	"okp4/nemeton-leaderboard/app/event"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PublishEventMessage struct {
	Event event.Event
}

type SubscribeEventMessage struct {
	PID  *actor.PID
	From *primitive.ObjectID
}

type UnsubscribeEventMessage struct {
	PID *actor.PID
}

type NewEventMessage struct {
	Event event.Event
}

type BrokenStreamMessage struct{}

// GetBlock Ask to requets a block at a given height.
type GetBlock struct {
	// Height of the block to get
	Height int64
}

// GetLatestBlock Request the latest block of the chain.
type GetLatestBlock struct{}

type GetBlockResponse struct {
	Block *tmservice.Block
}

type GetValidator struct {
	Valoper types.ValAddress
}

type GetValidatorResponse struct {
	Validator *stakingtypes.Validator
}

// SyncBlock used for ask synchronization to request new block on chain.
type SyncBlock struct{}

// SearchTweet message to ask actor to launch tweet research.
type SearchTweet struct{}
