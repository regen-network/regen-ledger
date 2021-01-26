package shacl

import (
	"fmt"

	"github.com/regen-network/regen-ledger/x/data/rdf"
)

type NodeKindConstraintParser struct{}

var _ ConstraintComponent = NodeKindConstraintParser{}

func (c NodeKindConstraintParser) Parse(_ rdf.Context, graph rdf.IndexedGraph, target rdf.Node) ([]ConstraintInstance, error) {
	it := graph.BySubject(target).ByPredicate(SHNodeKind).Iterator()
	if it.Next() {
		nodeKind := it.Object()
		if it.Next() {
			return nil, fmt.Errorf("only 1 %s nodes expected, got at least 2", SHNodeKind)
		}

		switch nodeKind {
		case ShIRI:
			return []ConstraintInstance{iriConstraint{}}, nil
		default:
			panic("TODO")
		}
	}
	return nil, nil
}

type iriConstraint struct{}

func (c iriConstraint) Validate(ctx rdf.ValidationContext, graph rdf.IndexedGraph, target rdf.Term) error {
	if _, ok := target.(rdf.IRI); !ok {
		return fmt.Errorf("expected IRI, got %+v", target)
	}
	return nil
}
