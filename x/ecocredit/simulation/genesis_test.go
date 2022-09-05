package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

	marketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation"
)

func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: math.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	bz := simState.GenState[ecocredit.ModuleName]

	db := dbm.NewMemDB()
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)

	ormCtx := ormtable.WrapContextDefault(backend)
	baseStore, err := api.NewStateStore(ormdb)
	require.NoError(t, err)

	marketStore, err := marketapi.NewStateStore(ormdb)
	require.NoError(t, err)

	jsonSource, err := ormjson.NewRawMessageSource(bz)
	require.NoError(t, err)

	err = ormdb.ImportJSON(ormCtx, jsonSource)
	require.NoError(t, err)

	allowListEnabled, err := baseStore.ClassCreatorAllowlistTable().Get(ormCtx)
	require.NoError(t, err)

	require.True(t, allowListEnabled.Enabled)

	creator, err := sdk.AccAddressFromBech32("regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4")
	require.NoError(t, err)

	_, err = baseStore.AllowedClassCreatorTable().Get(ormCtx, creator)
	require.NoError(t, err)

	classFee, err := baseStore.ClassFeeTable().Get(ormCtx)
	require.NoError(t, err)

	require.Equal(t, classFee.Fee.Denom, sdk.DefaultBondDenom)
	require.Equal(t, classFee.Fee.Amount, "8")

	allowedDenom, err := marketStore.AllowedDenomTable().Get(ormCtx, sdk.DefaultBondDenom)
	require.NoError(t, err)
	require.Equal(t, allowedDenom.BankDenom, sdk.DefaultBondDenom)
	require.Equal(t, allowedDenom.DisplayDenom, sdk.DefaultBondDenom)
	require.Equal(t, allowedDenom.Exponent, uint32(18))
}
