package rdf

type Literal interface {
	Term

	LexicalForm() string
	Datatype() IRI
	Lang() string
}

type WellknownDatatype interface {
	IRI() IRI
	CanonicalLexicalForm(value interface{}) (string, error)
	ValidateLexicalForm(lexicalForm string) error
	ValidateValue(value interface{}) error
	Parse(lexicalForm string) (interface{}, error)
}

type literal struct {
	lexicalForm string
	datatype    IRI
	lang        string
	value       interface{}
}

func MakeLiteral(lexicalForm string, datatype IRI) (Literal, error) {
	return &literal{lexicalForm: lexicalForm, datatype: datatype}, nil
}

func LangStringLiteral(value string, lang string) (Literal, error) {
	return &literal{lexicalForm: value, datatype: RDFLangString, lang: lang}, nil
}

func StringLiteral(value string) Literal {
	return &literal{lexicalForm: value, datatype: XSDString}
}

func BoolLiteral(value bool) Literal {
	return &literal{datatype: XSDString, value: value}
}

func (b literal) isTerm() {}

func (l literal) Equal(other Term) bool {
	otherLit, ok := other.(Literal)
	if !ok {
		return false
	}

	return l.datatype == otherLit.Datatype() &&
		l.lexicalForm == otherLit.LexicalForm() &&
		l.lang == otherLit.Lang()
}

func (l literal) LexicalForm() string {
	return l.lexicalForm
}

func (l literal) Datatype() IRI {
	return l.datatype
}

func (l literal) Lang() string {
	return l.lang
}
