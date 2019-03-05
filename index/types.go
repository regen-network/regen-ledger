// Package index provides an interface for indexing Tendermint blockchains
// in external databases by intercepting ABCI method handlers
package index

import (
	abci "github.com/tendermint/tendermint/abci/types"
)

// Indexer defines a type that performs custom indexing of a Tendermint
// blockchain where interceptor methods have been defined that forward
// the specified parameters to the methods below
type Indexer interface {
	// OnInitChain intercepts the InitChain ABCI method
	OnInitChain(abci.RequestInitChain, abci.ResponseInitChain)
	// OnBeginBlock intercepts the BeginBlock ABCI method
	OnBeginBlock(abci.RequestBeginBlock, abci.ResponseBeginBlock)
	// BeforeDeliverTx should be called on the ABCI DeliverTx method
	// before the actual handler for DeliverTx is called
	BeforeDeliverTx(txBytes []byte)
	// AfterDeliverTx should be called after ABCI DeliverTx is
	// handled
	AfterDeliverTx(txBytes []byte, res abci.ResponseDeliverTx)
	// OnEndBlock intercepts the EndBlock ABCI method
	OnEndBlock(abci.RequestEndBlock, abci.ResponseEndBlock)
	// OnCommit intercepts the Commit ABCI method
	OnCommit(abci.ResponseCommit)
	// TODO OnUpgrade
}
