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
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
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

func genAskDenoms() []*core.AskDenom {
	return []*core.AskDenom{
		{
			Denom:        "stake",
			DisplayDenom: "stake",
			Exponent:     18,
		},
	}
}

// genCreditClassFee randomized CreditClassFee
func genCreditClassFee(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simtypes.RandIntBetween(r, 1, 10)))))
}

// genAllowedClassCreators generate random set of creators
func genAllowedClassCreators(r *rand.Rand, accs []simtypes.Account) []string {
	max := 50

	switch len(accs) {
	case 0:
		return []string{}
	case 1:
		return []string{accs[0].Address.String()}
	default:
		if len(accs) < max {
			max = len(accs)
		}
	}
	n := simtypes.RandIntBetween(r, 1, max)
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
		BasketFee:            basketCreationFee,
		AllowedAskDenoms:     genAskDenoms(),
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
			if err := genGenesisState(ormCtx, r, simState, ss); err != nil {
				panic(err)
			}
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

func createClass(ctx context.Context, sStore api.StateStore, class *api.Class) (uint64, error) {
	cKey, err := sStore.ClassTable().InsertReturningID(ctx, class)
	if err != nil {
		return 0, err
	}

	seq, err := sStore.ClassSequenceTable().Get(ctx, class.CreditTypeAbbrev)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			if err := sStore.ClassSequenceTable().Insert(ctx, &api.ClassSequence{
				CreditTypeAbbrev: class.CreditTypeAbbrev,
				NextSequence:     2,
			}); err != nil {
				return 0, err
			}
			return cKey, nil
		}

		return 0, err
	} else {
		if err := sStore.ClassSequenceTable().Update(ctx, &api.ClassSequence{
			CreditTypeAbbrev: class.CreditTypeAbbrev,
			NextSequence:     seq.NextSequence + 1,
		}); err != nil {
			return 0, err
		}
	}

	return cKey, nil
}

func createProject(ctx context.Context, sStore api.StateStore, project *api.Project) (uint64, error) {
	pKey, err := sStore.ProjectTable().InsertReturningID(ctx, project)
	if err != nil {
		return 0, err
	}

	seq, err := sStore.ProjectSequenceTable().Get(ctx, project.ClassKey)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			if err := sStore.ProjectSequenceTable().Insert(ctx, &api.ProjectSequence{
				ClassKey:     project.ClassKey,
				NextSequence: 2,
			}); err != nil {
				return 0, err
			}
			return pKey, nil
		}

		return 0, err
	}

	if err := sStore.ProjectSequenceTable().Update(ctx, &api.ProjectSequence{
		ClassKey:     project.ClassKey,
		NextSequence: seq.NextSequence + 1,
	}); err != nil {
		return 0, err
	}

	return pKey, nil
}

func getBatchSequence(ctx context.Context, sStore api.StateStore, projectKey uint64) (uint64, error) {
	seq, err := sStore.BatchSequenceTable().Get(ctx, projectKey)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			if err := sStore.BatchSequenceTable().Insert(ctx, &api.BatchSequence{
				ProjectKey:   projectKey,
				NextSequence: 2,
			}); err != nil {
				return 0, err
			}
			return 1, nil
		}
		return 0, err
	}

	if err := sStore.BatchSequenceTable().Update(ctx, &api.BatchSequence{
		ProjectKey:   projectKey,
		NextSequence: seq.NextSequence + 1,
	}); err != nil {
		return 0, err
	}

	return seq.NextSequence, nil
}

