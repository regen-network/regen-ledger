package data

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

const DataCodespace = "regen.data"

var (
	ErrInvalidIRI                  = sdkerrors.Register(DataCodespace, 2, "invalid IRI")
	ErrInvalidMediaExtension       = sdkerrors.Register(DataCodespace, 3, "invalid media extension")
	ErrUnauthorizedResolverManager = sdkerrors.Register(DataCodespace, 4, "unauthorized resolver manager")
	ErrParseFailure                = sdkerrors.Register(DataCodespace, 5, "parse error")
)
