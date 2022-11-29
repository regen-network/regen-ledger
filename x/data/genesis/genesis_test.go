package genesis

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/x/data/v2"
)

func TestValidateGenesis(t *testing.T) {
	t.Parallel()

	// initial setup

	moduleDB, err := ormdb.NewModuleDB(&data.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)

	ss, err := api.NewStateStore(moduleDB)
	require.NoError(t, err)

	// valid state (all state messages)

	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	require.NoError(t, ss.DataIDTable().Insert(ormCtx, &api.DataID{
		Id:  []byte("foo"),
		Iri: "regen:13toVgf5aZqSVSeJQv562xkkeoe3rr3bJWa29PHVKVf77VAkVMcDvVd.rdf",
	}))

	timestamp, err := types.ParseDate("timestamp", "2020-01-01")
	require.NoError(t, err)

	require.NoError(t, ss.DataAnchorTable().Insert(ormCtx, &api.DataAnchor{
		Id:        []byte("foo"),
		Timestamp: timestamppb.New(timestamp),
	}))

	require.NoError(t, ss.DataAttestorTable().Insert(ormCtx, &api.DataAttestor{
		Id:        []byte("foo"),
		Attestor:  sdk.AccAddress("alice"),
		Timestamp: timestamppb.New(timestamp),
	}))

	require.NoError(t, ss.ResolverTable().Insert(ormCtx, &api.Resolver{
		Url:     "https://foo.bar",
		Manager: sdk.AccAddress("alice"),
	}))

	require.NoError(t, ss.DataResolverTable().Insert(ormCtx, &api.DataResolver{
		Id:         []byte("foo"),
		ResolverId: 1,
	}))

	target := ormjson.NewRawMessageTarget()
	require.NoError(t, moduleDB.ExportJSON(ormCtx, target))

	genesisJSON, err := target.JSON()
	require.NoError(t, err)

	err = ValidateGenesis(genesisJSON)
	require.NoError(t, err)

	// invalid DataID

	ormCtx = ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	require.NoError(t, ss.DataIDTable().Insert(ormCtx, &api.DataID{}))

	target = ormjson.NewRawMessageTarget()
	require.NoError(t, moduleDB.ExportJSON(ormCtx, target))

	genesisJSON, err = target.JSON()
	require.NoError(t, err)

	err = ValidateGenesis(genesisJSON)
	require.ErrorContains(t, err, "Error in JSON for table regen.data.v1.DataID")

	// invalid DataAnchor

	ormCtx = ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	require.NoError(t, ss.DataAnchorTable().Insert(ormCtx, &api.DataAnchor{}))

	target = ormjson.NewRawMessageTarget()
	require.NoError(t, moduleDB.ExportJSON(ormCtx, target))

	genesisJSON, err = target.JSON()
	require.NoError(t, err)

	err = ValidateGenesis(genesisJSON)
	require.ErrorContains(t, err, "Error in JSON for table regen.data.v1.DataAnchor")

	// invalid DataAttestor

	ormCtx = ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	require.NoError(t, ss.DataAttestorTable().Insert(ormCtx, &api.DataAttestor{}))

	target = ormjson.NewRawMessageTarget()
	require.NoError(t, moduleDB.ExportJSON(ormCtx, target))

	genesisJSON, err = target.JSON()
	require.NoError(t, err)

	err = ValidateGenesis(genesisJSON)
	require.ErrorContains(t, err, "Error in JSON for table regen.data.v1.DataAttestor")

	// invalid Resolver

	ormCtx = ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	require.NoError(t, ss.ResolverTable().Insert(ormCtx, &api.Resolver{}))

	target = ormjson.NewRawMessageTarget()
	require.NoError(t, moduleDB.ExportJSON(ormCtx, target))

	genesisJSON, err = target.JSON()
	require.NoError(t, err)

	err = ValidateGenesis(genesisJSON)
	require.ErrorContains(t, err, "Error in JSON for table regen.data.v1.Resolver")

	// invalid DataResolver

	ormCtx = ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	require.NoError(t, ss.DataResolverTable().Insert(ormCtx, &api.DataResolver{}))

	target = ormjson.NewRawMessageTarget()
	require.NoError(t, moduleDB.ExportJSON(ormCtx, target))

	genesisJSON, err = target.JSON()
	require.NoError(t, err)

	err = ValidateGenesis(genesisJSON)
	require.ErrorContains(t, err, "Error in JSON for table regen.data.v1.DataResolver")
}
