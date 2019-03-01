package postgresql

import (
	"database/sql"
	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/lib/pq"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewIndexer(connString string) (*Indexer, error) {
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &Indexer{conn: conn}, nil
}

func (indexer *Indexer) Index(query string, args ...interface{}) {
	tx := indexer.curTx
	if tx == nil {
		panic("not in a transaction")
	}
	_, err := tx.Exec(query, args...)
	if err != nil {
		panic(err)
	}
}

func (indexer *Indexer) OnInitChain(abci.RequestInitChain, abci.ResponseInitChain) {
	_, err := indexer.conn.Exec(initSchema)
	if err != nil {
		panic(err)
	}
}

func (indexer *Indexer) OnBeginBlock(req abci.RequestBeginBlock, res abci.ResponseBeginBlock) {
	tx, err := indexer.conn.Begin()
	if err != nil {
		panic(err)
	}
	indexer.curTx = tx
	indexer.Index("INSERT INTO block (height, time, hash) VALUES ($1, $2, $3)", req.Header.Height, req.Header.Time, req.Hash)
}

func (indexer *Indexer) BeforeDeliverTx(txBytes []byte) {
	// TODO avoid decoding Tx here because it has already been done elsewhere
	indexer.Index("SAVEPOINT msg_state")
}

func (indexer *Indexer) AfterDeliverTx(res abci.ResponseDeliverTx) {
	if res.Code == uint32(sdk.CodeOK) {
		indexer.Index("RELEASE SAVEPOINT msg_state")
	} else {
		indexer.Index("ROLLBACK TO SAVEPOINT msg_state")
	}
}

func (indexer *Indexer) OnEndBlock(abci.RequestEndBlock, abci.ResponseEndBlock) {
}

func (indexer *Indexer) OnCommit(abci.ResponseCommit) {
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
  data bytea
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
