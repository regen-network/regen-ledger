package data

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

const DataCodespace = "regen.data"

var (
	ErrHashVerificationFailed = sdkerrors.Register(DataCodespace, 1, "hash verification failed")
	ErrInvalidIRI             = sdkerrors.Register(DataCodespace, 2, "invalid IRI")
	ErrInvalidMediaExtension  = sdkerrors.Register(DataCodespace, 3, "invalid media extension")
)
