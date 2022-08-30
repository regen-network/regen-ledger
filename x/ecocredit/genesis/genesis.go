package genesis

import (
	"context"
	"encoding/json"
	"fmt"

	gogoproto "github.com/gogo/protobuf/proto"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	marketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

// ValidateGenesis performs basic validation for the following:
// - params are valid param types with valid properties
// - proto messages are valid proto messages
// - the credit type referenced in each credit class exists
// - the credit class referenced in each project exists
// - the tradable amount of each credit batch complies with the credit type precision
// - the retired amount of each credit batch complies with the credit type precision
// - the calculated total amount of each credit batch matches the total supply
// An error is returned if any of these validation checks fail.
func ValidateGenesis(data json.RawMessage, params core.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}

	db := dbm.NewMemDB()
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{
		JSONValidator: validateMsg,
	})
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

	if err := ormdb.ValidateJSON(jsonSource); err != nil {
		return err
	}

	abbrevToPrecision := make(map[string]uint32) // map of credit abbreviation to precision
	ctItr, err := ss.CreditTypeTable().List(ormCtx, &api.CreditTypePrimaryKey{})
	if err != nil {
		return err
	}
	for ctItr.Next() {
		ct, err := ctItr.Value()
		if err != nil {
			return err
		}
		abbrevToPrecision[ct.Abbreviation] = ct.Precision
	}
	ctItr.Close()

	cItr, err := ss.ClassTable().List(ormCtx, api.ClassPrimaryKey{})
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

		if _, ok := abbrevToPrecision[class.CreditTypeAbbrev]; !ok {
			return sdkerrors.ErrNotFound.Wrapf("credit type not exist for %s abbreviation", class.CreditTypeAbbrev)
		}
	}

	projectKeyToClassKey := make(map[uint64]uint64) // map of project key to class key
	pItr, err := ss.ProjectTable().List(ormCtx, api.ProjectPrimaryKey{})
	if err != nil {
		return err
	}
	defer pItr.Close()

	for pItr.Next() {
		project, err := pItr.Value()
		if err != nil {
			return err
		}

		if _, exists := projectKeyToClassKey[project.Key]; exists {
			continue
		}
		projectKeyToClassKey[project.Key] = project.ClassKey
	}

	batchIDToPrecision := make(map[uint64]uint32) // map of batchID to precision
	batchDenomToIDMap := make(map[string]uint64)  // map of batchDenom to batchID
	bItr, err := ss.BatchTable().List(ormCtx, api.BatchPrimaryKey{})
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

		batchDenomToIDMap[batch.Denom] = batch.Key

		if _, exists := batchIDToPrecision[batch.Key]; exists {
			continue
		}

		class, err := ss.ClassTable().Get(ormCtx, projectKeyToClassKey[batch.ProjectKey])
		if err != nil {
			return err
		}

		if class.Key == projectKeyToClassKey[batch.ProjectKey] {
			batchIDToPrecision[batch.Key] = abbrevToPrecision[class.CreditTypeAbbrev]
		}
	}

	batchIDToCalSupply := make(map[uint64]math.Dec) // map of batchID to calculated supply
	batchIDToSupply := make(map[uint64]math.Dec)    // map of batchID to actual supply
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
			tSupply, err = math.NewNonNegativeFixedDecFromString(batchSupply.TradableAmount, batchIDToPrecision[batchSupply.BatchKey])
			if err != nil {
				return err
			}
		}
		if batchSupply.RetiredAmount != "" {
			rSupply, err = math.NewNonNegativeFixedDecFromString(batchSupply.RetiredAmount, batchIDToPrecision[batchSupply.BatchKey])
			if err != nil {
				return err
			}
		}

		total, err := math.SafeAddBalance(tSupply, rSupply)
		if err != nil {
			return err
		}

		batchIDToSupply[batchSupply.BatchKey] = total
	}

	// calculate credit batch supply from genesis tradable, retired and escrowed balances
	if err := calculateSupply(ormCtx, batchIDToPrecision, ss, batchIDToCalSupply); err != nil {
		return err
	}

	basketStore, err := basketapi.NewStateStore(ormdb)
	if err != nil {
		return err
	}

	bBalanceItr, err := basketStore.BasketBalanceTable().List(ormCtx, basketapi.BasketBalancePrimaryKey{})
	if err != nil {
		return err
	}
	defer bBalanceItr.Close()

	for bBalanceItr.Next() {
		bBalance, err := bBalanceItr.Value()
		if err != nil {
			return err
		}
		batchID, ok := batchDenomToIDMap[bBalance.BatchDenom]
		if !ok {
			return fmt.Errorf("unknown credit batch %d in basket", batchID)
		}

		bb, err := math.NewNonNegativeDecFromString(bBalance.Balance)
		if err != nil {
			return err
		}

		if amount, ok := batchIDToCalSupply[batchID]; ok {
			result, err := math.SafeAddBalance(amount, bb)
			if err != nil {
				return err
			}
			batchIDToCalSupply[batchID] = result
		} else {
			return fmt.Errorf("unknown credit batch %d in basket", batchID)
		}
	}

	// verify calculated total amount of each credit batch matches the total supply
	if err := validateSupply(batchIDToCalSupply, batchIDToSupply); err != nil {
		return err
	}

	return nil
}

