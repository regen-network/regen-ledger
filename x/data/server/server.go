package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/modules/incubator/orm"
	_ "github.com/cosmos/modules/incubator/orm"
	"github.com/gogo/protobuf/types"
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

type Server interface {
	data.MsgServer
	data.QueryServer
}

func NewServer(storeKey sdk.StoreKey) Server {
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
