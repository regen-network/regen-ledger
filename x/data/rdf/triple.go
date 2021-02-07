package rdf

type Triple interface {
	Subject() IRIOrBNode
	Predicate() IRIOrBNode
	Object() Term
}

func NewTriple(subject IRIOrBNode, predicate IRIOrBNode, object Term) Triple {
	return &quad{subject: subject, predicate: predicate, object: object}
}
