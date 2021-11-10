package server

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	regenmath "github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s serverImpl) CreateBasket(goCtx context.Context, req *ecocredit.MsgCreateBasket) (*ecocredit.MsgCreateBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	basketDenom := getBasketDenom(req.Curator, req.Name)
	err := s.basketInfoTable.Create(ctx, &ecocredit.BasketInfo{
		BasketDenom:       basketDenom,
		DisplayName:       req.DisplayName,
		Exponent:          req.Exponent,
		AdmissionCriteria: req.AdmissionCriteria,
		RetireOnTake:      req.RetireOnTake,
		AllowPicking:      req.AllowPicking,
	})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateBasketResponse{BasketDenom: basketDenom}, nil
}

func (s serverImpl) AddToBasket(goCtx context.Context, req *ecocredit.MsgAddToBasket) (*ecocredit.MsgAddToBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.storeKey)

	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	if !s.basketInfoTable.Has(ctx, orm.RowID(req.BasketDenom)) {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid basket denom", req.BasketDenom)
	}

	var basket ecocredit.BasketInfo
	err = s.basketInfoTable.GetOne(ctx, orm.RowID(req.BasketDenom), &basket)
	if err != nil {
		return nil, err
	}

	amtReceived := regenmath.NewDecFromInt64(0)

	for _, credit := range req.Credits {
		batchDenom := batchDenomT(credit.BatchDenom)
		if !s.batchInfoTable.Has(ctx, orm.RowID(batchDenom)) {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid credit batch denom", batchDenom)
		}

		// TODO Add AdmissionCriteria validation here

		maxDecimalPlaces, err := s.getBatchPrecision(ctx, batchDenom)
		if err != nil {
			return nil, err
		}

		tradable, err := regenmath.NewNonNegativeFixedDecFromString(credit.TradableAmount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		amtReceived, err = amtReceived.Add(tradable)
		if err != nil {
			return nil, err
		}

		basketCreditsKey := BasketCreditsKey(basketDenomT(req.BasketDenom), owner, batchDenom)
		err = addAndSetDecimal(store, basketCreditsKey, tradable)
		if err != nil {
			return nil, err
		}

		// Send credits from owner to derived module account
		derivedKey := s.storeKey.Derive(basketCreditsKey)
		_, err = s.Send(goCtx, &ecocredit.MsgSend{
			Sender:    req.Owner,
			Recipient: derivedKey.Address().String(),
			Credits: []*ecocredit.MsgSend_SendCredits{
				{
					BatchDenom:     credit.BatchDenom,
					TradableAmount: credit.TradableAmount,
				},
			},
		})
		if err != nil {
			return nil, err
		}
	}

	// TODO Why 10?
	multiplier, err := regenmath.NewNonNegativeFixedDecFromString(basket.AdmissionCriteria[0].Multiplier, 10)
	if err != nil {
		return nil, err
	}
	multipliedAmtReceived, err := amtReceived.Mul(multiplier)
	if err != nil {
		return nil, err
	}

	// TODO Is there another way than to convert from regenmath.Dec to sdk.Int other than passing by int64?
	i, err := multipliedAmtReceived.Int64()
	if err != nil {
		return nil, err
	}
	amtAsInt := sdk.NewIntFromUint64(uint64(i))
	basketTokens := sdk.NewCoins(sdk.NewCoin(basket.BasketDenom, amtAsInt))

	// TODO don't hardcode ecocredit string.
	err = s.bankKeeper.MintCoins(ctx.Context, "ecocredit", basketTokens)
	if err != nil {
		return nil, err
	}
	err = s.bankKeeper.SendCoinsFromModuleToAccount(ctx.Context, "ecocredit", owner, basketTokens)
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgAddToBasketResponse{
		AmountReceived: amtReceived.String(),
	}, nil
}

func (s serverImpl) TakeFromBasket(goCtx context.Context, req *ecocredit.MsgTakeFromBasket) (*ecocredit.MsgTakeFromBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	if !s.basketInfoTable.Has(ctx, orm.RowID(req.BasketDenom)) {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid basket denom", req.BasketDenom)
	}

	// store := ctx.KVStore(s.storeKey)
	// it := sdk.KVStorePrefixIterator(store, address.MustLengthPrefix([]byte(req.BasketDenom)))
	// defer it.Close()

	// n := 0
	// var tradable []regenmath.Dec
	// // var batchDenoms []string
	// // var owners [][]byte
	// for ; it.Valid(); it.Next() {
	// 	// strip batchDenom and owner from key
	// 	// it.Key()
	// 	// TODO use NewNonNegativeFixedDecFromString with batch precision
	// 	v, err := regenmath.NewDecFromString(string(it.Value()))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	tradable = append(tradable, v)
	// 	n++
	// }
	// hashInt := new(big.Int)
	// hashInt.SetString(prevBlockHash, 16)
	// hashFloat := new(big.Float).SetInt(hashInt)
	// hash, _ := hashFloat.Float64()
	// idx := math.Mod(hash, float64(n))
	return &ecocredit.MsgTakeFromBasketResponse{}, nil
}

