package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
)

var _ sdk.MsgRequest = &MsgCreateGroupRequest{}

// GetSigners returns the expected signers for a MsgCreateGroupRequest.
func (m MsgCreateGroupRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgCreateGroupRequest) ValidateBasic() error {
	if m.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}
	if err := Members(m.Members).ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "members")
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

var _ sdk.MsgRequest = &MsgUpdateGroupAdminRequest{}

// GetSigners returns the expected signers for a MsgUpdateGroupAdminRequest.
func (m MsgUpdateGroupAdminRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupAdminRequest) ValidateBasic() error {
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

var _ sdk.MsgRequest = &MsgUpdateGroupCommentRequest{}

// GetSigners returns the expected signers for a MsgUpdateGroupCommentRequest.
func (m MsgUpdateGroupCommentRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupCommentRequest) ValidateBasic() error {
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

var _ sdk.MsgRequest = &MsgUpdateGroupMembersRequest{}

// GetSigners returns the expected signers for a MsgUpdateGroupMembersRequest.
func (m MsgUpdateGroupMembersRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupMembersRequest) ValidateBasic() error {
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
		return sdkerrors.Wrap(err, "members")
	}
	return nil
}

var _ sdk.MsgRequest = &MsgCreateGroupAccountRequest{}

// GetSigners returns the expected signers for a MsgCreateGroupAccountRequest.
func (m MsgCreateGroupAccountRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgCreateGroupAccountRequest) ValidateBasic() error {
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
		return sdkerrors.Wrap(ErrEmpty, "decision policy")
	}

	if err := policy.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "decision policy")
	}
	return nil
}

var _ sdk.MsgRequest = &MsgUpdateGroupAccountAdminRequest{}

// GetSigners returns the expected signers for a MsgUpdateGroupAccountAdminRequest.
func (m MsgUpdateGroupAccountAdminRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupAccountAdminRequest) ValidateBasic() error {
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

	if m.GroupAccount.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "group account")
	}
	if err := sdk.VerifyAddressFormat(m.GroupAccount); err != nil {
		return sdkerrors.Wrap(err, "group account")
	}

	if m.Admin.Equals(m.NewAdmin) {
		return sdkerrors.Wrap(ErrInvalid, "new and old admin are the same")
	}
	return nil
}

var _ sdk.MsgRequest = &MsgUpdateGroupAccountDecisionPolicyRequest{}
var _ types.UnpackInterfacesMessage = MsgUpdateGroupAccountDecisionPolicyRequest{}

// GetSigners returns the expected signers for a MsgUpdateGroupAccountDecisionPolicyRequest.
func (m MsgUpdateGroupAccountDecisionPolicyRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupAccountDecisionPolicyRequest) ValidateBasic() error {
	if m.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if m.GroupAccount.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "group account")
	}
	if err := sdk.VerifyAddressFormat(m.GroupAccount); err != nil {
		return sdkerrors.Wrap(err, "group account")
	}

	policy := m.GetDecisionPolicy()
	if policy == nil {
		return sdkerrors.Wrap(ErrEmpty, "decision policy")
	}

	if err := policy.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "decision policy")
	}

	return nil
}

func (m *MsgUpdateGroupAccountDecisionPolicyRequest) GetDecisionPolicy() DecisionPolicy {
	decisionPolicy, ok := m.DecisionPolicy.GetCachedValue().(DecisionPolicy)
	if !ok {
		return nil
	}
	return decisionPolicy
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m MsgUpdateGroupAccountDecisionPolicyRequest) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var decisionPolicy DecisionPolicy
	return unpacker.UnpackAny(m.DecisionPolicy, &decisionPolicy)
}

var _ sdk.MsgRequest = &MsgUpdateGroupAccountCommentRequest{}

// GetSigners returns the expected signers for a MsgUpdateGroupAccountCommentRequest.
func (m MsgUpdateGroupAccountCommentRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupAccountCommentRequest) ValidateBasic() error {
	if m.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if m.GroupAccount.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "group account")
	}
	if err := sdk.VerifyAddressFormat(m.GroupAccount); err != nil {
		return sdkerrors.Wrap(err, "group account")
	}

	return nil
}

var _ sdk.MsgRequest = &MsgCreateGroupAccountRequest{}
var _ types.UnpackInterfacesMessage = MsgCreateGroupAccountRequest{}

