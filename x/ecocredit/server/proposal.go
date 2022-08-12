package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// ProposalKeeper defines methods for ecocredit gov handlers.
type ProposalKeeper interface {
	AllowDenom(ctx sdk.Context, proposal *marketplace.AllowDenomProposal) error
}

func (s serverImpl) AllowDenom(ctx sdk.Context, proposal *marketplace.AllowDenomProposal) error {
	return s.marketplaceKeeper.AllowDenom(ctx, proposal)
}

func NewProposalHandler(k ProposalKeeper) govv1beta1.Handler {
	return func(ctx sdk.Context, content govv1beta1.Content) error {
		switch c := content.(type) {
		case *marketplace.AllowDenomProposal:
			return handleAllowDenomProposal(ctx, k, c)
		default:
			return sdkerrors.ErrUnknownRequest.Wrapf("unrecognized proposal content type: %T", c)
		}
	}
}

func handleAllowDenomProposal(ctx sdk.Context, k ProposalKeeper, proposal *marketplace.AllowDenomProposal) error {
	return k.AllowDenom(ctx, proposal)
}
