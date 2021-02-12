package group

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers all the necessary group module concrete
// types and interfaces with the provided codec reference.
// These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*DecisionPolicy)(nil), nil)
	cdc.RegisterConcrete(&ThresholdDecisionPolicy{}, "cosmos-sdk/ThresholdDecisionPolicy", nil)
	cdc.RegisterConcrete(&MsgCreateGroupRequest{}, "cosmos-sdk/MsgCreateGroup", nil)
	cdc.RegisterConcrete(&MsgUpdateGroupMembersRequest{}, "cosmos-sdk/MsgUpdateGroupMembers", nil)
	cdc.RegisterConcrete(&MsgUpdateGroupAdminRequest{}, "cosmos-sdk/MsgUpdateGroupAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateGroupMetadataRequest{}, "cosmos-sdk/MsgUpdateGroupMetadata", nil)
	cdc.RegisterConcrete(&MsgCreateGroupAccountRequest{}, "cosmos-sdk/MsgCreateGroupAccount", nil)
	cdc.RegisterConcrete(&MsgUpdateGroupAccountAdminRequest{}, "cosmos-sdk/MsgUpdateGroupAccountAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateGroupAccountDecisionPolicyRequest{}, "cosmos-sdk/MsgUpdateGroupAccountDecisionPolicy", nil)
	cdc.RegisterConcrete(&MsgUpdateGroupAccountMetadataRequest{}, "cosmos-sdk/MsgUpdateGroupAccountMetadata", nil)
	cdc.RegisterConcrete(&MsgCreateProposalRequest{}, "cosmos-sdk/group/MsgCreateProposal", nil)
	cdc.RegisterConcrete(&MsgVoteRequest{}, "cosmos-sdk/group/MsgVote", nil)
	cdc.RegisterConcrete(&MsgExecRequest{}, "cosmos-sdk/group/MsgExec", nil)
}

func RegisterTypes(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGroupRequest{},
		&MsgUpdateGroupMembersRequest{},
		&MsgUpdateGroupAdminRequest{},
		&MsgUpdateGroupMetadataRequest{},
		&MsgCreateGroupAccountRequest{},
		&MsgUpdateGroupAccountAdminRequest{},
		&MsgUpdateGroupAccountDecisionPolicyRequest{},
		&MsgUpdateGroupAccountMetadataRequest{},
		&MsgCreateProposalRequest{},
		&MsgVoteRequest{},
		&MsgExecRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &Msg_ServiceDesc)

	registry.RegisterInterface(
		"regen.group.v1alpha1.DecisionPolicy",
		(*DecisionPolicy)(nil),
		&ThresholdDecisionPolicy{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}
