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

func (c ClassConstraintComponent) Parse(_ rdf.Context, graph rdf.IndexedGraph, target rdf.Node) ([]Constraint, error) {
	acc := graph.BySubject(target).ByPredicate(ShClass)
	var res []Constraint
	it := acc.Iterator()
	for it.Next() {
		obj := acc.Iterator().Object()
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

func (c classConstraint) Validate(ctx rdf.ValidationContext, graph rdf.IndexedGraph, target rdf.Term) error {
	node, ok := target.(rdf.IRIOrBNode)
	if !ok {
		return fmt.Errorf("expected an IRI or blank node")
	}

	acc := graph.BySubject(node).ByPredicate(rdf.RDFType)
	if acc.HasValue(c.class) {
		return nil
	}
	it := acc.Iterator()
	for it.Next() {
		panic("TODO")
	}
	return nil
}
