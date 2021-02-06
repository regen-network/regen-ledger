package internal

import "github.com/regen-network/regen-ledger/x/data/rdf"

type QuadIterator struct {
	Chan chan QuadOrErr
}

type QuadOrErr struct {
	Quad rdf.Quad
	Err  error
}

func (q QuadIterator) Next() (rdf.Quad, error) {
	if q.Chan == nil {
		return nil, nil
	}

	quadOrErr, ok := <-q.Chan
	if !ok {
		return nil, nil
	}

	return quadOrErr.Quad, quadOrErr.Err
}

var _ rdf.QuadIterator = QuadIterator{}