// NewMsgCreateGroupAccount creates a new MsgCreateGroupAccountRequest.
func NewMsgCreateGroupAccount(admin sdk.AccAddress, group GroupID, comment string, decisionPolicy DecisionPolicy) (*MsgCreateGroupAccountRequest, error) {
	m := &MsgCreateGroupAccountRequest{
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

func (m *MsgCreateGroupAccountRequest) GetAdmin() sdk.AccAddress {
	return m.Admin
}

func (m *MsgCreateGroupAccountRequest) GetGroup() GroupID {
	return m.Group
}

func (m *MsgCreateGroupAccountRequest) GetComment() string {
	return m.Comment
}

func (m *MsgCreateGroupAccountRequest) GetDecisionPolicy() DecisionPolicy {
	decisionPolicy, ok := m.DecisionPolicy.GetCachedValue().(DecisionPolicy)
	if !ok {
		return nil
	}
	return decisionPolicy
}

func (m *MsgCreateGroupAccountRequest) SetDecisionPolicy(decisionPolicy DecisionPolicy) error {
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
func (m MsgCreateGroupAccountRequest) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var decisionPolicy DecisionPolicy
	return unpacker.UnpackAny(m.DecisionPolicy, &decisionPolicy)
}

var _ sdk.MsgRequest = &MsgCreateProposalRequest{}

// GetSigners returns the expected signers for a MsgCreateProposalRequest.
func (m MsgCreateProposalRequest) GetSigners() []sdk.AccAddress {
	return m.Proposers
}

// ValidateBasic does a sanity check on the provided data
func (m MsgCreateProposalRequest) ValidateBasic() error {
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
	for i, any := range m.Msgs {
		msg, ok := any.GetCachedValue().(sdk.Msg)
		if !ok {
			return sdkerrors.Wrapf(sdkerrors.ErrUnpackAny, "cannot unpack Any into sdk.Msg %T", any)
		}
		if err := msg.ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "msg %d", i)
		}
	}
	return nil
}

// SetMsgs packs msgs into Any's
func (m *MsgCreateProposalRequest) SetMsgs(msgs []sdk.Msg) error {
	anys := make([]*types.Any, len(msgs))

	for i, msg := range msgs {
		var err error
		anys[i], err = types.NewAnyWithValue(msg)
		if err != nil {
			return err
		}
	}
	m.Msgs = anys
	return nil
}

// GetMsgs unpacks m.Msgs Any's into sdk.Msg's
func (m MsgCreateProposalRequest) GetMsgs() []sdk.Msg {
	msgs := make([]sdk.Msg, len(m.Msgs))
	for i, any := range m.Msgs {
		msg, ok := any.GetCachedValue().(sdk.Msg)
		if !ok {
			return nil
		}
		msgs[i] = msg
	}
	return msgs
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m MsgCreateProposalRequest) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, m := range m.Msgs {
		err := types.UnpackInterfaces(m, unpacker)
		if err != nil {
			return err
		}
	}

	return nil
}

var _ sdk.MsgRequest = &MsgVoteRequest{}

// GetSigners returns the expected signers for a MsgVoteRequest.
func (m MsgVoteRequest) GetSigners() []sdk.AccAddress {
	return m.Voters
}

// ValidateBasic does a sanity check on the provided data
func (m MsgVoteRequest) ValidateBasic() error {
	if len(m.Voters) == 0 {
		return sdkerrors.Wrap(ErrEmpty, "voters")
	}
	if err := AccAddresses(m.Voters).ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "voters")
	}
	if m.Proposal == 0 {
		return sdkerrors.Wrap(ErrEmpty, "proposal")
	}
	if m.Choice == Choice_CHOICE_UNSPECIFIED {
		return sdkerrors.Wrap(ErrEmpty, "choice")
	}
	if _, ok := Choice_name[int32(m.Choice)]; !ok {
		return sdkerrors.Wrap(ErrInvalid, "choice")
	}
	return nil
}

var _ sdk.MsgRequest = &MsgExecRequest{}

// GetSigners returns the expected signers for a MsgExecRequest.
func (m MsgExecRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgExecRequest) ValidateBasic() error {
	if m.Signer.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "signer")
	}
	if err := sdk.VerifyAddressFormat(m.Signer); err != nil {
		return sdkerrors.Wrap(ErrInvalid, "signer")
	}
	if m.Proposal == 0 {
		return sdkerrors.Wrap(ErrEmpty, "proposal")
	}
	return nil
}