func genGenesisState(ctx context.Context, r *rand.Rand, simState *module.SimulationState, ss api.StateStore) error {
	accs := simState.Accounts
	metadata := simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 5, 100))

	if err := ss.CreditTypeTable().Insert(ctx, &api.CreditType{
		Abbreviation: "C",
		Name:         "carbon",
		Unit:         "metric ton c02",
		Precision:    6,
	}); err != nil {
		return err
	}

	if err := ss.CreditTypeTable().Insert(ctx, &api.CreditType{
		Abbreviation: "BIO",
		Name:         "biodiversity",
		Unit:         "acres",
		Precision:    6,
	}); err != nil {
		return err
	}

	// create few classes
	cKey1, err := createClass(ctx, ss, &api.Class{
		Id:               "C01",
		Admin:            accs[0].Address,
		Metadata:         metadata,
		CreditTypeAbbrev: "C",
	})
	if err != nil {
		return err
	}

	cKey2, err := createClass(ctx, ss, &api.Class{
		Id:               "C02",
		Admin:            accs[1].Address,
		Metadata:         metadata,
		CreditTypeAbbrev: "C",
	})
	if err != nil {
		return err
	}

	// create class issuers
	if err := ss.ClassIssuerTable().Insert(ctx,
		&api.ClassIssuer{ClassKey: cKey1, Issuer: accs[0].Address},
	); err != nil {
		return err
	}

	if err := ss.ClassIssuerTable().Insert(ctx,
		&api.ClassIssuer{ClassKey: cKey1, Issuer: accs[1].Address},
	); err != nil {
		return err
	}

	if err := ss.ClassIssuerTable().Insert(ctx,
		&api.ClassIssuer{ClassKey: cKey2, Issuer: accs[1].Address},
	); err != nil {
		return err
	}

	if err := ss.ClassIssuerTable().Insert(ctx,
		&api.ClassIssuer{ClassKey: cKey2, Issuer: accs[2].Address},
	); err != nil {
		return err
	}

	// create few projects
	pKey1, err := createProject(ctx, ss, &api.Project{
		ClassKey:     cKey1,
		Id:           "P01",
		Admin:        accs[0].Address,
		Jurisdiction: "AQ",
		Metadata:     metadata,
	})
	if err != nil {
		return err
	}

	pKey2, err := createProject(ctx, ss, &api.Project{
		ClassKey:     cKey2,
		Id:           "P02",
		Admin:        accs[1].Address,
		Jurisdiction: "AQ",
		Metadata:     metadata,
	})
	if err != nil {
		return err
	}

	// create few batches
	startDate := simState.GenTimestamp.UTC()
	endDate := simState.GenTimestamp.AddDate(0, 1, 0).UTC()
	batchSeq, err := getBatchSequence(ctx, ss, pKey1)
	if err != nil {
		return err
	}
	denom, err := core.FormatDenom("C01", batchSeq, &startDate, &endDate)
	if err != nil {
		return err
	}

	bKey1, err := ss.BatchTable().InsertReturningID(ctx,
		&api.Batch{
			Issuer:       accs[0].Address,
			ProjectKey:   pKey1,
			Denom:        denom,
			StartDate:    timestamppb.New(startDate),
			EndDate:      timestamppb.New(endDate),
			Metadata:     metadata,
			IssuanceDate: timestamppb.New(simtypes.RandTimestamp(r).UTC()),
		},
	)
	if err != nil {
		return err
	}

	batchSeq, err = getBatchSequence(ctx, ss, pKey2)
	if err != nil {
		return err
	}
	denom, err = core.FormatDenom("C02", batchSeq, &startDate, &endDate)
	if err != nil {
		return err
	}

	bKey2, err := ss.BatchTable().InsertReturningID(ctx,
		&api.Batch{
			Issuer:       accs[2].Address,
			ProjectKey:   pKey2,
			Denom:        denom,
			StartDate:    timestamppb.New(startDate.UTC()),
			EndDate:      timestamppb.New(endDate.UTC()),
			Metadata:     metadata,
			IssuanceDate: timestamppb.New(simtypes.RandTimestamp(r).UTC()),
		},
	)
	if err != nil {
		return err
	}

	// batch balances
	if err := ss.BatchBalanceTable().Insert(ctx, &api.BatchBalance{
		BatchKey: bKey1,
		Address:  accs[0].Address,
		Tradable: "100",
		Retired:  "10",
	}); err != nil {
		return err
	}

	if err := ss.BatchBalanceTable().Insert(ctx, &api.BatchBalance{
		BatchKey: bKey2,
		Address:  accs[2].Address,
		Tradable: "100",
		Retired:  "10",
	}); err != nil {
		return err
	}

	if err := ss.BatchBalanceTable().Insert(ctx, &api.BatchBalance{
		BatchKey: bKey2,
		Address:  accs[1].Address,
		Tradable: "100",
		Retired:  "0",
	}); err != nil {
		return err
	}

	// add batch supply
	if err := ss.BatchSupplyTable().Insert(ctx, &api.BatchSupply{
		BatchKey:       bKey1,
		TradableAmount: "100",
		RetiredAmount:  "10",
	}); err != nil {
		return err
	}

	if err := ss.BatchSupplyTable().Insert(ctx, &api.BatchSupply{
		BatchKey:       bKey2,
		TradableAmount: "200",
		RetiredAmount:  "10",
	}); err != nil {
		return err
	}

	return nil
}
