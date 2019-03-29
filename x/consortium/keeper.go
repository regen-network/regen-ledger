package consortium

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/proposal"
	"github.com/regen-network/regen-ledger/x/upgrade"
	abci "github.com/tendermint/tendermint/abci/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           *codec.Codec
	agentKeeper   group.Keeper
	upgradeKeeper upgrade.Keeper
}

var (
	consortiumGroupId = group.GroupAddrFromUint64(0)
	keyValidators     = []byte("validators")
)

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec, agentKeeper group.Keeper, upgradeKeeper upgrade.Keeper) Keeper {
	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		agentKeeper:   agentKeeper,
		upgradeKeeper: upgradeKeeper,
	}
}

func (keeper Keeper) CheckProposal(ctx sdk.Context, action proposal.ProposalAction) (bool, sdk.Result) {
	switch action.(type) {
	case ActionScheduleUpgrade:
		return true, sdk.Result{Code: sdk.CodeOK}
	case ActionCancelUpgrade:
		return true, sdk.Result{Code: sdk.CodeOK}
	case ActionChangeValidatorSet:
		return true, sdk.Result{Code: sdk.CodeOK}
	default:
		return false, sdk.Result{Code: sdk.CodeUnknownRequest}
	}
}

func (keeper Keeper) HandleProposal(ctx sdk.Context, action proposal.ProposalAction, voters []sdk.AccAddress) sdk.Result {
	switch action := action.(type) {
	case ActionScheduleUpgrade:
		return keeper.handleActionScheduleUpgrade(ctx, action, voters)
	case ActionCancelUpgrade:
		return keeper.handleActionCancelUpgrade(ctx, action, voters)
	case ActionChangeValidatorSet:
		return keeper.handleActionChangeValidatorSet(ctx, action, voters)
	default:
		errMsg := fmt.Sprintf("Unrecognized action type: %v", action.Type())
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
}

func (keeper Keeper) handleActionScheduleUpgrade(ctx sdk.Context, action ActionScheduleUpgrade, signers []sdk.AccAddress) sdk.Result {
	if !keeper.agentKeeper.Authorize(ctx, consortiumGroupId, signers) {
		return sdk.Result{Code: sdk.CodeUnauthorized}
	}
	err := keeper.upgradeKeeper.ScheduleUpgrade(ctx, action.Plan)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{}
}

func (keeper Keeper) handleActionCancelUpgrade(ctx sdk.Context, _ ActionCancelUpgrade, signers []sdk.AccAddress) sdk.Result {
	if !keeper.agentKeeper.Authorize(ctx, consortiumGroupId, signers) {
		return sdk.Result{Code: sdk.CodeUnauthorized}
	}
	keeper.upgradeKeeper.ClearUpgradePlan(ctx)
	return sdk.Result{}
}

func (keeper Keeper) handleActionChangeValidatorSet(ctx sdk.Context, action ActionChangeValidatorSet, signers []sdk.AccAddress) sdk.Result {
	keeper.SetValidators(ctx, action.Validators)
	return sdk.Result{}
}

func (keeper Keeper) GetValidators(ctx sdk.Context) []abci.ValidatorUpdate {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(keyValidators)
	if bz == nil {
		panic("Validators not set")
	}
	var validators []abci.ValidatorUpdate
	keeper.cdc.MustUnmarshalBinaryBare(bz, &validators)
	return validators
}

func (keeper Keeper) SetValidators(ctx sdk.Context, validators []abci.ValidatorUpdate) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryBare(validators)
	store.Set(keyValidators, bz)
}

func (keeper Keeper) EndBlocker(ctx sdk.Context) []abci.ValidatorUpdate {
	return keeper.GetValidators(ctx)
}
