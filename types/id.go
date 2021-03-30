package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/address"
)

type ModuleID struct {
	ModuleName string
	Path       []byte
}

func (m ModuleID) Address() sdk.AccAddress {
	return address.Module(m.ModuleName, m.Path)
}
