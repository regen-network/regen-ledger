package agent

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AgentId []byte

type AuthPolicy int

const (
	MultiSig AuthPolicy = 1
)

// An agent can be used to abstract over users and groups
// It could be used by a single user to manage multiple devices and setup a multisig policy
// It could be used to group individuals into a group or several groups/users into a larger group
type AgentInfo struct {
	AuthPolicy AuthPolicy
	// List of either agents or account addresses
	Agents            []AgentRef
	MultisigThreshold int
}

type AgentRefType int

const (
	AgentRef_Agent AgentRefType = 1
	AgentRef_Address AgentRefType = 2
)


type AgentRef struct {
	Type AgentRefType
	Agent   AgentId
	Address sdk.AccAddress
}

