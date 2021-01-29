package data

import (
	"crypto"
	_ "crypto/sha256"
	"fmt"
	"hash"
	"strings"
	"testing"

	_ "golang.org/x/crypto/blake2b"
	_ "golang.org/x/crypto/blake2s"

	"github.com/lazyledger/smt"
	"github.com/piprate/json-gold/ld"
	"github.com/zeebo/blake3"
)

const testCase1 = `
<http://vocab.getty.edu/aat/300014078> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.cidoc-crm.org/cidoc-crm/E57_Material> .
<http://vocab.getty.edu/aat/300014078> <http://www.w3.org/2000/01/rdf-schema#label> "canvas" .
<http://vocab.getty.edu/aat/300014844> <http://www.cidoc-crm.org/cidoc-crm/P2_has_type> <http://vocab.getty.edu/aat/300241583> .
<http://vocab.getty.edu/aat/300014844> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.cidoc-crm.org/cidoc-crm/E55_Type> .
<http://vocab.getty.edu/aat/300014844> <http://www.w3.org/2000/01/rdf-schema#label> "Support" .
<http://vocab.getty.edu/aat/300015045> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.cidoc-crm.org/cidoc-crm/E57_Material> .
<http://vocab.getty.edu/aat/300015045> <http://www.w3.org/2000/01/rdf-schema#label> "watercolors" .
<http://vocab.getty.edu/aat/300033618> <http://www.cidoc-crm.org/cidoc-crm/P2_has_type> <http://vocab.getty.edu/aat/300435443> .
<http://vocab.getty.edu/aat/300033618> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.cidoc-crm.org/cidoc-crm/E55_Type> .
<http://vocab.getty.edu/aat/300033618> <http://www.w3.org/2000/01/rdf-schema#label> "Painting" .
<http://vocab.getty.edu/aat/300133025> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.cidoc-crm.org/cidoc-crm/E55_Type> .
<http://vocab.getty.edu/aat/300133025> <http://www.w3.org/2000/01/rdf-schema#label> "Artwork" .
<http://vocab.getty.edu/aat/300241583> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.cidoc-crm.org/cidoc-crm/E55_Type> .
<http://vocab.getty.edu/aat/300241583> <http://www.w3.org/2000/01/rdf-schema#label> "Part Type" .
<http://vocab.getty.edu/aat/300435443> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.cidoc-crm.org/cidoc-crm/E55_Type> .
<http://vocab.getty.edu/aat/300435443> <http://www.w3.org/2000/01/rdf-schema#label> "Type of Work" .
<https://linked.art/example/object/2> <http://www.cidoc-crm.org/cidoc-crm/P2_has_type> <http://vocab.getty.edu/aat/300033618> .
<https://linked.art/example/object/2> <http://www.cidoc-crm.org/cidoc-crm/P2_has_type> <http://vocab.getty.edu/aat/300133025> .
<https://linked.art/example/object/2> <http://www.cidoc-crm.org/cidoc-crm/P45_consists_of> <http://vocab.getty.edu/aat/300015045> .
<https://linked.art/example/object/2> <http://www.cidoc-crm.org/cidoc-crm/P46_is_composed_of> _:b0 .
<https://linked.art/example/object/2> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.cidoc-crm.org/cidoc-crm/E22_Human-Made_Object> .
<https://linked.art/example/object/2> <http://www.w3.org/2000/01/rdf-schema#label> "Example Painting" .
_:b0 <http://www.cidoc-crm.org/cidoc-crm/P2_has_type> <http://vocab.getty.edu/aat/300014844> .
_:b0 <http://www.cidoc-crm.org/cidoc-crm/P45_consists_of> <http://vocab.getty.edu/aat/300014078> .
_:b0 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.cidoc-crm.org/cidoc-crm/E22_Human-Made_Object> .
_:b0 <http://www.w3.org/2000/01/rdf-schema#label> "Canvas Support" .
`

