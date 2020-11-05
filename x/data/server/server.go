package server

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/data"
)

const (
	AnchorTablePrefix  byte = 0x0
	SignersTablePrefix byte = 0x1
	DataTablePrefix    byte = 0x2
)

type serverImpl struct {
	storeKey sdk.StoreKey

	anchorTable orm.Table

	signersTable orm.Table
}

func newServer(storeKey sdk.StoreKey) serverImpl {
	s := serverImpl{storeKey: storeKey}

	anchorTable := orm.NewTableBuilder(AnchorTablePrefix, storeKey, &types.Timestamp{}, orm.Max255DynamicLengthIndexKeyCodec{})
	s.anchorTable = anchorTable.Build()

	signersTable := orm.NewTableBuilder(SignersTablePrefix, storeKey, &data.Signers{}, orm.Max255DynamicLengthIndexKeyCodec{})
	s.signersTable = signersTable.Build()

	return s
}

func DataKey(cid []byte) []byte {
	return append([]byte{DataTablePrefix}, cid...)
}

func RegisterServices(storeKey sdk.StoreKey, configurator module.Configurator) {
	impl := newServer(storeKey)
	data.RegisterMsgServer(configurator.MsgServer(), impl)
	data.RegisterQueryServer(configurator.QueryServer(), impl)
}
