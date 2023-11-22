package v1

import sdk "github.com/cosmos/cosmos-sdk/types"

func init() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
}