const testCase2 = `
_:b0 <http://rdf.data-vocabulary.org/#ingredients> "1 cup ice cubes" .
_:b0 <http://rdf.data-vocabulary.org/#ingredients> "1 tablespoons white sugar" .
_:b0 <http://rdf.data-vocabulary.org/#ingredients> "1/2 cup club soda" .
_:b0 <http://rdf.data-vocabulary.org/#ingredients> "1/2 lime, juiced with pulp" .
_:b0 <http://rdf.data-vocabulary.org/#ingredients> "12 fresh mint leaves" .
_:b0 <http://rdf.data-vocabulary.org/#ingredients> "2 fluid ounces white rum" .
_:b0 <http://rdf.data-vocabulary.org/#instructions> _:b1 .
_:b0 <http://rdf.data-vocabulary.org/#instructions> _:b2 .
_:b0 <http://rdf.data-vocabulary.org/#instructions> _:b3 .
_:b0 <http://rdf.data-vocabulary.org/#instructions> _:b4 .
_:b0 <http://rdf.data-vocabulary.org/#instructions> _:b5 .
_:b0 <http://rdf.data-vocabulary.org/#name> "Mojito" .
_:b0 <http://rdf.data-vocabulary.org/#yield> "1 cocktail" .
_:b1 <http://rdf.data-vocabulary.org/#description> "Crush lime juice, mint and sugar together in glass." .
_:b1 <http://rdf.data-vocabulary.org/#step> "1"^^<http://www.w3.org/2001/XMLSchema#integer> .
_:b2 <http://rdf.data-vocabulary.org/#description> "Fill glass to top with ice cubes." .
_:b2 <http://rdf.data-vocabulary.org/#step> "2"^^<http://www.w3.org/2001/XMLSchema#integer> .
_:b3 <http://rdf.data-vocabulary.org/#description> "Pour white rum over ice." .
_:b3 <http://rdf.data-vocabulary.org/#step> "3"^^<http://www.w3.org/2001/XMLSchema#integer> .
_:b4 <http://rdf.data-vocabulary.org/#description> "Fill the rest of glass with club soda, stir." .
_:b4 <http://rdf.data-vocabulary.org/#step> "4"^^<http://www.w3.org/2001/XMLSchema#integer> .
_:b5 <http://rdf.data-vocabulary.org/#description> "Garnish with a lime wedge." .
_:b5 <http://rdf.data-vocabulary.org/#step> "5"^^<http://www.w3.org/2001/XMLSchema#integer> .
`

const testCase3 = `
_:b0 <http://schema.org/description> "The Empire State Building is a 102-story landmark in New York City." .
_:b0 <http://schema.org/geo> _:b1 .
_:b0 <http://schema.org/image> <http://www.civil.usherbrooke.ca/cours/gci215a/empire-state-building.jpg> .
_:b0 <http://schema.org/name> "The Empire State Building" .
_:b1 <http://schema.org/latitude> "40.75"^^<http://www.w3.org/2001/XMLSchema#float> .
_:b1 <http://schema.org/longitude> "73.98"^^<http://www.w3.org/2001/XMLSchema#float> .
`

func BenchmarkNormalize1(b *testing.B) {
	for i, tc := range []string{testCase1, testCase2, testCase3} {
		b.Run("Normalize", func(b *testing.B) {
			benchmarkNormalize(b, tc)
		})
		for _, h := range []crypto.Hash{crypto.SHA256, crypto.BLAKE2s_256, crypto.BLAKE2b_256} {
			b.Run(fmt.Sprintf("SMT %s %d", h.String(), i+1), func(b *testing.B) {
				benchmarkSMT(b, tc, func() hash.Hash { return h.New() })
			})
		}
		b.Run(fmt.Sprintf("SMT Blake3 %d", i), func(b *testing.B) {
			benchmarkSMT(b, tc, func() hash.Hash { return blake3.New() })
		})
	}
}

func benchmarkNormalize(b *testing.B, txt string) {
	b.StopTimer()
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	doc, err := proc.FromRDF(txt, options)
	if err != nil {
		panic(err)
	}
	res, err := proc.ToRDF(doc, options)
	dataset := res.(*ld.RDFDataset)

	algOpts := ld.NewJsonLdOptions("")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		alg := ld.NewNormalisationAlgorithm("URDNA2015")
		res, err = alg.Main(dataset, algOpts)
		if err != nil {
			panic(err)
		}
	}
	b.StopTimer()
}

var tru = []byte{1}

func benchmarkSMT(b *testing.B, txt string, newHash func() hash.Hash) {
	lines := strings.Split(txt, "\n")
	numLines := len(lines)
	for i := 0; i < b.N; i++ {
		store := smt.NewSimpleMap()
		tree := smt.NewSparseMerkleTree(store, newHash())
		for j := 0; j < numLines; j++ {
			_, err := tree.Update([]byte(lines[j]), tru)
			if err != nil {
				panic(err)
			}
		}
		_ = tree.Root()
	}
}
