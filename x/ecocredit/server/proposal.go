package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// ProposalKeeper defines methods for ecocredit gov handlers.
type ProposalKeeper interface {
	AddCreditType(ctx sdk.Context, ctp *core.CreditTypeProposal) error
	AllowDenom(ctx sdk.Context, proposal *marketplace.AllowDenomProposal) error
}

func (s serverImpl) AllowDenom(ctx sdk.Context, proposal *marketplace.AllowDenomProposal) error {
	return s.marketplaceKeeper.AllowDenom(ctx, proposal)
}

func (s serverImpl) AddCreditType(ctx sdk.Context, ctp *core.CreditTypeProposal) error {
	return s.coreKeeper.AddCreditType(ctx, ctp)
}

func NewProposalHandler(k ProposalKeeper) govv1beta1.Handler {
	return func(ctx sdk.Context, content govv1beta1.Content) error {
		switch c := content.(type) {
		case *core.CreditTypeProposal:
			return handleAddCreditTypeProposal(ctx, k, c)
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

func handleAddCreditTypeProposal(ctx sdk.Context, k ProposalKeeper, proposal *core.CreditTypeProposal) error {
	return k.AddCreditType(ctx, proposal)
}
