package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// ProposalKeeper defines methods for ecocredit gov handlers.
type ProposalKeeper interface {
	AddCreditType(ctx sdk.Context, ctp *core.CreditTypeProposal) error
	AddAskDenom(ctx sdk.Context, proposal *marketplace.AskDenomProposal) error
}

func (s serverImpl) AddAskDenom(ctx sdk.Context, proposal *marketplace.AskDenomProposal) error {
	return s.marketplaceKeeper.AddAskDenom(ctx, proposal)
}

func (s serverImpl) AddCreditType(ctx sdk.Context, ctp *core.CreditTypeProposal) error {
	return s.coreKeeper.AddCreditType(ctx, ctp)
}

func NewProposalHandler(k ProposalKeeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *core.CreditTypeProposal:
			return handleCreditTypeProposal(ctx, k, c)
		case *marketplace.AskDenomProposal:
			return handleAskDenomProposal(ctx, k, c)
		default:
			return sdkerrors.ErrUnknownRequest.Wrapf("unrecognized proposal content type: %T", c)
		}
	}
}

func handleAskDenomProposal(ctx sdk.Context, k ProposalKeeper, proposal *marketplace.AskDenomProposal) error {
	return k.AddAskDenom(ctx, proposal)
}

func handleCreditTypeProposal(ctx sdk.Context, k ProposalKeeper, proposal *core.CreditTypeProposal) error {
	return k.AddCreditType(ctx, proposal)
}
