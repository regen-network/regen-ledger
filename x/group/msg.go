package group

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

const (
	msgTypeCreateGroup        = "create_group"
	msgTypeUpdateGroupAdmin   = "update_group_admin"
	msgTypeUpdateGroupComment = "update_group_comment"
	msgTypeUpdateGroupMembers = "update_group_members"
	msgTypeCreateGroupAccount = "create_group_account"
	msgTypeVote               = "vote"
	msgTypeExecProposal       = "exec_proposal"
)

var _ sdk.Msg = &MsgCreateGroup{}

func (m MsgCreateGroup) Route() string { return ModuleName }
func (m MsgCreateGroup) Type() string  { return msgTypeCreateGroup }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgCreateGroup) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgCreateGroup) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(errors.Wrap(err, "get sign bytes"))
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgCreateGroup) ValidateBasic() error {
	if m.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}
	if err := Members(m.Members).ValidateBasic(); err != nil {
		return errors.Wrap(err, "members")
	}
	for i := range m.Members {
		member := m.Members[i]
		if member.Power.Equal(sdk.ZeroDec()) {
			return sdkerrors.Wrap(ErrEmpty, "member power")
		}
	}
	return nil
}

func (m Member) ValidateBasic() error {
	if m.Address.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "address")
	}
	if m.Power.IsNil() || m.Power.LT(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalid, "power")
	}
	if err := sdk.VerifyAddressFormat(m.Address); err != nil {
		return sdkerrors.Wrap(err, "address")
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateGroupAdmin{}

func (m MsgUpdateGroupAdmin) Route() string { return ModuleName }
func (m MsgUpdateGroupAdmin) Type() string  { return msgTypeUpdateGroupAdmin }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgUpdateGroupAdmin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgUpdateGroupAdmin) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(errors.Wrap(err, "get sign bytes"))
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupAdmin) ValidateBasic() error {
	if m.Group == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group")

	}

	if m.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if m.NewAdmin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "new admin")
	}
	if err := sdk.VerifyAddressFormat(m.NewAdmin); err != nil {
		return sdkerrors.Wrap(err, "new admin")
	}

	if m.Admin.Equals(m.NewAdmin) {
		return sdkerrors.Wrap(ErrInvalid, "new and old admin are the same")
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateGroupComment{}

func (m MsgUpdateGroupComment) Route() string { return ModuleName }
func (m MsgUpdateGroupComment) Type() string  { return msgTypeUpdateGroupComment }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgUpdateGroupComment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgUpdateGroupComment) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(errors.Wrap(err, "get sign bytes"))
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupComment) ValidateBasic() error {
	if m.Group == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group")

	}
	if m.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateGroupMembers{}

func (m MsgUpdateGroupMembers) Route() string { return ModuleName }
func (m MsgUpdateGroupMembers) Type() string  { return msgTypeUpdateGroupMembers }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgUpdateGroupMembers) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgUpdateGroupMembers) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(errors.Wrap(err, "get sign bytes"))
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupMembers) ValidateBasic() error {
	if m.Group == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group")

	}
	if m.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if len(m.MemberUpdates) == 0 {
		return sdkerrors.Wrap(ErrEmpty, "member updates")
	}
	if err := Members(m.MemberUpdates).ValidateBasic(); err != nil {
		return errors.Wrap(err, "members")
	}
	return nil
}

func (m *MsgProposeBase) ValidateBasic() error {
	if m.GroupAccount.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "group account")
	}
	if err := sdk.VerifyAddressFormat(m.GroupAccount); err != nil {
		return sdkerrors.Wrap(err, "group account")
	}
	if len(m.Proposers) == 0 {
		return sdkerrors.Wrap(ErrEmpty, "proposers")
	}
	if err := AccAddresses(m.Proposers).ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "proposers")
	}
	return nil
}

var _ sdk.Msg = &MsgCreateGroupAccount{}

func (m MsgCreateGroupAccount) Route() string { return ModuleName }
func (m MsgCreateGroupAccount) Type() string  { return msgTypeCreateGroupAccount }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgCreateGroupAccount) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgCreateGroupAccount) GetSignBytes() []byte {
	bz := moduleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic does a sanity check on the provided data
func (m MsgCreateGroupAccount) ValidateBasic() error {
	if m.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}
	if m.Group == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group")
	}

	policy := m.GetDecisionPolicy()
	if policy == nil {
		return errors.Wrap(ErrEmpty, "decision policy")
	}

	if err := policy.ValidateBasic(); err != nil {
		return errors.Wrap(err, "decision policy")
	}
	return nil
}

var _ types.UnpackInterfacesMessage = MsgCreateGroupAccount{}

// NewMsgCreateGroupAccount creates a new MsgCreateGroupAccount.
func NewMsgCreateGroupAccount(admin sdk.AccAddress, group GroupID, comment string, decisionPolicy DecisionPolicy) (*MsgCreateGroupAccount, error) {
	m := &MsgCreateGroupAccount{
		Admin:   admin,
		Group:   group,
		Comment: comment,
	}
	err := m.SetDecisionPolicy(decisionPolicy)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MsgCreateGroupAccount) GetAdmin() sdk.AccAddress {
	return m.Admin
}

func (m *MsgCreateGroupAccount) GetGroup() GroupID {
	return m.Group
}

func (m *MsgCreateGroupAccount) GetComment() string {
	return m.Comment
}

func (m *MsgCreateGroupAccount) GetDecisionPolicy() DecisionPolicy {
	decisionPolicy, ok := m.DecisionPolicy.GetCachedValue().(DecisionPolicy)
	if !ok {
		return nil
	}
	return decisionPolicy
}

func (m *MsgCreateGroupAccount) SetDecisionPolicy(decisionPolicy DecisionPolicy) error {
	msg, ok := decisionPolicy.(proto.Message)
	if !ok {
		return fmt.Errorf("can't proto marshal %T", msg)
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return err
	}
	m.DecisionPolicy = any
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m MsgCreateGroupAccount) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var decisionPolicy DecisionPolicy
	return unpacker.UnpackAny(m.DecisionPolicy, &decisionPolicy)
}

var _ sdk.Msg = &MsgVote{}

func (m MsgVote) Route() string { return ModuleName }
func (m MsgVote) Type() string  { return msgTypeVote }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgVote) GetSigners() []sdk.AccAddress {
	return m.Voters
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgVote) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(errors.Wrap(err, "get sign bytes"))
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgVote) ValidateBasic() error {
	if len(m.Voters) == 0 {
		return errors.Wrap(ErrEmpty, "voters")
	}
	if err := AccAddresses(m.Voters).ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "voters")
	}
	if m.Proposal == 0 {
		return errors.Wrap(ErrEmpty, "proposal")
	}
	if m.Choice == Choice_UNKNOWN {
		return errors.Wrap(ErrEmpty, "choice")
	}
	if _, ok := Choice_name[int32(m.Choice)]; !ok {
		return errors.Wrap(ErrInvalid, "choice")
	}
	return nil
}

var _ sdk.Msg = &MsgExec{}

func (m MsgExec) Route() string { return ModuleName }
func (m MsgExec) Type() string  { return msgTypeExecProposal }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgExec) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgExec) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(errors.Wrap(err, "get sign bytes"))
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgExec) ValidateBasic() error {
	if m.Signer.Empty() {
		return errors.Wrap(ErrEmpty, "signer")
	}
	if err := sdk.VerifyAddressFormat(m.Signer); err != nil {
		return errors.Wrap(ErrInvalid, "signer")
	}
	if m.Proposal == 0 {
		return errors.Wrap(ErrEmpty, "proposal")
	}
	return nil
}
