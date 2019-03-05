package postgresql

import (
	"database/sql"
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// Imports Postgres driver
	_ "github.com/lib/pq"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
)

type indexer struct {
	conn        *sql.DB
	txDecoder   sdk.TxDecoder
	migrations  []string
	curTx       *sql.Tx
	blockHeader abci.Header
}

func (indexer indexer) AddMigration(ddl string) {
	indexer.migrations = append(indexer.migrations, ddl)
}

// NewIndexer creates a PostgreSQL indexer that does default
// block and transaction indexing and can be used by keepers
// for custom indexing.
func NewIndexer(connString string, txDecoder sdk.TxDecoder) (Indexer, error) {
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &indexer{conn: conn, txDecoder: txDecoder}, nil
}

func (indexer *indexer) Exec(query string, args ...interface{}) {
	tx := indexer.curTx
	if tx == nil {
		panic("not in a transaction")
	}
	_, err := tx.Exec(query, args...)
	if err != nil {
		panic(err)
	}
}

func (indexer *indexer) OnInitChain(abci.RequestInitChain, abci.ResponseInitChain) {
	_, err := indexer.conn.Exec(InitialSchema)
	if err != nil {
		panic(err)
	}
	for _, ddl := range indexer.migrations {
		_, err := indexer.conn.Exec(ddl)
		if err != nil {
			panic(err)
		}
	}
}

func (indexer *indexer) OnBeginBlock(req abci.RequestBeginBlock, res abci.ResponseBeginBlock) {
	tx, err := indexer.conn.Begin()
	if err != nil {
		panic(err)
	}
	indexer.curTx = tx
	height := req.Header.Height
	indexer.Exec("INSERT INTO block (height, time, hash) VALUES ($1, $2, $3)", height, req.Header.Time, req.Hash)
	for _, tag := range res.Tags {
		indexer.Exec("INSERT INTO block_tags (block, key, value) VALUES ($1, $2, $3)",
			height, string(tag.Key), string(tag.Value))
	}
	indexer.blockHeader = req.Header
}

func (indexer *indexer) BeforeDeliverTx(txBytes []byte) {
	// TODO avoid decoding Tx here because it has already been done elsewhere
	hash := crypto.Sha256(txBytes)
	tx, err := indexer.txDecoder(txBytes)
	var jsonStr interface{}
	if err == nil {
		j, err := json.Marshal(tx)
		if err == nil {
			jsonStr = string(j)
		}
	}
	indexer.Exec("INSERT INTO tx (hash, block, bytes, tx_json) VALUES ($1, $2, $3, $4)",
		hash, indexer.blockHeader.Height, txBytes, jsonStr)
	indexer.Exec("SAVEPOINT msg_state")

}

func (indexer *indexer) AfterDeliverTx(txBytes []byte, res abci.ResponseDeliverTx) {
	if res.Code == uint32(sdk.CodeOK) {
		indexer.Exec("RELEASE SAVEPOINT msg_state")
	} else {
		indexer.Exec("ROLLBACK TO SAVEPOINT msg_state")
	}
	hash := crypto.Sha256(txBytes)
	indexer.Exec("UPDATE tx SET code = $1, result = $2 WHERE hash = $3",
		res.Code, res.Data, hash)
	for _, tag := range res.Tags {
		indexer.Exec("INSERT INTO tx_tags (tx, key, value) VALUES ($1, $2, $3)",
			hash, string(tag.Key), string(tag.Value))
	}
}

func (indexer *indexer) OnEndBlock(req abci.RequestEndBlock, res abci.ResponseEndBlock) {
	height := req.Height
	for _, tag := range res.Tags {
		indexer.Exec("INSERT INTO block_tags (block, key, value) VALUES ($1, $2, $3)",
			height, string(tag.Key), string(tag.Value))
	}
}

func (indexer *indexer) OnCommit(abci.ResponseCommit) {
	tx := indexer.curTx
	if tx == nil {
		panic("not in a transaction")
	}
	err := tx.Commit()
	if err != nil {
		panic(err)
	}
	indexer.curTx = nil
}
