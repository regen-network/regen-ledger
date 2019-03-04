package postgresql

import "gitlab.com/regen-network/regen-ledger/index"

type Indexer interface {
	index.Indexer

	Exec(query string, args ...interface{})
}
