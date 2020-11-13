package server

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
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
func DoExecuteMsgs(ctx sdk.Context, msgServiceRouter *baseapp.MsgServiceRouter, groupAccount sdk.AccAddress, msgs []sdk.Msg) ([]sdk.Result, error) {
	results := make([]sdk.Result, len(msgs))
	if err := ensureMsgAuthZ(msgs, groupAccount); err != nil {
		return nil, err
	}

	for i, msg := range msgs {
		if svcMsg, ok := msg.(sdk.ServiceMsg); ok {
			msgFqName := svcMsg.MethodName
			handler := msgServiceRouter.Handler(msgFqName)
			if handler == nil {
				return nil, errors.Wrapf(errors.ErrUnknownRequest, "unrecognized message service method: %s; message index: %d", msgFqName, i)
			}
			msgResult, err := handler(ctx, svcMsg.Request)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to execute message; message index: %d", i)
			}
			results[i] = *msgResult
		} else {
			// Should we support legacy sdk.Msg routing?
			return nil, errors.Wrapf(errors.ErrUnknownRequest, "non ServiceMsg not supported: %s; message index: %d", msg, i)
		}
	}
	return results, nil
}
