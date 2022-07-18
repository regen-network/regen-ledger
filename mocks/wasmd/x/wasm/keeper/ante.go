package keeper

// https://github.com/CosmWasm/wasmd/blob/v0.22.0/x/wasm/keeper/ante.go

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmWasm/wasmd/x/wasm/types"
)

type CountTXDecorator struct {
	storeKey sdk.StoreKey
}

func NewCountTXDecorator(storeKey sdk.StoreKey) *CountTXDecorator {
	return &CountTXDecorator{storeKey: storeKey}
}

func (a CountTXDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if simulate {
		return next(ctx, tx, simulate)
	}
	store := ctx.KVStore(a.storeKey)
	currentHeight := ctx.BlockHeight()

	var txCounter uint32
	if bz := store.Get(types.TXCounterPrefix); bz != nil {
		lastHeight, val := decodeHeightCounter(bz)
		if currentHeight == lastHeight {
			txCounter = val
		}
	}
	store.Set(types.TXCounterPrefix, encodeHeightCounter(currentHeight, txCounter+1))

	return next(types.WithTXCounter(ctx, txCounter), tx, simulate)
}

func encodeHeightCounter(height int64, counter uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, counter)
	return append(sdk.Uint64ToBigEndian(uint64(height)), b...)
}

func decodeHeightCounter(bz []byte) (int64, uint32) {
	return int64(sdk.BigEndianToUint64(bz[0:8])), binary.BigEndian.Uint32(bz[8:])
}

type LimitSimulationGasDecorator struct {
	gasLimit *sdk.Gas
}

func NewLimitSimulationGasDecorator(gasLimit *sdk.Gas) *LimitSimulationGasDecorator {
	return &LimitSimulationGasDecorator{gasLimit: gasLimit}
}

func (l LimitSimulationGasDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if !simulate {
		return next(ctx, tx, simulate)
	}

	if l.gasLimit != nil {
		return next(ctx.WithGasMeter(sdk.NewGasMeter(*l.gasLimit)), tx, simulate)
	}

	if maxGas := ctx.ConsensusParams().GetBlock().MaxGas; maxGas > 0 {
		return next(ctx.WithGasMeter(sdk.NewGasMeter(sdk.Gas(maxGas))), tx, simulate)
	}

	return next(ctx, tx, simulate)
}
