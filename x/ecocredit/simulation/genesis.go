package simulation

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	dbm "github.com/tendermint/tm-db"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// Simulation parameter constants
const (
	classFee             = "credit_class_fee"
	allowedCreators      = "allowed_class_creators"
	typeAllowListEnabled = "allow_list_enabled"
	typeCreditTypes      = "credit_types"
)

// genCreditClassFee randomized CreditClassFee
func genCreditClassFee(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simtypes.RandIntBetween(r, 1, 10)))))
}

// genAllowedClassCreators generate random set of creators
func genAllowedClassCreators(r *rand.Rand, accs []simtypes.Account) []string {
	n := simtypes.RandIntBetween(r, 1, len(accs))
	creators := make([]string, n)

	for i := 0; i < n; i++ {
		creators[i] = accs[i].Address.String()
	}

	return creators
}

func genAllowListEnabled(r *rand.Rand) bool {
	return r.Int63n(101) <= 90
}

func genCreditTypes(r *rand.Rand) []*core.CreditType {
	return []*core.CreditType{
		{
			Name:         "carbon",
			Abbreviation: "C",
			Unit:         "metric ton CO2 equivalent",
			Precision:    6,
		},
		{
			Name:         "biodiversity",
			Abbreviation: "BIO",
			Unit:         "ton",
			Precision:    6, // TODO: randomize precision, precision is currently locked to 6
		},
	}
}

