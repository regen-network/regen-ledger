package shacl

import "github.com/regen-network/regen-ledger/x/data/rdf"

type ConstraintComponent interface {
	IRI() rdf.IRI
	//MandatoryParameters() []rdf.IRI
	//OptionalParameters() []rdf.IRI
	Parse(ctx rdf.Context, graph rdf.Graph, target rdf.Node) ([]Constraint, error)
}

type Constraint interface {
	Cost() uint64
}

type SimpleConstraint interface {
	Constraint

	Validate(ctx ValidationContext, graph rdf.Graph, value rdf.Term) error
}

type MultiValueConstraint interface {
	Constraint

	ValidateMany(ctx ValidationContext, graph rdf.Graph, value rdf.TermIterator) error
}
