package fixture

import (
	"crypto/sha256"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ModuleID struct {
	ModuleName string
	Path       []byte
}

func (m ModuleID) Address() sdk.AccAddress {
	return AddressHash(m.ModuleName, m.Path)
}

func AddressHash(prefix string, contents []byte) []byte {
	preImage := []byte(prefix)
	if len(contents) != 0 {
		preImage = append(preImage, 0)
		preImage = append(preImage, contents...)
	}
	sum := sha256.Sum256(preImage)
	return sum[:20]
}
