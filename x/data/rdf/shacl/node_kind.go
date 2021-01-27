package shacl

import (
	"fmt"

	"github.com/regen-network/regen-ledger/x/data/rdf"
)

type NodeKindConstraintComponent struct{}

func (c NodeKindConstraintComponent) IRI() rdf.IRI {
	return ShNodeKindConstraintComponent
}

var _ ConstraintComponent = NodeKindConstraintComponent{}

func (c NodeKindConstraintComponent) Parse(_ rdf.Context, graph rdf.Graph, target rdf.Node) ([]Constraint, error) {
	it := graph.FindBySubjectPredicate(target, ShNodeKind)
	defer it.Close()

	if it.Next() {
		nodeKind := it.Object()
		if it.Next() {
			return nil, fmt.Errorf("only 1 %s nodes expected, got at least 2", ShNodeKind)
		}

		switch nodeKind {
		case ShIRI:
			return []Constraint{iriConstraint{}}, nil
		default:
			panic("TODO")
		}
	}
	return nil, nil
}

type iriConstraint struct{}

func (c iriConstraint) Cost() uint64 {
	return SimpleOperationCost
}

func (c iriConstraint) Validate(_ ValidationContext, _ rdf.Graph, value rdf.Term) error {
	if _, ok := value.(rdf.IRI); !ok {
		return fmt.Errorf("expected IRI, got %+v", value)
	}
	return nil
}
