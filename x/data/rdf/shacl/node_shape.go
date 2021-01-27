package shacl

import "github.com/regen-network/regen-ledger/x/data/rdf"

type NodeShape struct {
	Target         NodeShapeTarget
	PropertyShapes []PropertyShape
}

type NodeShapeTarget interface {
	Cost() uint64
	Resolve(graph rdf.IndexedGraph) rdf.ObjectAccessor
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

func (ns NodeShape) Validate(graph rdf.IndexedGraph) {
	targetIterator := ns.Target.Resolve(graph).Iterator()
	for targetIterator.Next() {
		if iri, ok := targetIterator.(rdf.IRIOrBNode); ok {
			for _, ps := range ns.PropertyShapes {
				ps.Validate(nil, graph, iri)
			}
		}
	}
}
