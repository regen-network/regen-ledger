package datatypes

import (
	"fmt"

	"github.com/cockroachdb/apd/v2"
	"github.com/regen-network/regen-ledger/math"
	"github.com/regen-network/regen-ledger/x/data/rdf"
)

type Integer struct{}

func (d Integer) IRI() rdf.IRI {
	return rdf.XSDInteger
}

func (d Integer) CanonicalLexicalForm(value interface{}) (string, error) {
	dec, ok := value.(*apd.Decimal)

	if !ok {
		return "", fmt.Errorf("expected %T, got %T", &apd.Decimal{}, value)
	}

	return dec.Text('f'), nil
}

func (d Integer) Parse(lexicalForm string) (interface{}, error) {
	// TODO xsd:integer does not allow scientific exponents
	dec, _, err := apd.NewFromString(lexicalForm)
	if math.NumDecimalPlaces(dec) > 0 {
		return nil, fmt.Errorf("expected an integer, got %s", lexicalForm)
	}

	return dec, err
}

var _ rdf.WellknownDatatype = Integer{}
