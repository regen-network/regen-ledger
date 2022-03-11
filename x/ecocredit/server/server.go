package server

import (
	marketApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/ormutil"
)

const (
	TradableBalancePrefix    byte = 0x0
	TradableSupplyPrefix     byte = 0x1
	RetiredBalancePrefix     byte = 0x2
	RetiredSupplyPrefix      byte = 0x3
	CreditTypeSeqTablePrefix byte = 0x4
	ClassInfoTablePrefix     byte = 0x5
	BatchInfoTablePrefix     byte = 0x6

	ProjectInfoTablePrefix    byte = 0x10
	ProjectInfoTableSeqPrefix byte = 0x11
	ProjectsByClassIDIndex    byte = 0x12
	BatchesByProjectIndex     byte = 0x13

	// sell order table
	SellOrderTablePrefix             byte = 0x20
	SellOrderTableSeqPrefix          byte = 0x21
	SellOrderByExpirationIndexPrefix byte = 0x22
	SellOrderByAddressIndexPrefix    byte = 0x23
	SellOrderByBatchDenomIndexPrefix byte = 0x24

	// buy order table
	BuyOrderTablePrefix             byte = 0x25
	BuyOrderTableSeqPrefix          byte = 0x26
	BuyOrderByExpirationIndexPrefix byte = 0x27
	BuyOrderByAddressIndexPrefix    byte = 0x28

	AskDenomTablePrefix byte = 0x30
)

type serverImpl struct {
	storeKey sdk.StoreKey

	paramSpace    paramtypes.Subspace
	bankKeeper    ecocredit.BankKeeper
	accountKeeper ecocredit.AccountKeeper

	// Store sequence numbers per credit type
	creditTypeSeqTable orm.PrimaryKeyTable

	classInfoTable orm.PrimaryKeyTable
	batchInfoTable orm.PrimaryKeyTable

	// sell order table
	sellOrderTable             orm.AutoUInt64Table
	sellOrderByAddressIndex    orm.Index
	sellOrderByBatchDenomIndex orm.Index
	sellOrderByExpirationIndex orm.Index

	// buy order table
	buyOrderTable             orm.AutoUInt64Table
	buyOrderByAddressIndex    orm.Index
	buyOrderByExpirationIndex orm.Index

	askDenomTable orm.PrimaryKeyTable

	// project table
	projectInfoTable        orm.PrimaryKeyTable
	projectInfoSeq          orm.Sequence
	projectsByClassIDIndex  orm.Index
	batchesByProjectIDIndex orm.Index

	basketKeeper basket.Keeper

	db ormdb.ModuleDB
}

var ModuleSchema = ormdb.ModuleSchema{
	FileDescriptors: map[uint32]protoreflect.FileDescriptor{
		1: api.File_regen_ecocredit_v1_state_proto,
		2: basketapi.File_regen_ecocredit_basket_v1_state_proto,
		3: marketApi.File_regen_ecocredit_marketplace_v1_state_proto,
	},
	Prefix: []byte{ecocredit.ORMPrefix},
}

