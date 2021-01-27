package rdf

type Term interface {
	Equals(ctx Context, other Term) bool
}

type Context interface{}

type ValidationContext interface {
	Context
}

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

type IndexedGraph interface {
	HasTriple(Triple) bool
	Iterator() GraphIterator
	BySubject(subject Node) PredicateObjectAccessor
	//ByPredicate(predicate Node) SubjectObectAccessor
}

type Countable interface {
	Count() int
	CountGTE(int) bool
	CountLTE(int) bool
}

type GraphIterator interface {
	Next() bool
	Subject() Node
	Properties() PredicateObjectAccessor
}

type PredicateObjectAccessor interface {
	Countable

	ByPredicate(predicate Node) ObjectAccessor
	Iterator() PredicateObjectIterator
}

type PredicateObjectIterator interface {
	Next() bool
	Predicate() Node
	Object() ObjectAccessor
}

type ObjectAccessor interface {
	Countable

	HasValue(Term) bool
	Iterator() ObjectIterator
}

type ObjectIterator interface {
	Next() bool
	Object() Term
}

//type SubjectObectAccessor interface {
//	ObjectAccesor(subject Node) ObjectAccessor
//	Iterator() SubjectObectIterator
//}
//
//type SubjectObectIterator interface {
//	Next() bool
//	Subject() Node
//	Object() Term
//}

type GraphBuilder interface {
	IndexedGraph

	AddTriple(triple Triple)
	RemoveTriple(triple Triple)
	Merge(graph IndexedGraph)
	NewBNode() BNode
}
