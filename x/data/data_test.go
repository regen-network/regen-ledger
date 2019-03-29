package data_test

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/prop"
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/graph/binary"
	"github.com/regen-network/regen-ledger/graph/gen"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
	"github.com/regen-network/regen-ledger/x/schema"
	schematest "github.com/regen-network/regen-ledger/x/schema/test"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"testing"
)

type Suite struct {
	suite.Suite
	SchemaKeeper schema.Keeper
	Keeper       data.Keeper
	Handler      sdk.Handler
	Ctx          sdk.Context
	Cms          store.CommitMultiStore
	AnAddr       sdk.AccAddress
	Resolver     graph.SchemaResolver
}

func (s *Suite) SetupTest() {
	db := dbm.NewMemDB()
	s.Cms = store.NewCommitMultiStore(db)
	schemaKey := sdk.NewKVStoreKey("schema")
	dataKey := sdk.NewKVStoreKey("data")
	cdc := codec.New()
	schema.RegisterCodec(cdc)
	s.SchemaKeeper = schema.NewKeeper(schemaKey, cdc)
	s.Keeper = data.NewKeeper(dataKey, s.SchemaKeeper, cdc)
	s.Cms.MountStoreWithDB(schemaKey, sdk.StoreTypeIAVL, db)
	s.Cms.MountStoreWithDB(dataKey, sdk.StoreTypeIAVL, db)
	_ = s.Cms.LoadLatestVersion()
	s.Ctx = sdk.NewContext(s.Cms, abci.Header{}, false, log.NewNopLogger())
	s.AnAddr = sdk.AccAddress{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s.Handler = data.NewHandler(s.Keeper)
	s.Resolver = schema.NewOnChainSchemaResolver(s.SchemaKeeper, s.Ctx)
	schematest.CreateSampleSchema(s.Suite, s.SchemaKeeper, s.Ctx, s.AnAddr)
}

func (s *Suite) TestStoreDataGraph() {
	params := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(params)
	properties.Property("can round trip store/retrieve graphs",
		prop.ForAll(func(g1 graph.Graph) (bool, error) {
			buf := new(bytes.Buffer)
			err := binary.SerializeGraph(s.Resolver, g1, buf)
			if err != nil {
				return false, err
			}
			hash := graph.Hash(g1)

			// check if we have existing data (because the generator repeats values)
			addr := types.GetDataAddressOnChainGraph(hash)
			bz, err := s.Keeper.GetData(s.Ctx, addr)
			if bz == nil {
				res := s.Handler(s.Ctx, data.MsgStoreGraph{Hash: hash, Data: buf.Bytes(), Signer: s.AnAddr})
				if res.Code != sdk.CodeOK {
					return false, fmt.Errorf("%+v", res)
				}

				url := res.Tags[0].Value
				addr2 := types.MustDecodeDataURL(string(url))
				if !bytes.Equal(addr, addr2) {
					return false, fmt.Errorf("unexpected DataAddress %+v, %+v", []byte(addr), []byte(addr2))
				}

				// verify can't store same graph again
				res = s.Handler(s.Ctx, data.MsgStoreGraph{Hash: hash, Data: buf.Bytes(), Signer: s.AnAddr})
				if res.Code == sdk.CodeOK {
					return false, fmt.Errorf("shouldn't be able to store the same graph twice")
				}

				bz, err = s.Keeper.GetData(s.Ctx, addr)
				if err != nil {
					return false, err
				}
			}

			g2, err := binary.DeserializeGraph(s.Resolver, bytes.NewBuffer(bz))
			if err != nil {
				return false, err
			}

			hash2 := graph.Hash(g2)
			if !bytes.Equal(hash, hash2) {
				return false, fmt.Errorf("wrong hash")
			}

			s1, err := graph.CanonicalString(g1)
			if err != nil {
				return false, err
			}

			s2, err := graph.CanonicalString(g2)
			if err != nil {
				return false, err
			}

			if s1 != s2 {
				return false, fmt.Errorf("wrong canonical text")
			}

			return true, nil
		}, gen.Graph(s.Resolver)))
	properties.TestingRun(s.T())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
