package proposal

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateProposal{}, "proposal/MsgCreateProposal", nil)
	cdc.RegisterConcrete(MsgVote{}, "proposal/MsgVote", nil)
	cdc.RegisterConcrete(MsgTryExecuteProposal{}, "proposal/MsgTryExecuteProposal", nil)
	cdc.RegisterConcrete(MsgWithdrawProposal{}, "proposal/MsgWithdrawProposal", nil)
	cdc.RegisterConcrete(Proposal{}, "proposal/Proposal", nil)
	cdc.RegisterInterface((*ProposalAction)(nil), nil)
}

