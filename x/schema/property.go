package schema

import (
	"github.com/regen-network/regen-ledger/graph"
	"net/url"
)

// NewProperty wraps a PropertyDefinition as a Property
// TODO move this to the schema module
func NewProperty(propertyDefinition PropertyDefinition, uri *url.URL) graph.Property {
	return &property{PropertyDefinition: propertyDefinition, uri: uri}
}

type property struct {
	PropertyDefinition
	uri *url.URL
}

func (p *property) URI() *url.URL {
	return p.uri
}

func (p *property) Arity() graph.Arity {
	return p.PropertyDefinition.Arity
}

func (p *property) Type() graph.PropertyType {
	return p.PropertyDefinition.PropertyType
}
