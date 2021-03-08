package group

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (p *Proposal) GetMsgs() []sdk.Msg {
	msgs := make([]sdk.Msg, len(p.Msgs))
	for i, any := range p.Msgs {
		var msg sdk.Msg
		if isServiceMsg(any.TypeUrl) {
			req := any.GetCachedValue()
			if req == nil {
				panic("Any cached value is nil. Transaction messages must be correctly packed Any values.")
			}
			msg = sdk.ServiceMsg{
				MethodName: any.TypeUrl,
				Request:    any.GetCachedValue().(sdk.MsgRequest),
			}
		} else {
			msg = any.GetCachedValue().(sdk.Msg)
		}
		msgs[i] = msg
	}
	return msgs
}

func (p *Proposal) SetMsgs(msgs []sdk.Msg) error {
	anys := make([]*types.Any, len(msgs))
	for i, msg := range msgs {
		var err error
		switch msg := msg.(type) {
		case sdk.ServiceMsg:
			anys[i], err = types.NewAnyWithCustomTypeURL(msg.Request, msg.MethodName)
		default:
			anys[i], err = types.NewAnyWithValue(msg)
		}
		if err != nil {
			return err
		}
	}
	p.Msgs = anys
	return nil
}

func (p Proposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(p.GroupAccount)
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

func isServiceMsg(typeURL string) bool {
	return strings.Count(typeURL, "/") >= 2
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, any := range p.Msgs {
		// If the any's typeUrl contains 2 slashes, then we unpack the any into
		// a ServiceMsg struct as per ADR-031.
		if isServiceMsg(any.TypeUrl) {
			var req sdk.MsgRequest
			err := unpacker.UnpackAny(any, &req)
			if err != nil {
				return err
			}
		} else {
			var msg sdk.Msg
			err := unpacker.UnpackAny(any, &msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
