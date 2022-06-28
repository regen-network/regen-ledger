package testsuite

import (
	"crypto"
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/data"
)

type GenesisTestSuite struct {
	suite.Suite

	fixtureFactory testutil.FixtureFactory
	fixture        testutil.Fixture

	genesisCtx  types.Context
	msgClient   data.MsgClient
	queryClient data.QueryClient
	addr1       sdk.AccAddress
	addr2       sdk.AccAddress
	hash1       *data.ContentHash
	hash2       *data.ContentHash
}

func NewGenesisTestSuite(fixtureFactory testutil.FixtureFactory) *GenesisTestSuite {
	return &GenesisTestSuite{fixtureFactory: fixtureFactory}
}

func (s *GenesisTestSuite) SetupSuite() {
	require := s.Require()

	s.fixture = s.fixtureFactory.Setup()

	blockTime := time.Now().UTC()
	sdkCtx := s.fixture.Context().(types.Context).WithBlockTime(blockTime)
	s.genesisCtx = types.Context{Context: sdkCtx}

	s.msgClient = data.NewMsgClient(s.fixture.TxConn())
	s.queryClient = data.NewQueryClient(s.fixture.QueryConn())
	require.GreaterOrEqual(len(s.fixture.Signers()), 2)
	s.addr1 = s.fixture.Signers()[0]
	s.addr2 = s.fixture.Signers()[1]

	content := []byte("xyzabc123")
	hash := crypto.BLAKE2b_256.New()
	_, err := hash.Write(content)
	require.NoError(err)
	digest := hash.Sum(nil)

	graphHash := &data.ContentHash_Graph{
		Hash:                      digest,
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
	}
	s.hash1 = &data.ContentHash{Graph: graphHash}

	rawHash := &data.ContentHash_Raw{
		Hash:            digest,
		DigestAlgorithm: data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		MediaType:       data.RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
	}
	s.hash2 = &data.ContentHash{Raw: rawHash}
}

func (s *GenesisTestSuite) TestInitGenesis() {
	require := s.Require()

	idsJSON := `[{"id":"MQ==","iri":"regen:13toVhBfBKwiWB683CBZWFoYmH1KGU2umANuhiGCCZZTVv964SkBbCL.rdf"},{"id":"Mg==","iri":"regen:114DDL1RtVwKpfqgaPfAG153ckiKfuPEgTT7tEGs1Hic5sC9dCta.bin"}]`
	anchorsJSON := `[{"id":"YQ==","timestamp":"2022-04-05T07:03:19.464153411Z"},{"id":"Yg==","timestamp":"2022-04-05T06:52:42.106314060Z"}]`
	attestorsJSON := `[{"attestor":"CyzUKxKh0MHmBM5vlN0/L8suJzQ=","timestamp":"2022-04-05T07:06:59.400392064Z"},{"attestor":"hUjhdJPEILo2/U4kA3V65IXK4Cs=","timestamp":"2022-04-05T07:06:59.400392064Z"}]`
	resolverInfoJSON := `[{"id":"0","url":"https://foo.bar","manager":"XqdMDUBiSacEypUx5lmrYfxGgec="},{"id":"0","url":"https://foo1.bar","manager":"s8uqM3U2HfHgopDvaLq55Gsxnek="}]`

	wrapper := map[string]json.RawMessage{}
	wrapper[gogoproto.MessageName(&data.DataID{})] = []byte(idsJSON)
	wrapper[gogoproto.MessageName(&data.DataAnchor{})] = []byte(anchorsJSON)
	wrapper[gogoproto.MessageName(&data.DataAttestor{})] = []byte(attestorsJSON)
	wrapper[gogoproto.MessageName(&data.ResolverInfo{})] = []byte(resolverInfoJSON)

	bz, err := json.Marshal(wrapper)
	require.NoError(err)
	wrapper = map[string]json.RawMessage{}
	wrapper[data.ModuleName] = bz

	_, err = s.fixture.InitGenesis(s.genesisCtx.Context, wrapper)
	require.NoError(err)

	exported, err := s.fixture.ExportGenesis(s.genesisCtx.Context)
	require.NoError(err)
	require.NotNil(exported)

}

func (s *GenesisTestSuite) TearDownSuite() {
	s.fixture.Teardown()
}
