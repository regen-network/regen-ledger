package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Simulation parameter constants
const (
	class                = "classes"
	batch                = "batches"
	balance              = "balances"
	supply               = "supplies"
	classFee             = "credit_class_fee"
	allowedCreators      = "allowed_class_creators"
	typeAllowListEnabled = "allow_list_enabled"
	typeCreditTypes      = "credit_types"
)

// genCreditClassFee randomized CreditClassFee
func genCreditClassFee(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1)))
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

func genCreditTypes(r *rand.Rand) []*ecocredit.CreditType {
	return []*ecocredit.CreditType{
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
			Precision:    6,
		},
	}
}

func genClasses(r *rand.Rand, accounts []simtypes.Account) []*ecocredit.ClassInfo {
	classes := make([]*ecocredit.ClassInfo, 3)
	creditType := ecocredit.DefaultParams().CreditTypes[0]

	for i := 1; i < 4; i++ {
		classes[i-1] = &ecocredit.ClassInfo{
			ClassId:    ecocredit.FormatClassID(*creditType, uint64(i)),
			Admin:      accounts[0].Address.String(),
			Issuers:    []string{accounts[0].Address.String(), accounts[1].Address.String(), accounts[2].Address.String()},
			Metadata:   []byte(simtypes.RandStringOfLength(r, 10)),
			CreditType: ecocredit.DefaultParams().CreditTypes[0],
		}
	}
	return classes
}

func genBatches(r *rand.Rand, startTime, endTime time.Time) []*ecocredit.BatchInfo {
	batches := make([]*ecocredit.BatchInfo, 3)
	accounts := simtypes.RandomAccounts(r, 3)
	creditType := ecocredit.DefaultParams().CreditTypes[0]

	for i := 1; i < 4; i++ {
		classID := ecocredit.FormatClassID(*creditType, uint64(i))
		bd, _ := ecocredit.FormatDenom(classID, uint64(i), &startTime, &endTime)
		batches[i-1] = &ecocredit.BatchInfo{
			ClassId:         classID,
			BatchDenom:      bd,
			TotalAmount:     "100000",
			Metadata:        []byte(simtypes.RandStringOfLength(r, 10)),
			AmountCancelled: "100",
			StartDate:       &startTime,
			EndDate:         &endTime,
			Issuer:          accounts[i-1].Address.String(),
			ProjectLocation: "AB-CDE FG1 345",
		}
	}

	return batches
}

func genBalances(r *rand.Rand, startTime, endTime time.Time) []*ecocredit.Balance {
	var balances []*ecocredit.Balance
	accounts := simtypes.RandomAccounts(r, 4)
	creditType := ecocredit.DefaultParams().CreditTypes[0]

	for i := 0; i < 3; i++ {
		classID := ecocredit.FormatClassID(*creditType, uint64(i+1))
		bd, _ := ecocredit.FormatDenom(classID, uint64(i+1), &startTime, &endTime)
		balances = append(balances,
			&ecocredit.Balance{
				Address:         accounts[i].Address.String(),
				BatchDenom:      bd,
				TradableBalance: "987.123",
				RetiredBalance:  "123.123",
			},
			&ecocredit.Balance{
				Address:         accounts[i+1].Address.String(),
				BatchDenom:      bd,
				TradableBalance: "12.988",
				RetiredBalance:  "876.988",
			},
		)
	}

	return balances
}

func genSupplies(r *rand.Rand, startTime, endTime time.Time) []*ecocredit.Supply {
	supplies := make([]*ecocredit.Supply, 3)
	creditType := ecocredit.DefaultParams().CreditTypes[0]

	for i := 0; i < 3; i++ {
		classID := ecocredit.FormatClassID(*creditType, uint64(i+1))
		bd, _ := ecocredit.FormatDenom(classID, uint64(i+1), &startTime, &endTime)
		supplies[i] = &ecocredit.Supply{
			BatchDenom:     bd,
			TradableSupply: "1000.111",
			RetiredSupply:  "1000.111",
		}
	}

	return supplies
}

// RandomizedGenState generates a random GenesisState for the ecocredit module.
func RandomizedGenState(simState *module.SimulationState) {
	startTime := simtypes.RandTimestamp(simState.Rand)
	endTime := startTime.Add(24 * time.Hour)

	// params
	var (
		creditClassFee       sdk.Coins
		allowedClassCreators []string
		allowListEnabled     bool
		creditTypes          []*ecocredit.CreditType
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

	// classes
	var classes []*ecocredit.ClassInfo
	simState.AppParams.GetOrGenerate(
		simState.Cdc, class, &classes, simState.Rand,
		func(r *rand.Rand) { classes = genClasses(r, simState.Accounts) },
	)

	// batches
	var batches []*ecocredit.BatchInfo
	simState.AppParams.GetOrGenerate(
		simState.Cdc, batch, &batches, simState.Rand,
		func(r *rand.Rand) { batches = genBatches(r, startTime, endTime) },
	)

	// balances
	var balances []*ecocredit.Balance
	simState.AppParams.GetOrGenerate(
		simState.Cdc, balance, &balances, simState.Rand,
		func(r *rand.Rand) { balances = genBalances(r, startTime, endTime) },
	)

	// supplies
	var supplies []*ecocredit.Supply
	simState.AppParams.GetOrGenerate(
		simState.Cdc, supply, &supplies, simState.Rand,
		func(r *rand.Rand) { supplies = genSupplies(r, startTime, endTime) },
	)

	ecocreditGenesis := ecocredit.GenesisState{
		Params: ecocredit.Params{
			CreditClassFee:       creditClassFee,
			AllowedClassCreators: allowedClassCreators,
			AllowlistEnabled:     allowListEnabled,
			CreditTypes:          creditTypes,
		},
		ClassInfo: classes,
		BatchInfo: batches,
		Balances:  balances,
		Supplies:  supplies,
		Sequences: []*ecocredit.CreditTypeSeq{
			{
				Abbreviation: "C",
				SeqNumber:    4,
			},
		},
	}

	bz, err := json.MarshalIndent(ecocreditGenesis.Params, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated ecocredit parameters:\n%s\n", bz)

	simState.GenState[ecocredit.ModuleName] = simState.Cdc.MustMarshalJSON(&ecocreditGenesis)
}
