package graphql

import "encoding/json"

const GenTXSubmittedEventType = "gentx-submitted"

type GenTXSubmittedEvent struct {
	Twitter *string                `json:"twitter,omitempty"`
	Discord string                 `json:"discord"`
	Country string                 `json:"country"`
	GenTX   map[string]interface{} `json:"gentx"`
}

func (e *GenTXSubmittedEvent) Marshall() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func Unmarshall(data map[string]interface{}) (*GenTXSubmittedEvent, error) {
	var event *GenTXSubmittedEvent
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(d, &event)
	return event, err
}
