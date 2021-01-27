package shacl

import (
	"fmt"

	"github.com/regen-network/regen-ledger/x/data/rdf"
)

type ClassConstraintComponent struct{}

var _ ConstraintComponent = ClassConstraintComponent{}

func (c ClassConstraintComponent) IRI() rdf.IRI {
	return ShClassConstraintComponent
}

func (c ClassConstraintComponent) Parse(_ rdf.Context, graph rdf.Graph, target rdf.Node) ([]Constraint, error) {
	var res []Constraint

	it := graph.FindBySubjectPredicate(target, ShClass)
	defer it.Close()

	for it.Next() {
		obj := it.Object()
		iri, ok := obj.(rdf.IRI)
		if !ok {
			return nil, fmt.Errorf("expected an IRI, got %+v", obj)
		}
		res = append(res, classConstraint{class: iri})
	}
	return res, nil
}

type classConstraint struct {
	class rdf.IRI
}

func (c classConstraint) Cost() uint64 {
	panic("implement me")
}

func (c classConstraint) Validate(ctx ValidationContext, graph rdf.Graph, target rdf.Term) error {
	node, ok := target.(rdf.IRIOrBNode)
	if !ok {
		return fmt.Errorf("expected an IRI or blank node")
	}

	if graph.HasTriple(node, rdf.RDFType, c.class) {
		return nil
	}

	it := graph.FindBySubjectPredicate(node, rdf.RDFType)
	defer it.Close()

	for it.Next() {
		panic("TODO")
	}

	return nil
}
