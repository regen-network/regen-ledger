package rdf

import "fmt"

func GetOneObject(accessor ObjectAccessor) (Term, error) {
	it := accessor.Iterator()
	if !it.Next() {
		return nil, nil
	}

	val := it.Object()

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
