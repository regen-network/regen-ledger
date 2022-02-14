package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/v2/x/bond"
)

const (
	BondInfoTablePrefix          byte = 0x0
	BondInfoByEmissionDenomIndex byte = 0x1
)

type serverImpl struct {
	storeKey sdk.StoreKey

	paramSpace    paramtypes.Subspace
	bankKeeper    bond.BankKeeper
	accountKeeper bond.AccountKeeper

	// Bond table
	bondInfoTable           orm.PrimaryKeyTable
	bondInfoEmissionIDIndex orm.Index
	bondInfoSeq             orm.Sequence
}

func (s serverImpl) PruneOrders(ctx sdk.Context) error {
	//TODO implement me
	panic("implement me")
}

func newServer(storeKey sdk.StoreKey, paramSpace paramtypes.Subspace, accountKeeper bond.AccountKeeper, bankKeeper bond.BankKeeper, cdc codec.Codec) serverImpl {
	s := serverImpl{
		storeKey:      storeKey,
		paramSpace:    paramSpace,
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}

	bondInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(BondInfoTablePrefix, storeKey, &bond.BondInfo{}, cdc)
	if err != nil {
		panic(err.Error())
	}

	s.bondInfoEmissionIDIndex, err = orm.NewIndex(bondInfoTableBuilder, BondInfoByEmissionDenomIndex, func(value interface{}) ([]interface{}, error) {
		bondInfo, ok := value.(*bond.BondInfo)
		if !ok {
			return nil, sdkerrors.ErrInvalidType.Wrapf("expected %T got %T", bond.BondInfo{}, value)
		}
		return []interface{}{bondInfo.EmissionDenom}, nil
	}, bond.BondInfo{}.EmissionDenom)
	if err != nil {
		panic(err.Error())
	}

	s.bondInfoTable = bondInfoTableBuilder.Build()
	return s
}

func RegisterServices(
	configurator server.Configurator,
	paramSpace paramtypes.Subspace,
	accountKeeper bond.AccountKeeper,
	bankKeeper bond.BankKeeper) bond.Keeper {

	impl := newServer(configurator.ModuleKey(), paramSpace, accountKeeper, bankKeeper, configurator.Marshaler())
	bond.RegisterMsgServer(configurator.MsgServer(), impl)
	bond.RegisterQueryServer(configurator.QueryServer(), impl)
	configurator.RegisterGenesisHandlers(impl.InitGenesis, impl.ExportGenesis)
	//configurator.RegisterWeightedOperationsHandler(impl.WeightedOperations)
	//configurator.RegisterInvariantsHandler(impl.RegisterInvariants)
	return impl
}
