// Package ormtest contains utilities for testing modules built with the ORM.
package ormtest

import (
	"github.com/regen-network/regen-ledger/api/v2/orm/internal/testkv"
	"github.com/regen-network/regen-ledger/api/v2/orm/model/ormtable"
)

// NewMemoryBackend returns a new ORM memory backend which can be used for
// testing purposes independent of any storage layer.
//
// Example:
//
//	backend := ormtest.NewMemoryBackend()
//	ctx := ormtable.WrapContextDefault()
//	...
func NewMemoryBackend() ormtable.Backend {
	return testkv.NewSplitMemBackend()
}
