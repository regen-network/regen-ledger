package group

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/module/server"
)

func (p *Proposal) GetMsgs() []sdk.Msg {
	msgs, err := server.GetMsgs(p.Msgs)
	if err != nil {
		panic(err)
	}
	return msgs
}

func (p *Proposal) SetMsgs(msgs []sdk.Msg) error {
	anys, err := server.SetMsgs(msgs)
	if err != nil {
		return err
	}
	p.Msgs = anys
	return nil
}

func (p Proposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(p.Address)
	if err != nil {
		return sdkerrors.Wrap(err, "group account")
	}

	if len(p.Proposers) == 0 {
		return sdkerrors.Wrap(ErrEmpty, "proposers")
	}
	addrs := make([]sdk.AccAddress, len(p.Proposers))
	for i, proposer := range p.Proposers {
		addr, err := sdk.AccAddressFromBech32(proposer)
		if err != nil {
			return sdkerrors.Wrap(err, "proposers")
		}
		addrs[i] = addr
	}
	if err := AccAddresses(addrs).ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "proposers")
	}

	if p.SubmittedAt.Seconds == 0 && p.SubmittedAt.Nanos == 0 {
		return sdkerrors.Wrap(ErrEmpty, "submitted at")
	}
	if p.GroupVersion == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group version")
	}
	if p.GroupAccountVersion == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group account version")
	}
	if p.Status == ProposalStatusInvalid {
		return sdkerrors.Wrap(ErrEmpty, "status")
	}
	if _, ok := Proposal_Status_name[int32(p.Status)]; !ok {
		return sdkerrors.Wrap(ErrInvalid, "status")
	}
	if p.Result == ProposalResultInvalid {
		return sdkerrors.Wrap(ErrEmpty, "result")
	}
	if _, ok := Proposal_Result_name[int32(p.Result)]; !ok {
		return sdkerrors.Wrap(ErrInvalid, "result")
	}
	if p.ExecutorResult == ProposalExecutorResultInvalid {
		return sdkerrors.Wrap(ErrEmpty, "executor result")
	}
	if _, ok := Proposal_ExecutorResult_name[int32(p.ExecutorResult)]; !ok {
		return sdkerrors.Wrap(ErrInvalid, "executor result")
	}
	if err := p.VoteState.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "vote state")
	}
	if p.Timeout.Seconds == 0 && p.Timeout.Nanos == 0 {
		return sdkerrors.Wrap(ErrEmpty, "timeout")
	}
	msgs := p.GetMsgs()
	for i, msg := range msgs {
		if err := msg.ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "message %d", i)
		}
	}
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	return server.UnpackInterfaces(unpacker, p.Msgs)
}

func (p Proposal) PrimaryKeyFields() []interface{} {
	return []interface{}{p.ProposalId}
}
