package server

import (
	"context"
	"fmt"

	"github.com/regen-network/regen-ledger/types"
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

	// TODO
}

func (s serverImpl) TakeFromBasket(goCtx context.Context, req *ecocredit.MsgTakeFromBasket) (*ecocredit.MsgTakeFromBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	// TODO
}

func (s serverImpl) PickFromBasket(goCtx context.Context, req *ecocredit.MsgPickFromBasket) (*ecocredit.MsgPickFromBasketResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	// TODO
}

func getBasketDenom(curator, name string) string {
	return fmt.Sprintf("ecocredit:%s:%s", curator, name)
}
