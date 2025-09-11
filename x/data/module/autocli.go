package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	datav1beta1 "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
)

func (am Module) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              datav1beta1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: false, // use custom commands only until v0.51
			RpcCommandOptions:    []*autocliv1.RpcCommandOptions{},
		},
		Query: &autocliv1.ServiceCommandDescriptor{
			Service:           datav1beta1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
		},
	}
}
