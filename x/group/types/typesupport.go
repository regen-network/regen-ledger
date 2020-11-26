package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type Members []Member

func (ms Members) ValidateBasic() error {
	index := make(map[string]struct{}, len(ms))
	for i := range ms {
		member := ms[i]
		if err := member.ValidateBasic(); err != nil {
			return err
		}
		addr := string(member.Address)
		if _, exists := index[addr]; exists {
			return sdkerrors.Wrapf(ErrDuplicate, "address: %s", member.Address)
		}
		index[addr] = struct{}{}
	}
	return nil
}

type AccAddresses []sdk.AccAddress

func (a AccAddresses) ValidateBasic() error {
	index := make(map[string]struct{}, len(a))
	for i := range a {
		accAddr := a[i]
		if accAddr.Empty() {
			return sdkerrors.Wrap(ErrEmpty, "address")
		}
		if err := sdk.VerifyAddressFormat(accAddr); err != nil {
			return sdkerrors.Wrap(err, "address")
		}
		addr := string(accAddr)
		if _, exists := index[addr]; exists {
			return sdkerrors.Wrapf(ErrDuplicate, "address: %s", accAddr.String())
		}
		index[addr] = struct{}{}
	}
	return nil
}
