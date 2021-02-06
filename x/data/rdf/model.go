package rdf

type Term interface {
	Equal(other Term) bool
}

type Context interface{}

type IRI string

func (IRI) isTerm() {}

func (i IRI) String() string {
	return string(i)
}

func (IRI) isIRIOrBNode() {}

func (i IRI) Equal(other Term) bool {
	panic("implement me")
}

type BNode string

func (BNode) isIRIOrBNode() {}

func (b BNode) Equal(other Term) bool {
	panic("implement me")
}

var _, _ IRIOrBNode = (*IRI)(nil), (*BNode)(nil)

type IRIOrBNode interface {
	Term

	isIRIOrBNode()
}

type Node = IRIOrBNode

type Literal interface {
	Term

	LexicalForm() string
	Datatype() IRI
	Lang() string
}

type Triple interface {
	Subject() IRIOrBNode
	Predicate() IRIOrBNode
	Object() Term
}

type Quad interface {
	Subject() IRIOrBNode
	Predicate() IRIOrBNode
	Object() Term
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

type QuadIterator interface {
	Next() (Quad, error)
}
