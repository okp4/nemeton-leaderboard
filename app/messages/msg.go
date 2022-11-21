package messages

import "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"

type GetBlock struct {
	// Height of the block to get
	Height int64
}

// GetLatestBlock Request the latest block of the chain.
type GetLatestBlock struct{}

type GetBlockResponse struct {
	Block *tmservice.Block
}
