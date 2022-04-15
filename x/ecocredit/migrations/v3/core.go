package v3

import (
	"context"
	"sort"

	"github.com/cosmos/cosmos-sdk/codec"
	ormerrors "github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	orm "github.com/regen-network/regen-ledger/orm"
)

type batchMapT struct {
	Id              uint64
	AmountCancelled string
}

// MigrateState performs in-place store migrations from v3.0 to v4.0.
func MigrateState(sdkCtx sdk.Context, storeKey storetypes.StoreKey,
	cdc codec.Codec, ss api.StateStore) error {
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

	// migrate credit classes to ORM v1
	classItr, err := classInfoTable.PrefixScan(sdkCtx, nil, nil)
	if err != nil {
		return err
	}
	defer classItr.Close()

	classIDNameMap := make(map[uint64]string) // map of a credit classID to className
	classNameIDMap := make(map[string]uint64) // map of a credit className to classID
	ctx := sdk.WrapSDKContext(sdkCtx)
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
		dest := api.ClassInfo{
			Name:       classInfo.ClassId,
			Admin:      admin,
			Metadata:   string(classInfo.Metadata),
			CreditType: classInfo.CreditType.Abbreviation,
		}
		classID, err := ss.ClassInfoTable().InsertReturningID(ctx, &dest)
		if err != nil {
			return err
		}
		classIDNameMap[classID] = classInfo.ClassId
		classNameIDMap[classInfo.ClassId] = classID

		// migrate class issuers to ORM v1
		for _, issuer := range classInfo.Issuers {
			addr, err := sdk.AccAddressFromBech32(issuer)
			if err != nil {
				return err
			}

			if err := ss.ClassIssuerTable().Insert(ctx, &api.ClassIssuer{
				ClassId: classID,
				Issuer:  addr,
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
			CreditType:  ctype.Abbreviation,
			NextClassId: ctype.SeqNumber,
		}); err != nil {
			return err
		}

		// delete credit type sequence from old store
		if err := creditTypeSeqTable.Delete(sdkCtx, &ctype); err != nil {
			return err
		}
	}

	// migrate credit batches to ORM v1 and create projects for existing credit classes
	batchIDsMap := make(map[string]batchMapT) // map of a batch denom to batch-id and amount cancelled
	batchSeqMap := make(map[uint64]uint64)    // map of a project-id to batch sequence
	projectSeqMap := make(map[uint64]uint64)  // map of a credit classID to project sequence
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
		pItr, err := ss.ProjectInfoTable().List(ctx, api.ProjectInfoAdminIndexKey{}.WithAdmin(admin.Bytes()))
		if err != nil {
			return err
		}

		projectExists := false
		var projectID uint64
		for pItr.Next() {
			pInfo, err := pItr.Value()
			if err != nil {
				return err
			}

			if pInfo.ClassId == classNameIDMap[batchInfo.ClassId] && pInfo.ProjectLocation == batchInfo.ProjectLocation {
				projectExists = true
				projectID = pInfo.Id
				break
			}

		}
		pItr.Close()

		if !projectExists {
			classID := classNameIDMap[batchInfo.ClassId]
			var projectSeq uint64 = 1
			if val, ok := projectSeqMap[classID]; ok {
				projectSeqMap[classID] = val + 1
				projectSeq = val
			} else {
				projectSeqMap[classID] = 2
			}

			name := FormatProjectID(batchInfo.ClassId, projectSeq)
			id, err := ss.ProjectInfoTable().InsertReturningID(ctx,
				&api.ProjectInfo{
					Name:            name,
					Admin:           admin,
					ClassId:         classID,
					ProjectLocation: batchInfo.ProjectLocation,
					Metadata:        "",
				},
			)
			if err != nil {
				return err
			}
			projectID = id
		}

		bInfo := api.BatchInfo{
			ProjectId:    projectID,
			BatchDenom:   batchInfo.BatchDenom,
			Metadata:     string(batchInfo.Metadata),
			StartDate:    timestamppb.New(*batchInfo.StartDate),
			EndDate:      timestamppb.New(*batchInfo.EndDate),
			IssuanceDate: nil,
		}

		bID, err := ss.BatchInfoTable().InsertReturningID(ctx, &bInfo)
		if err != nil {
			return err
		}

		batchIDsMap[bInfo.BatchDenom] = batchMapT{
			Id:              bID,
			AmountCancelled: batchInfo.AmountCancelled,
		}

		if v, ok := batchSeqMap[bInfo.ProjectId]; ok {
			batchSeqMap[bInfo.ProjectId] = v + 1
		} else {
			batchSeqMap[bInfo.ProjectId] = 2
		}

		// delete credit batch from old store
		if err := batchInfoTable.Delete(sdkCtx, &batchInfo); err != nil {
			return err
		}
	}

	// add project sequence
	keys := make([]int, 0, len(projectSeqMap))
	for k := range projectSeqMap {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		key := uint64(k)
		err = ss.ProjectSequenceTable().Save(ctx, &api.ProjectSequence{
			ClassId:       key,
			NextProjectId: projectSeqMap[key],
		})
		if err != nil {
			return err
		}
	}

	// add batch sequence
	keys1 := make([]int, 0, len(batchSeqMap))
	for k := range batchSeqMap {
		keys1 = append(keys1, int(k))
	}
	sort.Ints(keys1)

	for _, k := range keys1 {
		pInfo, err := ss.ProjectInfoTable().Get(ctx, uint64(k))
		if err != nil {
			return err
		}

		if err := ss.BatchSequenceTable().Save(ctx, &api.BatchSequence{
			ProjectId:   pInfo.Name,
			NextBatchId: batchSeqMap[uint64(k)],
		}); err != nil {
			return err
		}
	}

	store := sdkCtx.KVStore(storeKey)
	if err = migrateBalances(store, ss, ctx, batchIDsMap); err != nil {
		return err
	}

	if err = migrateSupply(store, ss, ctx, batchIDsMap); err != nil {
		return err
	}

	return nil
}

