package memory

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/leanovate/gopter/gen"

	"github.com/leanovate/gopter"
	"github.com/regen-network/regen-ledger/x/data/rdf"
)

const (
	ab rdf.FullIRI = "a:b"
	cd rdf.FullIRI = "c:d"
	ef rdf.FullIRI = "e:f"
)

func TestMemoryGraph(t *testing.T) {
	//properties := gopter.NewProperties(nil)
	//
	//properties.Property("graph has all triples", prop.ForAll(func(triples []rdf.Triple) (bool, error) {
	//	graph := NewGraph()
	//	for _, t := range triples {
	//		rdf.AddTriple(graph, t)
	//	}
	//
	//	for _, t := range triples {
	//		if !rdf.HasTriple(graph, t) {
	//			return false, fmt.Errorf("missing triple %+v", t)
	//		}
	//	}
	//
	//	for _, t := range triples {
	//		rdf.RemoveTriple(graph, t)
	//	}
	//
	//	for _, t := range triples {
	//		if rdf.HasTriple(graph, t) {
	//			return false, fmt.Errorf("has unexpected triple %+v", t)
	//		}
	//	}
	//
	//	return true, nil
	//}, gen.SliceOf(TripleGen)))
	//
	//properties.TestingRun(t)

	graph := NewGraph()
	graph.AddTriple(ab, cd, ef)
	graph.AddTriple(cd, ab, ef)
	graph.AddTriple(ef, cd, ab)
	graph.AddTriple(ab, ab, ab)
	require.True(t, graph.HasTriple(ab, cd, ef))
	it := graph.FindAll()
	for it.Next() {
		t.Logf("%s %s %s", it.Subject(), it.Predicate(), it.Object())
	}
}

var IRIGen = gopter.DeriveGen(
	func(x string) rdf.FullIRI {
		return rdf.FullIRI(x)
	},
	func(iri rdf.FullIRI) string {
		return string(iri)
	},
	gen.AnyString(),
)

var BNodeGen = gopter.DeriveGen(
	func(x string) rdf.BNode {
		return rdf.BNode(x)
	},
	func(b rdf.BNode) string {
		return string(b)
	},
	gen.AnyString(),
)

var IRIOrBNodeGen = gen.OneGenOf(IRIGen, BNodeGen)

var TermGen = gen.OneGenOf(IRIOrBNodeGen)

var TripleGen = gopter.DeriveGen(
	func(sub rdf.IRIOrBNode, pred rdf.IRIOrBNode, obj rdf.Term) rdf.Triple {
		return rdf.Triple{
			Subject:   sub,
			Predicate: pred,
			Object:    obj,
		}
	},
	func(t rdf.Triple) (rdf.IRIOrBNode, rdf.IRIOrBNode, rdf.Term) {
		return t.Subject, t.Predicate, t.Object
	},
	IRIOrBNodeGen, IRIOrBNodeGen, TermGen,
)
