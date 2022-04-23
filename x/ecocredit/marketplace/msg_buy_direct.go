package marketplace

import (
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var _ legacytx.LegacyMsg = &MsgBuyDirect{}

func (m MsgBuyDirect) ValidateBasic() error {
	if _, err := types.AccAddressFromBech32(m.Buyer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}
	for _, order := range m.Orders {
		if order.SellOrderId == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("0 is not a valid sell order id")
		}
		if _, err := math.NewDecFromString(order.Quantity); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
		}
		if !order.DisableAutoRetire {
			if err := core.ValidateJurisdiction(order.RetirementJurisdiction); err != nil {
				return sdkerrors.Wrapf(err, "when DisableAutoRetire is false, a valid retirement jurisdiction must be provided")
			}
		}
		if order.BidPrice == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("must specify price per credit")
		}
		if err := order.BidPrice.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (m MsgBuyDirect) GetSigners() []types.AccAddress {
	addr, _ := types.AccAddressFromBech32(m.Buyer)
	return []types.AccAddress{addr}
}

func (m MsgBuyDirect) GetSignBytes() []byte {
	return types.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgBuyDirect) Route() string { return types.MsgTypeURL(&m) }

func (m MsgBuyDirect) Type() string { return types.MsgTypeURL(&m) }
