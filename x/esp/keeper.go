package esp

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/regen-network/regen-ledger/utils"
	"gitlab.com/regen-network/regen-ledger/x/agent"
	"gitlab.com/regen-network/regen-ledger/x/proposal"
	"gitlab.com/regen-network/regen-ledger/x/geo"
	"golang.org/x/crypto/blake2b"

	//"github.com/twpayne/go-geom/encoding/ewkb"
)

type Keeper struct {
	storeKey sdk.StoreKey

	agentKeeper agent.Keeper

	geoKeeper geo.Keeper

	cdc *codec.Codec
}

func (keeper Keeper) CheckProposal(ctx sdk.Context, action proposal.ProposalAction) (bool, sdk.Result) {
	switch action := action.(type) {
	case ActionRegisterESPVersion:
		return true, sdk.Result{
			Tags: sdk.EmptyTags().AppendTag("proposal.agent", []byte(agent.MustEncodeBech32AgentID(action.Curator))),
		}
	case ActionReportESPResult:
		return true, sdk.Result{
			Tags: sdk.EmptyTags().
				AppendTag("proposal.agent", []byte(agent.MustEncodeBech32AgentID(action.Verifier))).
				AppendTag("esp.id", []byte(espVersionId(action.Curator, action.Name, action.Version))),
		}
	default:
		return false, sdk.Result{Code: sdk.CodeUnknownRequest}
	}
}

func (keeper Keeper) HandleProposal(ctx sdk.Context, action proposal.ProposalAction, approvers []sdk.AccAddress) sdk.Result {
	switch action := action.(type) {
	case ActionRegisterESPVersion:
		return keeper.RegisterESPVersion(ctx, action.ESPVersionSpec, approvers)
	case ActionReportESPResult:
		return keeper.ReportESPResult(ctx, action.ESPResult, approvers)
	default:
		errMsg := fmt.Sprintf("Unrecognized action type: %v", action.Type())
		return sdk.ErrUnknownRequest(errMsg).Result()
	}
}

func NewKeeper(
	storeKey sdk.StoreKey,
	agentKeeper agent.Keeper,
	geoKeeper geo.Keeper,
	cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey:    storeKey,
		agentKeeper: agentKeeper,
		geoKeeper: geoKeeper,
		cdc:         cdc,
	}
}

func espVersionId(curator agent.AgentID, name string, version string) string {
	return fmt.Sprintf("esp:%s/%s/%s", agent.MustEncodeBech32AgentID(curator), name, version)
}

func (keeper Keeper) RegisterESPVersion(ctx sdk.Context, spec ESPVersionSpec, signers []sdk.AccAddress) sdk.Result {
	// TODO consume gas

	id := espVersionId(spec.Curator, spec.Name, spec.Version)
	store := ctx.KVStore(keeper.storeKey)
	if store.Has([]byte(id)) {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
		}
	}

	if !keeper.agentKeeper.Authorize(ctx, spec.Curator, signers) {
		return sdk.Result{
			Code: sdk.CodeUnauthorized,
		}
	}

	bz, err := keeper.cdc.MarshalBinaryBare(spec)
	if err != nil {
		panic(err)
	}

	store.Set([]byte(id), bz)

	return sdk.Result{
		Code: sdk.CodeOK,
		Tags: sdk.EmptyTags().AppendTag("esp.id", []byte(id)),
	}
}

func (keeper Keeper) GetESPVersion(ctx sdk.Context, curator agent.AgentID, name string, version string) (spec ESPVersionSpec, err sdk.Error) {
	key := espVersionId(curator, name, version)
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get([]byte(key))
	marshalErr := keeper.cdc.UnmarshalBinaryBare(bz, &spec)
	if marshalErr != nil {
		return spec, sdk.ErrUnknownRequest(marshalErr.Error())
	}
	return spec, nil
}

func espResultKey(resHash []byte) []byte {
	return []byte(fmt.Sprintf("result:%s", hex.EncodeToString(resHash)))
}

func (keeper Keeper) ReportESPResult(ctx sdk.Context, result ESPResult, signers []sdk.AccAddress) sdk.Result {
	// TODO consume gas
	spec, err := keeper.GetESPVersion(ctx, result.Curator, result.Name, result.Version)

	if err != nil {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
			Log:  "can't find ESP version",
		}
	}

	canVerify := false

	verifiers := spec.Verifiers

	n := len(verifiers)

	for i := 0; i < n; i++ {
		if result.Verifier == verifiers[i] {
			canVerify = true
			break
		}
	}

	if !canVerify {
		return sdk.Result{
			Code: sdk.CodeUnauthorized,
		}
	}

	if !keeper.agentKeeper.Authorize(ctx, result.Verifier, signers) {
		return sdk.Result{
			Code: sdk.CodeUnauthorized,
		}
	}

	// Verify geometry exists
<<<<<<< HEAD
	geoID := keeper.geoKeeper.GetGeometry(ctx, result.GeoID)

	if geoID == nil {
=======
	geo_id, err := keeper.geoKeeper.GetGeometry(ctx, result.GeoID)

	if err != nil {
>>>>>>> 857d8e3013a1eed402d38275ddda7509baec64ee
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
			Log:  "can't find geo",
		}
	}


	// TODO verify schema

	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryBare(result)
	hash, hashErr := blake2b.New256(nil)
	if hashErr != nil {
		panic(hashErr)
	}
	hash.Write(bz)
	hashBz := hash.Sum(nil)
	store.Set(espResultKey(hashBz), bz)

	return sdk.Result{
		Code: sdk.CodeOK,
		Tags: sdk.EmptyTags().
			AppendTag("esp.id", []byte(espVersionId(result.Curator, result.Name, result.Version))).
			AppendTag("esp.result", []byte(utils.MustEncodeBech32("espr", hashBz))),
	}
}
