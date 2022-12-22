package util

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

func ParseGenTX(genMap map[string]interface{}) (*types.MsgCreateValidator, error) {
	raw, err := json.Marshal(genMap)
	if err != nil {
		return nil, err
	}

	tx, err := simapp.MakeTestEncodingConfig().TxConfig.TxJSONDecoder()(raw)
	if err != nil {
		return nil, err
	}

	if len(tx.GetMsgs()) != 1 {
		return nil, fmt.Errorf("couldn't find 'MsgCreateValidator' in gentx")
	}

	if msg, ok := tx.GetMsgs()[0].(*types.MsgCreateValidator); ok {
		return msg, nil
	}
	return nil, fmt.Errorf("couldn't find 'MsgCreateValidator' in gentx")
}
