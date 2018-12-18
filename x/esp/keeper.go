package esp

import (
	"bytes"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/regen-network/regen-ledger/x/agent"

	//"github.com/twpayne/go-geom/encoding/ewkb"
)

type Keeper struct {
	espStoreKey sdk.StoreKey
	espResultStoreKey sdk.StoreKey

	agentKeeper agent.Keeper

	cdc *codec.Codec
}

func NewKeeper(
	espStoreKey sdk.StoreKey,
	espResultStoreKey sdk.StoreKey,
	cdc *codec.Codec) Keeper {
	return Keeper{
		espStoreKey:espStoreKey,
		espResultStoreKey:espResultStoreKey,
		cdc:          cdc,
	}
}


func espKey(curator agent.AgentId, name string, version string) string {
	k, err := bech32.Encode("", curator)
	if err != nil {
		panic(err)
	}
	return k + "/" + name + "/" + version
}

func (keeper Keeper) RegisterESPVersion(ctx sdk.Context, curator agent.AgentId, name string, version string, spec ESPVersionSpec, signers []sdk.AccAddress)  sdk.CodeType {
	// TODO consume gas

	key := espKey(curator, name, version)
	store := ctx.KVStore(keeper.espStoreKey)
	if store.Has([]byte(key)) {
		return sdk.CodeUnknownRequest
	}

	if !keeper.agentKeeper.Authorize(ctx, curator, signers) {
		return sdk.CodeUnauthorized
	}

	bz, err := keeper.cdc.MarshalBinaryBare(spec)
	if err != nil {
		panic(err)
	}

	store.Set([]byte(key), bz)

	return sdk.CodeOK
}

func (keeper Keeper) GetESPVersion(ctx sdk.Context, curator agent.AgentId, name string, version string) (spec ESPVersionSpec, err sdk.Error) {
	key := espKey(curator, name, version)
	store := ctx.KVStore(keeper.espStoreKey)
	bz := store.Get([]byte(key))
	marshalErr := keeper.cdc.UnmarshalBinaryBare(bz, &spec)
	if marshalErr != nil {
		return spec, sdk.ErrUnknownRequest(marshalErr.Error())
	}
	return spec, nil
}

func (keeper Keeper) ReportESPResult(ctx sdk.Context, curator agent.AgentId, name string, version string, verifier agent.AgentId, result ESPResult, signers []sdk.AccAddress)  sdk.CodeType {
	// TODO consume gas
	spec, err := keeper.GetESPVersion(ctx, curator, name, version)

	if err != nil {
		return sdk.CodeUnknownRequest
	}

	canVerify := false

	verifiers := spec.Verifiers

	n := len(verifiers)

	for i := 0; i < n; i++ {
		if bytes.Compare(verifier, verifiers[i]) == 0 {
			canVerify = true
			break
		}
	}

	if !canVerify {
		return sdk.CodeUnauthorized
	}

	if !keeper.agentKeeper.Authorize(ctx, verifier, signers) {
		return sdk.CodeUnauthorized
	}

	// TODO verify geometry
	// TODO verify schema

	return sdk.CodeOK
}
