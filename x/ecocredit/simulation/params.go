// DONTCOVER
package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	allowListEnabled := false

	return []simtypes.ParamChange{
		simulation.NewSimParamChange(ecocredit.ModuleName, string(ecocredit.KeyCreditClassFee),
			func(r *rand.Rand) string {
				bz, err := json.Marshal(genCreditClassFee(r))
				if err != nil {
					panic(err)
				}

				return string(bz)
			},
		),

		simulation.NewSimParamChange(ecocredit.ModuleName, string(ecocredit.KeyAllowlistEnabled),
			func(r *rand.Rand) string {
				allowListEnabled = genAllowListEnabled(r)
				return fmt.Sprintf("%v", allowListEnabled)
			},
		),

		simulation.NewSimParamChange(ecocredit.ModuleName, string(ecocredit.KeyAllowedClassDesigners),
			func(r *rand.Rand) string {
				if allowListEnabled {
					accs := simtypes.RandomAccounts(r, 10)
					bz, err := json.Marshal(genAllowedClassDesigners(r, accs))
					if err != nil {
						panic(err)
					}
					return string(bz)
				} else {
					return ""
				}
			},
		),

		simulation.NewSimParamChange(ecocredit.ModuleName, string(ecocredit.KeyCreditTypes),
			func(r *rand.Rand) string {
				bz, err := json.Marshal(genCreditTypes(r))
				if err != nil {
					panic(err)
				}
				return string(bz)
			},
		),
	}
}
