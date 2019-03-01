package postgresql

import (
	"database/sql"
)

type Indexer struct {
	conn  *sql.DB
	curTx *sql.Tx
}
