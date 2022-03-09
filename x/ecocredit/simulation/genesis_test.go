package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

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
		InitialStake: 1000,
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var wrapper map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(simState.GenState[ecocredit.ModuleName], &wrapper))

	var ecocreditGenesis ecocredit.GenesisState
	simState.Cdc.MustUnmarshalJSON(wrapper[proto.MessageName(&ecocreditGenesis)], &ecocreditGenesis)

	require.Equal(t, ecocreditGenesis.Params.AllowedClassCreators, []string{"cosmos1tnh2q55v8wyygtt9srz5safamzdengsnqeycj3"})
	require.Equal(t, ecocreditGenesis.Params.AllowlistEnabled, true)
	require.Equal(t, ecocreditGenesis.Params.CreditClassFee, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(9))))
	require.Equal(t, ecocreditGenesis.Params.AllowlistEnabled, true)

	require.Len(t, ecocreditGenesis.ClassInfo, 3)
	require.Len(t, ecocreditGenesis.BatchInfo, 3)
	require.Len(t, ecocreditGenesis.Balances, 6)
	require.Len(t, ecocreditGenesis.Supplies, 3)

	require.Equal(t, ecocreditGenesis.Sequences, []*ecocredit.CreditTypeSeq{
		{
			Abbreviation: "C",
			SeqNumber:    4,
		},
		{
			Abbreviation: "BIO",
			SeqNumber:    4,
		},
	})
}
