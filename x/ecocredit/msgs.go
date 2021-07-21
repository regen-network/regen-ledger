package ecocredit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/math"
)

var (
	_, _, _, _, _, _ sdk.Msg = &MsgCreateClass{}, &MsgCreateBatch{}, &MsgSend{},
		&MsgRetire{}, &MsgCancel{}, &MsgSetPrecision{}
)

func (m *MsgCreateClass) ValidateBasic() error {
	if len(m.Issuers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "issuers cannot be empty")
	}

	return nil
}

func (m *MsgCreateClass) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Designer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgCreateBatch) ValidateBasic() error {
	if m.StartDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("Must provide a start date for the credit batch")
	}
	if m.EndDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("Must provide an end date for the credit batch")
	}
	if m.EndDate.Before(*m.StartDate) {
		return sdkerrors.ErrInvalidRequest.Wrapf("The batch end date (%s) must be the same as or after the batch start date (%s)", m.EndDate.Format("2006-01-02"), m.StartDate.Format("2006-01-02"))
	}

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

func (m *MsgCreateBatch) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Issuer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgSend) ValidateBasic() error {
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

func (m *MsgSend) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgRetire) ValidateBasic() error {
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

func (m *MsgRetire) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Holder)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgCancel) ValidateBasic() error {
	for _, iss := range m.Credits {
		_, err := math.ParsePositiveDecimal(iss.Amount)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MsgCancel) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Holder)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgSetPrecision) ValidateBasic() error {
	if len(m.BatchDenom) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing batch_denom")
	}
	return nil
}

func (m *MsgSetPrecision) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Issuer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}
