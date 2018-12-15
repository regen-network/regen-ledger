package geo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgRegisterGeometry struct {
	Geometry Geometry
	Signer sdk.AccAddress
}
