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
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
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

	var params core.Params
	simState.Cdc.MustUnmarshalJSON(wrapper[proto.MessageName(&core.Params{})], &params)

	require.Equal(t, params.AllowedClassCreators, []string{"regen1tnh2q55v8wyygtt9srz5safamzdengsnlm0yy4"})
	require.Equal(t, params.AllowlistEnabled, true)
	require.Equal(t, params.CreditClassFee, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(9))))
	require.Equal(t, params.AllowlistEnabled, true)
}
