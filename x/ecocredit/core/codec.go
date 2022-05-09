package core

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterTypes(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*govtypes.Content)(nil), &CreditTypeProposal{})
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateClass{}, "regen.core/MsgCreateClass", nil)
	cdc.RegisterConcrete(&MsgCreateProject{}, "regen.core/MsgCreateProject", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "regen.core/MsgCancel", nil)
	cdc.RegisterConcrete(&MsgCreateBatch{}, "regen.core/MsgCreateBatch", nil)
	cdc.RegisterConcrete(&MsgMintBatchCredits{}, "regen.core/MsgMintBatchCredits", nil)
	cdc.RegisterConcrete(&MsgSealBatch{}, "regen.core/MsgSealBatch", nil)
	cdc.RegisterConcrete(&MsgRetire{}, "regen.core/MsgRetire", nil)
	cdc.RegisterConcrete(&MsgSend{}, "regen.core/MsgSend", nil)
	cdc.RegisterConcrete(&MsgUpdateClassAdmin{}, "regen.core/MsgUpdateClassAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateClassMetadata{}, "regen.core/MsgUpdateClassMetadata", nil)
	cdc.RegisterConcrete(&MsgUpdateClassIssuers{}, "regen.core/MsgUpdateClassIssuers", nil)
	cdc.RegisterConcrete(&MsgUpdateProjectAdmin{}, "regen.core/MsgUpdateProjectAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateProjectMetadata{}, "regen.core/MsgUpdateProjectMetadata", nil)
	cdc.RegisterConcrete(&CreditTypeProposal{}, "regen.core/CreditTypeProposal", nil)
}

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	RegisterLegacyAminoCodec(amino)
}
