package proposal

import (
"encoding/json"
sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateProposal struct {
	Proposer sdk.AccAddress `json:"proposer"`
	Action ProposalAction `json:"action"`
}

type MsgVote struct {
	ProposalId []byte `json:"proposal_id"`
	Voter sdk.AccAddress `json:"voter"`
	Vote bool `json:"vote"`
}

type MsgTryExecuteProposal struct {
	ProposalId []byte `json:"proposal_id"`
	Signer sdk.AccAddress `json:"signer"`
}

type MsgWithdrawProposal struct {
	ProposalId []byte `json:"proposal_id"`
	Proposer sdk.AccAddress `json:"proposer"`
}

func (msg MsgCreateProposal) Route() string { return "proposal" }

func (msg MsgCreateProposal) Type() string { return "proposal.create" }

func (msg MsgCreateProposal) ValidateBasic() sdk.Error {
	return msg.Action.ValidateBasic()
}

func (msg MsgCreateProposal) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgCreateProposal) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Proposer}
}

func (msg MsgVote) Route() string { return "proposal" }

func (msg MsgVote) Type() string { return "proposal.vote" }

func (msg MsgVote) ValidateBasic() sdk.Error { return nil }

func (msg MsgVote) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgVote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Voter}
}

func (msg MsgTryExecuteProposal) Route() string { return "proposal" }

func (msg MsgTryExecuteProposal) Type() string { return "proposal.exec" }

func (msg MsgTryExecuteProposal) ValidateBasic() sdk.Error {
	return nil
}

func (msg MsgTryExecuteProposal) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgTryExecuteProposal) GetSigners() []sdk.AccAddress {
    return []sdk.AccAddress{msg.Signer}
}

func (msg MsgWithdrawProposal) Route() string { return "proposal" }

func (msg MsgWithdrawProposal) Type() string { return "proposal.withdraw" }

func (msg MsgWithdrawProposal) ValidateBasic() sdk.Error {
    return nil
}

func (msg MsgWithdrawProposal) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgWithdrawProposal) GetSigners() []sdk.AccAddress {
    return []sdk.AccAddress{msg.Proposer}
}
