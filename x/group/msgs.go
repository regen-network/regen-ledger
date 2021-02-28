package group

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	"github.com/regen-network/regen-ledger/math"
)

// Group message types and routes
const (
	TypeMsgCreateGroup                      = "create_group"
	TypeMsgUpdateGroupAdmin                 = "update_group_admin"
	TypeMsgUpdateGroupComment               = "update_group_comment"
	TypeMsgUpdateGroupMembers               = "update_group_members"
	TypeMsgCreateGroupAccount               = "create_group_account"
	TypeMsgUpdateGroupAccountAdmin          = "update_group_account_admin"
	TypeMsgUpdateGroupAccountDecisionPolicy = "update_group_account_decision_policy"
	TypeMsgUpdateGroupAccountComment        = "update_group_account_comment"
	TypeMsgCreateProposal                   = "create_proposal"
	TypeMsgVote                             = "vote"
	TypeMsgExec                             = "exec"
)

var _ sdk.Msg = &MsgCreateGroupRequest{}

// Route Implements Msg.
func (m MsgCreateGroupRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgCreateGroupRequest) Type() string { return TypeMsgCreateGroup }

// GetSignBytes Implements Msg.
func (m MsgCreateGroupRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgCreateGroupRequest.
func (m MsgCreateGroupRequest) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgCreateGroupRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	members := Members{Members: m.Members}
	if err := members.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "members")
	}
	for i := range m.Members {
		member := m.Members[i]
		if _, err := math.ParsePositiveDecimal(member.Weight); err != nil {
			return sdkerrors.Wrap(err, "member weight")
		}
	}
	return nil
}

func (m Member) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrap(err, "address")
	}
	if _, err := math.ParseNonNegativeDecimal(m.Weight); err != nil {
		return sdkerrors.Wrap(err, "weight")
	}

	return nil
}

var _ sdk.Msg = &MsgUpdateGroupAdminRequest{}

// Route Implements Msg.
func (m MsgUpdateGroupAdminRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgUpdateGroupAdminRequest) Type() string { return TypeMsgUpdateGroupAdmin }

// GetSignBytes Implements Msg.
func (m MsgUpdateGroupAdminRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgUpdateGroupAdminRequest.
func (m MsgUpdateGroupAdminRequest) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupAdminRequest) ValidateBasic() error {
	if m.GroupId == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group")
	}

	admin, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	newAdmin, err := sdk.AccAddressFromBech32(m.NewAdmin)
	if err != nil {
		return sdkerrors.Wrap(err, "new admin")
	}

	if admin.Equals(newAdmin) {
		return sdkerrors.Wrap(ErrInvalid, "new and old admin are the same")
	}
	return nil
}

func (m *MsgUpdateGroupAdminRequest) GetGroupID() uint64 {
	return m.GroupId
}

var _ sdk.Msg = &MsgUpdateGroupMetadataRequest{}

// Route Implements Msg.
func (m MsgUpdateGroupMetadataRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgUpdateGroupMetadataRequest) Type() string { return TypeMsgUpdateGroupComment }

// GetSignBytes Implements Msg.
func (m MsgUpdateGroupMetadataRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgUpdateGroupMetadataRequest.
func (m MsgUpdateGroupMetadataRequest) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupMetadataRequest) ValidateBasic() error {
	if m.GroupId == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group")

	}
	_, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		return sdkerrors.Wrap(err, "admin")
	}
	return nil
}

func (m *MsgUpdateGroupMetadataRequest) GetGroupID() uint64 {
	return m.GroupId
}

var _ sdk.Msg = &MsgUpdateGroupMembersRequest{}

// Route Implements Msg.
func (m MsgUpdateGroupMembersRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgUpdateGroupMembersRequest) Type() string { return TypeMsgUpdateGroupMembers }

// GetSignBytes Implements Msg.
func (m MsgUpdateGroupMembersRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

var _ sdk.MsgRequest = &MsgUpdateGroupMembersRequest{}

// GetSigners returns the expected signers for a MsgUpdateGroupMembersRequest.
func (m MsgUpdateGroupMembersRequest) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupMembersRequest) ValidateBasic() error {
	if m.GroupId == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group")

	}
	_, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if len(m.MemberUpdates) == 0 {
		return sdkerrors.Wrap(ErrEmpty, "member updates")
	}
	members := Members{Members: m.MemberUpdates}
	if err := members.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "members")
	}
	return nil
}

