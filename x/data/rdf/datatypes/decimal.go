package datatypes

import (
	"fmt"

	"github.com/cockroachdb/apd/v2"
	"github.com/regen-network/regen-ledger/x/data/rdf"
)

type Decimal struct{}

func (d Decimal) IRI() rdf.IRI {
	return rdf.XSDDecimal
}

func (d Decimal) CanonicalLexicalForm(value interface{}) (string, error) {
	dec, ok := value.(*apd.Decimal)

	if !ok {
		return "", fmt.Errorf("expected %T, got %T", &apd.Decimal{}, value)
	}

	return dec.Text('f'), nil
}

func (d Decimal) Parse(lexicalForm string) (interface{}, error) {
	// TODO xsd:decimal does not allow scientific exponents
	dec, _, err := apd.NewFromString(lexicalForm)
	return dec, err
}

var _ rdf.WellknownDatatype = Decimal{}
