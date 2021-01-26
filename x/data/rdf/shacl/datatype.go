package shacl

import (
	"fmt"

	"github.com/regen-network/regen-ledger/x/data/rdf"
)

type DatatypeConstraintComponent struct{}

var _ ConstraintComponent = DatatypeConstraintComponent{}

func (d DatatypeConstraintComponent) IRI() rdf.IRI {
	return ShDatatypeConstraintComponent
}

func (d DatatypeConstraintComponent) Parse(ctx rdf.Context, graph rdf.IndexedGraph, target rdf.Node) ([]ConstraintInstance, error) {
	val, err := rdf.GetOneObject(graph.BySubject(target).ByPredicate(ShDatatype))
	if err != nil {
		return nil, fmt.Errorf("expected only one value for %s", ShDatatype)
	}

	iri, ok := val.(rdf.IRI)
	if !ok {
		return nil, fmt.Errorf("expected an IRI for %s, got %s", ShDatatype, val)
	}

	return []ConstraintInstance{datatypeConstraint{datatype: iri}}, nil
}

type datatypeConstraint struct {
	datatype rdf.IRI
}

var _ ConstraintInstance = datatypeConstraint{}

func (d datatypeConstraint) Validate(ctx rdf.ValidationContext, graph rdf.IndexedGraph, value rdf.Term) error {
	literal, ok := value.(rdf.Literal)
	if !ok {
		return fmt.Errorf("expected a literal, got %")
	}

	if !literal.Datatype().Equals(ctx, d.datatype) {
		return fmt.Errorf("expected %s, got %s", d.datatype, literal.Datatype())
	}

	return nil
}
