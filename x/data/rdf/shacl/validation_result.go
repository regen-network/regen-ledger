package shacl

import (
	"github.com/regen-network/regen-ledger/x/data/rdf"
	"github.com/regen-network/regen-ledger/x/data/rdf/memory"
)

func ValidationResult(value rdf.Term) rdf.NodeBuilder {
	return rdf.NewNodeBuilder(memory.NewGraph(),
		map[rdf.IRIOrBNode][]rdf.Term{
			rdf.RDFType: {ShValidationResult},
			ShValue:     {value},
		},
	)
}