func (s serverImpl) PickFromBasket(goCtx context.Context, req *ecocredit.MsgPickFromBasket) (*ecocredit.MsgPickFromBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	var basketInfo ecocredit.BasketInfo
	err := s.basketInfoTable.GetOne(ctx, orm.RowID(req.BasketDenom), &basketInfo)
	if err != nil {
		return nil, err
	}

	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	store := ctx.KVStore(s.storeKey)
	basketDenom := basketDenomT(req.BasketDenom)
	for _, c := range req.Credits {
		batchDenom := batchDenomT(c.BatchDenom)
		maxDecimalPlaces, err := s.getBatchPrecision(ctx, batchDenom)
		if err != nil {
			return nil, err
		}
		tradableAmount, err := regenmath.NewNonNegativeFixedDecFromString(c.TradableAmount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		// Only an address which deposited credits in the basket can pick those credits
		if !basketInfo.AllowPicking {
			basketCreditKey := BasketCreditsKey(basketDenom, owner, batchDenom)
			totalTradable, err := getDecimal(store, basketCreditKey)
			if err != nil {
				return nil, err
			}

			totalTradable, err = regenmath.SafeSubBalance(totalTradable, tradableAmount)
			if err != nil {
				return nil, err
			}
			// Update basket credit
			setDecimal(store, basketCreditKey, totalTradable)
			// Retire if needed
			if basketInfo.RetireOnTake {
				err := retireUpdateBalanceSupply(ctx, store, owner, batchDenom, tradableAmount, req.RetirementLocation)
				if err != nil {
					return nil, err
				}
			} else {
				// Send credits from corresponding sub module account to req.Owner
				derivedKey := s.storeKey.Derive(basketCreditKey)
				_, err = s.Send(goCtx, &ecocredit.MsgSend{
					Recipient: derivedKey.Address().String(),
					Sender:    req.Owner,
					Credits: []*ecocredit.MsgSend_SendCredits{
						{
							BatchDenom:     c.BatchDenom,
							TradableAmount: c.TradableAmount,
						},
					},
				})
				if err != nil {
					return nil, err
				}
			}
		} else {
			prefix := BasketCreditsKey(basketDenom, []byte{}, batchDenom)
			it := sdk.KVStorePrefixIterator(store, prefix)
			defer it.Close()

			basketCredits := make(map[string]regenmath.Dec)
			batchTotalTradable := regenmath.NewDecFromInt64(0)
			for ; it.Valid(); it.Next() {
				value, err := regenmath.NewDecFromString(string(it.Value()))
				if err != nil {
					return nil, err
				}
				batchTotalTradable, err = batchTotalTradable.Add(value)
				if err != nil {
					return nil, err
				}
				basketCredits[string(it.Key())] = value
			}
			if batchTotalTradable.Cmp(tradableAmount) == -1 {
				return nil, ecocredit.ErrInsufficientFunds
			}
			var nextTradableAmount regenmath.Dec
			for key, value := range basketCredits {
				sub, err := value.Sub(tradableAmount)
				if err != nil {
					return nil, err
				}
				// Update basket credit
				if sub.IsPositive() || sub.IsZero() {
					setDecimal(store, []byte(key), sub)
					nextTradableAmount = regenmath.NewDecFromInt64(0)
				} else if sub.IsNegative() {
					setDecimal(store, []byte(key), regenmath.NewDecFromInt64(0))
					nextTradableAmount, err = sub.Mul(regenmath.NewDecFromInt64(-1))
					if err != nil {
						return nil, err
					}
				}

				err = s.sendOrRetire(basketInfo.RetireOnTake, goCtx, ctx, store, owner, batchDenom, basketCreditKey, tradableAmount, req.RetirementLocation)
				if err != nil {
					return nil, err
				}
				// Retire if needed
				if basketInfo.RetireOnTake {
					err = retireUpdateBalanceSupply(ctx, store, owner, batchDenom, tradableAmount, req.RetirementLocation)
					if err != nil {
						return nil, err
					}
					// Send credits from corresponding sub module account to req.Owner
				} else {
					derivedKey := s.storeKey.Derive(basketCreditKey)
					_, err = s.Send(goCtx, &ecocredit.MsgSend{
						Recipient: derivedKey.Address().String(),
						Sender:    req.Owner,
						Credits: []*ecocredit.MsgSend_SendCredits{
							{
								BatchDenom:     c.BatchDenom,
								TradableAmount: c.TradableAmount,
							},
						},
					})
					if err != nil {
						return nil, err
					}
				}
				if nextTradableAmount.IsZero() {
					break
				}
				tradableAmount = nextTradableAmount
			}
		}
	}
	return &ecocredit.MsgPickFromBasketResponse{}, nil
}

func (s serverImpl) sendOrRetire(retireOnTake bool, goCtx context.Context, ctx types.Context,
	store sdk.KVStore, recipient sdk.AccAddress, batchDenom batchDenomT, basketCreditKey []byte, amount math.Dec, location string) error {
	if retireOnTake {
		err := retireUpdateBalanceSupply(ctx, store, recipient, batchDenom, amount, location)
		if err != nil {
			return err
		}
	} else {
		// Send credits from corresponding sub module account to req.Owner
		derivedKey := s.storeKey.Derive(basketCreditKey)
		_, err := s.Send(goCtx, &ecocredit.MsgSend{
			Recipient: derivedKey.Address().String(),
			Sender:    recipient.String(),
			Credits: []*ecocredit.MsgSend_SendCredits{
				{
					BatchDenom:     string(batchDenom),
					TradableAmount: amount.String(),
				},
			},
		})
		if err != nil {
			return err
		}
	}
}

func getBasketDenom(curator, name string) string {
	return fmt.Sprintf("ecocredit:%s:%s", curator, name)
}
