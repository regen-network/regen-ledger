package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type EventReducer struct {
	ecocreditv1.StateStore
}

func (er EventReducer) Emit(ctx context.Context, evt proto.Message) error {
	err := er.reduce(ctx, evt)
	if err != nil {
		return err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.EventManager().EmitTypedEvent(evt)
}

func (er EventReducer) reduce(ctx context.Context, evt proto.Message) error {
	switch evt := evt.(type) {
	case *types.EventCreateClass:
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
		if err = er.ClassSequenceTable().Save(ctx, &api.ClassSequence{
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

		return er.ClassTable().Insert(ctx, &ecocreditv1.Class{
			Id:               evt.ClassId,
			Admin:            admin,
			Metadata:         evt.Metadata,
			CreditTypeAbbrev: evt.CreditType,
		})
	case *types.EventUpdateClassIssuers:
		class, err := er.ClassTable().GetById(ctx, evt.ClassId)
		if err != nil {
			return err
		}

		issuer, err := sdk.AccAddressFromBech32(evt.Issuer)
		if err != nil {
			return err
		}

		if evt.Removed {
			return er.ClassIssuerTable().Delete(ctx, &api.ClassIssuer{
				ClassKey: class.Key,
				Issuer:   issuer,
			})
		} else {
			return er.ClassIssuerTable().Save(ctx, &api.ClassIssuer{
				ClassKey: class.Key,
				Issuer:   issuer,
			})
		}
	}
	return nil
}
