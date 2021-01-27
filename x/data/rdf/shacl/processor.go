package shacl

import "github.com/regen-network/regen-ledger/x/data/rdf"

type ShapesGraph struct {
	nodeShapes []NodeShape
}

func (shg ShapesGraph) Validate(resolver ImportResolver, dataGraph rdf.Graph) rdf.Graph {
	for _, ns := range shg.nodeShapes {
		ns.Validate(dataGraph)
	}
	return nil
}
