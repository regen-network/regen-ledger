package shacl

import "github.com/regen-network/regen-ledger/x/data/rdf"

type NodeShape struct {
	Target         NodeShapeTarget
	PropertyShapes []PropertyShape
}

type NodeShapeTarget interface {
	Cost() uint64
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
