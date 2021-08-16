package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Simulation parameter constants
const (
	class                = "classes"
	classFee             = "credit_class_fee"
	allowedDesigners     = "allowed_class_designers"
	typeAllowListEnabled = "allow_list_enabled"
	typeCreditTypes      = "credit_types"
)

// genCreditClassFee randomized CreditClassFee
func genCreditClassFee(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, simtypes.RandomAmount(r, sdk.NewInt(1000))))
}

// genAllowedClassDesigners generate random set of designers
func genAllowedClassDesigners(r *rand.Rand) []string {
	accounts := simtypes.RandomAccounts(r, 3)
	designers := make([]string, 3)
	for i, account := range accounts {
		designers[i] = account.Address.String()
	}

	return designers
}

func genAllowListEnabled(r *rand.Rand) bool {
	// 90% chance of credit creation being enable or P(a) = 0.9 for success
	return r.Int63n(101) <= 90
}

func genCreditTypes(r *rand.Rand) []*ecocredit.CreditType {
	return ecocredit.DefaultParams().CreditTypes
}

func genClasses(r *rand.Rand) []*ecocredit.ClassInfo {
	classes := make([]*ecocredit.ClassInfo, 3)
	accounts := simtypes.RandomAccounts(r, 3)
	for i := 0; i < 3; i++ {
		classes[i] = &ecocredit.ClassInfo{
			ClassId:    fmt.Sprintf("C%d", i),
			Designer:   accounts[0].Address.String(),
			Issuers:    []string{accounts[0].Address.String(), accounts[1].Address.String(), accounts[2].Address.String()},
			Metadata:   []byte(simtypes.RandStringOfLength(r, 10)),
			CreditType: ecocredit.DefaultParams().CreditTypes[0],
		}
	}
	return classes
}

// func genbatches(r *rand.Rand) []*ecocredit.BatchInfo {
// 	batches := make([]*ecocredit.BatchInfo, 3)
// 	accounts := simtypes.RandomAccounts(r, 3)
// 	for i := 0; i < 3; i++ {
// 		batches[i] = &ecocredit.BatchInfo{
// 			ClassId:    fmt.Sprintf("C%d", i),
			
// 			Designer:   accounts[0].Address.String(),
// 			Issuers:    []string{accounts[0].Address.String(), accounts[1].Address.String(), accounts[2].Address.String()},
// 			Metadata:   []byte(simtypes.RandStringOfLength(r, 10)),
// 			CreditType: ecocredit.DefaultParams().CreditTypes[0],
// 		}
// 	}
// 	return classes
// }

// RandomizedGenState generates a random GenesisState for the ecocredit module.
func RandomizedGenState(simState *module.SimulationState) {

	//params
	var (
		creditClassFee        sdk.Coins
		allowedClassDesigners []string
		allowListEnabled      bool
		creditTypes           []*ecocredit.CreditType
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, classFee, &creditClassFee, simState.Rand,
		func(r *rand.Rand) { creditClassFee = genCreditClassFee(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, allowedDesigners, &allowedClassDesigners, simState.Rand,
		func(r *rand.Rand) { allowedClassDesigners = genAllowedClassDesigners(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, typeAllowListEnabled, &allowListEnabled, simState.Rand,
		func(r *rand.Rand) { allowListEnabled = genAllowListEnabled(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, typeCreditTypes, &creditTypes, simState.Rand,
		func(r *rand.Rand) { creditTypes = genCreditTypes(r) },
	)

	params := ecocredit.Params{
		CreditClassFee:        creditClassFee,
		AllowedClassDesigners: allowedClassDesigners,
		AllowlistEnabled:      allowListEnabled,
		CreditTypes:           creditTypes,
	}

	// classes
	var classes []*ecocredit.ClassInfo
	simState.AppParams.GetOrGenerate(
		simState.Cdc, class, &classes, simState.Rand,
		func(r *rand.Rand) { classes = genClasses(r) },
	)

	//batches
	// var batches []*ecocredit.BatchInfo
	// simState.AppParams.GetOrGenerate(
	// 	simState.Cdc, batch, &batches, simState.Rand,
	// 	func(r *rand.Rand) { batches = genBatches(r) },
	// )

	ecocreditGenesis := ecocredit.GenesisState{
		Params: params,
		
	}

	bz, err := json.MarshalIndent(params, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated ecocredit parameters:\n%s\n", bz)

	simState.GenState[ecocredit.ModuleName] = simState.Cdc.MustMarshalJSON(&ecocreditGenesis)
}
