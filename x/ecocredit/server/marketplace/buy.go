package marketplace

import (
	"context"

	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Buy creates a buy order that matches sell orders based on filter criteria.
func (k Keeper) Buy(ctx context.Context, req *marketplace.MsgBuy) (*marketplace.MsgBuyResponse, error) {
	return nil, sdkerrors.ErrInvalidRequest.Wrap("only direct buy orders are enabled at this time")
}
