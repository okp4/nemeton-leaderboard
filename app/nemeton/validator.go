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
	RPCEndpoint *url.URL                     `bson:"rpcEndpoint,omitempty"`
	Snapshot    *url.URL                     `bson:"snapshot,omitempty"`
	Dashboard   *url.URL                     `bson:"dashboard,omitempty"`
	Status      string                       `bson:"status"`
	Points      *uint64                      `bson:"points,omitempty"`
	Tasks       map[int]map[string]TaskState `bson:"tasks,omitempty"`
	BonusPoints *[]BonusPoints               `bson:"bonusPoints,omitempty"`
}

type BonusPoints struct {
	Points uint64 `bson:"points"`
	Reason string `bson:"reason"`
}

func MakeValidatorFromMsg(createMsg *stakingtypes.MsgCreateValidator, discord, country string, twitter *string) (*Validator, error) {
	valoper, err := types.ValAddressFromBech32(createMsg.ValidatorAddress)
	if err != nil {
		return nil, err
	}
	delegator, err := types.AccAddressFromBech32(createMsg.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	pubkey, ok := createMsg.Pubkey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, fmt.Errorf("couldn't parse public key")
	}

	return NewValidator(valoper, delegator, types.GetConsAddress(pubkey), createMsg.Description, discord, country, twitter)
}

func NewValidator(
	valoper types.ValAddress,
	delegator types.AccAddress,
	valcons types.ConsAddress,
	description stakingtypes.Description,
	discord, country string,
	twitter *string,
) (*Validator, error) {
	var website *url.URL
	var err error
	if len(description.Website) > 0 {
		website, err = url.Parse(description.Website)
		if err != nil {
			return nil, err
		}
	}

	var identity *string
	if len(description.Identity) > 0 {
		identity = &description.Identity
	}

	var details *string
	if len(description.Details) > 0 {
		details = &description.Details
	}

	return &Validator{
		Moniker:   description.Moniker,
		Identity:  identity,
		Details:   details,
		Valoper:   valoper,
		Delegator: delegator,
		Valcons:   valcons,
		Twitter:   twitter,
		Website:   website,
		Discord:   discord,
		Country:   country,
		Status:    "inactive",
	}, nil
}

func (v *Validator) Cursor() *Cursor {
	pts := uint64(0)
	if v.Points != nil {
		pts = *v.Points
	}
	return &Cursor{
		points:   pts,
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
