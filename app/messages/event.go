package messages

import "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"

const (
	ReceiveNewBlockEvent = "receive-new-block"
)

// Event represent an event managed by EventHandler.
type Event interface {
	// Name return the string name of event
	Name() string
}

// ReceiveNewBlock is an event when a new block has been received.
type ReceiveNewBlock struct {
	Height int64
	Block  *tmservice.Block
}

func (e ReceiveNewBlock) Name() string {
	return ReceiveNewBlockEvent
}
