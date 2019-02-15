package agent

import "github.com/cosmos/cosmos-sdk/codec"


func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateAgent{}, "agent/MsgCreateAgent", nil)
	cdc.RegisterConcrete(MsgUpdateAgent{}, "agent/MsgUpdateAgent", nil)
	cdc.RegisterConcrete(AgentInfo{}, "agent/AgentInfo", nil)
}
