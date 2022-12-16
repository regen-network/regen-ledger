package v1

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcodec "github.com/cosmos/cosmos-sdk/x/gov/codec"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	groupcodec "github.com/cosmos/cosmos-sdk/x/group/codec"
)

func RegisterTypes(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*govv1beta1.Content)(nil), &AllowDenomProposal{})
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgBuyDirect{}, "regen.marketplace/MsgBuyDirect", nil)
	cdc.RegisterConcrete(&MsgSell{}, "regen.marketplace/MsgSell", nil)
	cdc.RegisterConcrete(&MsgCancelSellOrder{}, "regen.marketplace/MsgCancelSellOrder", nil)
	cdc.RegisterConcrete(&MsgUpdateSellOrders{}, "regen.marketplace/MsgUpdateSellOrders", nil)
	cdc.RegisterConcrete(&MsgRemoveAllowedDenom{}, "regen.marketplace/MsgRemoveAllowedDenom", nil)
	cdc.RegisterConcrete(&MsgAddAllowedDenom{}, "regen.marketplace/MsgAddAllowedDenom", nil)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz, gov, and group Amino codec so that
	// this can later be used to properly serialize MsgGrant, MsgExec and MsgSubmitProposal instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
	RegisterLegacyAminoCodec(groupcodec.Amino)
}
