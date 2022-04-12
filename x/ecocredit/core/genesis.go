package core

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/reflect/protoreflect"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// ValidateGenesis performs basic validation for the following:
// - params are valid param types with valid properties
// - the credit type referenced in each credit class exists
// - the credit class referenced in each project exists
// - the tradable amount of each credit batch complies with the credit type precision
// - the retired amount of each credit batch complies with the credit type precision
// - the calculated total amount of each credit batch matches the total supply
// An error is returned if any of these validation checks fail.
func ValidateGenesis(data json.RawMessage, params Params) error {
	if err := params.Validate(); err != nil {
		return err
	}

	db := dbm.NewMemDB()
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		return err
	}

	ormCtx := ormtable.WrapContextDefault(backend)
	ss, err := api.NewStateStore(ormdb)
	if err != nil {
		return err
	}

	jsonSource, err := ormjson.NewRawMessageSource(data)
	if err != nil {
		return err
	}

	err = ormdb.ImportJSON(ormCtx, jsonSource)
	if err != nil {
		return err
	}

	abbrevToPrecision := make(map[string]uint32) // map of credit abbreviation to precision
	for _, ct := range params.CreditTypes {
		abbrevToPrecision[ct.Abbreviation] = ct.Precision
	}

	cItr, err := ss.ClassInfoTable().List(ormCtx, api.ClassInfoPrimaryKey{})
	if err != nil {
		return err
	}
	defer cItr.Close()

	// make sure credit type exist for class abbreviation in params
	for cItr.Next() {
		class, err := cItr.Value()
		if err != nil {
			return err
		}

		if _, ok := abbrevToPrecision[class.CreditType]; !ok {
			return sdkerrors.ErrNotFound.Wrapf("credit type not exist for %s abbreviation", class.CreditType)
		}
	}

	projectIdToClassId := make(map[uint64]uint64) // map of projectID to classID
	pItr, err := ss.ProjectInfoTable().List(ormCtx, api.ProjectInfoPrimaryKey{})
	if err != nil {
		return err
	}
	defer pItr.Close()

	for pItr.Next() {
		project, err := pItr.Value()
		if err != nil {
			return err
		}

		if _, exists := projectIdToClassId[project.Id]; exists {
			continue
		}
		projectIdToClassId[project.Id] = project.ClassId
	}

	batchIdToPrecision := make(map[uint64]uint32) // map of batchID to precision
	bItr, err := ss.BatchInfoTable().List(ormCtx, api.BatchInfoPrimaryKey{})
	if err != nil {
		return err
	}
	defer bItr.Close()

	// create index batchID => precision for faster lookup
	for bItr.Next() {
		batch, err := bItr.Value()
		if err != nil {
			return err
		}

		if _, exists := batchIdToPrecision[batch.Id]; exists {
			continue
		}

		class, err := ss.ClassInfoTable().Get(ormCtx, projectIdToClassId[batch.ProjectId])
		if err != nil {
			return err
		}

		if class.Id == projectIdToClassId[batch.ProjectId] {
			batchIdToPrecision[batch.Id] = abbrevToPrecision[class.CreditType]
		}
	}

	batchIdToCalSupply := make(map[uint64]math.Dec) // map of batchID to calculated supply
	batchIdToSupply := make(map[uint64]math.Dec)    // map of batchID to actual supply
	bsItr, err := ss.BatchSupplyTable().List(ormCtx, api.BatchSupplyPrimaryKey{})
	if err != nil {
		return err
	}
	defer bsItr.Close()

	// calculate total supply for each credit batch (tradable + retired supply)
	for bsItr.Next() {
		batchSupply, err := bsItr.Value()
		if err != nil {
			return err
		}

		tSupply := math.NewDecFromInt64(0)
		rSupply := math.NewDecFromInt64(0)
		if batchSupply.TradableAmount != "" {
			tSupply, err = math.NewNonNegativeFixedDecFromString(batchSupply.TradableAmount, batchIdToPrecision[batchSupply.BatchId])
			if err != nil {
				return err
			}
		}
		if batchSupply.RetiredAmount != "" {
			rSupply, err = math.NewNonNegativeFixedDecFromString(batchSupply.RetiredAmount, batchIdToPrecision[batchSupply.BatchId])
			if err != nil {
				return err
			}
		}

		total, err := math.SafeAddBalance(tSupply, rSupply)
		if err != nil {
			return err
		}

		batchIdToSupply[batchSupply.BatchId] = total
	}

	// calculate credit batch supply from genesis tradable and retired balances
	if err := calculateSupply(ormCtx, batchIdToPrecision, ss, batchIdToCalSupply); err != nil {
		return err
	}

	// verify calculated total amount of each credit batch matches the total supply
	if err := validateSupply(batchIdToCalSupply, batchIdToSupply); err != nil {
		return err
	}

	return nil
}

func calculateSupply(ctx context.Context, decimalPlaces map[uint64]uint32, ss api.StateStore, calSupply map[uint64]math.Dec) error {
	bbItr, err := ss.BatchBalanceTable().List(ctx, api.BatchBalancePrimaryKey{})
	if err != nil {
		return err
	}
	defer bbItr.Close()

	for bbItr.Next() {
		tBalance := math.NewDecFromInt64(0)
		rBalance := math.NewDecFromInt64(0)

		balance, err := bbItr.Value()
		if err != nil {
			return err
		}

		if _, ok := decimalPlaces[balance.BatchId]; !ok {
			return sdkerrors.ErrInvalidType.Wrapf("credit type not exist for %d batch", balance.BatchId)
		}

		if balance.Tradable != "" {
			tBalance, err = math.NewNonNegativeFixedDecFromString(balance.Tradable, decimalPlaces[balance.BatchId])
			if err != nil {
				return err
			}
		}

		if balance.Retired != "" {
			rBalance, err = math.NewNonNegativeFixedDecFromString(balance.Retired, decimalPlaces[balance.BatchId])
			if err != nil {
				return err
			}
		}

		total, err := math.Add(tBalance, rBalance)
		if err != nil {
			return err
		}

		if supply, ok := calSupply[balance.BatchId]; ok {
			result, err := math.SafeAddBalance(supply, total)
			if err != nil {
				return err
			}
			calSupply[balance.BatchId] = result
		} else {
			calSupply[balance.BatchId] = total
		}
	}

	return nil
}

func validateSupply(calSupply, supply map[uint64]math.Dec) error {
	for denom, cs := range calSupply {
		if s, ok := supply[denom]; ok {
			if s.Cmp(cs) != math.EqualTo {
				return sdkerrors.ErrInvalidCoins.Wrapf("supply is incorrect for %d credit batch, expected %v, got %v", denom, s, cs)
			}
		} else {
			return sdkerrors.ErrNotFound.Wrapf("supply is not found for %d credit batch", denom)
		}
	}

	return nil
}

// MergeParamsIntoTarget merges params message into the ormjson.WriteTarget.
func MergeParamsIntoTarget(cdc codec.JSONCodec, message proto.Message, target ormjson.WriteTarget) error {
	w, err := target.OpenWriter(protoreflect.FullName(proto.MessageName(message)))
	if err != nil {
		return err
	}

	bz, err := cdc.MarshalJSON(message)
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	if err != nil {
		return err
	}

	return w.Close()
}
