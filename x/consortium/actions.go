package consortium

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/upgrade"
	abci "github.com/tendermint/tendermint/abci/types"
)

type ActionScheduleUpgrade struct {
	Plan upgrade.UpgradePlan `json:"upgrade_plan"`
}

type ActionCancelUpgrade struct {
}

/*
This action is a temporary placeholder for more
thorough validator delegation and staking support
*/
type ActionChangeValidatorSet struct {
	Validators []abci.ValidatorUpdate `json:"validators"`
}

// TODO token inflation rate
// TODO token minting rewards

func (action ActionScheduleUpgrade) Route() string { return "consortium" }

func (action ActionScheduleUpgrade) Type() string { return "consortium.upgrade" }

func (action ActionScheduleUpgrade) ValidateBasic() sdk.Error {
	return action.Plan.ValidateBasic()
}

func (action ActionScheduleUpgrade) GetSignBytes() []byte {
	b, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (ActionCancelUpgrade) Route() string { return "consortium" }

func (ActionCancelUpgrade) Type() string { return "consortium.cancel-upgrade" }

func (ActionCancelUpgrade) ValidateBasic() sdk.Error {
	return nil
}

func (action ActionCancelUpgrade) GetSignBytes() []byte {
	b, err := json.Marshal(action)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (action ActionChangeValidatorSet) Route() string { return "consortium" }

func (action ActionChangeValidatorSet) Type() string { return "consortium.changeValidatorSet" }

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
