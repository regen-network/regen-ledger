package group

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
)

func NewMsgCreateGroup(group Group, signer sdk.AccAddress) MsgCreateGroup {
	return MsgCreateGroup{
		Data:   group,
		Signer: signer,
	}
}

func (msg MsgCreateGroup) Route() string { return "group" }

func (msg MsgCreateGroup) Type() string { return "group.create" }

func (info Group) ValidateBasic() sdk.Error {
	if len(info.Members) <= 0 {
		return sdk.ErrUnknownRequest("Group must reference a non-empty set of members")
	}
	if info.DecisionThreshold.Cmp(big.NewInt(0)) <= 0 {
		return sdk.ErrUnknownRequest("DecisionThreshold must be a positive integer")
	}
	return nil
}

func (msg MsgCreateGroup) ValidateBasic() sdk.Error {
	// TODO what are valid group ID's
	return msg.Data.ValidateBasic()
}

func (msg MsgCreateGroup) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgCreateGroup) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
