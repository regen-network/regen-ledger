package ecocredit

import (
	ormapi "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
)

const (
	// ModuleName is the module name constant used in many places
	ModuleName        = "ecocredit"
	RouterKey         = ModuleName
	DefaultParamspace = ModuleName

	ORMPrefix byte = 0x7
)

var ModuleSchema = ormapi.ModuleSchemaDescriptor{
	SchemaFile: []*ormapi.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: basketapi.File_regen_ecocredit_basket_v1_state_proto.Path()},
		{Id: 2, ProtoFileName: api.File_regen_ecocredit_v1_state_proto.Path()},
		{Id: 3, ProtoFileName: marketApi.File_regen_ecocredit_marketplace_v1_state_proto.Path()},
	},
	Prefix: []byte{ORMPrefix},
}
