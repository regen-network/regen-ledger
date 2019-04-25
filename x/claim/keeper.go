package claim

import (
	"bytes"
	"fmt"
	"github.com/campoy/unique"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
)

// Keeper is that claim module keeper
type Keeper struct {
	storeKey   sdk.StoreKey
	dataKeeper data.Keeper
	cdc        *codec.Codec
}

// NewKeeper creates a new claim Keeper
func NewKeeper(storeKey sdk.StoreKey, dataKeeper data.Keeper, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, dataKeeper: dataKeeper, cdc: cdc}
}

// KeySignatures returns a story key pointing to the signatures for the content
func KeySignatures(content types.DataAddress) []byte {
	return []byte(fmt.Sprintf("%x/sigs", content))
}

// KeySignatureEvidence returns a story key pointing to the evidence for a given content signature
func KeySignatureEvidence(content types.DataAddress, signer sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("%x/sig/%x", content, signer))
}

// SignClaim asserts that the signers sign the content with the provided evidence
func (keeper Keeper) SignClaim(ctx sdk.Context, content types.DataAddress, evidence []types.DataAddress, signers []sdk.AccAddress) sdk.Error {
	if !types.IsGraphDataAddress(content) {
		return sdk.ErrUnknownRequest("Content isn't a graph")
	}
	if !keeper.dataKeeper.HasData(ctx, content) {
		return sdk.ErrUnknownRequest("Content data can't be found")
	}

	store := ctx.KVStore(keeper.storeKey)
	sigsToStore := signers
	existingSigs := keeper.GetSigners(ctx, content)
	if len(existingSigs) != 0 {
		sigsToStore = append(sigsToStore, existingSigs...)
		unique.Slice(&sigsToStore, func(i, j int) bool {
			return bytes.Compare(sigsToStore[i], sigsToStore[j]) < 0
		})
	}
	store.Set(KeySignatures(content), keeper.cdc.MustMarshalBinaryBare(sigsToStore))

	for _, signer := range signers {
		evidenceToStore := evidence
		existingEvidence := keeper.GetEvidence(ctx, content, signer)
		if len(existingEvidence) != 0 {
			evidenceToStore = append(evidenceToStore, existingEvidence...)
			unique.Slice(&sigsToStore, func(i, j int) bool {
				return bytes.Compare(evidenceToStore[i], evidenceToStore[j]) < 0
			})
		}
		store.Set(KeySignatureEvidence(content, signer), keeper.cdc.MustMarshalBinaryBare(evidenceToStore))
	}

	return nil
}

// GetSigners gets the signers for the given content
func (keeper Keeper) GetSigners(ctx sdk.Context, content types.DataAddress) []sdk.AccAddress {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeySignatures(content))

	if bz == nil {
		return nil
	}

	var sigs []sdk.AccAddress
	keeper.cdc.MustUnmarshalBinaryBare(bz, &sigs)
	return sigs
}

// GetEvidence returns the evidence provided by the given signer for the given content
func (keeper Keeper) GetEvidence(ctx sdk.Context, content types.DataAddress, signer sdk.AccAddress) []types.DataAddress {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeySignatureEvidence(content, signer))

	if bz == nil {
		return nil
	}

	var evidence []types.DataAddress
	keeper.cdc.MustUnmarshalBinaryBare(bz, &evidence)
	return evidence
}
