package basket

import (
	"context"

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

	err = k.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, sender, feeModuleAccName, fee)
	if err != nil {
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
		Description: "",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    msg.DisplayName,
				Exponent: msg.Exponent,
				Aliases:  nil,
			},
		},
		Base:    denom,
		Display: msg.DisplayName,
		Name:    msg.Name,
		Symbol:  denom,
	})

	err = sdkCtx.EventManager().EmitTypedEvent(&basket.EventCreate{
		BasketDenom: denom,
		Curator:     msg.Curator,
	})

	return &basket.MsgCreateResponse{BasketDenom: denom}, err
}
