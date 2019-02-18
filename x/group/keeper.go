package group

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	groupStoreKey sdk.StoreKey

	cdc *codec.Codec
}

func NewKeeper(groupStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		groupStoreKey: groupStoreKey,
		cdc:           cdc,
	}
}

var (
	keyNewGroupID = []byte("newGroupID")
)

func keyGroupID(id sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("#%d", id))
}

func (keeper Keeper) GetGroupInfo(ctx sdk.Context, id sdk.AccAddress) (info Group, err sdk.Error) {
	if len(id) < 1 || id[0] != 'G' {
		return info, sdk.ErrUnknownRequest("Not a valid group")
	}
	store := ctx.KVStore(keeper.groupStoreKey)
	bz := store.Get(keyGroupID(id))
	if bz == nil {
		return info, sdk.ErrUnknownRequest("Not found")
	}
	info = Group{}
	marshalErr := keeper.cdc.UnmarshalBinaryBare(bz, &info)
	if marshalErr != nil {
		return info, sdk.ErrUnknownRequest(marshalErr.Error())
	}
	return info, nil
}

func GroupAddrFromUint64(id uint64) sdk.AccAddress {
	addr := make([]byte, binary.MaxVarintLen64+1)
	addr[0] = 'G'
	n := binary.PutUvarint(addr[1:], id)
	return addr[:n+1]
}

func (keeper Keeper) getNewGroupId(ctx sdk.Context) sdk.AccAddress {
	store := ctx.KVStore(keeper.groupStoreKey)
	bz := store.Get(keyNewGroupID)
	var groupId uint64 = 0
	if bz != nil {
		keeper.cdc.MustUnmarshalBinaryBare(bz, &groupId)
	}
	bz = keeper.cdc.MustMarshalBinaryBare(groupId + 1)
	store.Set(keyNewGroupID, bz)
	return GroupAddrFromUint64(groupId)
}

func (keeper Keeper) CreateGroup(ctx sdk.Context, info Group) sdk.AccAddress {
	id := keeper.getNewGroupId(ctx)
	keeper.setGroupInfo(ctx, id, info)
	return id
}

func (keeper Keeper) setGroupInfo(ctx sdk.Context, id sdk.AccAddress, info Group) {
	store := ctx.KVStore(keeper.groupStoreKey)
	bz, err := keeper.cdc.MarshalBinaryBare(info)
	if err != nil {
		panic(err)
	}
	store.Set(keyGroupID(id), bz)
}

//func (keeper Keeper) UpdateGroupInfo(ctx sdk.Context, id GroupID, signers []sdk.AccAddress, info GroupInfo) bool {
//	if !keeper.Authorize(ctx, id, signers) {
//		return false
//	}
//	keeper.setGroupInfo(ctx, id, info)
//	return true
//}

func (keeper Keeper) Authorize(ctx sdk.Context, address sdk.AccAddress, signers []sdk.AccAddress) bool {
	info, err := keeper.GetGroupInfo(ctx, address)
	if err != nil {
		return false
	}
	ctx.GasMeter().ConsumeGas(10, "group auth")
	return keeper.AuthorizeGroupInfo(ctx, &info, signers)
}

func (keeper Keeper) AuthorizeGroupInfo(ctx sdk.Context, info *Group, signers []sdk.AccAddress) bool {
	voteCount := big.NewInt(0)
	sigThreshold := info.DecisionThreshold

	nMembers := len(info.Members)
	nSigners := len(signers)
	for i := 0; i < nMembers; i++ {
		mem := info.Members[i]
		// TODO Use a hash map to optimize this
		for j := 0; j < nSigners; j++ {
			ctx.GasMeter().ConsumeGas(10, "check addr")
			if bytes.Compare(mem.Address, signers[j]) == 0 || keeper.Authorize(ctx, mem.Address, signers) {
				voteCount = voteCount.Add(voteCount, &mem.Weight)
				if voteCount.Cmp(&sigThreshold) >= 0 {
					return true
				}
				break
			}
		}
	}
	return false
}
