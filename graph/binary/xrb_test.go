package binary_test

import (
	"bytes"
	"fmt"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/graph/binary"
	gengraph "github.com/regen-network/regen-ledger/graph/gen"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/util"
	schematest "github.com/regen-network/regen-ledger/x/schema/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

// TODO graph generator
// TODO verify graph can be serialized and deserialized and is equivalent and has same hash

type TestSuite struct {
	schematest.Harness
}

func (s *TestSuite) SetupSuite() {
	s.Harness.Setup()
	s.Harness.CreateSampleSchema()
}

func (s *TestSuite) TestGenGraph() {
	gs, ok := gen.SliceOfN(3, gengraph.Graph(s.Resolver)).Sample()
	if ok {
		for _, g := range gs.([]graph.Graph) {
			s.T().Logf("Graph %s:\n %s",
				util.MustEncodeBech32(types.Bech32DataAddressPrefix, graph.Hash(g)),
				g.String(),
			)
		}
	}
}

func (s *TestSuite) TestProperties() {
	params := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(params)
	properties.Property("can round trip serialize/deserialize graphs with same hashes",
		prop.ForAll(func(g1 graph.Graph) (bool, error) {
			txt1, err := graph.CanonicalString(g1)
			if err != nil {
				return false, err
			}
			hash1 := graph.Hash(g1)
			w := new(bytes.Buffer)
			err = binary.SerializeGraph(s.Resolver, g1, w)
			if err != nil {
				return false, err
			}
			g2, err := binary.DeserializeGraph(s.Resolver, w)
			if err != nil {
				return false, err
			}
			txt2, err := graph.CanonicalString(g1)
			if err != nil {
				return false, err
			}
			hash2 := graph.Hash(g2)
			if txt1 != txt2 {
				return false, fmt.Errorf("canonical strings do not match")
			}
			if !bytes.Equal(hash1, hash2) {
				return false, fmt.Errorf("hashes do not match")
			}

			// TODO actually compare the contents of the graphs, not just their hashes
			return true, nil
		}, gengraph.Graph(s.Resolver)))

	properties.TestingRun(s.T())
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
