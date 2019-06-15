package action

import cosmos "github.com/cosmos/cosmos-sdk/types"

type dispatcher struct {
	Keeper
}

func NewDispatcher(k Keeper) Dispatcher {
	return &dispatcher{k}
}

func (dispatcher dispatcher) DispatchAction(ctx cosmos.Context, actor cosmos.AccAddress, action Action) cosmos.Result {
	//caps := action.RequiredCapabilities()
	panic("TODO")
}
