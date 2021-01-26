package rdf

import (
	"math/big"

	"github.com/cockroachdb/apd/v2"
)

type String struct {
	string
}

func (s String) Equals(ctx Context, other Term) bool {
	panic("implement me")
}

func (s String) LexicalForm() string {
	return s.string
}

func (s String) Datatype() IRI {
	panic("implement me")
}

var _ Literal = String{}

type Decimal struct {
	*apd.Decimal
}

type Integer struct {
	*big.Int
}

var _, _ Literal = (*Decimal)(nil), (*Integer)(nil)

func (d Decimal) Equals(ctx Context, other Term) bool {
	panic("implement me")
}

func (d Decimal) LexicalForm() string {
	panic("implement me")
}

func (d Decimal) Datatype() IRI {
	panic("implement me")
}

func (i Integer) Equals(ctx Context, other Term) bool {
	panic("implement me")
}

func (i Integer) LexicalForm() string {
	panic("implement me")
}

func (i Integer) Datatype() IRI {
	panic("implement me")
}
