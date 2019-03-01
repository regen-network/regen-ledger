package index

import abci "github.com/tendermint/tendermint/abci/types"

type Indexer interface {
	OnInitChain(abci.RequestInitChain, abci.ResponseInitChain)
	// TODO OnUpgrade
	OnBeginBlock(abci.RequestBeginBlock, abci.ResponseBeginBlock)
	OnDeliverTx(tx []byte, res abci.ResponseDeliverTx)
	OnEndBlock(abci.RequestEndBlock, abci.ResponseEndBlock)
	OnCommit(abci.ResponseCommit)
}
