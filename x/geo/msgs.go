package geo

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/twpayne/go-geom/encoding/ewkb"

)

type MsgStoreGeometry struct {
	Data Geometry `json:"data"`
	Signer sdk.AccAddress `json:"signer"`
}

func (MsgStoreGeometry) Route() string { return "geo" }

func (MsgStoreGeometry) Type() string { return "geo.store" }

func (msg MsgStoreGeometry) ValidateBasic() sdk.Error {
	if len(msg.Data.EWKB) <= 0 {
		return sdk.ErrUnknownRequest("GeometryEWKB cannot be empty")
	}

	_, err := ewkb.Unmarshal(msg.Data.EWKB)

	if err != nil {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Geometry is not in EWKB format: %+v", err))
	}

	// TODO validate geometry type

	return nil
}

func (msg MsgStoreGeometry) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgStoreGeometry) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
