// Package postgresql implements an Indexer for PostgreSQL that does
// default indexing of blocks and transactions and can be used in
// keepers for custom indexing
package postgresql

import "gitlab.com/regen-network/regen-ledger/index"

// Indexer is a PostgreSQL based Indexer
type Indexer interface {
	index.Indexer

	// Exec executes a PostgreSQL statement against the database.
	// Can be used in keepers for custom indexing
	Exec(query string, args ...interface{})
}
