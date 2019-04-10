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

type Keeper struct {
	storeKey   sdk.StoreKey
	dataKeeper data.Keeper
	cdc        *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, dataKeeper data.Keeper, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, dataKeeper: dataKeeper, cdc: cdc}
}

func KeySignatures(content types.DataAddress) []byte {
	return []byte(fmt.Sprintf("%x/sigs", content))
}

func KeySignatureEvidence(content types.DataAddress, signer sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("%x/sig/%x", content, signer))
}

func (keeper Keeper) SignClaim(ctx sdk.Context, claim types.DataAddress, evidence []types.DataAddress, signers []sdk.AccAddress) sdk.Error {
	if !keeper.dataKeeper.HasData(ctx, claim) {
		return sdk.ErrUnknownRequest("Content data can't be found")
	}
	if !types.IsGraphDataAddress(claim) {
		return sdk.ErrUnknownRequest("Content isn't a graph")
	}

	store := ctx.KVStore(keeper.storeKey)
	sigsToStore := signers
	existingSigs := keeper.GetSignatures(ctx, claim)
	if len(existingSigs) != 0 {
		sigsToStore = append(sigsToStore, existingSigs...)
		unique.Slice(&sigsToStore, func(i, j int) bool {
			return bytes.Compare(sigsToStore[i], sigsToStore[j]) < 0
		})
	}
	store.Set(KeySignatures(claim), keeper.cdc.MustMarshalBinaryBare(sigsToStore))

	for _, signer := range signers {
		evidenceToStore := evidence
		existingEvidence := keeper.GetEvidence(ctx, claim, signer)
		if len(existingEvidence) != 0 {
			evidenceToStore = append(evidenceToStore, existingEvidence...)
			unique.Slice(&sigsToStore, func(i, j int) bool {
				return bytes.Compare(evidenceToStore[i], evidenceToStore[j]) < 0
			})
		}
		store.Set(KeySignatureEvidence(claim, signer), keeper.cdc.MustMarshalBinaryBare(evidenceToStore))
	}

	return nil
}

func (keeper Keeper) GetSignatures(ctx sdk.Context, content types.DataAddress) []sdk.AccAddress {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeySignatures(content))

	if bz == nil {
		return nil
	}

	var sigs []sdk.AccAddress
	keeper.cdc.MustUnmarshalBinaryBare(bz, &sigs)
	return sigs
}

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
