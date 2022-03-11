package server

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// NOTE: currently we have ORM + non-ORM genesis in parallel. We will remove
// the non-ORM genesis soon, but for now, we merge both genesis JSON's into
// the same map.

// InitGenesis performs genesis initialization for the ecocredit module. It
// returns no validator updates.
func (s serverImpl) InitGenesis(ctx types.Context, cdc codec.Codec, data json.RawMessage) ([]abci.ValidatorUpdate, error) {
	jsonSource, err := ormjson.NewRawMessageSource(data)
	if err != nil {
		return nil, err
	}

	err = s.db.ImportJSON(ctx, jsonSource)
	if err != nil {
		return nil, err
	}

	var genesisState ecocredit.GenesisState
	r, err := jsonSource.OpenReader(protoreflect.FullName(proto.MessageName(&genesisState)))
	if err != nil {
		return nil, err
	}

	if r == nil { // r is nil when theres no table data, so we can just unmarshal the data given
		bz := bytes.NewBuffer(data)
		err = (&jsonpb.Unmarshaler{AllowUnknownFields: true}).Unmarshal(bz, &genesisState)
		if err != nil {
			return nil, err
		}
	} else { // r is not nil, so there is table data and we can just use r.
		err = (&jsonpb.Unmarshaler{AllowUnknownFields: true}).Unmarshal(r, &genesisState)
		if err != nil {
			return nil, err
		}
	}

	s.paramSpace.SetParamSet(ctx.Context, &genesisState.Params)

	if err := s.creditTypeSeqTable.Import(ctx, genesisState.Sequences, 0); err != nil {
		return nil, errors.Wrap(err, "sequences")
	}

	if err := s.classInfoTable.Import(ctx, genesisState.ClassInfo, 0); err != nil {
		return nil, errors.Wrap(err, "class-info")
	}

	if err := s.projectInfoTable.Import(ctx, genesisState.ProjectInfo, 0); err != nil {
		return nil, errors.Wrap(err, "project-info")
	}

	if err := s.projectInfoSeq.InitVal(ctx, genesisState.ProjectSeqNum); err != nil {
		return nil, errors.Wrap(err, "project seq")
	}

	if err := s.batchInfoTable.Import(ctx, genesisState.BatchInfo, 0); err != nil {
		return nil, errors.Wrap(err, "batch-info")
	}

	store := ctx.KVStore(s.storeKey)
	if err := setBalanceAndSupply(store, genesisState.Balances); err != nil {
		return nil, err
	}

	if err := validateSupplies(store, genesisState.Supplies); err != nil {
		return nil, err
	}

	return []abci.ValidatorUpdate{}, nil
}

// validateSupplies returns an error if credit batch genesis supply does not equal to calculated supply.
func validateSupplies(store sdk.KVStore, supplies []*ecocredit.Supply) error {
	var denomT ecocredit.BatchDenomT
	for _, supply := range supplies {
		denomT = ecocredit.BatchDenomT(supply.BatchDenom)
		tradableSupply := math.NewDecFromInt64(0)
		retiredSupply := math.NewDecFromInt64(0)
		var err error
		if supply.TradableSupply != "" {
			tradableSupply, err = math.NewNonNegativeDecFromString(supply.TradableSupply)
			if err != nil {
				return err
			}
		}

		tradable, err := ecocredit.GetDecimal(store, ecocredit.TradableSupplyKey(denomT))
		if err != nil {
			return err
		}

		if tradableSupply.Cmp(tradable) != 0 {
			return sdkerrors.ErrInvalidCoins.Wrapf("tradable supply is incorrect for %s credit batch, expected %v, got %v", supply.BatchDenom, tradable, tradableSupply)
		}

		if supply.RetiredSupply != "" {
			retiredSupply, err = math.NewNonNegativeDecFromString(supply.RetiredSupply)
			if err != nil {
				return err
			}
		}

		retired, err := ecocredit.GetDecimal(store, ecocredit.RetiredSupplyKey(denomT))
		if err != nil {
			return err
		}

		if retiredSupply.Cmp(retired) != 0 {
			return sdkerrors.ErrInvalidCoins.Wrapf("retired supply is incorrect for %s credit batch, expected %v, got %v", supply.BatchDenom, retired, retiredSupply)
		}
	}

	return nil
}

