package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterTypes(registry cdctypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	registry.RegisterInterface(
		"regen.group.v1alpha1.DecisionPolicy",
		(*DecisionPolicy)(nil),
		&ThresholdDecisionPolicy{},
	)
}
