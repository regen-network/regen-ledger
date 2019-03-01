package postgresql

import (
	"database/sql"
	"database/sql/driver"
	"github.com/lib/pq"
	abci "github.com/tendermint/tendermint/abci/types"
)

type TxWrapper interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type IndexerContext interface {
	Tx() TxWrapper
}

type Indexer struct {
	conn driver.Conn
}

func NewIndexer(connString string) (*Indexer, error) {
	conn, err := pq.Open(connString)
	if err != nil {
		return nil, err
	}
	return &Indexer{conn}, nil
}

func (indexer *Indexer) OnInitChain(abci.RequestInitChain, abci.ResponseInitChain) {
	panic("implement me")
}

func (indexer *Indexer) OnBeginBlock(abci.RequestBeginBlock, abci.ResponseBeginBlock) {
	panic("implement me")
}

func (indexer *Indexer) OnDeliverTx(tx []byte, res abci.ResponseDeliverTx) {
	panic("implement me")
}

func (indexer *Indexer) OnEndBlock(abci.RequestEndBlock, abci.ResponseEndBlock) {
	panic("implement me")
}

func (indexer *Indexer) OnCommit(abci.ResponseCommit) {
	panic("implement me")
}

func (indexer *Indexer) Context() IndexerContext {
	panic("TODO")
}
