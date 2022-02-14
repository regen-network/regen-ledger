package basket

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) Create(ctx context.Context, msg *basket.MsgCreate) (*basket.MsgCreateResponse, error) {
	rgCtx := types.UnwrapSDKContext(ctx)
	fee := k.ecocreditKeeper.GetCreateBasketFee(ctx)
	if err := basket.ValidateCreateFee(msg, fee); err != nil {
		return nil, err

	}
	sender, err := sdk.AccAddressFromBech32(msg.Curator)
	if err != nil {
		return nil, err
	}

	err = k.distKeeper.FundCommunityPool(rgCtx.Context, fee, sender)
	if err != nil {
		return nil, err
	}
	if err = assertCreditTypeExists(ctx, k.ecocreditKeeper, msg.CreditTypeName, msg.Exponent); err != nil {
		return nil, err
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

	if err = k.indexAllowedClasses(rgCtx, id, msg.AllowedClasses); err != nil {
		return nil, err
	}

	k.bankKeeper.SetDenomMetaData(rgCtx.Context, banktypes.Metadata{
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

	err = rgCtx.Context.EventManager().EmitTypedEvent(&basket.EventCreate{
		BasketDenom: denom,
		Curator:     msg.Curator,
	})

	return &basket.MsgCreateResponse{BasketDenom: denom}, err
}

// assertCreditTypeExists returns error if a given credit type doesn't exist or it's precision
// is bigger then the requested exponent.
func assertCreditTypeExists(ctx context.Context, k EcocreditKeeper, creditType string, exponent uint32) error {
	res, err := k.CreditTypes(ctx, &ecocredit.QueryCreditTypesRequest{})
	if err != nil {
		return err
	}

	found := false
	for _, c := range res.CreditTypes {
		if c.Abbreviation == creditType {
			found = true
			if c.Precision > exponent {
				return sdkerrors.ErrInvalidRequest.Wrapf(
					"exponent %d must be >= credit type precision %d",
					exponent,
					c.Precision,
				)
			}
			break
		}
	}
	if !found {
		return sdkerrors.ErrInvalidRequest.Wrapf("can't find credit type %s", creditType)
	}
	return nil
}

func (k Keeper) indexAllowedClasses(ctx types.Context, basketID uint64, allowedClasses []string) error {
	for _, class := range allowedClasses {
		// TODO: we should be able to check if class exists without deserializing it.
		if !k.ecocreditKeeper.HasClassInfo(ctx, class) {
			return sdkerrors.ErrInvalidRequest.Wrapf("credit class %q doesn't exist", class)
		}

		err := k.stateStore.BasketClassStore().Insert(ctx,
			&basketv1.BasketClass{
				BasketId: basketID,
				ClassId:  class,
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
