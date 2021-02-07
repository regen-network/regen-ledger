package rdf

type BNode string

func (b BNode) isTerm() {}

func (BNode) isIRIOrBNode() {}

func (b BNode) Equal(other Term) bool {
	panic("implement me")
}
