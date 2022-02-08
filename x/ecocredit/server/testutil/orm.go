package testutil

import (
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	orderbookv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/orderbook/v1beta1"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var TestModuleSchema = ormdb.ModuleSchema{
	FileDescriptors: map[uint32]protoreflect.FileDescriptor{
		1: ecocreditv1beta1.File_regen_ecocredit_v1beta1_state_proto,
		2: marketplacev1beta1.File_regen_ecocredit_marketplace_v1beta1_state_proto,
		3: orderbookv1beta1.File_regen_ecocredit_orderbook_v1beta1_memory_proto,
	},
}
