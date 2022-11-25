package nemeton

import (
	"net/url"

	"github.com/cosmos/cosmos-sdk/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Validator struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Moniker   string             `bson:"moniker"`
	Identity  *string            `bson:"identity,omitempty"`
	Valoper   types.ValAddress   `bson:"valoper"`
	Delegator types.AccAddress   `bson:"delegator"`
	Twitter   *string            `bson:"twitter,omitempty"`
	Website   *url.URL           `bson:"website,omitempty"`
	Discord   string             `bson:"discord"`
	Country   string             `bson:"country"`
	Status    string             `bson:"status"`
	Points    uint64             `bson:"points"`
}

func (v *Validator) Cursor() *Cursor {
	return &Cursor{
		points:   v.Points,
		objectID: v.ID,
	}
}
