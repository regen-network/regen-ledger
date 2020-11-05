package testdata

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group"
)

type Keeper struct {
	groupKeeper group.Keeper
	key         sdk.StoreKey
}

func NewKeeper(storeKey sdk.StoreKey, groupKeeper group.Keeper) Keeper {
	k := Keeper{
		groupKeeper: groupKeeper,
		key:         storeKey,
	}

	return k
}

func (k Keeper) CreateProposal(ctx sdk.Context, accountAddress sdk.AccAddress,
	proposers []sdk.AccAddress, comment string, msgs []sdk.Msg) (group.ProposalID, error) {
	return k.groupKeeper.CreateProposal(ctx, accountAddress, comment, proposers, msgs)
}

var (
	storageKey = []byte("key")
	counterKey = []byte("counter")
)

func (k Keeper) SetValue(ctx sdk.Context, value string) {
	ctx.KVStore(k.key).Set(storageKey, []byte(value))
}

func (k Keeper) GetValue(ctx sdk.Context) string {
	return string(ctx.KVStore(k.key).Get(storageKey))
}

func (k Keeper) IncCounter(ctx sdk.Context) []byte {
	i := k.GetCounter(ctx)
	i++
	v := orm.EncodeSequence(i)
	ctx.KVStore(k.key).Set(counterKey, v)
	return v
}

func (k Keeper) GetCounter(ctx sdk.Context) uint64 {
	return orm.DecodeSequence(ctx.KVStore(k.key).Get(counterKey))
}
