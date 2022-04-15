package core

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogoproto "github.com/gogo/protobuf/proto"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit"
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
func ValidateGenesis(data json.RawMessage, params Params) error {
	if err := params.Validate(); err != nil {
		return err
	}

	db := dbm.NewMemDB()
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{
		JSONValidator: func(m proto.Message) error {
			return validateMsg(m)
		},
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

	// calculate credit batch supply from genesis tradable, retired and escrowed balances
	if err := calculateSupply(ormCtx, batchIdToPrecision, ss, batchIdToCalSupply); err != nil {
		return err
	}

	// verify calculated total amount of each credit batch matches the total supply
	if err := validateSupply(batchIdToCalSupply, batchIdToSupply); err != nil {
		return err
	}

	return nil
}

func validateMsg(m proto.Message) error {
	switch m.(type) {
	case *api.ClassInfo:
		msg := &ClassInfo{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}

		return msg.Validate()
	case *api.ClassIssuer:
		msg := &ClassIssuer{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}

		return msg.Validate()
	case *api.ProjectInfo:
		msg := &ProjectInfo{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}

		return msg.Validate()
	case *api.BatchInfo:
		msg := &BatchInfo{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	case *api.CreditType:
		msg := &CreditType{}
		if err := ormutil.PulsarToGogoSlow(m, msg); err != nil {
			return err
		}
		return msg.Validate()
	}

	return nil
}

func calculateSupply(ctx context.Context, batchIdToPrecision map[uint64]uint32, ss api.StateStore, batchIdToSupply map[uint64]math.Dec) error {
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

		if _, ok := batchIdToPrecision[balance.BatchId]; !ok {
			return sdkerrors.ErrInvalidType.Wrapf("credit type not exist for %d batch", balance.BatchId)
		}

		if balance.Tradable != "" {
			tradable, err = math.NewNonNegativeFixedDecFromString(balance.Tradable, batchIdToPrecision[balance.BatchId])
			if err != nil {
				return err
			}
		}

		if balance.Retired != "" {
			retired, err = math.NewNonNegativeFixedDecFromString(balance.Retired, batchIdToPrecision[balance.BatchId])
			if err != nil {
				return err
			}
		}

		if balance.Escrowed != "" {
			escrowed, err = math.NewNonNegativeFixedDecFromString(balance.Retired, batchIdToPrecision[balance.BatchId])
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

		if supply, ok := batchIdToSupply[balance.BatchId]; ok {
			result, err := math.SafeAddBalance(supply, total)
			if err != nil {
				return err
			}
			batchIdToSupply[balance.BatchId] = result
		} else {
			batchIdToSupply[balance.BatchId] = total
		}
	}

	return nil
}

func validateSupply(batchIdToSupplyCal, batchIdToSupply map[uint64]math.Dec) error {
	for denom, cs := range batchIdToSupplyCal {
		if s, ok := batchIdToSupply[denom]; ok {
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

// Validate performs a basic validation of credit class
func (c ClassInfo) Validate() error {
	if len(c.Metadata) > ecocredit.MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrap("credit class metadata")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(c.Admin).String()); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if len(c.Name) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class name cannot be empty")
	}

	if err := ecocredit.ValidateClassID(c.Name); err != nil {
		return err
	}

	if len(c.CreditType) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("must specify a credit type abbreviation")
	}

	return nil
}

// Validate performs a basic validation of credit class issuers
func (c ClassIssuer) Validate() error {
	if c.ClassId == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class id cannot be zero")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(c.Issuer).String()); err != nil {
		return sdkerrors.Wrap(err, "issuer")
	}

	return nil
}

// Validate performs a basic validation of project
func (p ProjectInfo) Validate() error {
	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(p.Admin).String()); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if p.ClassId == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class id cannot be zero")
	}

	if err := ValidateLocation(p.ProjectLocation); err != nil {
		return err
	}

	if len(p.Metadata) > ecocredit.MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrap("project metadata")
	}

	if err := ValidateProjectID(p.Name); err != nil {
		return err
	}

	return nil
}

// Validate performs a basic validation of credit batch
func (b BatchInfo) Validate() error {
	if err := ValidateDenom(b.BatchDenom); err != nil {
		return err
	}

	if b.ProjectId == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("project id cannot be zero")
	}

	if b.StartDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide a start date for the credit batch")
	}
	if b.EndDate == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("must provide an end date for the credit batch")
	}
	if b.EndDate.Compare(*b.StartDate) != 1 {
		return sdkerrors.ErrInvalidRequest.Wrapf("the batch end date (%s) must be the same as or after the batch start date (%s)", b.EndDate.String(), b.StartDate.String())
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(b.Issuer).String()); err != nil {
		return sdkerrors.Wrap(err, "issuer")
	}

	return nil
}

// Validate performs a basic validation of credit type
func (c CreditType) Validate() error {
	if err := ValidateCreditTypeAbbreviation(c.Abbreviation); err != nil {
		return err
	}
	if len(c.Name) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("name cannot be empty")
	}
	if len(c.Unit) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("unit cannot be empty")
	}
	if c.Precision != PRECISION {
		return sdkerrors.ErrInvalidRequest.Wrapf("credit type precision is currently locked to %d", PRECISION)
	}

	return nil
}
