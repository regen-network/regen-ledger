package ecocredit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/types/math"
)

var (
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ sdk.Msg = &MsgCreateClass{}, &MsgCreateBatch{}, &MsgSend{},
		&MsgRetire{}, &MsgCancel{}, &MsgUpdateClassAdmin{}, &MsgUpdateClassIssuers{}, &MsgUpdateClassMetadata{},
		&MsgSell{}, &MsgUpdateSellOrders{}, &MsgBuy{}, &MsgAllowAskDenom{}, &MsgCreateProject{}, &MsgCreateBasket{}, &MsgPickFromBasket{},
		&MsgAddToBasket{}, &MsgTakeFromBasket{}

	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ legacytx.LegacyMsg = &MsgCreateClass{}, &MsgCreateBatch{}, &MsgSend{},
		&MsgRetire{}, &MsgCancel{}, &MsgUpdateClassAdmin{}, &MsgUpdateClassIssuers{}, &MsgUpdateClassMetadata{},
		&MsgSell{}, &MsgUpdateSellOrders{}, &MsgBuy{}, &MsgAllowAskDenom{}, &MsgCreateProject{},
		&MsgCreateBasket{}, &MsgPickFromBasket{}, &MsgAddToBasket{}, &MsgTakeFromBasket{}
)

// MaxMetadataLength defines the max length of the metadata bytes field
// for the credit-class & credit-batch.
// TODO: This could be used as params once x/params is upgraded to use protobuf
const MaxMetadataLength = 256

// Route implements the LegacyMsg interface.
func (m MsgCreateClass) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCreateClass) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateClass) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateClass) ValidateBasic() error {

	if len(m.Metadata) > MaxMetadataLength {
		return ErrMaxLimit.Wrap("credit class metadata")
	}

	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if len(m.Issuers) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("issuers cannot be empty")
	}

	if len(m.CreditTypeName) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credit class must have a credit type")
	}
	for _, issuer := range m.Issuers {

		if _, err := sdk.AccAddressFromBech32(issuer); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateClass.
func (m *MsgCreateClass) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgCreateBatch) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCreateBatch) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateBatch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateBatch) ValidateBasic() error {

	if len(m.Metadata) > MaxMetadataLength {
		return ErrMaxLimit.Wrap("credit batch metadata")
	}

	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
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

	if err := ValidateProjectID(m.ProjectId); err != nil {
		return err
	}

	for _, iss := range m.Issuance {

		if _, err := sdk.AccAddressFromBech32(iss.Recipient); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
		}

		if iss.TradableAmount != "" {
			if _, err := math.NewNonNegativeDecFromString(iss.TradableAmount); err != nil {
				return err
			}
		}

		if iss.RetiredAmount != "" {
			retiredAmount, err := math.NewNonNegativeDecFromString(iss.RetiredAmount)
			if err != nil {
				return err
			}

			if !retiredAmount.IsZero() {
				if err = ValidateJurisdiction(iss.RetirementLocation); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateBatch.
func (m *MsgCreateBatch) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgSend) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgSend) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgSend) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSend) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return sdkerrors.Wrap(err, "sender")
	}

	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap(err.Error())
	}

	if len(m.Credits) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credits should not be empty")
	}

	for _, credit := range m.Credits {
		if err := ValidateDenom(credit.BatchDenom); err != nil {
			return err
		}

		if _, err := math.NewNonNegativeDecFromString(credit.TradableAmount); err != nil {
			return err
		}

		retiredAmount, err := math.NewNonNegativeDecFromString(credit.RetiredAmount)
		if err != nil {
			return err
		}

		if !retiredAmount.IsZero() {
			if err = ValidateJurisdiction(credit.RetirementLocation); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetSigners returns the expected signers for MsgSend.
func (m *MsgSend) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgRetire) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgRetire) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgRetire) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgRetire) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Holder); err != nil {
		return sdkerrors.Wrap(err, "holder")
	}

	if len(m.Credits) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credits should not be empty")
	}

	for _, credit := range m.Credits {
		if err := ValidateDenom(credit.BatchDenom); err != nil {
			return err
		}

		if _, err := math.NewPositiveDecFromString(credit.Amount); err != nil {
			return err
		}
	}

	if err := ValidateJurisdiction(m.Location); err != nil {
		return err
	}

	return nil
}

// GetSigners returns the expected signers for MsgRetire.
func (m *MsgRetire) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Holder)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgCancel) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCancel) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCancel) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCancel) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(m.Holder); err != nil {
		return sdkerrors.Wrap(err, "holder")
	}

	if len(m.Credits) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credits should not be empty")
	}

	for _, credit := range m.Credits {
		if err := ValidateDenom(credit.BatchDenom); err != nil {
			return err
		}

		if _, err := math.NewPositiveDecFromString(credit.Amount); err != nil {
			return err
		}
	}
	return nil
}

// GetSigners returns the expected signers for MsgCancel.
func (m *MsgCancel) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Holder)
	return []sdk.AccAddress{addr}
}

func (m MsgUpdateClassAdmin) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassAdmin) Type() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m *MsgUpdateClassAdmin) ValidateBasic() error {
	if m.Admin == m.NewAdmin {
		return sdkerrors.ErrInvalidAddress.Wrap("new admin should be a different address from the signer")
	}

	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	if _, err := sdk.AccAddressFromBech32(m.NewAdmin); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	if err := ValidateClassID(m.ClassId); err != nil {
		return err
	}

	return nil
}