// migrateBalances migrates ecocredit tradable and retired balances to orm v1
func migrateBalances(store storetypes.KVStore, ss api.StateStore, ctx context.Context, batchIDsMap map[string]batchMapT) error {
	// migrate tradable balances to ORM v1
	if err := IterateBalances(store, TradableBalancePrefix, func(address, denom, balance string) (bool, error) {
		addr, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return true, err
		}

		if err := ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
			Address:  addr,
			BatchId:  batchIDsMap[denom].Id,
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

		b, err := ss.BatchBalanceTable().Get(ctx, addr, batchIDsMap[denom].Id)
		if err != nil {
			if ormerrors.IsNotFound(err) {
				if err := ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
					Address: addr,
					BatchId: batchIDsMap[denom].Id,
					Retired: balance,
				}); err != nil {
					return true, err
				}
				return false, nil
			}

			return true, err
		}

		if err := ss.BatchBalanceTable().Update(ctx, &api.BatchBalance{
			Address:  addr,
			BatchId:  batchIDsMap[denom].Id,
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
func migrateSupply(store storetypes.KVStore, ss api.StateStore, ctx context.Context, batchIDsMap map[string]batchMapT) error {
	// migrate tradable supply to ORM v1
	if err := IterateSupplies(store, TradableSupplyPrefix, func(denom, supply string) (bool, error) {
		if err := ss.BatchSupplyTable().Save(ctx, &api.BatchSupply{
			BatchId:         batchIDsMap[denom].Id,
			CancelledAmount: batchIDsMap[denom].AmountCancelled,
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
		bs, err := ss.BatchSupplyTable().Get(ctx, batchIDsMap[denom].Id)
		if err != nil {
			if ormerrors.IsNotFound(err) {
				if err := ss.BatchSupplyTable().Save(ctx, &api.BatchSupply{
					BatchId:         batchIDsMap[denom].Id,
					CancelledAmount: batchIDsMap[denom].AmountCancelled,
					RetiredAmount:   supply,
				}); err != nil {
					return false, err
				}
				return false, nil
			}
			return true, err

		}

		if err := ss.BatchSupplyTable().Update(ctx, &api.BatchSupply{
			BatchId:         batchIDsMap[denom].Id,
			CancelledAmount: batchIDsMap[denom].AmountCancelled,
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
