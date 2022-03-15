package ormstore

import (
	"context"
	ormv1alpha1 "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	"github.com/regen-network/regen-ledger/types"

	queryv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/query/v1beta1"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// NewStoreKeyDB creates an ormdb.ModuleDB from an ormdb.ModuleDB and a StoreKey.
// It is an interim solution for using the ORM in existing Cosmos SDK modules
// before fuller integration has been done.
func NewStoreKeyDB(desc *ormv1alpha1.ModuleSchemaDescriptor, key storetypes.StoreKey, options ormdb.ModuleDBOptions) (ormdb.ModuleDB, error) {
	backEndResolver := func(_ ormv1alpha1.StorageType) (ormtable.BackendResolver, error) {
		getBackend := func(ctx context.Context) (ormtable.ReadBackend, error) {
			sdkCtx := types.UnwrapSDKContext(ctx)
			store := sdkCtx.KVStore(key)
			wrapper := storeWrapper{store}
			return ormtable.NewBackend(ormtable.BackendOptions{
				CommitmentStore: wrapper,
				IndexStore:      wrapper,
			}), nil
		}
		return getBackend, nil
	}
	options.GetBackendResolver = backEndResolver
	return ormdb.NewModuleDB(desc, options)
}

func GogoPageReqToPulsarPageReq(request *query.PageRequest) *queryv1beta1.PageRequest {
	return &queryv1beta1.PageRequest{
		Key:        request.Key,
		Offset:     request.Offset,
		Limit:      request.Limit,
		CountTotal: request.CountTotal,
		Reverse:    request.Reverse,
	}
}

func PulsarPageResToGogoPageRes(response *queryv1beta1.PageResponse) *query.PageResponse {
	return &query.PageResponse{
		NextKey: response.NextKey,
		Total:   response.Total,
	}
}
