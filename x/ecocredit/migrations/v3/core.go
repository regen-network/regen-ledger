package v3

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type batchMapT struct {
	Id              uint64
	AmountCancelled string
}

// MigrateState performs in-place store migrations from v3.0 to v4.0.
func MigrateState(sdkCtx sdk.Context, storeKey storetypes.StoreKey,
	cdc codec.Codec, ss api.StateStore, subspace paramtypes.Subspace) error {
	classInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(ClassInfoTablePrefix, storeKey, &ClassInfo{}, cdc)
	if err != nil {
		return err
	}
	classInfoTable := classInfoTableBuilder.Build()

	batchInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(BatchInfoTablePrefix, storeKey, &BatchInfo{}, cdc)
	if err != nil {
		return err
	}
	batchInfoTable := batchInfoTableBuilder.Build()

	creditTypeSeqTableBuilder, err := orm.NewPrimaryKeyTableBuilder(CreditTypeSeqTablePrefix, storeKey, &CreditTypeSeq{}, cdc)
	if err != nil {
		return err
	}
	creditTypeSeqTable := creditTypeSeqTableBuilder.Build()

	// migrate credit types from params to ORM table
	bz := subspace.GetRaw(sdkCtx, KeyCreditTypes)
	if bz == nil {
		return fmt.Errorf("credit types empty")
	}

	var creditTypes []CreditType
	if err := json.Unmarshal(bz, &creditTypes); err != nil {
		return err
	}

	ctx := sdk.WrapSDKContext(sdkCtx)
	for _, creditType := range creditTypes {
		if err := ss.CreditTypeTable().Insert(ctx, &api.CreditType{
			Abbreviation: creditType.Abbreviation,
			Name:         creditType.Name,
			Unit:         creditType.Unit,
			Precision:    creditType.Precision,
		}); err != nil {
			return err
		}
	}

	// migrate credit classes to ORM v1
	classItr, err := classInfoTable.PrefixScan(sdkCtx, nil, nil)
	if err != nil {
		return err
	}
	defer classItr.Close()

	classKeyToClassId := make(map[uint64]string) // map of a class key to a class id
	classIdToClassKey := make(map[string]uint64) // map of a class id to a class key
	for {
		var classInfo ClassInfo
		if _, err := classItr.LoadNext(&classInfo); err != nil {
			if orm.ErrIteratorDone.Is(err) {
				break
			}
			return err
		}

		admin, err := sdk.AccAddressFromBech32(classInfo.Admin)
		if err != nil {
			return err
		}
		dest := api.Class{
			Id:               classInfo.ClassId,
			Admin:            admin,
			Metadata:         string(classInfo.Metadata),
			CreditTypeAbbrev: classInfo.CreditType.Abbreviation,
		}
		classKey, err := ss.ClassTable().InsertReturningID(ctx, &dest)
		if err != nil {
			return err
		}
		classKeyToClassId[classKey] = classInfo.ClassId
		classIdToClassKey[classInfo.ClassId] = classKey

		// migrate class issuers to ORM v1
		for _, issuer := range classInfo.Issuers {
			addr, err := sdk.AccAddressFromBech32(issuer)
			if err != nil {
				return err
			}

			if err := ss.ClassIssuerTable().Insert(ctx, &api.ClassIssuer{
				ClassKey: classKey,
				Issuer:   addr,
			}); err != nil {
				return err
			}
		}

		// delete class info from old store
		if err := classInfoTable.Delete(sdkCtx, &ClassInfo{ClassId: classInfo.ClassId}); err != nil {
			return err
		}
	}

	// migrate credit type sequence to ORM v1
	cItr, err := creditTypeSeqTable.PrefixScan(sdkCtx, nil, nil)
	if err != nil {
		return err
	}
	defer cItr.Close()

	for {
		var ctype CreditTypeSeq
		if _, err := cItr.LoadNext(&ctype); err != nil {
			if orm.ErrIteratorDone.Is(err) {
				break
			}

			return err
		}
		if err := ss.ClassSequenceTable().Save(ctx, &api.ClassSequence{
			CreditTypeAbbrev: ctype.Abbreviation,
			NextSequence:     ctype.SeqNumber,
		}); err != nil {
			return err
		}

		// delete credit type sequence from old store
		if err := creditTypeSeqTable.Delete(sdkCtx, &ctype); err != nil {
			return err
		}
	}

	// migrate credit batches to ORM v1 and create projects for existing credit classes
	batchDenomToBatchMap := make(map[string]batchMapT) // map of a batch denom to batch id and amount cancelled
	projectKeyToBatchSeq := make(map[uint64]uint64)    // map of a project key to batch sequence
	classKeyToProjectSeq := make(map[uint64]uint64)    // map of a class key to project sequence
	batchItr, err := batchInfoTable.PrefixScan(sdkCtx, nil, nil)
	if err != nil {
		return err
	}
	defer batchItr.Close()

	for {
		var batchInfo BatchInfo
		if _, err := batchItr.LoadNext(&batchInfo); err != nil {
			if orm.ErrIteratorDone.Is(err) {
				break
			}
			return err
		}

		admin, err := sdk.AccAddressFromBech32(batchInfo.Issuer)
		if err != nil {
			return err
		}
		pItr, err := ss.ProjectTable().List(ctx, api.ProjectAdminIndexKey{}.WithAdmin(admin.Bytes()))
		if err != nil {
			return err
		}

		projectExists := false
		var projectKey uint64
		var projectId string
		for pItr.Next() {
			project, err := pItr.Value()
			if err != nil {
				return err
			}

			if project.ClassKey == classIdToClassKey[batchInfo.ClassId] && project.Jurisdiction == batchInfo.ProjectLocation {
				projectExists = true
				projectKey = project.Key
				projectId = project.Id
				break
			}

		}
		pItr.Close()

		if !projectExists {
			classKey := classIdToClassKey[batchInfo.ClassId]
			var projectSeq uint64 = 1
			if val, ok := classKeyToProjectSeq[classKey]; ok {
				classKeyToProjectSeq[classKey] = val + 1
				projectSeq = val
			} else {
				classKeyToProjectSeq[classKey] = 2
			}

			id := core.FormatProjectId(batchInfo.ClassId, projectSeq)
			key, err := ss.ProjectTable().InsertReturningID(ctx,
				&api.Project{
					Id:           id,
					Admin:        admin,
					ClassKey:     classKey,
					Jurisdiction: batchInfo.ProjectLocation,
					Metadata:     "",
				},
			)
			if err != nil {
				return err
			}
			projectKey = key
			projectId = id
		}

		startDate, endDate, err := parseBatchDenom(batchInfo.BatchDenom)
		if err != nil {
			return err
		}

		var batchSeq uint64 = 1
		if v, ok := projectKeyToBatchSeq[projectKey]; ok {
			batchSeq = v
		}

		batchDenom, err := core.FormatBatchDenom(projectId, batchSeq, startDate, endDate)
		if err != nil {
			return err
		}

		batch := api.Batch{
			ProjectKey:   projectKey,
			Denom:        batchDenom,
			Metadata:     string(batchInfo.Metadata),
			StartDate:    timestamppb.New(*batchInfo.StartDate),
			EndDate:      timestamppb.New(*batchInfo.EndDate),
			IssuanceDate: nil,
		}

		bID, err := ss.BatchTable().InsertReturningID(ctx, &batch)
		if err != nil {
			return err
		}

		batchDenomToBatchMap[batchInfo.BatchDenom] = batchMapT{
			Id:              bID,
			AmountCancelled: batchInfo.AmountCancelled,
		}

		if v, ok := projectKeyToBatchSeq[batch.ProjectKey]; ok {
			projectKeyToBatchSeq[batch.ProjectKey] = v + 1
		} else {
			projectKeyToBatchSeq[batch.ProjectKey] = 2
		}

		// delete credit batch from old store
		if err := batchInfoTable.Delete(sdkCtx, &batchInfo); err != nil {
			return err
		}
	}

	// add project sequence
	keys := make([]int, 0, len(classKeyToProjectSeq))
	for k := range classKeyToProjectSeq {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		key := uint64(k)
		err = ss.ProjectSequenceTable().Save(ctx, &api.ProjectSequence{
			ClassKey:     key,
			NextSequence: classKeyToProjectSeq[key],
		})
		if err != nil {
			return err
		}
	}

	// add batch sequence
	keys1 := make([]int, 0, len(projectKeyToBatchSeq))
	for k := range projectKeyToBatchSeq {
		keys1 = append(keys1, int(k))
	}
	sort.Ints(keys1)

	for _, k := range keys1 {
		if err := ss.BatchSequenceTable().Save(ctx, &api.BatchSequence{
			ProjectKey:   uint64(k),
			NextSequence: projectKeyToBatchSeq[uint64(k)],
		}); err != nil {
			return err
		}
	}

	store := sdkCtx.KVStore(storeKey)
	if err = migrateBalances(store, ss, ctx, batchDenomToBatchMap); err != nil {
		return err
	}

	if err = migrateSupply(store, ss, ctx, batchDenomToBatchMap); err != nil {
		return err
	}

	return nil
}

