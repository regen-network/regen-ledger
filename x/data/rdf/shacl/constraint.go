package shacl

import "github.com/regen-network/regen-ledger/x/data/rdf"

type ConstraintComponent interface {
	IRI() rdf.IRI
	Parse(ctx rdf.Context, graph rdf.IndexedGraph, target rdf.Node) ([]ConstraintInstance, error)
}

type ConstraintInstance interface {
	Validate(ctx rdf.ValidationContext, graph rdf.IndexedGraph, target rdf.Term) error
}
