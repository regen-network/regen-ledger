package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/group"
)

// ensureMsgAuthZ checks that if a message requires signers that all of them are equal to the given group account.
func ensureMsgAuthZ(msgs []sdk.Msg, groupAccount sdk.AccAddress) error {
	for i := range msgs {
		for _, acct := range msgs[i].GetSigners() {
			if !groupAccount.Equals(acct) {
				return errors.Wrap(errors.ErrUnauthorized, "msg does not have group account authorization")
			}
		}
	}
	return nil
}

// DoExecuteMsgs routes the messages to the registered handlers. Messages are limited to those that require no authZ or
// by the group account only. Otherwise this gives access to other peoples accounts as the sdk ant handler is bypassed
func DoExecuteMsgs(ctx sdk.Context, router sdk.Router, groupAccount sdk.AccAddress, msgs []sdk.Msg) ([]sdk.Result, error) {
	results := make([]sdk.Result, len(msgs))
	if err := ensureMsgAuthZ(msgs, groupAccount); err != nil {
		return nil, err
	}
	for i, msg := range msgs {
		handler := router.Route(ctx, msg.(legacytx.LegacyMsg).Route())
		if handler == nil {
			return nil, errors.Wrapf(group.ErrInvalid, "no message handler found for %q", msg.(legacytx.LegacyMsg).Route())
		}
		r, err := handler(ctx, msg)
		if err != nil {
			return nil, errors.Wrapf(err, "message %q at position %d", msg.(legacytx.LegacyMsg).Type(), i)
		}
		if r != nil {
			results[i] = *r
		}
	}
	return results, nil
}