// migrateBalances migrates ecocredit tradable and retired balances to orm v1
func migrateBalances(store storetypes.KVStore, ss api.StateStore, ctx context.Context, batchDenomToBatchMap map[string]batchMapT) error {
	// migrate tradable balances to ORM v1
	if err := IterateBalances(store, TradableBalancePrefix, func(address, denom, balance string) (bool, error) {
		addr, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return true, err
		}

		if err := ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
			BatchKey: batchDenomToBatchMap[denom].Id,
			Address:  addr,
			Tradable: balance,
		}); err != nil {
			return true, err
		}

		// delete tradable balance from old store
		store.Delete(TradableBalanceKey(addr, BatchDenomT(denom)))

		return false, nil
	}); err != nil {
		return err
	}

	// migrate retired balances to ORM v1
	err := IterateBalances(store, RetiredBalancePrefix, func(address, denom, balance string) (bool, error) {
		addr, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return false, err
		}

		b, err := ss.BatchBalanceTable().Get(ctx, addr, batchDenomToBatchMap[denom].Id)
		if err != nil {
			if ormerrors.IsNotFound(err) {
				if err := ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
					BatchKey: batchDenomToBatchMap[denom].Id,
					Address:  addr,
					Retired:  balance,
				}); err != nil {
					return true, err
				}
				return false, nil
			}

			return true, err
		}

		if err := ss.BatchBalanceTable().Update(ctx, &api.BatchBalance{
			BatchKey: batchDenomToBatchMap[denom].Id,
			Address:  addr,
			Tradable: b.Tradable,
			Retired:  balance,
		}); err != nil {
			return true, err
		}

		// delete retired balance from old store
		store.Delete(RetiredBalanceKey(addr, BatchDenomT(denom)))

		return false, nil
	})

	return err
}