func (m *MsgUpdateGroupMembersRequest) GetGroupID() uint64 {
	return m.GroupId
}

var _ sdk.Msg = &MsgCreateGroupAccountRequest{}

// Route Implements Msg.
func (m MsgCreateGroupAccountRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgCreateGroupAccountRequest) Type() string { return TypeMsgCreateGroupAccount }

// GetSignBytes Implements Msg.
func (m MsgCreateGroupAccountRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgCreateGroupAccountRequest.
func (m MsgCreateGroupAccountRequest) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgCreateGroupAccountRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		return sdkerrors.Wrap(err, "admin")
	}
	if m.GroupId == 0 {
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

var _ sdk.Msg = &MsgUpdateGroupAccountAdminRequest{}

// Route Implements Msg.
func (m MsgUpdateGroupAccountAdminRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgUpdateGroupAccountAdminRequest) Type() string { return TypeMsgUpdateGroupAccountAdmin }

// GetSignBytes Implements Msg.
func (m MsgUpdateGroupAccountAdminRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgUpdateGroupAccountAdminRequest.
func (m MsgUpdateGroupAccountAdminRequest) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupAccountAdminRequest) ValidateBasic() error {
	admin, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	newAdmin, err := sdk.AccAddressFromBech32(m.NewAdmin)
	if err != nil {
		return sdkerrors.Wrap(err, "new admin")
	}

	_, err = sdk.AccAddressFromBech32(m.GroupAccount)
	if err != nil {
		return sdkerrors.Wrap(err, "group account")
	}

	if admin.Equals(newAdmin) {
		return sdkerrors.Wrap(ErrInvalid, "new and old admin are the same")
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateGroupAccountDecisionPolicyRequest{}
var _ types.UnpackInterfacesMessage = MsgUpdateGroupAccountDecisionPolicyRequest{}

func NewMsgUpdateGroupAccountDecisionPolicyRequest(admin sdk.AccAddress, groupAccount sdk.AccAddress, decisionPolicy DecisionPolicy) (*MsgUpdateGroupAccountDecisionPolicyRequest, error) {
	m := &MsgUpdateGroupAccountDecisionPolicyRequest{
		Admin:        admin.String(),
		GroupAccount: groupAccount.String(),
	}
	err := m.SetDecisionPolicy(decisionPolicy)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MsgUpdateGroupAccountDecisionPolicyRequest) SetDecisionPolicy(decisionPolicy DecisionPolicy) error {
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

// Route Implements Msg.
func (m MsgUpdateGroupAccountDecisionPolicyRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgUpdateGroupAccountDecisionPolicyRequest) Type() string {
	return TypeMsgUpdateGroupAccountDecisionPolicy
}

// GetSignBytes Implements Msg.
func (m MsgUpdateGroupAccountDecisionPolicyRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgUpdateGroupAccountDecisionPolicyRequest.
func (m MsgUpdateGroupAccountDecisionPolicyRequest) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupAccountDecisionPolicyRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	_, err = sdk.AccAddressFromBech32(m.GroupAccount)
	if err != nil {
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

var _ sdk.Msg = &MsgUpdateGroupAccountMetadataRequest{}

// Route Implements Msg.
func (m MsgUpdateGroupAccountMetadataRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgUpdateGroupAccountMetadataRequest) Type() string { return TypeMsgUpdateGroupAccountComment }

// GetSignBytes Implements Msg.
func (m MsgUpdateGroupAccountMetadataRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgUpdateGroupAccountMetadataRequest.
func (m MsgUpdateGroupAccountMetadataRequest) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgUpdateGroupAccountMetadataRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	_, err = sdk.AccAddressFromBech32(m.GroupAccount)
	if err != nil {
		return sdkerrors.Wrap(err, "group account")
	}

	return nil
}

var _ sdk.Msg = &MsgCreateGroupAccountRequest{}
var _ types.UnpackInterfacesMessage = MsgCreateGroupAccountRequest{}

// NewMsgCreateGroupAccountRequest creates a new MsgCreateGroupAccountRequest.
func NewMsgCreateGroupAccountRequest(admin sdk.AccAddress, group uint64, metadata []byte, decisionPolicy DecisionPolicy) (*MsgCreateGroupAccountRequest, error) {
	m := &MsgCreateGroupAccountRequest{
		Admin:    admin.String(),
		GroupId:  group,
		Metadata: metadata,
	}
	err := m.SetDecisionPolicy(decisionPolicy)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MsgCreateGroupAccountRequest) GetAdmin() string {
	return m.Admin
}

func (m *MsgCreateGroupAccountRequest) GetGroupID() uint64 {
	return m.GroupId
}

func (m *MsgCreateGroupAccountRequest) GetMetadata() []byte {
	return m.Metadata
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

var _ sdk.Msg = &MsgCreateProposalRequest{}

// NewMsgCreateProposalRequest creates a new MsgCreateProposalRequest.
func NewMsgCreateProposalRequest(acc string, proposers []string, msgs []sdk.Msg, metadata []byte) (*MsgCreateProposalRequest, error) {
	m := &MsgCreateProposalRequest{
		GroupAccount: acc,
		Proposers:    proposers,
		Metadata:     metadata,
	}
	err := m.SetMsgs(msgs)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Route Implements Msg.
func (m MsgCreateProposalRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgCreateProposalRequest) Type() string { return TypeMsgCreateProposal }

// GetSignBytes Implements Msg.
func (m MsgCreateProposalRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgCreateProposalRequest.
func (m MsgCreateProposalRequest) GetSigners() []sdk.AccAddress {
	addrs := make([]sdk.AccAddress, len(m.Proposers))
	for i, proposer := range m.Proposers {
		addr, err := sdk.AccAddressFromBech32(proposer)
		if err != nil {
			panic(err)
		}
		addrs[i] = addr
	}
	return addrs
}

// ValidateBasic does a sanity check on the provided data
func (m MsgCreateProposalRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.GroupAccount)
	if err != nil {
		return sdkerrors.Wrap(err, "group account")
	}

	if len(m.Proposers) == 0 {
		return sdkerrors.Wrap(ErrEmpty, "proposers")
	}
	addrs := make([]sdk.AccAddress, len(m.Proposers))
	for i, proposer := range m.Proposers {
		addr, err := sdk.AccAddressFromBech32(proposer)
		if err != nil {
			return sdkerrors.Wrap(err, "proposers")
		}
		addrs[i] = addr
	}
	if err := AccAddresses(addrs).ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "proposers")
	}

	msgs := m.GetMsgs()
	for i, msg := range msgs {
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
	// for _, m := range m.Msgs {
	// 	err := types.UnpackInterfaces(m, unpacker)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	for _, any := range m.Msgs {
		var msg sdk.Msg
		err := unpacker.UnpackAny(any, &msg)
		if err != nil {
			return err
		}
	}

	return nil
}

var _ sdk.Msg = &MsgVoteRequest{}

// Route Implements Msg.
func (m MsgVoteRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgVoteRequest) Type() string { return TypeMsgVote }

// GetSignBytes Implements Msg.
func (m MsgVoteRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgVoteRequest.
func (m MsgVoteRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Voter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgVoteRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Voter)
	if err != nil {
		return sdkerrors.Wrap(err, "voter")
	}
	if m.ProposalId == 0 {
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

var _ sdk.Msg = &MsgExecRequest{}

// Route Implements Msg.
func (m MsgExecRequest) Route() string { return RouterKey }

// Type Implements Msg.
func (m MsgExecRequest) Type() string { return TypeMsgExec }

// GetSignBytes Implements Msg.
func (m MsgExecRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners returns the expected signers for a MsgExecRequest.
func (m MsgExecRequest) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgExecRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		return sdkerrors.Wrap(err, "signer")
	}
	if m.ProposalId == 0 {
		return sdkerrors.Wrap(ErrEmpty, "proposal")
	}
	return nil
}
