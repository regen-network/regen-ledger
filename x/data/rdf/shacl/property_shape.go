package shacl

import "github.com/regen-network/regen-ledger/x/data/rdf"

type PropertyShape struct {
	Path        PropertyPath
	Constraints []constraintInstance
}

type constraintInstance struct {
	componentIRI rdf.IRI
	Constraint   Constraint
}

func (ps PropertyShape) Cost() uint64 {
	cost := ps.Path.Cost()
	for _, constraint := range ps.Constraints {
		cost += constraint.Constraint.Cost()
	}
	return cost
}

func (ps PropertyShape) Validate(ctx rdf.ValidationContext, graph rdf.IndexedGraph, target rdf.Term) {
}
