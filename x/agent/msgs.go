package agent

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateAgent struct {
	Data AgentInfo
	Signer sdk.AccAddress
}

type MsgUpdateAgent struct {
	Data AgentInfo
	Signers []AgentSig
}