func (m *MsgUpdateClassAdmin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}

func (m MsgUpdateClassIssuers) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassIssuers) Type() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassIssuers) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m *MsgUpdateClassIssuers) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	if err := ValidateClassID(m.ClassId); err != nil {
		return err
	}

	if len(m.Issuers) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("issuers cannot be empty")
	}

	for _, addr := range m.Issuers {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return sdkerrors.ErrInvalidAddress
		}
	}

	return nil
}

func (m *MsgUpdateClassIssuers) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}

func (m MsgUpdateClassMetadata) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassMetadata) Type() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassMetadata) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m *MsgUpdateClassMetadata) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	if err := ValidateClassID(m.ClassId); err != nil {
		return err
	}

	if len(m.Metadata) > MaxMetadataLength {
		return ErrMaxLimit.Wrap("credit class metadata")
	}

	return nil
}

func (m *MsgUpdateClassMetadata) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgSell) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgSell) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgSell) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSell) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	for _, order := range m.Orders {
		if err := ValidateDenom(order.BatchDenom); err != nil {
			return err
		}

		if _, err := math.NewPositiveDecFromString(order.Quantity); err != nil {
			return sdkerrors.Wrapf(err, "quantity must be positive decimal: %s", order.Quantity)
		}

		if order.AskPrice == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("ask price cannot be empty")
		}

		if err := order.AskPrice.Validate(); err != nil {
			return err
		}

		if !order.AskPrice.Amount.IsPositive() {
			return sdkerrors.ErrInvalidRequest.Wrap("ask price must be positive amount")
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgSell.
func (m *MsgSell) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgUpdateSellOrders) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgUpdateSellOrders) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgUpdateSellOrders) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgUpdateSellOrders) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	for _, update := range m.Updates {

		if _, err := math.NewPositiveDecFromString(update.NewQuantity); err != nil {
			return sdkerrors.Wrapf(err, "quantity must be positive decimal: %s", update.NewQuantity)
		}

		if update.NewAskPrice == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("new ask price cannot be empty")
		}

		if err := update.NewAskPrice.Validate(); err != nil {
			return err
		}

		if !update.NewAskPrice.Amount.IsPositive() {
			return sdkerrors.ErrInvalidRequest.Wrap("ask price must be positive amount")
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgUpdateSellOrders.
func (m *MsgUpdateSellOrders) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgBuy) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgBuy) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgBuy) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgBuy) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Buyer); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	for _, order := range m.Orders {

		if _, err := math.NewPositiveDecFromString(order.Quantity); err != nil {
			return sdkerrors.Wrapf(err, "quantity must be positive decimal: %s", order.Quantity)
		}

		if order.BidPrice == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("bid price cannot be empty")
		}

		if err := order.BidPrice.Validate(); err != nil {
			return err
		}

		if !order.BidPrice.Amount.IsPositive() {
			return sdkerrors.ErrInvalidRequest.Wrap("bid price must be positive amount")
		}

		if order.RetirementLocation != "" {
			if err := ValidateJurisdiction(order.RetirementLocation); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgBuy.
func (m *MsgBuy) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Buyer)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgAllowAskDenom) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgAllowAskDenom) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgAllowAskDenom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgAllowAskDenom) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(m.RootAddress); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return err
	}

	return nil
}

// GetSigners returns the expected signers for MsgAllowAskDenom.
func (m *MsgAllowAskDenom) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.RootAddress)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgCreateProject) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCreateProject) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateProject) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateProject) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(m.Issuer); err != nil {
		return sdkerrors.ErrInvalidAddress
	}

	if err := ValidateClassID(m.ClassId); err != nil {
		return err
	}

	if len(m.Metadata) > MaxMetadataLength {
		return ErrMaxLimit.Wrap("create project metadata")
	}

	if err := ValidateJurisdiction(m.ProjectLocation); err != nil {
		return err
	}

	if m.ProjectId != "" {
		if err := ValidateProjectID(m.ProjectId); err != nil {
			return err
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreateProject.
func (m *MsgCreateProject) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgCreateBasket) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgCreateBasket) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgCreateBasket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgCreateBasket) ValidateBasic() error {
	// TODO: add MsgCreateBasket validation
	return nil
}

// GetSigners returns the expected signers for MsgCreateBasket.
func (m *MsgCreateBasket) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Curator)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgAddToBasket) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgAddToBasket) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgAddToBasket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgAddToBasket) ValidateBasic() error {
	// TODO: add MsgAddToBasket validation
	return nil
}

// GetSigners returns the expected signers for MsgAddToBasket.
func (m *MsgAddToBasket) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgPickFromBasket) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgPickFromBasket) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgPickFromBasket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgPickFromBasket) ValidateBasic() error {
	// TODO: add MsgPickFromBasket validation
	return nil
}

// GetSigners returns the expected signers for MsgPickFromBasket.
func (m *MsgPickFromBasket) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgTakeFromBasket) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgTakeFromBasket) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgTakeFromBasket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgTakeFromBasket) ValidateBasic() error {
	// TODO: add MsgTakeFromBasket validation
	return nil
}

// GetSigners returns the expected signers for MsgTakeFromBasket.
func (m *MsgTakeFromBasket) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}