func validateMsg(m proto.Message) error {
	switch m.(type) {

	// ecocredit core
	case *api.CreditType:
		msg := &core.CreditType{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.Class:
		msg := &core.Class{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.ClassIssuer:
		msg := &core.ClassIssuer{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.Project:
		msg := &core.Project{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.Batch:
		msg := &core.Batch{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.ClassSequence:
		msg := &core.ClassSequence{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.ProjectSequence:
		msg := &core.ProjectSequence{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.BatchSequence:
		msg := &core.BatchSequence{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.BatchBalance:
		msg := &core.BatchBalance{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.BatchSupply:
		msg := &core.BatchSupply{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.OriginTxIndex:
		msg := &core.OriginTxIndex{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.BatchContract:
		msg := &core.BatchContract{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()

	// basket submodule
	case *basketapi.Basket:
		msg := &baskettypes.Basket{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *basketapi.BasketClass:
		msg := &baskettypes.BasketClass{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *basketapi.BasketBalance:
		msg := &baskettypes.BasketBalance{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()

	// marketplace submodule
	case *marketapi.SellOrder:
		msg := &markettypes.SellOrder{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *marketapi.AllowedDenom:
		msg := &markettypes.AllowedDenom{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *marketapi.Market:
		msg := &markettypes.Market{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	}

	return nil
}

func calculateSupply(ctx context.Context, batchIDToPrecision map[uint64]uint32, ss api.StateStore, batchIDToSupply map[uint64]math.Dec) error {
	bbItr, err := ss.BatchBalanceTable().List(ctx, api.BatchBalancePrimaryKey{})
	if err != nil {
		return err
	}
	defer bbItr.Close()

	for bbItr.Next() {
		tradable := math.NewDecFromInt64(0)
		retired := math.NewDecFromInt64(0)
		escrowed := math.NewDecFromInt64(0)

		balance, err := bbItr.Value()
		if err != nil {
			return err
		}

		precision, ok := batchIDToPrecision[balance.BatchKey]
		if !ok {
			return sdkerrors.ErrInvalidType.Wrapf("credit type not exist for %d batch", balance.BatchKey)
		}

		if balance.TradableAmount != "" {
			tradable, err = math.NewNonNegativeFixedDecFromString(balance.TradableAmount, precision)
			if err != nil {
				return err
			}
		}

		if balance.RetiredAmount != "" {
			retired, err = math.NewNonNegativeFixedDecFromString(balance.RetiredAmount, precision)
			if err != nil {
				return err
			}
		}

		if balance.EscrowedAmount != "" {
			escrowed, err = math.NewNonNegativeFixedDecFromString(balance.EscrowedAmount, precision)
			if err != nil {
				return err
			}
		}

		total, err := math.Add(tradable, retired)
		if err != nil {
			return err
		}

		total, err = math.Add(total, escrowed)
		if err != nil {
			return err
		}

		if supply, ok := batchIDToSupply[balance.BatchKey]; ok {
			result, err := math.SafeAddBalance(supply, total)
			if err != nil {
				return err
			}
			batchIDToSupply[balance.BatchKey] = result
		} else {
			batchIDToSupply[balance.BatchKey] = total
		}
	}

	return nil
}

func validateSupply(batchIDToSupplyCal, batchIDToSupply map[uint64]math.Dec) error {
	if len(batchIDToSupplyCal) == 0 && len(batchIDToSupply) > 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("batch supply was given but no balances were found")
	}
	if len(batchIDToSupply) == 0 && len(batchIDToSupplyCal) > 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("batch balances were given but no supplies were found")
	}
	for denom, cs := range batchIDToSupplyCal {
		if s, ok := batchIDToSupply[denom]; ok {
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
func MergeParamsIntoTarget(cdc codec.JSONCodec, message gogoproto.Message, target ormjson.WriteTarget) error {
	w, err := target.OpenWriter(protoreflect.FullName(gogoproto.MessageName(message)))
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

// MergeCreditTypesIntoTarget merges params message into the ormjson.WriteTarget.
func MergeCreditTypesIntoTarget(messages []core.CreditType, target ormjson.WriteTarget) error {
	w, err := target.OpenWriter(protoreflect.FullName(gogoproto.MessageName(&messages[0])))
	if err != nil {
		return err
	}

	// using json package because array is not a proto message
	bz, err := json.Marshal(messages)
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	if err != nil {
		return err
	}

	return w.Close()
}

// MergeAllowedDenomsIntoTarget merges params message into the ormjson.WriteTarget.
func MergeAllowedDenomsIntoTarget(messages []markettypes.AllowedDenom, target ormjson.WriteTarget) error {
	w, err := target.OpenWriter(protoreflect.FullName(gogoproto.MessageName(&messages[0])))
	if err != nil {
		return err
	}

	// using json package because array is not a proto message
	bz, err := json.Marshal(messages)
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	if err != nil {
		return err
	}

	return w.Close()
}

// MergeCreditClassFeesIntoTarget merges params message into the ormjson.WriteTarget.
func MergeCreditClassFeesIntoTarget(
	cdc codec.JSONCodec,
	creditClassFees core.ClassFees,
	target ormjson.WriteTarget) error {
	w, err := target.OpenWriter(protoreflect.FullName(gogoproto.MessageName(&creditClassFees)))
	if err != nil {
		return err
	}

	bz, err := cdc.MarshalJSON(&creditClassFees)
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	if err != nil {
		return err
	}

	return w.Close()
}

// MergeBasketFeesIntoTarget merges params message into the ormjson.WriteTarget.
func MergeBasketFeesIntoTarget(
	cdc codec.JSONCodec,
	basketFees baskettypes.BasketFees,
	target ormjson.WriteTarget) error {
	w, err := target.OpenWriter(protoreflect.FullName(gogoproto.MessageName(&basketFees)))
	if err != nil {
		return err
	}

	bz, err := cdc.MarshalJSON(&basketFees)
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	if err != nil {
		return err
	}

	return w.Close()
}
