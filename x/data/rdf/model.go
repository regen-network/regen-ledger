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
}

type Triple struct {
	Subject   IRIOrBNode
	Predicate IRIOrBNode
	Object    Term
}

type Quad struct {
	Subject   IRIOrBNode
	Predicate IRIOrBNode
	Object    Term
	Graph     IRIOrBNode
}
