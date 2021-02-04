package data

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

const DataCodespace = "regen.data"

var (
	ErrHashVerificationFailed = sdkerrors.Register(DataCodespace, 1, "hash verification failed")
)
