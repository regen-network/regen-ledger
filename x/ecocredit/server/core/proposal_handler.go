package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	coretypes "github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ProposalKeeper defines the expected interface for ecocredit module proposals.
type ProposalKeeper interface {
	AddCreditType(ctx sdk.Context, ctp *coretypes.CreditTypeProposal) error
}

func NewCreditTypeProposalHandler(k ProposalKeeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *coretypes.CreditTypeProposal:
			return handleCreditTypeProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized credit type proposal content type: %T", c)
		}
	}
}

func handleCreditTypeProposal(ctx sdk.Context, k ProposalKeeper, proposal *coretypes.CreditTypeProposal) error {
	return k.AddCreditType(ctx, proposal)
}
