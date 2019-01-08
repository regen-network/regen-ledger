package esp

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/regen-network/regen-ledger/x/agent"
	"gitlab.com/regen-network/regen-ledger/x/proposal"

	//"github.com/twpayne/go-geom/encoding/ewkb"
)

type Keeper struct {
	storeKey sdk.StoreKey

	agentKeeper agent.Keeper

	cdc *codec.Codec
}

func (keeper Keeper) CanHandle(action proposal.ProposalAction) bool {
	switch action.(type) {
	case ActionRegisterESPVersion:
		return true
	case ActionReportESPResult:
		return true
	default:
		return false
	}
}

func (keeper Keeper) Handle(ctx sdk.Context, action proposal.ProposalAction, approvers []sdk.AccAddress) sdk.Result {
	switch action := action.(type) {
	case ActionRegisterESPVersion:
		return keeper.RegisterESPVersion(ctx, action.Curator, action.Name, action.Version, action.Spec, approvers)
	case ActionReportESPResult:
		return keeper.ReportESPResult(ctx, action.Curator, action.Name, action.Version, action.Verifier, action.Result, approvers)
	default:
		errMsg := fmt.Sprintf("Unrecognized action type: %v", action.Type())
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
}

func NewKeeper(
	storeKey sdk.StoreKey,
	agentKeeper agent.Keeper,
	cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey:    storeKey,
		agentKeeper: agentKeeper,
		cdc:         cdc,
	}
}

func espKey(curator agent.AgentID, name string, version string) []byte {
	return []byte(fmt.Sprintf("esp:%d/%s/%s", curator, name, version))
}

func (keeper Keeper) RegisterESPVersion(ctx sdk.Context, curator agent.AgentID, name string, version string, spec ESPVersionSpec, signers []sdk.AccAddress) sdk.Result {
	// TODO consume gas

	key := espKey(curator, name, version)
	store := ctx.KVStore(keeper.storeKey)
	if store.Has(key) {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
		}
	}

	if !keeper.agentKeeper.Authorize(ctx, curator, signers) {
		return sdk.Result{
			Code: sdk.CodeUnauthorized,
		}
	}

	bz, err := keeper.cdc.MarshalBinaryBare(spec)
	if err != nil {
		panic(err)
	}

	store.Set([]byte(key), bz)

	return sdk.Result{Code: sdk.CodeOK}
}

func (keeper Keeper) GetESPVersion(ctx sdk.Context, curator agent.AgentID, name string, version string) (spec ESPVersionSpec, err sdk.Error) {
	key := espKey(curator, name, version)
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get([]byte(key))
	marshalErr := keeper.cdc.UnmarshalBinaryBare(bz, &spec)
	if marshalErr != nil {
		return spec, sdk.ErrUnknownRequest(marshalErr.Error())
	}
	return spec, nil
}

func (keeper Keeper) ReportESPResult(ctx sdk.Context, curator agent.AgentID, name string, version string, verifier agent.AgentID, result ESPResult, signers []sdk.AccAddress) sdk.Result {
	// TODO consume gas
	spec, err := keeper.GetESPVersion(ctx, curator, name, version)

	if err != nil {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
		}
	}

	canVerify := false

	verifiers := spec.Verifiers

	n := len(verifiers)

	for i := 0; i < n; i++ {
		if verifier == verifiers[i] {
			canVerify = true
			break
		}
	}

	if !canVerify {
		return sdk.Result{
			Code: sdk.CodeUnauthorized,
		}
	}

	if !keeper.agentKeeper.Authorize(ctx, verifier, signers) {
		return sdk.Result{
			Code: sdk.CodeUnauthorized,
		}
	}

	// TODO verify geometry
	// TODO verify schema
	// TODO store result
	return sdk.Result{Code: sdk.CodeUnknownRequest, Log: "not implemented"}
}
