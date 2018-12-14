package org

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AuthPolicy int

const (
	Multisig AuthPolicy = 1
)

type AgentRef struct {
	Agent []byte
	Address sdk.AccAddress
}

type AgentInfo struct {
	AuthPolicy AuthPolicy
	Agents []AgentRef
	MultisigThreshold int
}
