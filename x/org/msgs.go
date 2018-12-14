package org

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateAgent struct {
	Data AgentInfo
	Signers []sdk.AccAddress
}
