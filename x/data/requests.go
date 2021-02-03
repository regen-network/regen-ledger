package data

import (
	"bytes"
	"crypto"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_, _, _ sdk.MsgRequest = &MsgAnchorDataRequest{}, &MsgSignDataRequest{}, &MsgStoreRawDataRequest{}
)

func (m *MsgAnchorDataRequest) ValidateBasic() error {
	return m.Hash.Validate()
}

func (m *MsgAnchorDataRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgSignDataRequest) ValidateBasic() error {
	return m.Hash.Validate()
}

func (m *MsgSignDataRequest) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, len(m.Signers))

	for i, signer := range m.Signers {
		addr, err := sdk.AccAddressFromBech32(signer)
		if err != nil {
			panic(err)
		}
		addrs[i] = addr
	}

	return addrs
}

func (m *MsgStoreRawDataRequest) ValidateBasic() error {
	err := m.Hash.Validate()
	if err != nil {
		return err
	}

	digestAlgorithm := m.Hash.DigestAlgorithm
	switch digestAlgorithm {
	case DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256:
		hash := crypto.BLAKE2b_256.New()
		_, err = hash.Write(m.Content)
		if err != nil {
			return sdkerrors.Wrap(ErrHashVerificationFailed, err.Error())
		}

		digest := hash.Sum(nil)
		if !bytes.Equal(m.Hash.Hash, digest) {
			return ErrHashVerificationFailed
		}

		return nil
	default:
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("unsupported %T %s", digestAlgorithm, digestAlgorithm))
	}
}

func (m *MsgStoreRawDataRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}
