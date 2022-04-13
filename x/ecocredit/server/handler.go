package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	coretypes "github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func NewCreditTypeProposalHandler(k ProposalKeeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *coretypes.CreditTypeProposal:
			return handleCreditTypeProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized software upgrade proposal content type: %T", c)
		}
	}
}

func handleCreditTypeProposal(ctx sdk.Context, k ProposalKeeper, proposal *coretypes.CreditTypeProposal) error {
	return k.NewCreditType(ctx, proposal)
}
