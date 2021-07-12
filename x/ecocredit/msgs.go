package ecocredit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/math"
)

var (
	_, _, _, _, _, _ sdk.Msg = &MsgCreateClassRequest{}, &MsgCreateBatchRequest{}, &MsgSendRequest{},
		&MsgRetireRequest{}, &MsgCancelRequest{}, &MsgSetPrecisionRequest{}
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
		_, err := math.ParseNonNegativeDecimal(iss.TradableAmount)
		if err != nil {
			return err
		}

		retiredAmount, err := math.ParseNonNegativeDecimal(iss.RetiredAmount)
		if err != nil {
			return err
		}

		if !retiredAmount.IsZero() {
			err = validateLocation(iss.RetirementLocation)
			if err != nil {
				return err
			}
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
		_, err := math.ParseNonNegativeDecimal(iss.TradableAmount)
		if err != nil {
			return err
		}

		retiredAmount, err := math.ParseNonNegativeDecimal(iss.RetiredAmount)
		if err != nil {
			return err
		}

		if !retiredAmount.IsZero() {
			err = validateLocation(iss.RetirementLocation)
			if err != nil {
				return err
			}
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
		_, err := math.ParsePositiveDecimal(iss.Amount)
		if err != nil {
			return err
		}
	}

	err := validateLocation(m.Location)
	if err != nil {
		return err
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

func (m *MsgCancelRequest) ValidateBasic() error {
	for _, iss := range m.Credits {
		_, err := math.ParsePositiveDecimal(iss.Amount)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MsgCancelRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Holder)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgSetPrecisionRequest) ValidateBasic() error {
	if len(m.BatchDenom) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing batch_denom")
	}
	return nil
}

func (m *MsgSetPrecisionRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Issuer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}
