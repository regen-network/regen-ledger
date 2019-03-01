package index

import (
	abci "github.com/tendermint/tendermint/abci/types"
)

type Indexer interface {
	OnInitChain(abci.RequestInitChain, abci.ResponseInitChain)
	// TODO OnUpgrade
	OnBeginBlock(abci.RequestBeginBlock, abci.ResponseBeginBlock)
	BeforeDeliverTx(tx []byte)
	// index some data in the index where operation is an insert/update
	// operation in the index's underlying data language, that may potentially
	// be parameterized by args
	Index(operation string, args ...interface{})
	AfterDeliverTx(abci.ResponseDeliverTx)
	OnEndBlock(abci.RequestEndBlock, abci.ResponseEndBlock)
	OnCommit(abci.ResponseCommit)
}
