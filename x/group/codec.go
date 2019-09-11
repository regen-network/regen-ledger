package group

import "github.com/cosmos/cosmos-sdk/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateGroup{}, "group/MsgCreateGroup", nil)
	cdc.RegisterConcrete(Group{}, "group/Group", nil)
	cdc.RegisterConcrete(GroupAccount{}, "group/GroupAccount", nil)
}
