package server

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) execMsgs(ctx context.Context, path []byte, proposal group.Proposal) error {
	msgs := proposal.GetMsgs()
	for _, msg := range msgs {
		svcMsg, ok := msg.(sdk.ServiceMsg)
		if !ok {
			return fmt.Errorf("expected sdk.ServiceMsg, got %T", msg)
		}
		var reply interface{}
		derivedKey := s.key.Derive(path)
		// Execute the message using the derived key,
		// this will verify that the message signer is the group account.
		err := derivedKey.Invoke(ctx, svcMsg.Route(), svcMsg.Request, reply)
		if err != nil {
			return err
		}
	}
	return nil
}

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
