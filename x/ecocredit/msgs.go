package ecocredit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/math"
)

var (
	_, _, _, _, _ sdk.Msg = &MsgCreateClass{}, &MsgCreateBatch{}, &MsgSend{},
		&MsgRetire{}, &MsgCancel{}
)

func (m *MsgCreateClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Designer)
	if err != nil {
		return sdkerrors.Wrap(err, "designer")
	}

	if len(m.Issuers) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("issuers cannot be empty")
	}

	if len(m.CreditType) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credit class must have a credit type")
	}
	for _, issuer := range m.Issuers {
		_, err := sdk.AccAddressFromBech32(issuer)
		if err != nil {
			return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
		}
	}

	for _, issuer := range m.Issuers {
		_, err := sdk.AccAddressFromBech32(issuer)
		if err != nil {
			return sdkerrors.Wrapf(err, "issuer: %s", issuer)
		}
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
	_, err := sdk.AccAddressFromBech32(m.Issuer)
	if err != nil {
		return sdkerrors.Wrap(err, "issuer")
	}

	if m.StartDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide a start date for the credit batch")
	}
	if m.EndDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide an end date for the credit batch")
	}
	if m.EndDate.Before(*m.StartDate) {
		return sdkerrors.ErrInvalidRequest.Wrapf("the batch end date (%s) must be the same as or after the batch start date (%s)", m.EndDate.Format("2006-01-02"), m.StartDate.Format("2006-01-02"))
	}
	if m.ClassId == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("class ID should not be empty")
	}

	err = validateLocation(m.ProjectLocation)
	if err != nil {
		return err
	}

	for _, iss := range m.Issuance {
		_, err := sdk.AccAddressFromBech32(iss.Recipient)
		if err != nil {
			return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
		}

		if iss.TradableAmount != "" {
			_, err := math.ParseNonNegativeDecimal(iss.TradableAmount)
			if err != nil {
				return err
			}
		}

		if iss.RetiredAmount != "" {
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
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "sender")
	}

	_, err = sdk.AccAddressFromBech32(m.Recipient)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	if len(m.Credits) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credits should not be empty")
	}

	for _, credit := range m.Credits {
		if credit.BatchDenom == "" {
			return sdkerrors.ErrInvalidRequest.Wrap("batch denom should not be empty")
		}

		_, err := math.ParseNonNegativeDecimal(credit.TradableAmount)
		if err != nil {
			return err
		}

		retiredAmount, err := math.ParseNonNegativeDecimal(credit.RetiredAmount)
		if err != nil {
			return err
		}

		if !retiredAmount.IsZero() {
			err = validateLocation(credit.RetirementLocation)
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
	_, err := sdk.AccAddressFromBech32(m.Holder)
	if err != nil {
		return sdkerrors.Wrap(err, "holder")
	}

	if len(m.Credits) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credits should not be empty")
	}

	for _, credit := range m.Credits {
		if credit.BatchDenom == "" {
			return sdkerrors.ErrInvalidRequest.Wrap("batch denom should not be empty")
		}
		_, err := math.ParsePositiveDecimal(credit.Amount)
		if err != nil {
			return err
		}
	}

	err = validateLocation(m.Location)
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
	_, err := sdk.AccAddressFromBech32(m.Holder)
	if err != nil {
		return sdkerrors.Wrap(err, "holder")
	}

	if len(m.Credits) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credits should not be empty")
	}

	for _, credit := range m.Credits {
		if credit.BatchDenom == "" {
			return sdkerrors.ErrInvalidRequest.Wrap("batch denom should not be empty")
		}

		_, err := math.ParsePositiveDecimal(credit.Amount)
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