func newServer(storeKey sdk.StoreKey, paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper, bankKeeper ecocredit.BankKeeper, distKeeper ecocredit.DistributionKeeper, cdc codec.Codec) serverImpl {
	s := serverImpl{
		storeKey:      storeKey,
		paramSpace:    paramSpace,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}

	creditTypeSeqTable, err := orm.NewPrimaryKeyTableBuilder(ecocredit.CreditTypeSeqTablePrefix, storeKey, &ecocredit.CreditTypeSeq{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	s.creditTypeSeqTable = creditTypeSeqTable.Build()

	classInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(ecocredit.ClassInfoTablePrefix, storeKey, &ecocredit.ClassInfo{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	s.classInfoTable = classInfoTableBuilder.Build()

	batchInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(ecocredit.BatchInfoTablePrefix, storeKey, &ecocredit.BatchInfo{}, cdc)
	if err != nil {
		panic(err.Error())
	}

	s.batchesByProjectIDIndex, err = orm.NewIndex(batchInfoTableBuilder, BatchesByProjectIndex, func(value interface{}) ([]interface{}, error) {
		batchInfo, ok := value.(*ecocredit.BatchInfo)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("expected %T got %T", ecocredit.BatchInfo{}, value)
		}
		return []interface{}{batchInfo.ProjectId}, nil
	}, ecocredit.BatchInfo{}.ProjectId)
	if err != nil {
		panic(err.Error())
	}

	s.batchInfoTable = batchInfoTableBuilder.Build()

	sellOrderTableBuilder, err := orm.NewAutoUInt64TableBuilder(SellOrderTablePrefix, SellOrderTableSeqPrefix, storeKey, &ecocredit.SellOrder{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	s.sellOrderByAddressIndex, err = orm.NewIndex(sellOrderTableBuilder, SellOrderByAddressIndexPrefix, func(value interface{}) ([]interface{}, error) {
		order, ok := value.(*ecocredit.SellOrder)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("expected %T got %T", ecocredit.SellOrder{}, value)
		}
		addr, err := sdk.AccAddressFromBech32(order.Owner)
		if err != nil {
			return nil, err
		}
		return []interface{}{addr.Bytes()}, nil
	}, []byte{})
	if err != nil {
		panic(err.Error())
	}
	s.sellOrderByBatchDenomIndex, err = orm.NewIndex(sellOrderTableBuilder, SellOrderByBatchDenomIndexPrefix, func(value interface{}) ([]interface{}, error) {
		order, ok := value.(*ecocredit.SellOrder)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("expected %T got %T", ecocredit.SellOrder{}, value)
		}
		return []interface{}{order.BatchDenom}, nil
	}, ecocredit.SellOrder{}.BatchDenom)
	if err != nil {
		panic(err.Error())
	}
	s.sellOrderByExpirationIndex, err = orm.NewIndex(sellOrderTableBuilder, SellOrderByExpirationIndexPrefix, func(value interface{}) ([]interface{}, error) {
		order, ok := value.(*ecocredit.SellOrder)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("expected %T got %T", ecocredit.SellOrder{}, value)
		}
		if order.Expiration == nil {
			return nil, nil
		}
		return []interface{}{uint64(order.Expiration.UnixNano())}, nil
	}, uint64(0))
	if err != nil {
		panic(err.Error())
	}
	s.sellOrderTable = sellOrderTableBuilder.Build()

	buyOrderTableBuilder, err := orm.NewAutoUInt64TableBuilder(BuyOrderTablePrefix, BuyOrderTableSeqPrefix, storeKey, &ecocredit.BuyOrder{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	s.buyOrderByAddressIndex, err = orm.NewIndex(buyOrderTableBuilder, BuyOrderByAddressIndexPrefix, func(value interface{}) ([]interface{}, error) {
		order, ok := value.(*ecocredit.BuyOrder)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("expected %T got %T", ecocredit.BuyOrder{}, value)
		}
		addr, err := sdk.AccAddressFromBech32(order.Buyer)
		if err != nil {
			return nil, err
		}
		return []interface{}{addr.Bytes()}, nil
	}, []byte{})
	if err != nil {
		panic(err.Error())
	}
	s.buyOrderByExpirationIndex, err = orm.NewIndex(buyOrderTableBuilder, BuyOrderByExpirationIndexPrefix, func(value interface{}) ([]interface{}, error) {
		order, ok := value.(*ecocredit.BuyOrder)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("expected %T got %T", ecocredit.BuyOrder{}, value)
		}
		if order.Expiration == nil {
			return nil, nil
		}
		return []interface{}{uint64(order.Expiration.UnixNano())}, nil
	}, uint64(0))
	if err != nil {
		panic(err.Error())
	}
	s.buyOrderTable = buyOrderTableBuilder.Build()

	askDenomTableBuilder, err := orm.NewPrimaryKeyTableBuilder(AskDenomTablePrefix, storeKey, &ecocredit.AskDenom{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	s.askDenomTable = askDenomTableBuilder.Build()

	s.projectInfoSeq = orm.NewSequence(storeKey, ProjectInfoTableSeqPrefix)
	projectInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(ProjectInfoTablePrefix, storeKey, &ecocredit.ProjectInfo{}, cdc)
	if err != nil {
		panic(err.Error())
	}

	s.projectsByClassIDIndex, err = orm.NewIndex(projectInfoTableBuilder, ProjectsByClassIDIndex, func(value interface{}) ([]interface{}, error) {
		projectInfo, ok := value.(*ecocredit.ProjectInfo)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("expected %T got %T", ecocredit.ProjectInfo{}, value)
		}
		return []interface{}{projectInfo.ClassId}, nil
	}, ecocredit.ProjectInfo{}.ClassId)
	if err != nil {
		panic(err.Error())
	}

	s.projectInfoTable = projectInfoTableBuilder.Build()

	s.db, err = ormutil.NewStoreKeyDB(ModuleSchema, storeKey, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	s.basketKeeper = basket.NewKeeper(s.db, s, bankKeeper, distKeeper, storeKey)

	return s
}

func RegisterServices(
	configurator server.Configurator,
	paramSpace paramtypes.Subspace,
	accountKeeper ecocredit.AccountKeeper,
	bankKeeper ecocredit.BankKeeper,
	distKeeper ecocredit.DistributionKeeper,
) ecocredit.Keeper {
	impl := newServer(configurator.ModuleKey(), paramSpace, accountKeeper, bankKeeper, distKeeper, configurator.Marshaler())
	ecocredit.RegisterMsgServer(configurator.MsgServer(), impl)
	ecocredit.RegisterQueryServer(configurator.QueryServer(), impl)
	baskettypes.RegisterMsgServer(configurator.MsgServer(), impl.basketKeeper)
	baskettypes.RegisterQueryServer(configurator.QueryServer(), impl.basketKeeper)

	configurator.RegisterGenesisHandlers(impl.InitGenesis, impl.ExportGenesis)
	configurator.RegisterWeightedOperationsHandler(impl.WeightedOperations)
	configurator.RegisterInvariantsHandler(impl.RegisterInvariants)
	return impl
}
