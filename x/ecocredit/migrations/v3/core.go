package v3

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	orm "github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// TODO: add projects info
const projectsJSON = `{"projects":[{"issuer":"cosmos154hmhstk3gpkv2ndec8zkjkc5c3svutcqcswne","class_id":"A00","metadata":"hello","project_location":"AB-CDE FG1 345","project_id":"A0"}]}`

func MigrateStore(sdkCtx sdk.Context, storeKey storetypes.StoreKey,
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

	creditTypeSeqTableBuilder, err := orm.NewPrimaryKeyTableBuilder(CreditTypeSeqTablePrefix, storeKey, &ecocredit.CreditTypeSeq{}, cdc)
	if err != nil {
		return err
	}
	creditTypeSeqTable := creditTypeSeqTableBuilder.Build()

	// migrate credit classes
	classItr, err := classInfoTable.PrefixScan(sdkCtx, nil, nil)
	if err != nil {
		return err
	}
	defer classItr.Close()

	classIDsMap := make(map[uint64]string)
	projectIDsMap := make(map[string]uint64)
	var classID uint64 = 1
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
			Id:         classID,
			Name:       classInfo.ClassId,
			Admin:      admin,
			Metadata:   string(classInfo.Metadata),
			CreditType: classInfo.CreditType.Abbreviation,
		}
		classIDsMap[classID] = classInfo.ClassId
		if err := ss.ClassInfoTable().Insert(ctx, &dest); err != nil {
			return err
		}

		// migrate class issuers
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

		classID++
	}

	// migrate credit type sequence
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
		ss.ClassSequenceTable().Save(ctx, &api.ClassSequence{
			CreditType:  ctype.Abbreviation,
			NextClassId: ctype.SeqNumber,
		})
	}

	// migrate projects
	// TODO: manually add projects for existing credit classes
	projects := []api.ProjectInfo{}

	var pID uint64 = 1
	for _, p := range projects {
		dest := api.ProjectInfo{
			Id:              pID,
			Name:            p.Name,
			Admin:           p.Admin,
			ClassId:         p.ClassId,
			ProjectLocation: p.ProjectLocation,
			Metadata:        p.Metadata,
		}
		if err := ss.ProjectInfoTable().Insert(ctx, &dest); err != nil {
			return err
		}

		projectIDsMap[classIDsMap[p.ClassId]] = pID
		pID++
	}

	// TODO: migrate sequence
	ss.ProjectSequenceTable().Save(ctx, &api.ProjectSequence{})

	// migrate batches
	batchIDsMap := make(map[string]uint64)
	batchItr, err := batchInfoTable.PrefixScan(sdkCtx, nil, nil)
	if err != nil {
		return err
	}
	defer batchItr.Close()

	var batchID uint64 = 1
	for {
		var batchInfo BatchInfo
		if _, err := batchItr.LoadNext(&batchInfo); err != nil {
			if orm.ErrIteratorDone.Is(err) {
				break
			}
			return err
		}

		bInfo := api.BatchInfo{
			Id:           batchID,
			ProjectId:    projectIDsMap[batchInfo.ClassId],
			BatchDenom:   batchInfo.BatchDenom,
			Metadata:     string(batchInfo.Metadata),
			StartDate:    timestamppb.New(*batchInfo.StartDate),
			EndDate:      timestamppb.New(*batchInfo.EndDate),
			IssuanceDate: nil, // TODO: add issuance date
		}

		if err := ss.BatchInfoTable().Insert(ctx, &bInfo); err != nil {
			return err
		}

		batchIDsMap[bInfo.BatchDenom] = batchID
		batchID++
	}

	store := sdkCtx.KVStore(storeKey)
	// migrate tradable balances
	IterateBalances(store, TradableBalancePrefix, func(address, denom, balance string) (bool, error) {
		addr, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return true, err
		}

		if err := ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
			Address:  addr,
			BatchId:  batchIDsMap[denom],
			Tradable: balance,
		}); err != nil {
			return true, err
		}

		return false, nil
	})

	// migrate retired balances
	IterateBalances(store, RetiredBalancePrefix, func(address, denom, balance string) (bool, error) {
		addr, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return false, err
		}

		b, err := ss.BatchBalanceTable().Get(ctx, addr, batchIDsMap[denom])
		if err != nil {
			if orm.ErrNotFound.Is(err) {
				if err := ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
					Address: addr,
					BatchId: batchIDsMap[denom],
					Retired: balance,
				}); err != nil {
					return true, err
				}
			}

			return true, err
		}

		if err := ss.BatchBalanceTable().Update(ctx, &api.BatchBalance{
			Address:  addr,
			BatchId:  batchIDsMap[denom],
			Tradable: b.Tradable,
			Retired:  balance,
		}); err != nil {
			return true, err
		}

		return false, nil
	})

	// migrate tradable supply
	IterateSupplies(store, TradableSupplyPrefix, func(denom, supply string) (bool, error) {
		if err := ss.BatchSupplyTable().Save(ctx, &api.BatchSupply{
			BatchId:        batchIDsMap[denom],
			TradableAmount: supply,
		}); err != nil {
			return false, err
		}

		return false, nil
	})

	// migrate retired supply
	IterateSupplies(store, RetiredSupplyPrefix, func(denom, supply string) (bool, error) {
		bs, err := ss.BatchSupplyTable().Get(ctx, batchIDsMap[denom])
		if err != nil {
			if orm.ErrNotFound.Is(err) {
				if err := ss.BatchSupplyTable().Save(ctx, &api.BatchSupply{
					BatchId:       batchIDsMap[denom],
					RetiredAmount: supply,
				}); err != nil {
					return false, err
				}
			}
			return false, err
		}

		if err := ss.BatchSupplyTable().Update(ctx, &api.BatchSupply{
			BatchId:        batchIDsMap[denom],
			RetiredAmount:  supply,
			TradableAmount: bs.TradableAmount,
		}); err != nil {
			return false, err
		}

		return false, nil
	})

	return nil
}
