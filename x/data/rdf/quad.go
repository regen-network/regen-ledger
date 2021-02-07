package rdf

type Quad interface {
	Triple
	Graph() IRIOrBNode
}

func NewQuad(subject IRIOrBNode, predicate IRIOrBNode, object Term, graph IRIOrBNode) Quad {
	return &quad{subject: subject, predicate: predicate, object: object, graph: graph}
}

type quad struct {
	subject   IRIOrBNode
	predicate IRIOrBNode
	object    Term
	graph     IRIOrBNode
}

func (q quad) Subject() IRIOrBNode {
	return q.subject
}

func (q quad) Predicate() IRIOrBNode {
	return q.predicate
}

func (q quad) Object() Term {
	return q.object
}

func (q quad) Graph() IRIOrBNode {
	return q.graph
}
