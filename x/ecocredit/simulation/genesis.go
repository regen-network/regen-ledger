package simulation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/regen-network/regen-ledger/types/math"
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
	return sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(simtypes.RandIntBetween(r, 1, 100)))))
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

func genClasses(r *rand.Rand, accounts []simtypes.Account, creditTypes []*ecocredit.CreditType) []*ecocredit.ClassInfo {
	classes := make([]*ecocredit.ClassInfo, 3)

	for i := 1; i < 4; i++ {
		creditType := creditTypes[r.Intn(len(creditTypes))]
		classes[i-1] = &ecocredit.ClassInfo{
			ClassId:    ecocredit.FormatClassID(*creditType, uint64(i)),
			Admin:      accounts[0].Address.String(),
			Issuers:    []string{accounts[0].Address.String(), accounts[1].Address.String(), accounts[2].Address.String()},
			Metadata:   []byte(simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, 100))),
			CreditType: creditType,
		}
	}
	return classes
}

func genBatches(r *rand.Rand, classes []*ecocredit.ClassInfo,
	creditTypes []*ecocredit.CreditType) []*ecocredit.BatchInfo {
	batches := make([]*ecocredit.BatchInfo, 3)
	accounts := simtypes.RandomAccounts(r, 3)

	for i := 1; i < 4; i++ {
		startTime := simtypes.RandTimestamp(r)
		endTime := startTime.Add(24 * time.Hour)
		creditType := classes[i-1].CreditType
		classID := ecocredit.FormatClassID(*creditType, uint64(i))
		bd, _ := ecocredit.FormatDenom(classID, uint64(i), &startTime, &endTime)

		batches[i-1] = &ecocredit.BatchInfo{
			ClassId:         classID,
			BatchDenom:      bd,
			TotalAmount:     fmt.Sprintf("%d", simtypes.RandIntBetween(r, 500, 100000)),
			Metadata:        []byte(simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, 100))),
			AmountCancelled: fmt.Sprintf("%d", simtypes.RandIntBetween(r, 1, 50)),
			StartDate:       &startTime,
			EndDate:         &endTime,
			Issuer:          accounts[i-1].Address.String(),
			ProjectLocation: "AB-CDE FG1 345",
		}
	}

	return batches
}

func genBalances(r *rand.Rand,
	classes []*ecocredit.ClassInfo,
	batches []*ecocredit.BatchInfo,
	creditTypes []*ecocredit.CreditType) []*ecocredit.Balance {
	var result []*ecocredit.Balance
	accounts := simtypes.RandomAccounts(r, 4)

	for i := 0; i < 3; i++ {
		creditType := classes[i].CreditType
		classID := ecocredit.FormatClassID(*creditType, uint64(i+1))
		batch := getBatchbyClassID(batches, classID)
		bd, _ := ecocredit.FormatDenom(classID, uint64(i+1), batch.StartDate, batch.EndDate)
		balances := genRandomBalances(r, batch.TotalAmount, 4)
		result = append(result,
			&ecocredit.Balance{
				Address:         accounts[i].Address.String(),
				BatchDenom:      bd,
				TradableBalance: balances[0],
				RetiredBalance:  balances[1],
			},
			&ecocredit.Balance{
				Address:         accounts[i+1].Address.String(),
				BatchDenom:      bd,
				TradableBalance: balances[2],
				RetiredBalance:  balances[3],
			},
		)
	}

	return result
}

func getBatchbyClassID(batches []*ecocredit.BatchInfo, classID string) *ecocredit.BatchInfo {
	for _, batch := range batches {
		if batch.ClassId == classID {
			return batch
		}
	}

	return nil
}

func genSupplies(r *rand.Rand, classes []*ecocredit.ClassInfo,
	batches []*ecocredit.BatchInfo,
	balances []*ecocredit.Balance, creditTypes []*ecocredit.CreditType) []*ecocredit.Supply {
	supplies := make([]*ecocredit.Supply, 3)

	for i := 0; i < 3; i++ {
		creditType := classes[i].CreditType
		classID := ecocredit.FormatClassID(*creditType, uint64(i+1))
		batch := getBatchbyClassID(batches, classID)
		bd, _ := ecocredit.FormatDenom(classID, uint64(i+1), batch.StartDate, batch.EndDate)
		supply, _ := getBatchSupplyByDenom(balances, bd)
		supplies[i] = supply
	}

	return supplies
}

func getBatchSupplyByDenom(balances []*ecocredit.Balance, denom string) (*ecocredit.Supply, error) {
	tradableSupply := math.NewDecFromInt64(0)
	retiredSupply := math.NewDecFromInt64(0)

	for _, balance := range balances {
		if balance.BatchDenom == denom {
			tradable, err := math.NewDecFromString(balance.TradableBalance)
			if err != nil {
				return nil, err
			}

			retired, err := math.NewDecFromString(balance.RetiredBalance)
			if err != nil {
				return nil, err
			}

			tradableSupply, err = tradableSupply.Add(tradable)
			if err != nil {
				return nil, err
			}
			retiredSupply, err = retiredSupply.Add(retired)
			if err != nil {
				return nil, err
			}
		}
	}

	return &ecocredit.Supply{
		BatchDenom:     denom,
		TradableSupply: tradableSupply.String(),
		RetiredSupply:  retiredSupply.String(),
	}, nil

}

// RandomizedGenState generates a random GenesisState for the ecocredit module.
func RandomizedGenState(simState *module.SimulationState) {
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
		func(r *rand.Rand) { classes = genClasses(r, simState.Accounts, creditTypes) },
	)

	// batches
	var batches []*ecocredit.BatchInfo
	simState.AppParams.GetOrGenerate(
		simState.Cdc, batch, &batches, simState.Rand,
		func(r *rand.Rand) { batches = genBatches(r, classes, creditTypes) },
	)

	// balances
	var balances []*ecocredit.Balance
	simState.AppParams.GetOrGenerate(
		simState.Cdc, balance, &balances, simState.Rand,
		func(r *rand.Rand) { balances = genBalances(r, classes, batches, creditTypes) },
	)

	// supplies
	var supplies []*ecocredit.Supply
	simState.AppParams.GetOrGenerate(
		simState.Cdc, supply, &supplies, simState.Rand,
		func(r *rand.Rand) { supplies = genSupplies(r, classes, batches, balances, creditTypes) },
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
			{
				Abbreviation: "BIO",
				SeqNumber:    4,
			},
		},
	}

	bz := simState.Cdc.MustMarshalJSON(&ecocreditGenesis)
	var out bytes.Buffer
	if err := json.Indent(&out, bz, "", " "); err != nil {
		panic(err)
	}

	fmt.Printf("Selected randomly generated ecocredit parameters:\n%s\n", out.String())

	simState.GenState[ecocredit.ModuleName] = simState.Cdc.MustMarshalJSON(&ecocreditGenesis)
}

func genRandomBalances(r *rand.Rand, total string, n int) []string {
	res := make([]string, n)
	sum, _ := strconv.Atoi(total)

	for i := 0; i < n; i++ {
		j := simtypes.RandIntBetween(r, 0, sum)
		res[i] = fmt.Sprintf("%d", j)
		sum -= j
	}

	return res
}
