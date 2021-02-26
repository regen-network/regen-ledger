package datatypes

import "github.com/regen-network/regen-ledger/x/data/rdf"

type DateTime struct{}

func (d DateTime) IRI() rdf.IRI {
	return rdf.XSDDateTime
}

func (d DateTime) CanonicalLexicalForm(value interface{}) (string, error) {
	panic("implement me")
}

func (d DateTime) Parse(lexicalForm string) (interface{}, error) {
	panic("implement me")
}

var _ rdf.WellknownDatatype = DateTime{}
