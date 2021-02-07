package internal

import "github.com/regen-network/regen-ledger/x/data/rdf"

type ChanIterator struct {
	Chan chan QuadOrErr
}

type QuadOrErr struct {
	Quad rdf.Quad
	Err  error
}

func (q ChanIterator) NextTriple() (rdf.Triple, error) {
	return q.NextQuad()
}

func (q ChanIterator) NextQuad() (rdf.Quad, error) {
	if q.Chan == nil {
		return nil, nil
	}

	quadOrErr, ok := <-q.Chan
	if !ok {
		return nil, nil
	}

	return quadOrErr.Quad, quadOrErr.Err
}

var _ rdf.QuadIterator = ChanIterator{}
