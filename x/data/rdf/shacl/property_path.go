package shacl

import "github.com/regen-network/regen-ledger/x/data/rdf"

type PropertyPath interface {
	Cost() uint64
	Traverse(graph rdf.IndexedGraph, target rdf.IRIOrBNode) rdf.ObjectAccessor
}

type PredicatePath struct {
	IRI rdf.IRI
}

func (p PredicatePath) Cost() uint64 {
	return ReadOperationCost
}

func (p PredicatePath) Traverse(graph rdf.IndexedGraph, target rdf.IRIOrBNode) rdf.ObjectAccessor {
	return graph.BySubject(target).ByPredicate(p.IRI)
}

var _ PropertyPath = PredicatePath{}

type SequencePath struct {
	IRIs []rdf.IRI
}

func (p SequencePath) Cost() uint64 {
	return ReadOperationCost * uint64(len(p.IRIs))
}

func (p SequencePath) Traverse(graph rdf.IndexedGraph, target rdf.IRIOrBNode) rdf.ObjectAccessor {
	props := graph.BySubject(target)
	for _, iri := range p.IRIs {
		_ = props.ByPredicate(iri)
	}
	panic("TODO")
}

var _ PropertyPath = SequencePath{}
