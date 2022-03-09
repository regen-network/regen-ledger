package simulation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Simulation parameter constants
const (
	class                = "classes"
	project              = "projects"
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
			Precision:    6, // TODO: randomize precision, precision is currently locked to 6
		},
	}
}

func genClasses(r *rand.Rand, accounts []simtypes.Account, creditTypes []*ecocredit.CreditType) []*ecocredit.ClassInfo {
	classes := make([]*ecocredit.ClassInfo, 3)

	for i := 1; i < 4; i++ {
		creditType := creditTypes[r.Intn(len(creditTypes))]
		classes[i-1] = &ecocredit.ClassInfo{
			ClassId:    ecocredit.FormatClassID(creditType.Abbreviation, uint64(i)),
			Admin:      accounts[0].Address.String(),
			Issuers:    []string{accounts[0].Address.String(), accounts[1].Address.String(), accounts[2].Address.String()},
			Metadata:   []byte(simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, 100))),
			CreditType: creditType,
		}
	}
	return classes
}

func genProjects(r *rand.Rand, classes []*ecocredit.ClassInfo) []*ecocredit.ProjectInfo {
	projects := make([]*ecocredit.ProjectInfo, 3)

	for i := 0; i < 3; i++ {
		projects[i] = &ecocredit.ProjectInfo{
			ProjectId:       ecocredit.FormatProjectID(classes[i].ClassId, uint64(i)),
			ClassId:         classes[i].ClassId,
			Issuer:          classes[i].Issuers[0],
			ProjectLocation: "AB-CDE FG1 345",
			Metadata:        []byte(simtypes.RandStringOfLength(r, 10)),
		}
	}

	return projects
}

func genBatches(r *rand.Rand, projects []*ecocredit.ProjectInfo) []*ecocredit.BatchInfo {
	batches := make([]*ecocredit.BatchInfo, 3)

	for i := 1; i < 4; i++ {
		project := projects[i-1]
		startTime := simtypes.RandTimestamp(r)
		endTime := startTime.Add(24 * time.Hour)
		bd, _ := ecocredit.FormatDenom(project.ClassId, uint64(i), &startTime, &endTime)

		batches[i-1] = &ecocredit.BatchInfo{
			ProjectId:       project.ProjectId,
			BatchDenom:      bd,
			TotalAmount:     fmt.Sprintf("%d", simtypes.RandIntBetween(r, 500, 100000)),
			Metadata:        []byte(simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, 100))),
			AmountCancelled: fmt.Sprintf("%d", simtypes.RandIntBetween(r, 1, 50)),
			StartDate:       &startTime,
			EndDate:         &endTime,
		}
	}

	return batches
}

func genBalances(r *rand.Rand,
	projects []*ecocredit.ProjectInfo,
	batches []*ecocredit.BatchInfo,
	creditTypes []*ecocredit.CreditType) []*ecocredit.Balance {
	var result []*ecocredit.Balance
	accounts := simtypes.RandomAccounts(r, 4)

	for i := 0; i < 3; i++ {
		project := projects[i]
		batch := getBatchbyProjectID(batches, project.ProjectId)
		bd, _ := ecocredit.FormatDenom(project.ClassId, uint64(i+1), batch.StartDate, batch.EndDate)
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

func getBatchbyProjectID(batches []*ecocredit.BatchInfo, projectID string) *ecocredit.BatchInfo {
	for _, batch := range batches {
		if batch.ProjectId == projectID {
			return batch
		}
	}

	return nil
}

func genSupplies(r *rand.Rand, projects []*ecocredit.ProjectInfo,
	batches []*ecocredit.BatchInfo,
	balances []*ecocredit.Balance, creditTypes []*ecocredit.CreditType) []*ecocredit.Supply {
	supplies := make([]*ecocredit.Supply, 3)

	for i := 0; i < 3; i++ {
		project := projects[i]
		batch := getBatchbyProjectID(batches, project.ProjectId)
		bd, _ := ecocredit.FormatDenom(project.ClassId, uint64(i+1), batch.StartDate, batch.EndDate)
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

	// classes
	var classes []*ecocredit.ClassInfo
	simState.AppParams.GetOrGenerate(
		simState.Cdc, class, &classes, simState.Rand,
		func(r *rand.Rand) { classes = genClasses(r, simState.Accounts, creditTypes) },
	)

	// projects
	var projects []*ecocredit.ProjectInfo
	simState.AppParams.GetOrGenerate(
		simState.Cdc, project, &projects, simState.Rand,
		func(r *rand.Rand) { projects = genProjects(r, classes) },
	)

	// batches
	var batches []*ecocredit.BatchInfo
	simState.AppParams.GetOrGenerate(
		simState.Cdc, batch, &batches, simState.Rand,
		func(r *rand.Rand) { batches = genBatches(r, projects) },
	)

	// balances
	var balances []*ecocredit.Balance
	simState.AppParams.GetOrGenerate(
		simState.Cdc, balance, &balances, simState.Rand,
		func(r *rand.Rand) { balances = genBalances(r, projects, batches, creditTypes) },
	)

	// supplies
	var supplies []*ecocredit.Supply
	simState.AppParams.GetOrGenerate(
		simState.Cdc, supply, &supplies, simState.Rand,
		func(r *rand.Rand) { supplies = genSupplies(r, projects, batches, balances, creditTypes) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, "basket_fee", &basketCreationFee, simState.Rand,
		func(r *rand.Rand) {
			basketCreationFee = sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(simtypes.RandIntBetween(r, 0, 10))))
		},
	)

	ecocreditGenesis := ecocredit.GenesisState{
		Params: ecocredit.Params{
			CreditClassFee:       creditClassFee,
			AllowedClassCreators: allowedClassCreators,
			AllowlistEnabled:     allowListEnabled,
			CreditTypes:          creditTypes,
			BasketCreationFee:    basketCreationFee,
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

	wrapper := map[string]json.RawMessage{
		proto.MessageName(&ecocreditGenesis): bz,
	}

	bz, err := json.Marshal(wrapper)
	if err != nil {
		panic(err)
	}

	simState.GenState[ecocredit.ModuleName] = bz
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
