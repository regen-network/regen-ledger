package consortium

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"gitlab.com/regen-network/regen-ledger/x/upgrade"
)

type ActionScheduleUpgrade struct {
	upgradeInfo upgrade.UpgradeInfo
}

type ActionChangeValidatorSet struct {
	validators []abci.ValidatorUpdate
}

// TODO token inflation rate
// TODO token minting rewards

func (action ActionScheduleUpgrade) Route() string {
	return "consortium"
}

func (action ActionScheduleUpgrade) Type() string {
    return "upgrade"
}

func (action ActionScheduleUpgrade) ValidateBasic() sdk.Error {
	return nil
}

func (action ActionScheduleUpgrade) GetSignBytes() []byte {
	b, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}



