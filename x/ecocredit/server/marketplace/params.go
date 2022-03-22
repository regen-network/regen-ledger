package marketplace

import (
	"context"

	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AllowAskDenom(ctx context.Context, req *v1.MsgAllowAskDenom) (*v1.MsgAllowAskDenomResponse, error) {
	govAcc, err := sdk.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = server.AssertGovernance(govAcc, k.accountKeeper); err != nil {
		return nil, err
	}
	return nil, nil
}
