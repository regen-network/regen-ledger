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

func (d DatatypeConstraintComponent) Parse(ctx rdf.Context, graph rdf.Graph, target rdf.Node) ([]Constraint, error) {
	val, err := rdf.GetOneTerm(rdf.ObjectIterator(graph.FindBySubjectPredicate(target, ShDatatype)))
	if err != nil {
		return nil, fmt.Errorf("expected only one value for %s", ShDatatype)
	}

	iri, ok := val.(rdf.IRI)
	if !ok {
		return nil, fmt.Errorf("expected an IRI for %s, got %s", ShDatatype, val)
	}

	return []Constraint{datatypeConstraint{datatype: iri}}, nil
}

type datatypeConstraint struct {
	datatype rdf.IRI
}

var _ Constraint = datatypeConstraint{}

func (d datatypeConstraint) Cost() uint64 {
	return SimpleOperationCost
}

func (d datatypeConstraint) Validate(ctx ValidationContext, _ rdf.Graph, value rdf.Term) error {
	literal, ok := value.(rdf.Literal)
	if !ok {
		return fmt.Errorf("expected a literal, got %")
	}

	if !literal.Datatype().Equals(ctx, d.datatype) {
		return fmt.Errorf("expected %s, got %s", d.datatype, literal.Datatype())
	}

	return nil
}
