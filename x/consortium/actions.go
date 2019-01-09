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

/*
This action is a temporary placeholder for more
thorough validator delegation and staking support
 */
type ActionChangeValidatorSet struct {
	validators []abci.ValidatorUpdate
}

// TODO token inflation rate
// TODO token minting rewards

func (action ActionScheduleUpgrade) Route() string { return "consortium" }

func (action ActionScheduleUpgrade) Type() string { return "upgrade" }

func (action ActionScheduleUpgrade) ValidateBasic() sdk.Error {
	if action.upgradeInfo.Height <= 0 {
		return sdk.ErrUnknownRequest("Upgrade height must be greater than 0")
	}
	return nil
}

func (action ActionScheduleUpgrade) GetSignBytes() []byte {
	b, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (action ActionChangeValidatorSet) Route() string { return "consortium" }

func (action ActionChangeValidatorSet) Type() string { return "changeValidatorSet" }

func (action ActionChangeValidatorSet) ValidateBasic() sdk.Error {
	panic("implement me")
}

func (action ActionChangeValidatorSet) GetSignBytes() []byte {
	b, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}
