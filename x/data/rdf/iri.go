package rdf

type IRI string

func (IRI) isTerm() {}

func (i IRI) String() string {
	return string(i)
}

func (IRI) isIRIOrBNode() {}

func (i IRI) Equal(other Term) bool {
	panic("implement me")
}
