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
	AuthPolicy AuthPolicy `json:"auth_policy"`
	// An Agent can have either addresses or other agents as members
	Addresses []sdk.AccAddress `json:"addresses"`
	Agents []AgentId `json:"agents"`
	MultisigThreshold int `json:"multisig_threshold"`
}

