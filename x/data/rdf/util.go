package rdf

import "fmt"

func GetOneTerm(it TermIterator) (Term, error) {
	defer it.Close()

	if !it.Next() {
		return nil, nil
	}

	val := it.Term()

	if it.Next() {
		return nil, fmt.Errorf("expected only one value")
	}

	return val, nil
}

type NodeBuilder struct {
	Builder GraphBuilder
	Node    IRIOrBNode
}

func NewNodeBuilder(builder GraphBuilder, props map[IRIOrBNode][]Term) NodeBuilder {
	bNode := builder.NewBNode()
	for p, objs := range props {
		for _, obj := range objs {
			builder.AddTriple(Triple{
				Subject:   bNode,
				Predicate: p,
				Object:    obj,
			})
		}
	}
	return NodeBuilder{
		Builder: builder,
		Node:    bNode,
	}
}

func (builder *NodeBuilder) AddProps(props map[IRIOrBNode][]Term) {
	for p, objs := range props {
		for _, obj := range objs {
			builder.Builder.AddTriple(Triple{
				Subject:   builder.Node,
				Predicate: p,
				Object:    obj,
			})
		}
	}
}

func SubjectIterator(iterator TripleIterator) TermIterator {
	return subIterator{iterator}
}

type subIterator struct{ TripleIterator }

func (s subIterator) Term() Term { return s.Subject() }

func PredicateIterator(iterator TripleIterator) TermIterator {
	return predIterator{iterator}
}

type predIterator struct{ TripleIterator }

func (s predIterator) Term() Term { return s.Predicate() }

func ObjectIterator(iterator TripleIterator) TermIterator {
	return objIterator{iterator}
}

type objIterator struct{ TripleIterator }

func (s objIterator) Term() Term { return s.Object() }

func Filter(iterator TripleIterator, filterFn func(sub IRIOrBNode, pred IRIOrBNode, obj Term) bool) TripleIterator {
	return filterIterator{TripleIterator: iterator, filterFn: filterFn}
}

type filterIterator struct {
	TripleIterator
	filterFn func(sub IRIOrBNode, pred IRIOrBNode, obj Term) bool
}

func (f filterIterator) Count() int {
	panic("implement me")
}

func (f filterIterator) CountGTE(i int) bool {
	panic("implement me")
}

func (f filterIterator) CountLTE(i int) bool {
	panic("implement me")
}

func (f filterIterator) Next() bool {
	for {
		if !f.TripleIterator.Next() {
			return false
		}

		if f.filterFn(f.Subject(), f.Predicate(), f.Object()) {
			return true
		}
	}
}