// migrateSupply migrates tradable and retired supply to orm v1
func migrateSupply(store storetypes.KVStore, ss api.StateStore, ctx context.Context, batchDenomToBatchMap map[string]batchMapT) error {
	// migrate tradable supply to ORM v1
	if err := IterateSupplies(store, TradableSupplyPrefix, func(denom, supply string) (bool, error) {
		if err := ss.BatchSupplyTable().Save(ctx, &api.BatchSupply{
			BatchKey:        batchDenomToBatchMap[denom].Id,
			CancelledAmount: batchDenomToBatchMap[denom].AmountCancelled,
			TradableAmount:  supply,
		}); err != nil {
			return false, err
		}

		// delete tradable supply from old store
		store.Delete(TradableSupplyKey(BatchDenomT(denom)))

		return false, nil
	}); err != nil {
		return err
	}

	// migrate retired supply to ORM v1
	err := IterateSupplies(store, RetiredSupplyPrefix, func(denom, supply string) (bool, error) {
		bs, err := ss.BatchSupplyTable().Get(ctx, batchDenomToBatchMap[denom].Id)
		if err != nil {
			if ormerrors.IsNotFound(err) {
				if err := ss.BatchSupplyTable().Save(ctx, &api.BatchSupply{
					BatchKey:        batchDenomToBatchMap[denom].Id,
					CancelledAmount: batchDenomToBatchMap[denom].AmountCancelled,
					RetiredAmount:   supply,
				}); err != nil {
					return false, err
				}
				return false, nil
			}
			return true, err

		}

		if err := ss.BatchSupplyTable().Update(ctx, &api.BatchSupply{
			BatchKey:        batchDenomToBatchMap[denom].Id,
			CancelledAmount: batchDenomToBatchMap[denom].AmountCancelled,
			RetiredAmount:   supply,
			TradableAmount:  bs.TradableAmount,
		}); err != nil {
			return false, err
		}

		// delete retired supply from old store
		store.Delete(RetiredSupplyKey(BatchDenomT(denom)))

		return false, nil
	})

	return err
}
