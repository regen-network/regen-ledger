package group

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterLegacyAminoCodec registers all the necessary group module concrete types and
// interfaces with the provided LegacyAmino codec reference.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*DecisionPolicy)(nil), nil)
	cdc.RegisterConcrete(MsgCreateGroupRequest{}, "cosmos-sdk/MsgCreateGroupRequest", nil)
	cdc.RegisterConcrete(MsgUpdateGroupMembersRequest{}, "cosmos-sdk/MsgUpdateGroupMembersRequest", nil)
	cdc.RegisterConcrete(MsgUpdateGroupAdminRequest{}, "cosmos-sdk/MsgUpdateGroupAdminRequest", nil)
	cdc.RegisterConcrete(MsgUpdateGroupCommentRequest{}, "cosmos-sdk/MsgUpdateGroupCommentRequest", nil)
	cdc.RegisterConcrete(MsgCreateGroupAccountRequest{}, "cosmos-sdk/MsgCreateGroupAccountRequest", nil)
	cdc.RegisterConcrete(MsgVoteRequest{}, "cosmos-sdk/group/MsgVoteRequest", nil)
	cdc.RegisterConcrete(MsgExecRequest{}, "cosmos-sdk/group/MsgExecRequest", nil)

	// oh man... amino
	cdc.RegisterConcrete(&ThresholdDecisionPolicy{}, "cosmos-sdk/ThresholdDecisionPolicy", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGroupRequest{},
		&MsgUpdateGroupMembersRequest{},
		&MsgUpdateGroupCommentRequest{},
		&MsgCreateGroupAccountRequest{},
		&MsgVoteRequest{},
		&MsgExecRequest{},
	)
	registry.RegisterInterface(
		"regen.group.v1alpha1.DecisionPolicy",
		(*DecisionPolicy)(nil),
		&ThresholdDecisionPolicy{},
	)
}

var (
	amino = codec.NewLegacyAmino()

	// moduleCdc references the global group module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding as Amino
	// is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to group and
	// defined at the application level.
	moduleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
