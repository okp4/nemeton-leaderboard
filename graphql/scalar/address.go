package scalar

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cosmos/cosmos-sdk/types"
)

func UnmarshalAccAddress(v interface{}) (types.AccAddress, error) {
	strAddress, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("expect type 'string' and got type '%T'", v)
	}

	return types.AccAddressFromBech32(strAddress)
}

func MarshalAccAddress(addr types.AccAddress) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write(addr)
	})
}

func UnmarshalValoperAddress(v interface{}) (types.ValAddress, error) {
	strAddress, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("expect type 'string' and got type '%T'", v)
	}

	return types.ValAddressFromBech32(strAddress)
}

func MarshalValoperAddress(addr types.ValAddress) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write(addr)
	})
}
