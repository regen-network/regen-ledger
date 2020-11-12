package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

var _ sdk.MsgRequest = &MsgCreateGroupRequest{}

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
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

var _ sdk.MsgRequest = &MsgUpdateGroupAdminRequest{}

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgUpdateGroupAdminRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupAdminRequest) ValidateBasic() error {
	if m.GroupId == 0 {
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

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgUpdateGroupCommentRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupCommentRequest) ValidateBasic() error {
	if m.GroupId == 0 {
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

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgUpdateGroupMembersRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupMembersRequest) ValidateBasic() error {
	if m.GroupId == 0 {
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

func (m *MsgProposeBaseRequest) ValidateBasic() error {
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

var _ sdk.MsgRequest = &MsgCreateGroupAccountRequest{}

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
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
	if m.GroupId == 0 {
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

var _ types.UnpackInterfacesMessage = MsgCreateGroupAccountRequest{}

// NewMsgCreateGroupAccount creates a new MsgCreateGroupAccountRequest.
func NewMsgCreateGroupAccount(admin sdk.AccAddress, group ID, comment string, decisionPolicy DecisionPolicy) (*MsgCreateGroupAccountRequest, error) {
	m := &MsgCreateGroupAccountRequest{
		Admin:   admin,
		GroupId: group,
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

func (m *MsgCreateGroupAccountRequest) GetGroupId() ID {
	return m.GroupId
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

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m MsgUpdateGroupAccountDecisionPolicyRequest) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var decisionPolicy DecisionPolicy
	return unpacker.UnpackAny(m.DecisionPolicy, &decisionPolicy)
}

var _ sdk.MsgRequest = &MsgVoteRequest{}

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgVoteRequest) GetSigners() []sdk.AccAddress {
	return m.Voters
}

// ValidateBasic does a sanity check on the provided data
func (m MsgVoteRequest) ValidateBasic() error {
	if len(m.Voters) == 0 {
		return errors.Wrap(ErrEmpty, "voters")
	}
	if err := AccAddresses(m.Voters).ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "voters")
	}
	if m.ProposalId == 0 {
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

var _ sdk.MsgRequest = &MsgExecRequest{}

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgExecRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Signer}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgExecRequest) ValidateBasic() error {
	if m.Signer.Empty() {
		return errors.Wrap(ErrEmpty, "signer")
	}
	if err := sdk.VerifyAddressFormat(m.Signer); err != nil {
		return errors.Wrap(ErrInvalid, "signer")
	}
	if m.ProposalId == 0 {
		return errors.Wrap(ErrEmpty, "proposal")
	}
	return nil
}
