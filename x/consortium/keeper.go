package consortium

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"gitlab.com/regen-network/regen-ledger/x/agent"
	"gitlab.com/regen-network/regen-ledger/x/proposal"
	"gitlab.com/regen-network/regen-ledger/x/upgrade"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           *codec.Codec
	agentKeeper   agent.Keeper
	upgradeKeeper upgrade.Keeper
}

const (
	consortiumAgentId = 0
)

var (
	keyValidators = []byte("validators")
)

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec, agentKeeper agent.Keeper, upgradeKeeper upgrade.Keeper) Keeper {
	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		agentKeeper:   agentKeeper,
		upgradeKeeper: upgradeKeeper,
	}
}

func (keeper Keeper) CanHandle(action proposal.ProposalAction) bool {
	switch action.(type) {
	case ActionScheduleUpgrade:
		return true
	default:
		return false
	}
}

func (keeper Keeper) Handle(ctx sdk.Context, action proposal.ProposalAction, voters []sdk.AccAddress) sdk.Result {
	switch action := action.(type) {
	case ActionScheduleUpgrade:
		return keeper.handleActionScheduleUpgrade(ctx, action, voters)
	default:
		errMsg := fmt.Sprintf("Unrecognized action type: %v", action.Type())
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
}

func (keeper Keeper) handleActionScheduleUpgrade(ctx sdk.Context, action ActionScheduleUpgrade, signers []sdk.AccAddress) sdk.Result {
	if !keeper.agentKeeper.Authorize(ctx, consortiumAgentId, signers) {
		return sdk.Result{Code: sdk.CodeUnauthorized}
	}
	keeper.upgradeKeeper.ScheduleUpgrade(ctx, action.upgradeInfo)
	return sdk.Result{Code: sdk.CodeOK}
}

func (keeper Keeper) SetValidators(ctx sdk.Context, validators []abci.ValidatorUpdate) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(validators)
	store.Set(keyValidators, bz)
}

func (keeper Keeper) EndBlocker(context sdk.Context) []abci.ValidatorUpdate {
	return abci.ValidatorUpdates{}
}

