package postgresql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/lib/pq"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
)

type indexer struct {
	conn        *sql.DB
	curTx       *sql.Tx
	txDecoder   sdk.TxDecoder
	blockHeader abci.Header
}

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
	fmt.Println(initSchema)
	_, err := indexer.conn.Exec(initSchema)
	if err != nil {
		panic(err)
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
	var jsonStr interface{} = nil
	if err == nil {
		j, err := json.Marshal(tx)
		if err != nil {
			jsonStr = string(j)
		}
	}
	indexer.Exec("INSERT INTO tx (hash, block, bytes, tx_json) VALUES ($1, $2, $3, $4)",
		hash, txBytes, jsonStr, indexer.blockHeader.Height)
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
		indexer.Exec("INSERT INTO tx_tags (block, key, value) VALUES ($1, $2, $3)",
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

var initSchema = `
CREATE TABLE block (
  height BIGINT NOT NULL PRIMARY KEY,
  time timestamptz NOT NULL,
  hash bytea
);

CREATE TABLE tx (
  hash bytea NOT NULL PRIMARY KEY,
  block BIGINT NOT NULL REFERENCES block,
  bytes bytea NOT NULL,
  tx_json jsonb NOT NULL,
  code int NOT NULL,
  result bytea
);

CREATE TABLE block_tags (
  block bigint NOT NULL REFERENCES block,
  key text not null,
  value text not null
);

CREATE TABLE tx_tags (
  tx bytea NOT NULL REFERENCES tx,
  key text not null,
  value text not null
);
`
