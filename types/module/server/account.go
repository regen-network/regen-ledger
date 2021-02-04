package server

import (
	"context"

	"github.com/regen-network/regen-ledger/types"
)

const (
	createNewAccountMethod    = "$$CreateNewAccount"
	ensureAccountExistsMethod = "$$EnsureAccountExists"
	accountExistsMethod       = "$$AccountExistsMethod"
)

// CreateNewAccount attempts to initialize a new account for this key and
// returns an error if a new account could not be created or one already existed
func CreateNewAccount(moduleKey ModuleKey, ctx context.Context) error {
	return moduleKey.Invoke(ctx, createNewAccountMethod, nil, nil)
}

// EnsureAccountExists initializes a new account for this key if one does not
// exist and returns an error if a new account could not be created.
func EnsureAccountExists(moduleKey ModuleKey, ctx context.Context) error {
	return moduleKey.Invoke(ctx, ensureAccountExistsMethod, nil, nil)
}

// AccountExists returns true if an account for this key exists in state.
func AccountExists(moduleKey ModuleKey, ctx types.Context) bool {
	res := &AccountExistsResponse{}
	err := moduleKey.Invoke(ctx, accountExistsMethod, nil, nil)
	if err != nil {
		return false
	}

	return res.Exists
}

type AccountExistsResponse struct {
	Exists bool
}
