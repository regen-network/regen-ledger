package trig

import "github.com/regen-network/regen-ledger/x/data/rdf"

type state struct {
	baseURI      rdf.IRI
	namespaces   map[string]rdf.IRI
	bnodeLabels  map[string]rdf.BNode
	curSubject   rdf.Term
	curPredicate rdf.Term
	curGraph     rdf.Term
}

type trigDoc struct {
}
