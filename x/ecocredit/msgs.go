package ecocredit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/x/ecocredit/math"
)

var (
	_, _, _, _ sdk.MsgRequest = &MsgCreateClassRequest{}, &MsgCreateBatchRequest{}, &MsgSendRequest{}, &MsgRetireRequest{}
)

func (m *MsgCreateClassRequest) ValidateBasic() error {
	if len(m.Issuers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "issuers cannot be empty")
	}

	return nil
}

func (m *MsgCreateClassRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Designer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgCreateBatchRequest) ValidateBasic() error {
	for _, iss := range m.Issuance {
		_, err := math.ParseNonNegativeDecimal(iss.TradableUnits)
		if err != nil {
			return err
		}

		_, err = math.ParseNonNegativeDecimal(iss.RetiredUnits)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MsgCreateBatchRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Issuer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgSendRequest) ValidateBasic() error {
	for _, iss := range m.Credits {
		_, err := math.ParseNonNegativeDecimal(iss.TradableUnits)
		if err != nil {
			return err
		}

		_, err = math.ParseNonNegativeDecimal(iss.RetiredUnits)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MsgSendRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgRetireRequest) ValidateBasic() error {
	for _, iss := range m.Credits {
		_, err := math.ParsePositiveDecimal(iss.Units)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MsgRetireRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Holder)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}
