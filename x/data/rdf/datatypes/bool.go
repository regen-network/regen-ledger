package datatypes

import (
	"fmt"

	"github.com/regen-network/regen-ledger/x/data/rdf"
)

type Bool struct{}

func (b Bool) IRI() rdf.IRI {
	return rdf.XSDBoolean
}

func (b Bool) CanonicalLexicalForm(value interface{}) (string, error) {
	bVal, ok := value.(bool)
	if !ok {
		return "", fmt.Errorf("got %T, expected %T", value, true)
	}

	if bVal {
		return "true", nil
	} else {
		return "false", nil
	}
}

func (b Bool) ValidateLexicalForm(s string) error {
	switch s {
	case "true":
	case "false":
	case "1":
	case "0":
	default:
		return fmt.Errorf("expected true, false, 1 or 0, got %s", s)
	}
	return nil
}

func (b Bool) ValidateValue(value interface{}) error {
	_, ok := value.(bool)
	if !ok {
		return fmt.Errorf("got %T, expected %T", value, true)
	}
	return nil
}

func (b Bool) Parse(s string) (interface{}, error) {
	switch s {
	case "true":
		return true, nil
	case "false":
		return false, nil
	case "1":
		return true, nil
	case "0":
		return false, nil
	default:
		return nil, fmt.Errorf("expected true, false, 1 or 0, got %s", s)
	}
}

var _ rdf.WellknownDatatype = Bool{}
