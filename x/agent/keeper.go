package agent

import (
	"bytes"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	agentStoreKey sdk.StoreKey

	cdc *codec.Codec
}

func NewKeeper(agentStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		agentStoreKey: agentStoreKey,
		cdc:           cdc,
	}
}

func (keeper Keeper) GetAgentInfo(ctx sdk.Context, id AgentId) (info AgentInfo, err sdk.Error) {
	store := ctx.KVStore(keeper.agentStoreKey)
	bz := store.Get(id)
	if bz == nil {
		return info, sdk.ErrUnknownRequest("Not found")
	}
	marshalErr := keeper.cdc.UnmarshalBinaryBare(bz, &info)
	if marshalErr != nil {
        return info, sdk.ErrUnknownRequest(marshalErr.Error())
	}
	return info, nil
}

func (keeper Keeper) CreateAgent(ctx sdk.Context, id AgentId, info AgentInfo) {
	store := ctx.KVStore(keeper.agentStoreKey)
	if store.Has(id) {
		panic("Agent ID already exists")
	}
	keeper.setAgentInfo(ctx, id, info)
}

func (keeper Keeper) setAgentInfo(ctx sdk.Context, id AgentId, info AgentInfo) {
	store := ctx.KVStore(keeper.agentStoreKey)
	bz, err := keeper.cdc.MarshalBinaryBare(info)
	if err != nil {
		panic(err)
	}
	store.Set(id, bz)
}

func (keeper Keeper) UpdateAgentInfo(ctx sdk.Context, id AgentId, signers []sdk.AccAddress, info AgentInfo) bool {
	if !keeper.Authorize(ctx, id, signers) {
		return false
	}
	keeper.setAgentInfo(ctx, id, info)
	return true
}

func (keeper Keeper) Authorize(ctx sdk.Context, id AgentId, signers []sdk.AccAddress) bool {
	info, err := keeper.getAgentInfo(ctx, id)
	if err != nil {
		return false
	}
	if info.AuthPolicy != MultiSig {
        panic("Unknown auth policy")
	}
	nAgents := len(info.Agents)
	nSigners := len(signers)
	for i := 0; i < nAgents; i++ {
		agentRef := info.Agents[i]
		switch agentRef.Type {
		case AgentRef_Agent:
			if !keeper.Authorize(ctx, agentRef.Agent, signers) {
				return false
			}
		case AgentRef_Address:
			addr := agentRef.Address
			foundSig := false
			// TODO Use a hash map to optimize this
			for j := 0; j < nSigners; j++ {
				if bytes.Compare(addr, signers[j]) == 0 {
                    foundSig = true
                    break
				}
			}
			if !foundSig {
				return false
			}
		default:
			panic("Unknown AgentRef type")
		}
	}
	return true
}
