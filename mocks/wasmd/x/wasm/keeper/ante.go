package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CountTXDecorator struct{}

func (c CountTXDecorator) AnteHandle(sdk.Context, sdk.Tx, bool, sdk.AnteHandler) (sdk.Context, error) {
	return sdk.Context{}, nil
}

func NewCountTXDecorator(sdk.StoreKey) *CountTXDecorator {
	return &CountTXDecorator{}
}

type LimitSimulationGasDecorator struct{}

func (l LimitSimulationGasDecorator) AnteHandle(sdk.Context, sdk.Tx, bool, sdk.AnteHandler) (sdk.Context, error) {
	return sdk.Context{}, nil
}

func NewLimitSimulationGasDecorator(*sdk.Gas) *LimitSimulationGasDecorator {
	return &LimitSimulationGasDecorator{}
}
