package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/x/bank/math"
)

var _, _, _, _, _, _ sdk.MsgRequest = &MsgMintRequest{}, &MsgSendRequest{}, &MsgBurnRequest{}, &MsgSetPrecisionRequest{}, &MsgMoveRequest{}, &MsgCreateDenomRequest{}

func (m *MsgMintRequest) ValidateBasic() error {
	for _, issuance := range m.Issuance {
		_, err := sdk.AccAddressFromBech32(issuance.Recipient)
		if err != nil {
			panic(err)
		}

		return ValidateCoins(issuance.Coins)
	}
	return nil
}

func (m *MsgMintRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.MinterAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgSendRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.ToAddress)
	if err != nil {
		return err
	}

	return ValidateCoins(m.Amount)
}

func (m *MsgSendRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.FromAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgBurnRequest) ValidateBasic() error {
	return ValidateCoins(m.Coins)
}

func (m *MsgBurnRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.BurnerAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func ValidateCoins(coins []*Coin) error {
	if len(coins) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "coins cannot be empty")
	}

	for _, coin := range coins {
		if coin == nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "coin cannot be empty")
		}

		err := sdk.ValidateDenom(coin.Denom)
		if err != nil {
			return err
		}

		_, err = math.ParsePositiveDecimal(coin.Amount)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MsgSetPrecisionRequest) ValidateBasic() error {
	return sdk.ValidateDenom(m.Denom)
}

func (m *MsgSetPrecisionRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.DenomAdmin)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgMoveRequest) ValidateBasic() error {
	return nil
}

func (m *MsgMoveRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.MoverAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgCreateDenomRequest) ValidateBasic() error {
	return nil
}

func (m *MsgCreateDenomRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.NamespaceAdmin)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}
