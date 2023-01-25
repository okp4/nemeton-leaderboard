package synchronization

import (
	"encoding/json"
	"time"

	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/tendermint/tendermint/proto/tendermint/types"
)

const NewBlockEventType = "new-block"

type NewBlockEvent struct {
	Height     int64             `json:"height"`
	Time       time.Time         `json:"time"`
	Signatures []types.CommitSig `json:"signatures"`
	MsgVotes   []v1.MsgVote      `json:"msgVotes,omitempty"`
}

func (e *NewBlockEvent) Marshal() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func (e *NewBlockEvent) Unmarshal(data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, e)
}
