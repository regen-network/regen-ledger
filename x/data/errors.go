package data

import "cosmossdk.io/errors"

const DataCodespace = "regen.data"

var (
	ErrInvalidIRI                  = errors.Register(DataCodespace, 2, "invalid IRI")
	ErrInvalidMediaExtension       = errors.Register(DataCodespace, 3, "invalid media extension")
	ErrUnauthorizedResolverManager = errors.Register(DataCodespace, 4, "unauthorized resolver manager")
	ErrParseFailure                = errors.Register(DataCodespace, 5, "parse error")
)

