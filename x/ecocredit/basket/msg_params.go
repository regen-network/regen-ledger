package basket

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

var _ legacytx.LegacyMsg = &MsgUpdateBasketFee{}

func (m MsgUpdateBasketFee) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateBasketFee) GetSigners() []types.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateBasketFee) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateBasketFee) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m MsgUpdateBasketFee) Type() string {
	//TODO implement me
	panic("implement me")
}
