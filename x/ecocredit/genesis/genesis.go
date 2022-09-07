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
	baseapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
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
func ValidateGenesis(data json.RawMessage) error {
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
	ss, err := baseapi.NewStateStore(ormdb)
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
	ctItr, err := ss.CreditTypeTable().List(ormCtx, &baseapi.CreditTypePrimaryKey{})
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

	cItr, err := ss.ClassTable().List(ormCtx, baseapi.ClassPrimaryKey{})
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
	pItr, err := ss.ProjectTable().List(ormCtx, baseapi.ProjectPrimaryKey{})
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
	bItr, err := ss.BatchTable().List(ormCtx, baseapi.BatchPrimaryKey{})
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
	bsItr, err := ss.BatchSupplyTable().List(ormCtx, baseapi.BatchSupplyPrimaryKey{})
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

	// ecocredit base
	case *baseapi.CreditType:
		msg := &basetypes.CreditType{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.Class:
		msg := &basetypes.Class{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.ClassIssuer:
		msg := &basetypes.ClassIssuer{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.Project:
		msg := &basetypes.Project{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.Batch:
		msg := &basetypes.Batch{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.ClassSequence:
		msg := &basetypes.ClassSequence{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.ProjectSequence:
		msg := &basetypes.ProjectSequence{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.BatchSequence:
		msg := &basetypes.BatchSequence{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.BatchBalance:
		msg := &basetypes.BatchBalance{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.BatchSupply:
		msg := &basetypes.BatchSupply{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.OriginTxIndex:
		msg := &basetypes.OriginTxIndex{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.BatchContract:
		msg := &basetypes.BatchContract{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.ClassCreatorAllowlist:
		msg := &basetypes.ClassCreatorAllowlist{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.AllowedClassCreator:
		msg := &basetypes.AllowedClassCreator{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.ClassFee:
		msg := &basetypes.ClassFee{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *baseapi.AllowedBridgeChain:
		msg := &basetypes.AllowedBridgeChain{}
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
	case *basketapi.BasketFee:
		msg := &baskettypes.BasketFee{}
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

func calculateSupply(ctx context.Context, batchIDToPrecision map[uint64]uint32, ss baseapi.StateStore, batchIDToSupply map[uint64]math.Dec) error {
	bbItr, err := ss.BatchBalanceTable().List(ctx, baseapi.BatchBalancePrimaryKey{})
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

// MergeCreditTypesIntoTarget merges params message into the ormjson.WriteTarget.
func MergeCreditTypesIntoTarget(messages []basetypes.CreditType, target ormjson.WriteTarget) error {
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

// MergeClassFeeIntoTarget merges params message into the ormjson.WriteTarget.
func MergeClassFeeIntoTarget(
	cdc codec.JSONCodec,
	classFee basetypes.ClassFee,
	target ormjson.WriteTarget) error {
	w, err := target.OpenWriter(protoreflect.FullName(gogoproto.MessageName(&classFee)))
	if err != nil {
		return err
	}

	bz, err := cdc.MarshalJSON(&classFee)
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	if err != nil {
		return err
	}

	return w.Close()
}

// MergeBasketFeeIntoTarget merges params message into the ormjson.WriteTarget.
func MergeBasketFeeIntoTarget(
	cdc codec.JSONCodec,
	basketFee baskettypes.BasketFee,
	target ormjson.WriteTarget) error {
	w, err := target.OpenWriter(protoreflect.FullName(gogoproto.MessageName(&basketFee)))
	if err != nil {
		return err
	}

	bz, err := cdc.MarshalJSON(&basketFee)
	if err != nil {
		return err
	}

	_, err = w.Write(bz)
	if err != nil {
		return err
	}

	return w.Close()
}
