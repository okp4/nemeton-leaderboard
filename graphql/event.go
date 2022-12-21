package graphql

import (
	"encoding/json"
	"net/url"

	"github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	GenTXSubmittedEventType      = "gentx-submitted"
	ValidatorRegisteredEventType = "validator-registered"
	RegisterRPCEndpointEventType = "register-rpc-endpoint"
)

type GenTXSubmittedEvent struct {
	Twitter *string                `json:"twitter,omitempty"`
	Discord string                 `json:"discord"`
	Country string                 `json:"country"`
	GenTX   map[string]interface{} `json:"gentx"`
}

func (e *GenTXSubmittedEvent) Marshal() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func (e *GenTXSubmittedEvent) Unmarshal(data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, e)
}

type ValidatorRegisteredEvent struct {
	Twitter     *string                  `json:"twitter,omitempty"`
	Discord     string                   `json:"discord"`
	Country     string                   `json:"country"`
	Valoper     types.ValAddress         `json:"valoper"`
	Delegator   types.AccAddress         `json:"delegator"`
	Valcons     types.ConsAddress        `json:"valcons"`
	Description stakingtypes.Description `json:"description"`
}

func (e *ValidatorRegisteredEvent) Marshal() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func (e *ValidatorRegisteredEvent) Unmarshal(data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, e)
}

type RegisterRPCEndpointEvent struct {
	Validator types.ValAddress `json:"validator"`
	URL       *url.URL         `json:"url"`
}

func (e *RegisterRPCEndpointEvent) Marshal() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func (e *RegisterRPCEndpointEvent) Unmarshal(data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, e)
}
