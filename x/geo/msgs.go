package geo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateGeometry struct {
	Geometry Geometry
	Signer sdk.AccAddress
}
