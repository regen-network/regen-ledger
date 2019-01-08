package agent

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateAgent struct {
	Data   AgentInfo      `json:"data"`
	Signer sdk.AccAddress `json:"signer"`
}

type MsgUpdateAgent struct {
	Id      AgentID `json:"id"`
	Data    AgentInfo `json:"data"`
	Signers []sdk.AccAddress `json:"signers"`
}

func NewMsgCreateAgent(info AgentInfo, signer sdk.AccAddress) MsgCreateAgent {
	return MsgCreateAgent{
        Data: info,
		Signer:signer,
	}
}

func (msg MsgCreateAgent) Route() string { return "agent" }

func (msg MsgCreateAgent) Type() string { return "create" }

func (info AgentInfo) ValidateBasic() sdk.Error {
	if len(info.Agents) <= 0 && len(info.Addresses) <= 0 {
		return sdk.ErrUnknownRequest("Agent info must reference a non-empty set of members")
	}
	if info.AuthPolicy != MultiSig {
		return sdk.ErrUnknownRequest("Only multi-sig auth policies are currently supported")
	}
	if info.MultisigThreshold <= 0 {
		return sdk.ErrUnknownRequest("MultisigThreshold must be a positive integer")
	}
	return nil
}

func (msg MsgCreateAgent) ValidateBasic() sdk.Error {
	// TODO what are valid agent ID's
	return msg.Data.ValidateBasic()
}

func (msg MsgCreateAgent) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgCreateAgent) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

func (msg MsgUpdateAgent) Route() string { return "agent" }

func (msg MsgUpdateAgent) Type() string { return "update" }

func (msg MsgUpdateAgent) ValidateBasic() sdk.Error {
	return msg.Data.ValidateBasic()
}

func (msg MsgUpdateAgent) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgUpdateAgent) GetSigners() []sdk.AccAddress {
    return msg.Signers
}