// RandomizedGenState generates a random GenesisState for the ecocredit module.
func RandomizedGenState(simState *module.SimulationState) {
	// params
	var (
		creditClassFee       sdk.Coins
		allowedClassCreators []string
		allowListEnabled     bool
		creditTypes          []*core.CreditType
		basketCreationFee    sdk.Coins
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, classFee, &creditClassFee, simState.Rand,
		func(r *rand.Rand) { creditClassFee = genCreditClassFee(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, typeAllowListEnabled, &allowListEnabled, simState.Rand,
		func(r *rand.Rand) { allowListEnabled = genAllowListEnabled(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, allowedCreators, &allowedClassCreators, simState.Rand,
		func(r *rand.Rand) {
			if allowListEnabled {
				allowedClassCreators = genAllowedClassCreators(r, simState.Accounts)
			} else {
				allowedClassCreators = []string{}
			}
		},
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, typeCreditTypes, &creditTypes, simState.Rand,
		func(r *rand.Rand) { creditTypes = genCreditTypes(r) },
	)

	params := &core.Params{
		CreditClassFee:       creditClassFee,
		AllowedClassCreators: allowedClassCreators,
		AllowlistEnabled:     allowListEnabled,
		CreditTypes:          creditTypes,
		BasketFee:            basketCreationFee,
	}

	db := dbm.NewMemDB()
	backend := ormtable.NewBackend(ormtable.BackendOptions{
		CommitmentStore: db,
		IndexStore:      db,
	})

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	if err != nil {
		panic(err)
	}

	ormCtx := ormtable.WrapContextDefault(backend)
	ss, err := api.NewStateStore(ormdb)
	if err != nil {
		panic(err)
	}

	simState.AppParams.GetOrGenerate(
		simState.Cdc, typeCreditTypes, &creditTypes, simState.Rand,
		func(r *rand.Rand) {
			genGenesisState(ormCtx, r, simState, ss)
		})

	paramsBz := simState.Cdc.MustMarshalJSON(params)
	var out bytes.Buffer
	if err := json.Indent(&out, paramsBz, "", " "); err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated ecocredit parameters:\n%s\n", out.String())

	jsonTarget := ormjson.NewRawMessageTarget()
	if err := ormdb.ExportJSON(ormCtx, jsonTarget); err != nil {
		panic(err)
	}

	if err := core.MergeParamsIntoTarget(simState.Cdc, params, jsonTarget); err != nil {
		panic(err)
	}

	rawJson, err := jsonTarget.JSON()
	if err != nil {
		panic(err)
	}

	bz, err := json.Marshal(rawJson)
	if err != nil {
		panic(err)
	}

	simState.GenState[ecocredit.ModuleName] = bz
}

func genGenesisState(ctx context.Context, r *rand.Rand, simState *module.SimulationState, ss api.StateStore) error {
	accs := simState.Accounts
	metadata := simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 5, 100))

	// create few classes
	cID1, err := ss.ClassInfoTable().InsertReturningID(ctx,
		&api.ClassInfo{Name: "C01", Admin: accs[0].Address.Bytes(), Metadata: simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 5, 100))},
	)
	if err != nil {
		return err
	}

	cID2, err := ss.ClassInfoTable().InsertReturningID(ctx,
		&api.ClassInfo{Name: "C02", Admin: accs[1].Address.Bytes(), Metadata: metadata},
	)
	if err != nil {
		return err
	}

	// create class issuers
	if err := ss.ClassIssuerTable().Save(ctx,
		&api.ClassIssuer{ClassId: cID1, Issuer: accs[0].Address.Bytes()},
	); err != nil {
		return err
	}

	if err := ss.ClassIssuerTable().Save(ctx,
		&api.ClassIssuer{ClassId: cID1, Issuer: accs[1].Address.Bytes()},
	); err != nil {
		return err
	}

	if err := ss.ClassIssuerTable().Save(ctx,
		&api.ClassIssuer{ClassId: cID2, Issuer: accs[1].Address.Bytes()},
	); err != nil {
		return err
	}

	if err := ss.ClassIssuerTable().Save(ctx,
		&api.ClassIssuer{ClassId: cID2, Issuer: accs[2].Address.Bytes()},
	); err != nil {
		return err
	}

	// create few projects
	pID1, err := ss.ProjectInfoTable().InsertReturningID(ctx,
		&api.ProjectInfo{ClassId: cID1, Name: "Project1", Admin: accs[0].Address.Bytes(), ProjectLocation: "AQ", Metadata: metadata},
	)
	if err != nil {
		return err
	}

	pID2, err := ss.ProjectInfoTable().InsertReturningID(ctx,
		&api.ProjectInfo{ClassId: cID2, Admin: accs[1].Address.Bytes(), ProjectLocation: "AQ", Metadata: metadata},
	)
	if err != nil {
		return err
	}

	// create few batches
	startDate := simState.GenTimestamp
	endDate := simState.GenTimestamp.AddDate(0, 1, 0)
	denom, err := ecocredit.FormatDenom("C001", 1, &startDate, &endDate)
	if err != nil {
		return err
	}

	bID1, err := ss.BatchInfoTable().InsertReturningID(ctx,
		&api.BatchInfo{
			Issuer: accs[0].Address.Bytes(), ProjectId: pID1, BatchDenom: denom,
			StartDate: timestamppb.New(startDate.UTC()), EndDate: timestamppb.New(endDate.UTC()),
			Metadata: metadata, IssuanceDate: timestamppb.New(simtypes.RandTimestamp(r).UTC()),
		},
	)
	if err != nil {
		return err
	}

	denom, err = ecocredit.FormatDenom("C002", 2, &startDate, &endDate)
	if err != nil {
		return err
	}

	bID2, err := ss.BatchInfoTable().InsertReturningID(ctx,
		&api.BatchInfo{
			Issuer: accs[1].Address.Bytes(), ProjectId: pID1, BatchDenom: denom,
			StartDate: timestamppb.New(startDate.UTC()), EndDate: timestamppb.New(endDate.UTC()),
			Metadata: metadata, IssuanceDate: timestamppb.New(simtypes.RandTimestamp(r).UTC()),
		},
	)
	if err != nil {
		return err
	}

	denom, err = ecocredit.FormatDenom("C003", 3, &startDate, &endDate)
	if err != nil {
		return err
	}

	bID3, err := ss.BatchInfoTable().InsertReturningID(ctx,
		&api.BatchInfo{
			Issuer: accs[2].Address.Bytes(), ProjectId: pID2, BatchDenom: denom,
			StartDate: timestamppb.New(startDate.UTC()), EndDate: timestamppb.New(endDate.UTC()),
			Metadata: metadata, IssuanceDate: timestamppb.New(simtypes.RandTimestamp(r).UTC()),
		},
	)
	if err != nil {
		return err
	}

	// batch balances
	if err := ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
		Address:  accs[0].Address,
		BatchId:  bID1,
		Tradable: "100",
		Retired:  "10",
	}); err != nil {
		return err
	}

	if err := ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
		Address:  accs[1].Address,
		BatchId:  bID2,
		Tradable: "100",
		Retired:  "10",
	}); err != nil {
		return err
	}

	if err := ss.BatchBalanceTable().Save(ctx, &api.BatchBalance{
		Address:  accs[2].Address,
		BatchId:  bID3,
		Tradable: "100",
		Retired:  "10",
	}); err != nil {
		return err
	}

	// add batch supply
	if err := ss.BatchSupplyTable().Save(ctx, &api.BatchSupply{
		BatchId:        bID1,
		TradableAmount: "100",
		RetiredAmount:  "10",
	}); err != nil {
		return err
	}

	if err := ss.BatchSupplyTable().Save(ctx, &api.BatchSupply{
		BatchId:        bID2,
		TradableAmount: "100",
		RetiredAmount:  "10",
	}); err != nil {
		return err
	}

	if err := ss.BatchSupplyTable().Save(ctx, &api.BatchSupply{
		BatchId:        bID3,
		TradableAmount: "100",
		RetiredAmount:  "10",
	}); err != nil {
		return err
	}

	return nil
}
