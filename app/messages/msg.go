package messages

import (
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
)

// GetBlock is the message sent to actor that will request blockchain to get the block
// at a given height.
type GetBlock struct {
	// Height of the block to get
	Height int64
}

// GetLatestBlock Request the latest block of the chain.
type GetLatestBlock struct{}

// GetBlockResponse is the response of GetBlock or GetLatestBlock message.
type GetBlockResponse struct {
	Block *tmservice.Block
}

// NewEvent is the message sent to actor that will handle event.
type NewEvent[E Event] struct {
	Event *E
}
