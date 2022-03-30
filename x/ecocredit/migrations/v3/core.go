package v3

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	orm "github.com/regen-network/regen-ledger/orm"
)

// TODO: add projects info
const projectsJSON = `[{"issuer":"cosmos154hmhstk3gpkv2ndec8zkjkc5c3svutcqcswne","class_id":"A00","metadata":"hello","project_location":"AB-CDE FG1 345","project_id":"A0"}]`

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.Codec) error {
	classInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(ClassInfoTablePrefix, storeKey, &ClassInfo{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	classInfoTable := classInfoTableBuilder.Build()

	batchInfoTableBuilder, err := orm.NewPrimaryKeyTableBuilder(BatchInfoTablePrefix, storeKey, &BatchInfo{}, cdc)
	if err != nil {
		panic(err.Error())
	}
	batchInfoTable := batchInfoTableBuilder.Build()

	// classes
	classItr, err := classInfoTable.PrefixScan(ctx, nil, nil)
	if err != nil {
		return err
	}
	defer classItr.Close()
	projectIDsMap := make(map[string]uint64)
	var classID uint64 = 1
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
		classID++
		// TODO: insert into new orm table

		// TODO: also create project
		projectIDsMap[classInfo.ClassId] = 1

	}

	// batches
	batchItr, err := batchInfoTable.PrefixScan(ctx, nil, nil)
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

		// TODO: insert into new orm table
	}

	return nil
}
