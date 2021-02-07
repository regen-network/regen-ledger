package rdf

type TripleIterator interface {
	NextTriple() (Quad, error)
}

type QuadIterator interface {
	NextQuad() (Quad, error)
}
