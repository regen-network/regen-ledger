package basket

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"

	sdk "github.com/cosmos/cosmos-sdk/types"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) Create(ctx context.Context, msg *basket.MsgCreate) (*basket.MsgCreateResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	fee := k.ecocreditKeeper.GetCreateBasketFee(ctx)
	if err := basket.ValidateCreateFee(msg, fee); err != nil {
		return nil, err

	}
	sender, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}

	err = k.distKeeper.FundCommunityPool(sdk.UnwrapSDKContext(ctx), fee, sender)
	if err != nil {
		return nil, err
	}

	res, err := k.ecocreditKeeper.CreditTypes(ctx, &ecocredit.QueryCreditTypesRequest{})
	if err != nil {
		return nil, err
	}

	found := false
	for _, creditType := range res.CreditTypes {
		if creditType.Abbreviation == msg.CreditTypeName {
			found = true
			if creditType.Precision > msg.Exponent {
				return nil, sdkerrors.ErrInvalidRequest.Wrapf(
					"basket exponent %d must be >= credit type precision %d",
					msg.Exponent,
					creditType.Precision,
				)
			}
			break
		}
	}
	if !found {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("can't find credit type %s", msg.CreditTypeName)
	}

	// TODO: need to decide about the denom creation
	denom := msg.Name

	id, err := k.stateStore.BasketStore().InsertReturningID(ctx, &basketv1.Basket{
		BasketDenom:       denom,
		DisableAutoRetire: msg.DisableAutoRetire,
		CreditTypeName:    msg.CreditTypeName,
		DateCriteria:      msg.DateCriteria.ToApi(),
		Exponent:          msg.Exponent,
	})
	if err != nil {
		return nil, err
	}

	for _, class := range msg.AllowedClasses {
		_, err := k.ecocreditKeeper.ClassInfo(ctx, &ecocredit.QueryClassInfoRequest{ClassId: class})
		if err != nil {
			return nil, err
		}

		err = k.stateStore.BasketClassStore().Insert(ctx,
			&basketv1.BasketClass{
				BasketId: id,
				ClassId:  class,
			},
		)

		if err != nil {
			return nil, err
		}
	}

	k.bankKeeper.SetDenomMetaData(sdk.UnwrapSDKContext(ctx), banktypes.Metadata{
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    msg.DisplayName,
				Exponent: msg.Exponent,
				Aliases:  nil,
			},
		},
		Base:    denom,
		Display: msg.DisplayName,
		Name:    msg.DisplayName,
		Symbol:  msg.DisplayName,
	})

	err = sdkCtx.EventManager().EmitTypedEvent(&basket.EventCreate{
		BasketDenom: denom,
		Curator:     msg.Curator,
	})

	return &basket.MsgCreateResponse{BasketDenom: denom}, err
}
