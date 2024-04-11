package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

func (er EventReducer) reduceEventCreateClass(ctx context.Context, evt *types.EventCreateClass) error {
	creditType, err := er.CreditTypeTable().Get(ctx, evt.CreditType)
	if err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("could not get credit type with abbreviation %s: %s", evt.CreditType, err)
	}

	var seq uint64 = 1
	classSeq, err := er.ClassSequenceTable().Get(ctx, creditType.Abbreviation)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return err
		}
	} else {
		seq = classSeq.NextSequence
	}
	if err = er.ClassSequenceTable().Save(ctx, &ecocreditv1.ClassSequence{
		CreditTypeAbbrev: creditType.Abbreviation,
		NextSequence:     seq + 1,
	}); err != nil {
		return err
	}

	evt.ClassId = base.FormatClassID(creditType.Abbreviation, seq)

	admin, err := sdk.AccAddressFromBech32(evt.Admin)
	if err != nil {
		return err
	}

	classKey, err := er.ClassTable().InsertReturningID(ctx, &ecocreditv1.Class{
		Id:               evt.ClassId,
		Admin:            admin,
		Metadata:         evt.Metadata,
		CreditTypeAbbrev: evt.CreditType,
	})
	if err != nil {
		return err
	}

	for _, issuerStr := range evt.Issuers {
		issuer, err := sdk.AccAddressFromBech32(issuerStr)
		if err != nil {
			return err
		}

		err = er.ClassIssuerTable().Insert(ctx, &ecocreditv1.ClassIssuer{
			ClassKey: classKey,
			Issuer:   issuer,
		})
		if err != nil {
			return err
		}

		sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/MsgCreateClass issuer iteration")
	}

	return nil
}
