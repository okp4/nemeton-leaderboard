package messages

import "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"

type GetBlock struct {
	// Height of the block to get
	Height int64
}

type GetBlockResponse struct {
	Block *tmservice.Block
}