// setBalanceAndSupply sets the tradable and retired balance for an account and update supply for batch denom.
func setBalanceAndSupply(store sdk.KVStore, balances []*ecocredit.Balance) error {
	for _, balance := range balances {
		addr, err := sdk.AccAddressFromBech32(balance.Address)
		if err != nil {
			return err
		}
		denomT := ecocredit.BatchDenomT(balance.BatchDenom)

		// set tradable balance and update supply
		if balance.TradableBalance != "" {
			d, err := math.NewNonNegativeDecFromString(balance.TradableBalance)
			if err != nil {
				return err
			}
			key := ecocredit.TradableBalanceKey(addr, denomT)
			ecocredit.SetDecimal(store, key, d)

			key = ecocredit.TradableSupplyKey(denomT)
			ecocredit.AddAndSetDecimal(store, key, d)
		}

		// set retired balance and update supply
		if balance.RetiredBalance != "" {
			d, err := math.NewNonNegativeDecFromString(balance.RetiredBalance)
			if err != nil {
				return err
			}
			key := ecocredit.RetiredBalanceKey(addr, denomT)
			ecocredit.SetDecimal(store, key, d)

			key = ecocredit.RetiredSupplyKey(denomT)
			ecocredit.AddAndSetDecimal(store, key, d)
		}
	}

	return nil
}

// ExportGenesis will dump the ecocredit module state into a serializable GenesisState.
func (s serverImpl) ExportGenesis(ctx types.Context, cdc codec.Codec) (json.RawMessage, error) {
	// Get Params from the store and put them in the genesis state
	var params ecocredit.Params
	s.paramSpace.GetParamSet(ctx.Context, &params)

	store := ctx.KVStore(s.storeKey)
	var classInfo []*ecocredit.ClassInfo
	if _, err := s.classInfoTable.Export(ctx, &classInfo); err != nil {
		return nil, errors.Wrap(err, "class-info")
	}

	var projectInfo []*ecocredit.ProjectInfo
	if _, err := s.projectInfoTable.Export(ctx, &projectInfo); err != nil {
		return nil, errors.Wrap(err, "project-info")
	}

	var batchInfo []*ecocredit.BatchInfo
	if _, err := s.batchInfoTable.Export(ctx, &batchInfo); err != nil {
		return nil, errors.Wrap(err, "batch-info")
	}

	var sequences []*ecocredit.CreditTypeSeq
	if _, err := s.creditTypeSeqTable.Export(ctx, &sequences); err != nil {
		return nil, errors.Wrap(err, "batch-info")
	}

	suppliesMap := make(map[string]*ecocredit.Supply)
	ecocredit.IterateSupplies(store, ecocredit.TradableSupplyPrefix, func(denom, supply string) (bool, error) {
		suppliesMap[denom] = &ecocredit.Supply{
			BatchDenom:     denom,
			TradableSupply: supply,
		}

		return false, nil
	})

	ecocredit.IterateSupplies(store, ecocredit.RetiredSupplyPrefix, func(denom, supply string) (bool, error) {
		if _, exists := suppliesMap[denom]; exists {
			suppliesMap[denom].RetiredSupply = supply
		} else {
			suppliesMap[denom] = &ecocredit.Supply{
				BatchDenom:    denom,
				RetiredSupply: supply,
			}
		}

		return false, nil
	})

	supplies := make([]*ecocredit.Supply, len(suppliesMap))
	index := 0
	for _, supply := range suppliesMap {
		supplies[index] = supply
		index++
	}

	balancesMap := make(map[string]*ecocredit.Balance)
	ecocredit.IterateBalances(store, ecocredit.TradableBalancePrefix, func(address, denom, balance string) bool {
		balancesMap[fmt.Sprintf("%s%s", address, denom)] = &ecocredit.Balance{
			Address:         address,
			BatchDenom:      denom,
			TradableBalance: balance,
		}

		return false
	})

	ecocredit.IterateBalances(store, ecocredit.RetiredBalancePrefix, func(address, denom, balance string) bool {
		index := fmt.Sprintf("%s%s", address, denom)
		if _, exists := balancesMap[index]; exists {
			balancesMap[index].RetiredBalance = balance
		} else {
			balancesMap[index] = &ecocredit.Balance{
				Address:        address,
				BatchDenom:     denom,
				RetiredBalance: balance,
			}
		}

		return false
	})

	balances := make([]*ecocredit.Balance, len(balancesMap))
	index = 0
	for _, balance := range balancesMap {
		balances[index] = balance
		index++
	}

	gs := &ecocredit.GenesisState{
		Params:      params,
		ClassInfo:   classInfo,
		BatchInfo:   batchInfo,
		Sequences:   sequences,
		Balances:    balances,
		Supplies:    supplies,
		ProjectInfo: projectInfo,
	}

	gs.ProjectSeqNum = s.projectInfoSeq.CurVal(ctx)

	jsonTarget := ormjson.NewRawMessageTarget()
	err := s.db.ExportJSON(ctx, jsonTarget)
	if err != nil {
		return nil, err
	}

	err = MergeLegacyJSONIntoTarget(cdc, gs, jsonTarget)
	if err != nil {
		return nil, err
	}

	return jsonTarget.JSON()
}

// MergeLegacyJSONIntoTarget merges legacy genesis JSON in message into the
// ormjson.WriteTarget under key which has the name of the legacy message.
func MergeLegacyJSONIntoTarget(cdc codec.JSONCodec, message proto.Message, target ormjson.WriteTarget) error {
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
