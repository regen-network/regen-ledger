package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (er EcocreditEventReducer) ReduceEventUpdateClassIssuers(ctx context.Context, evt *types.EventUpdateClassIssuers) error {
	class, err := er.ClassTable().GetById(ctx, evt.ClassId)
	if err != nil {
		return err
	}

	// remove issuers
	for _, issuer := range evt.Removed {
		issuerAcc, err := sdk.AccAddressFromBech32(issuer)
		if err != nil {
			return err
		}
		if err = er.ClassIssuerTable().Delete(ctx, &api.ClassIssuer{
			ClassKey: class.Key,
			Issuer:   issuerAcc,
		}); err != nil {
			return err
		}

		sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/MsgUpdateClassIssuers issuer iteration")
	}

	// add the new issuers
	for _, issuer := range evt.Added {
		issuerAcc, err := sdk.AccAddressFromBech32(issuer)
		if err != nil {
			return err
		}
		if err = er.ClassIssuerTable().Insert(ctx, &api.ClassIssuer{
			ClassKey: class.Key,
			Issuer:   issuerAcc,
		}); err != nil {
			return err
		}

		sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/MsgUpdateClassIssuers issuer iteration")
	}

	return nil
}
