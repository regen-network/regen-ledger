package data

import (
	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	api "github.com/regen-network/regen-ledger/api/regen/data/v1"
)

const (
	ModuleName = "data"
)

const (
	ORMPrefix byte = iota
)

var ModuleSchema = ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: api.File_regen_data_v1_state_proto.Path()},
	},
	Prefix: []byte{ORMPrefix},
}
