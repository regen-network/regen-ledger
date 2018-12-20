package proposal

import (
"encoding/json"
sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateProposal struct {
	Proposer sdk.AccAddress
	Action ProposalAction
}

type MsgVote struct {
	ProposalId []byte
	Voter sdk.AccAddress
	Vote bool
}

type MsgTryExecuteProposal struct {
	ProposalId []byte
	Signer sdk.AccAddress
}

type MsgWithdrawProposal struct {
	ProposalId []byte
	Proposer sdk.AccAddress
}

func NewMsgCreateProposal(proposer sdk.AccAddress, action ProposalAction) MsgCreateProposal {
	return MsgCreateProposal{
		Proposer:proposer,
		Action:action,
	}
}

func (msg MsgCreateProposal) Route() string { return "proposal" }

func (msg MsgCreateProposal) Type() string { return "create" }

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

func NewMsgVote(proposalId []byte, voter sdk.AccAddress, vote bool) MsgVote {
	return MsgVote{ProposalId: proposalId, Voter: voter, Vote: vote}
}

func (msg MsgVote) Route() string { return "proposal" }

func (msg MsgVote) Type() string { return "vote" }

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

func (msg MsgTryExecuteProposal) Type() string { return "try-exec" }

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

func NewMsgWithdrawProposal(proposalId []byte, proposer sdk.AccAddress) MsgWithdrawProposal {
	return MsgWithdrawProposal{ProposalId: proposalId, Proposer: proposer}
}

func (msg MsgWithdrawProposal) Route() string {
	return "proposal"
}

func (msg MsgWithdrawProposal) Type() string {
	return "withdraw"
}

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
