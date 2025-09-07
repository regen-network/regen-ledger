package v1

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterTypes(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*govv1beta1.Content)(nil), &CreditTypeProposal{})
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateClass{}, "regen/MsgCreateClass", nil)
	cdc.RegisterConcrete(&MsgCreateProject{}, "regen/MsgCreateProject", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "regen/MsgCancel", nil)
	cdc.RegisterConcrete(&MsgCreateBatch{}, "regen/MsgCreateBatch", nil)
	cdc.RegisterConcrete(&MsgMintBatchCredits{}, "regen/MsgMintBatchCredits", nil)
	cdc.RegisterConcrete(&MsgSealBatch{}, "regen/MsgSealBatch", nil)
	cdc.RegisterConcrete(&MsgRetire{}, "regen/MsgRetire", nil)
	cdc.RegisterConcrete(&MsgSend{}, "regen/MsgSend", nil)
	cdc.RegisterConcrete(&MsgUpdateClassAdmin{}, "regen/MsgUpdateClassAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateClassMetadata{}, "regen/MsgUpdateClassMetadata", nil)
	cdc.RegisterConcrete(&MsgUpdateClassIssuers{}, "regen/MsgUpdateClassIssuers", nil)
	cdc.RegisterConcrete(&MsgUpdateProjectAdmin{}, "regen/MsgUpdateProjectAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateProjectMetadata{}, "regen/MsgUpdateProjectMetadata", nil)
	cdc.RegisterConcrete(&MsgUpdateBatchMetadata{}, "regen/MsgUpdateBatchMetadata", nil)
	cdc.RegisterConcrete(&CreditTypeProposal{}, "regen/CreditTypeProposal", nil)
	cdc.RegisterConcrete(&MsgBridgeReceive{}, "regen/MsgBridgeReceive", nil)
	cdc.RegisterConcrete(&MsgBridge{}, "regen/MsgBridge", nil)
	cdc.RegisterConcrete(&MsgAddCreditType{}, "regen/MsgAddCreditType", nil)
	cdc.RegisterConcrete(&MsgAddClassCreator{}, "regen/MsgAddClassCreator", nil)
	cdc.RegisterConcrete(&MsgRemoveClassCreator{}, "regen/MsgRemoveClassCreator", nil)
	cdc.RegisterConcrete(&MsgSetClassCreatorAllowlist{}, "regen/MsgSetClassCreatorAllowlist", nil)
	cdc.RegisterConcrete(&MsgUpdateClassFee{}, "regen/MsgUpdateClassFee", nil)
	cdc.RegisterConcrete(&MsgAddAllowedBridgeChain{}, "regen/MsgAddAllowedBridgeChain", nil)
	cdc.RegisterConcrete(&MsgRemoveAllowedBridgeChain{}, "regen/MsgRemoveAllowedBridgeChain", nil)
}
