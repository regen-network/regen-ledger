package index

import (
	abci "github.com/tendermint/tendermint/abci/types"
)

type nilIndexer struct{}

func NewNilIndexer() Indexer {
	return &nilIndexer{}
}

func (nilIndexer) OnInitChain(abci.RequestInitChain, abci.ResponseInitChain) {}

func (nilIndexer) OnBeginBlock(abci.RequestBeginBlock, abci.ResponseBeginBlock) {}

func (nilIndexer) BeforeDeliverTx(tx []byte) {}

func (nilIndexer) Index(operation string, args ...interface{}) {}

func (nilIndexer) AfterDeliverTx(abci.ResponseDeliverTx) {}

func (nilIndexer) OnEndBlock(abci.RequestEndBlock, abci.ResponseEndBlock) {}

func (nilIndexer) OnCommit(abci.ResponseCommit) {}
