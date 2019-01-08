package agent

import (
	"bytes"
	"fmt"
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

var (
	keyNewAgentID = []byte("newAgentID")
)

func keyAgentID(id AgentID) []byte {
	return []byte(fmt.Sprintf("#%d", id))
}

func (keeper Keeper) GetAgentInfo(ctx sdk.Context, id AgentID) (info AgentInfo, err sdk.Error) {
	store := ctx.KVStore(keeper.agentStoreKey)
	bz := store.Get(keyAgentID(id))
	if bz == nil {
		return info, sdk.ErrUnknownRequest("Not found")
	}
    info = AgentInfo{}
	marshalErr := keeper.cdc.UnmarshalBinaryBare(bz, &info)
	if marshalErr != nil {
		return info, sdk.ErrUnknownRequest(marshalErr.Error())
	}
	return info, nil
}

func (keeper Keeper) getNewAgentId(ctx sdk.Context) (agentId AgentID) {
	store := ctx.KVStore(keeper.agentStoreKey)
	bz := store.Get(keyNewAgentID)
	if bz == nil {
		agentId = 0
	} else {
		keeper.cdc.MustUnmarshalBinaryBare(bz, &agentId)
	}
	bz = keeper.cdc.MustMarshalBinaryBare(agentId + 1)
	store.Set(keyNewAgentID, bz)
	return agentId
}

func (keeper Keeper) CreateAgent(ctx sdk.Context, info AgentInfo) AgentID {
	id := keeper.getNewAgentId(ctx)
	keeper.setAgentInfo(ctx, id, info)
	return id
}

func (keeper Keeper) setAgentInfo(ctx sdk.Context, id AgentID, info AgentInfo) {
	store := ctx.KVStore(keeper.agentStoreKey)
	bz, err := keeper.cdc.MarshalBinaryBare(info)
	if err != nil {
		panic(err)
	}
	store.Set(keyAgentID(id), bz)
}

func (keeper Keeper) UpdateAgentInfo(ctx sdk.Context, id AgentID, signers []sdk.AccAddress, info AgentInfo) bool {
	if !keeper.Authorize(ctx, id, signers) {
		return false
	}
	keeper.setAgentInfo(ctx, id, info)
	return true
}

func (keeper Keeper) Authorize(ctx sdk.Context, id AgentID, signers []sdk.AccAddress) bool {
	ctx.GasMeter().ConsumeGas(10, "agent auth")
	info, err := keeper.GetAgentInfo(ctx, id)
	if err != nil {
		return false
	}
	return keeper.AuthorizeAgentInfo(ctx, &info, signers)
}

func (keeper Keeper) AuthorizeAgentInfo(ctx sdk.Context, info *AgentInfo, signers []sdk.AccAddress) bool {
	if info.AuthPolicy != MultiSig {
		panic("Unknown auth policy")
	}

	sigCount := 0
	sigThreshold := info.MultisigThreshold

	nAddrs := len(info.Addresses)
	nSigners := len(signers)
	for i := 0; i < nAddrs; i++ {
		addr := info.Addresses[i]
		// TODO Use a hash map to optimize this
		for j := 0; j < nSigners; j++ {
			ctx.GasMeter().ConsumeGas(10, "check addr")
			if bytes.Compare(addr, signers[j]) == 0 {
				sigCount++
				if sigCount >= sigThreshold {
					return true
				}
				break
			}
		}
	}

	nAgents := len(info.Agents)
	for i := 0; i < nAgents; i++ {
		agentId := info.Agents[i]
		if keeper.Authorize(ctx, agentId, signers) {
			sigCount++
			if sigCount >= sigThreshold {
				return true
			}
		}
	}
	return false
}
