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
	cdc.RegisterConcrete(MsgCreateGroup{}, "cosmos-sdk/MsgCreateGroup", nil)
	cdc.RegisterConcrete(MsgUpdateGroupMembers{}, "cosmos-sdk/MsgUpdateGroupMembers", nil)
	cdc.RegisterConcrete(MsgUpdateGroupAdmin{}, "cosmos-sdk/MsgUpdateGroupAdmin", nil)
	cdc.RegisterConcrete(MsgUpdateGroupComment{}, "cosmos-sdk/MsgUpdateGroupComment", nil)
	cdc.RegisterConcrete(MsgCreateGroupAccount{}, "cosmos-sdk/MsgCreateGroupAccount", nil)
	cdc.RegisterConcrete(MsgVote{}, "cosmos-sdk/group/MsgVote", nil)
	cdc.RegisterConcrete(MsgExec{}, "cosmos-sdk/group/MsgExec", nil)

	// oh man... amino
	// cdc.RegisterConcrete(StdDecisionPolicy{}, "cosmos-sdk/StdDecisionPolicy", nil)
	// cdc.RegisterConcrete(&StdDecisionPolicy_Threshold{}, "cosmos-sdk/StdDecisionPolicy_Threshold", nil)
	cdc.RegisterConcrete(&ThresholdDecisionPolicy{}, "cosmos-sdk/ThresholdDecisionPolicy", nil)
	// cdc.RegisterInterface((*isStdDecisionPolicy_Sum)(nil), nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGroup{},
		&MsgUpdateGroupMembers{},
		&MsgUpdateGroupComment{},
		&MsgCreateGroupAccount{},
		&MsgVote{},
		&MsgExec{},
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
