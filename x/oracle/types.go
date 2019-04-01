package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
)

type ResultCode int

const (
	CodeOK ResultCode = iota
	CodeError
	CodeTimedOut
)

// MsgMakeComputeClaim is used by an oracle to make a compute claim for a computation they have made themselves
type MsgMakeComputeClaim struct {
	Oracle          sdk.Address
	Invocation      types.DataAddress
	Code            ResultCode
	Result          types.DataAddress
	DependencyGraph types.DataAddress
}

// MsgApproveComputeClaim is used by an oracle to approve a computation that another oracle has made even if that
// result differs slightly from their own (eg. for floating point non-determinism)
type MsgApproveComputeClaim struct {
	Oracle     sdk.Address
	Code       ResultCode
	Invocation types.DataAddress
	Result     types.DataAddress
}
