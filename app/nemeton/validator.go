package nemeton

import (
	"fmt"
	"net/url"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Validator struct {
	ID          primitive.ObjectID           `bson:"_id,omitempty"`
	Moniker     string                       `bson:"moniker"`
	Identity    *string                      `bson:"identity,omitempty"`
	Details     *string                      `bson:"details,omitempty"`
	Valoper     types.ValAddress             `bson:"valoper"`
	Delegator   types.AccAddress             `bson:"delegator"`
	Valcons     types.ConsAddress            `bson:"valcons"`
	Twitter     *string                      `bson:"twitter,omitempty"`
	Website     *url.URL                     `bson:"website,omitempty"`
	Discord     string                       `bson:"discord"`
	Country     string                       `bson:"country"`
	RPCEndpoint *url.URL                     `bson:"rpcEndpoint"`
	Status      string                       `bson:"status"`
	Points      uint64                       `bson:"points"`
	Tasks       map[int]map[string]TaskState `bson:"tasks"`
}

func MakeValidator(createMsg *stakingtypes.MsgCreateValidator, discord, country string, twitter *string) (*Validator, error) {
	valoper, err := types.ValAddressFromBech32(createMsg.ValidatorAddress)
	if err != nil {
		return nil, err
	}
	delegator, err := types.AccAddressFromBech32(createMsg.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	var website *url.URL
	if len(createMsg.Description.Website) > 0 {
		website, err = url.Parse(createMsg.Description.Website)
		if err != nil {
			return nil, err
		}
	}

	pubkey, ok := createMsg.Pubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, fmt.Errorf("couldn't parse public key")
	}

	var identity *string
	if len(createMsg.Description.Identity) > 0 {
		identity = &createMsg.Description.Identity
	}
	var details *string
	if len(createMsg.Description.Details) > 0 {
		details = &createMsg.Description.Details
	}

	return &Validator{
		Moniker:   createMsg.Description.Moniker,
		Identity:  identity,
		Details:   details,
		Valoper:   valoper,
		Delegator: delegator,
		Valcons:   types.GetConsAddress(pubkey),
		Twitter:   twitter,
		Website:   website,
		Discord:   discord,
		Country:   country,
		Status:    "inactive",
	}, nil
}

func (v *Validator) Cursor() *Cursor {
	return &Cursor{
		points:   v.Points,
		objectID: v.ID,
	}
}

func (v *Validator) Task(phase int, id string) *TaskState {
	if tasks, ok := v.Tasks[phase]; ok {
		task := tasks[id]
		return &task
	}
	return nil
}
