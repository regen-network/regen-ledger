package data

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

const DataCodespace = "regen.data"

var (
	ErrHashVerificationFailed      = sdkerrors.Register(DataCodespace, 1, "hash verification failed")
	ErrInvalidIRI                  = sdkerrors.Register(DataCodespace, 2, "invalid IRI")
	ErrInvalidMediaExtension       = sdkerrors.Register(DataCodespace, 3, "invalid media extension")
	ErrResolverURLExists           = sdkerrors.Register(DataCodespace, 4, "resolver URL already exists")
	ErrResolverUndefined           = sdkerrors.Register(DataCodespace, 5, "resolver undefined")
	ErrUnauthorizedResolverManager = sdkerrors.Register(DataCodespace, 6, "unauthorized resolver manager")
)
