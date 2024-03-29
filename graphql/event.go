package graphql

import (
	"encoding/json"
	"net/url"

	"github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	GenTXSubmittedEventType       = "gentx-submitted"
	ValidatorRegisteredEventType  = "validator-registered"
	ValidatorUpdatedEventType     = "validator-updated"
	ValidatorRemovedEventType     = "validator-removed"
	TaskCompletedEventType        = "task-completed"
	TaskSubmittedEventType        = "task-submitted"
	RegisterURLEventType          = "register-url-endpoint"
	BonusPointsSubmittedEventType = "bonus-points-submitted"
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

type ValidatorUpdatedEvent struct {
	Delegator   types.AccAddress         `json:"delegator"`
	Twitter     *string                  `json:"twitter,omitempty"`
	Discord     string                   `json:"discord"`
	Country     string                   `json:"country"`
	Valoper     types.ValAddress         `json:"valoper"`
	Valcons     types.ConsAddress        `json:"valcons"`
	Description stakingtypes.Description `json:"description"`
}

func (e *ValidatorUpdatedEvent) Marshal() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func (e *ValidatorUpdatedEvent) Unmarshal(data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, e)
}

type ValidatorRemovedEvent struct {
	Validator types.ValAddress `json:"validator"`
}

func (e *ValidatorRemovedEvent) Marshal() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func (e *ValidatorRemovedEvent) Unmarshal(data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, e)
}

type RegisterURLEvent struct {
	Type      string           `json:"type"`
	Validator types.ValAddress `json:"validator"`
	URL       *url.URL         `json:"url"`
	Points    *uint64          `json:"rewards,omitempty"`
}

func (e *RegisterURLEvent) Marshal() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func (e *RegisterURLEvent) Unmarshal(data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, e)
}

type TaskSubmittedEvent struct {
	Validator types.ValAddress `json:"validator"`
	Phase     int              `json:"phase"`
	Task      string           `json:"task"`
}

func (e *TaskSubmittedEvent) Marshal() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func (e *TaskSubmittedEvent) Unmarshal(data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, e)
}

type TaskCompletedEvent struct {
	Validator types.ValAddress `json:"validator"`
	Phase     int              `json:"phase"`
	Task      string           `json:"task"`
	Points    *uint64          `json:"points,omitempty"`
	Override  bool             `json:"override"`
}

func (e *TaskCompletedEvent) Marshal() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func (e *TaskCompletedEvent) Unmarshal(data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, e)
}

type BonusPointsSubmittedEvent struct {
	Validator types.ValAddress `json:"validator"`
	Points    uint64           `json:"points"`
	Reason    string           `json:"reason"`
}

func (e *BonusPointsSubmittedEvent) Marshal() (map[string]interface{}, error) {
	var event map[string]interface{}
	data, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &event)
	return event, err
}

func (e *BonusPointsSubmittedEvent) Unmarshal(data map[string]interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(d, e)
}
