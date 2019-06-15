package action

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	cosmos "github.com/cosmos/cosmos-sdk/types"
)

type keeper struct {
	storeKey cosmos.StoreKey
	cdc      *codec.Codec
}

type capabilityGrant struct {
	// set if this is actor should be granted the capability at a "root" level
	root bool
	// all the actors that delegated this capability to the actor
	// the capability should be cleared if root is false and this array is cleared
	delegatedBy []cosmos.AccAddress

	// whenever this capability is undelegated or revoked, these delegations
	// need to be cleared recursively
	delegatedTo []cosmos.AccAddress
}

func NewKeeper(storeKey cosmos.StoreKey, cdc *codec.Codec) Keeper {
	return &keeper{storeKey: storeKey, cdc: cdc}
}

func ActorCapabilityKey(capability Capability, actor cosmos.AccAddress) []byte {
	return []byte(fmt.Sprintf("c/%s/%x", capability.CapabilityKey(), actor))
}

func (k keeper) GrantRootCapability(ctx cosmos.Context, actor cosmos.AccAddress, capability Capability) {
	store := ctx.KVStore(k.storeKey)
	store.Set(ActorCapabilityKey(capability, actor), k.cdc.MustMarshalBinaryBare(capabilityGrant{root: true}))
}

func (k keeper) RevokeRootCapability(ctx cosmos.Context, actor cosmos.AccAddress, capability Capability) {
	store := ctx.KVStore(k.storeKey)
	grant, found := k.getCapabilityGrant(ctx, actor, capability)
	// TODO can two actors be granted the same root capability? In this model yes
	if !found || !grant.root {
		return
	}
	// Handle the case where this capability was also granted from somwhere else
	if len(grant.delegatedBy) > 0 {
		grant.root = false
		store.Set(ActorCapabilityKey(capability, actor), k.cdc.MustMarshalBinaryBare(grant))
		return
	}
	for _, addr := range grant.delegatedTo {
		k.Undelegate(ctx, actor, addr, capability)
	}
	store.Delete(ActorCapabilityKey(capability, actor))
}

func (k keeper) getCapabilityGrant(ctx cosmos.Context, actor cosmos.AccAddress, capability Capability) (grant capabilityGrant, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(ActorCapabilityKey(capability, actor))
	if bz == nil {
		return grant, false
	}
	k.cdc.MustUnmarshalBinaryBare(bz, &grant)
	return grant, true
}

func (k keeper) Delegate(ctx cosmos.Context, grantor cosmos.AccAddress, grantee cosmos.AccAddress, capability Capability) bool {
	store := ctx.KVStore(k.storeKey)
	grantorGrant, found := k.getCapabilityGrant(ctx, grantor, capability)
	if !found {
		return false
	}
	grantorGrant.delegatedTo = append(grantorGrant.delegatedTo, grantee)
	store.Set(ActorCapabilityKey(capability, grantor), k.cdc.MustMarshalBinaryBare(grantorGrant))
	granteeGrant, _ := k.getCapabilityGrant(ctx, grantee, capability)
	granteeGrant.delegatedBy = append(granteeGrant.delegatedBy, grantor)
	store.Set(ActorCapabilityKey(capability, grantee), k.cdc.MustMarshalBinaryBare(granteeGrant))
	return true
}

func (k keeper) Undelegate(ctx cosmos.Context, grantor cosmos.AccAddress, grantee cosmos.AccAddress, capability Capability) {
	panic("implement me")
}

func (k keeper) HasCapability(ctx cosmos.Context, actor cosmos.AccAddress, capability Capability) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(ActorCapabilityKey(capability, actor))
}
