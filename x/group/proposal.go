package group

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/orm"
)

func (p *Proposal) GetMsgs() []sdk.Msg {
	msgs := make([]sdk.Msg, len(p.Msgs))
	for i, any := range p.Msgs {
		msg, ok := any.GetCachedValue().(sdk.Msg)
		if !ok {
			return nil
		}
		msgs[i] = msg
	}
	return msgs
}

func (p *Proposal) SetMsgs(new []sdk.Msg) error {
	p.Msgs = make([]*codectypes.Any, len(new))
	for i := range new {
		if new[i] == nil {
			return sdkerrors.Wrap(ErrInvalid, "msg must not be nil")
		}
		any, err := codectypes.NewAnyWithValue(new[i])
		if err != nil {
			return err
		}
		p.Msgs[i] = any
	}
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
func (p Proposal) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, any := range p.Msgs {
		var msg sdk.Msg
		err := unpacker.UnpackAny(any, &msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p Proposal) PrimaryKey() []byte {
	return orm.EncodeSequence(p.ProposalId)
}
