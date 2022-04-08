// DONTCOVER
package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	allowListEnabled := false

	return []simtypes.ParamChange{
		simulation.NewSimParamChange(ecocredit.ModuleName, string(core.KeyCreditClassFee),
			func(r *rand.Rand) string {
				bz, err := json.Marshal(genCreditClassFee(r))
				if err != nil {
					panic(err)
				}

				return string(bz)
			},
		),

		simulation.NewSimParamChange(ecocredit.ModuleName, string(core.KeyAllowlistEnabled),
			func(r *rand.Rand) string {
				allowListEnabled = genAllowListEnabled(r)
				return fmt.Sprintf("%v", allowListEnabled)
			},
		),

		simulation.NewSimParamChange(ecocredit.ModuleName, string(core.KeyAllowedClassCreators),
			func(r *rand.Rand) string {
				var bz []byte
				var err error
				if allowListEnabled {
					accs := simtypes.RandomAccounts(r, 10)
					bz, err = json.Marshal(genAllowedClassCreators(r, accs))
					if err != nil {
						panic(err)
					}
				} else {
					bz, err = json.Marshal([]string{})
					if err != nil {
						panic(err)
					}
				}

				return string(bz)
			},
		),

		simulation.NewSimParamChange(ecocredit.ModuleName, string(core.KeyCreditTypes),
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
