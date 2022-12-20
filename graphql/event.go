package graphql

import (
	"encoding/json"
	"net/url"
)

const (
	GenTXSubmittedEventType      = "gentx-submitted"
	RegisterRPCEndpointEventType = "register-rpc-endpoint"
)

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

type RegisterRPCEndpointEvent struct {
	Moniker string   `json:"moniker"`
	URL     *url.URL `json:"url"`
}

func (e *RegisterRPCEndpointEvent) Marshall() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func UnmarshallRegisterRPCEndpointEvent(data map[string]interface{}) (*RegisterRPCEndpointEvent, error) {
	var event *RegisterRPCEndpointEvent
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(d, &event)
	return event, err
}
