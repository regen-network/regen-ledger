package testsuite

import (
	"crypto"
	"encoding/json"
	"time"

	gogoproto "github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/types/v2/testutil/fixture"
	"github.com/regen-network/regen-ledger/x/data/v3"
)

type GenesisTestSuite struct {
	suite.Suite

	fixtureFactory fixture.Factory
	fixture        fixture.Fixture

	genesisCtx  sdk.Context
	msgClient   data.MsgClient
	queryClient data.QueryClient
	addr1       sdk.AccAddress
	addr2       sdk.AccAddress
	hash1       *data.ContentHash
	hash2       *data.ContentHash
}

func NewGenesisTestSuite(fixtureFactory fixture.Factory) *GenesisTestSuite {
	return &GenesisTestSuite{fixtureFactory: fixtureFactory}
}

func (s *GenesisTestSuite) SetupSuite() {
	require := s.Require()

	s.fixture = s.fixtureFactory.Setup()

	blockTime := time.Now().UTC()
	sdkCtx := sdk.UnwrapSDKContext(s.fixture.Context()).WithBlockTime(blockTime)
	s.genesisCtx = sdkCtx

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
		DigestAlgorithm:           uint32(data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256),
		CanonicalizationAlgorithm: uint32(data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_RDFC_1_0),
	}
	s.hash1 = &data.ContentHash{Graph: graphHash}

	rawHash := &data.ContentHash_Raw{
		Hash:            digest,
		DigestAlgorithm: uint32(data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256),
		FileExtension:   "bin",
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

	_, err = s.fixture.InitGenesis(s.genesisCtx, wrapper)
	require.NoError(err)

	exported, err := s.fixture.ExportGenesis(s.genesisCtx)
	require.NoError(err)
	require.NotNil(exported)

}

func (s *GenesisTestSuite) TearDownSuite() {
	s.fixture.Teardown()
}
