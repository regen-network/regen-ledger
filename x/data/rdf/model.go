package rdf

type Term interface {
	Equals(ctx Context, other Term) bool
}

type Context interface{}

type FullIRI string

var _ IRI = FullIRI("")

func (I FullIRI) isTerm() {}

func (I FullIRI) String(Context) string {
	return string(I)
}

type IRI interface {
	Term
	IRIOrBNode

	isIRI()
	String(ctx Context) string
}

func (I FullIRI) isIRI()        {}
func (I FullIRI) isIRIOrBNode() {}

func (I FullIRI) Equals(ctx Context, other Term) bool {
	panic("implement me")
}

type BNode string

func (B BNode) isIRIOrBNode() {}

func (B BNode) Equals(ctx Context, other Term) bool {
	panic("implement me")
}

var _, _ IRIOrBNode = (*FullIRI)(nil), (*BNode)(nil)

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

type Graph interface {
	HasTriple(sub IRIOrBNode, pred IRIOrBNode, obj Term) bool
	FindBySubject(sub IRIOrBNode) TripleIterator
	FindByPredicate(pred IRIOrBNode) TripleIterator
	FindByObject(obj Term) TripleIterator
	FindBySubjectPredicate(sub IRIOrBNode, pred IRIOrBNode) TripleIterator
	FindBySubjectObject(sub IRIOrBNode, obj Term) TripleIterator
	FindByPredicateObject(pred IRIOrBNode, obj Term) TripleIterator
	FindAll() TripleIterator
}

type TripleIterator interface {
	Countable
	Next() bool
	Subject() IRIOrBNode
	Predicate() IRIOrBNode
	Object() Term
	Close()
}

type Countable interface {
	Count() int
	CountGTE(int) bool
	CountLTE(int) bool
}

type GraphBuilder interface {
	Graph

	AddTriple(sub IRIOrBNode, pred IRIOrBNode, obj Term)
	RemoveTriple(sub IRIOrBNode, pred IRIOrBNode, obj Term)
	Merge(graph Graph)
	NewBNode() BNode
}

type TermIterator interface {
	Countable
	Next() bool
	Term() Term
	Close()
}
