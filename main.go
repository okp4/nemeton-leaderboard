package main

import (
	"okp4/nemeton-leaderboard/cmd"

	"github.com/cosmos/cosmos-sdk/types"
)

const (
	accprefix  = "okp4"
	valprefix  = "okp4valoper"
	consprefix = "okp4valcons"
)

func main() {
	conf := types.GetConfig()
	conf.SetBech32PrefixForAccount(accprefix, accprefix)
	conf.SetBech32PrefixForValidator(valprefix, valprefix)
	conf.SetBech32PrefixForConsensusNode(consprefix, consprefix)

	cmd.Execute()
}
