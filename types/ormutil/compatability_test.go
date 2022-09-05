package ormutil

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"gotest.tools/v3/assert"
)

func TestGogoPageReqToPulsarPageReq(t *testing.T) {
	// nil page request should set default limit
	var pr *query.PageRequest
	pg, err := GogoPageReqToPulsarPageReq(pr)
	assert.NilError(t, err)
	assert.Equal(t, pg.Limit, uint64(query.DefaultLimit))

	// when limit is set, it shouldn't be overridden.
	pr = new(query.PageRequest)
	pr.Limit = 50
	pg, err = GogoPageReqToPulsarPageReq(pr)
	assert.NilError(t, err)
	assert.Equal(t, pr.Limit, pg.Limit)
}
