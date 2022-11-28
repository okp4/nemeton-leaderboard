package synchronization

import (
	"encoding/json"
	"time"

	"github.com/tendermint/tendermint/proto/tendermint/types"
)

const NewBlockEventType = "new-block"

type NewBlockEvent struct {
	Height     int64             `json:"height"`
	Time       time.Time         `json:"time"`
	Signatures []types.CommitSig `json:"signatures"`
}

func (e *NewBlockEvent) Marshall() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func Unmarshall(data map[string]interface{}) (*NewBlockEvent, error) {
	var event *NewBlockEvent
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(d, &event)
	return event, err
}
