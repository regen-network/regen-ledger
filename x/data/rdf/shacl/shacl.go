package shacl

import "github.com/regen-network/regen-ledger/x/data/rdf"

type Parser struct {
	constraintComponents map[rdf.IRI]ConstraintComponent
}

func NewParser() *Parser {
	return &Parser{constraintComponents: map[rdf.IRI]ConstraintComponent{}}
}

func (p *Parser) RegisterConstraintComponents(components ...ConstraintComponent) {
	for _, c := range components {
		p.constraintComponents[c.IRI()] = c
	}
}

func (p Parser) Parse(resolver ImportResolver, shapesGraph rdf.IndexedGraph) (ShapesGraph, error) {
	panic("TODO")
}

type ImportResolver interface {
	ResolveImport(rdf.IRI) rdf.IndexedGraph
}
