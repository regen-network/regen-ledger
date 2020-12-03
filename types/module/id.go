package module

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
)

type ModuleID struct {
	ModuleName string
	Path       []byte
}

func (m ModuleID) Address() sdk.AccAddress {
	return types.AddressHash(m.ModuleName, m.Path)
}

func RootModuleID(moduleName string) ModuleID {
	return ModuleID{ModuleName: moduleName}
}
