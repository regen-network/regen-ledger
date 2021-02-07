package rdf

type Term interface {
	isTerm()

	Equal(other Term) bool
}

type IRIOrBNode interface {
	Term

	isIRIOrBNode()
}

var _, _ IRIOrBNode = (*IRI)(nil), (*BNode)(nil)
