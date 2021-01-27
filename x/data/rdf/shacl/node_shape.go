package shacl

import "github.com/regen-network/regen-ledger/x/data/rdf"

type NodeShape struct {
	Target         NodeShapeTarget
	PropertyShapes []PropertyShape
}

type NodeShapeTarget interface {
	Cost() uint64
	Resolve(graph rdf.Graph) rdf.TermIterator
}

type TargetNode struct {
	Node rdf.Term
}

func (ns NodeShape) Cost() uint64 {
	cost := ns.Target.Cost()
	for _, ps := range ns.PropertyShapes {
		cost += ps.Cost()
	}
	return cost
}

func (ns NodeShape) Validate(graph rdf.Graph) {
	it := ns.Target.Resolve(graph)
	defer it.Close()

	for it.Next() {
		if iri, ok := it.Term().(rdf.IRIOrBNode); ok {
			for _, ps := range ns.PropertyShapes {
				ps.Validate(nil, graph, iri)
			}
		}
	}
}
