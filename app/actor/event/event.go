package event

import "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"

type ReceiveNewBlock struct {
	Height int64
	Block  *tmservice.Block
}
