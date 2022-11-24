package nemeton

import (
	"net/url"

	"github.com/cosmos/cosmos-sdk/types"
)

type Validator struct {
	Moniker   string           `bson:"moniker"`
	Identity  *string          `bson:"identity,omitempty"`
	Valoper   types.ValAddress `bson:"valoper"`
	Delegator types.AccAddress `bson:"delegator"`
	Twitter   *string          `bson:"twitter,omitempty"`
	Website   *url.URL         `bson:"website,omitempty"`
	Discord   string           `bson:"discord"`
	Country   string           `bson:"country"`
	Status    string           `bson:"status"`
	Points    int              `bson:"points"`
}
