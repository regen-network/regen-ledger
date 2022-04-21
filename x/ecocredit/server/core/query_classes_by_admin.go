package core

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ClassesByAdmin queries for all classes with a specific admin address.
func (k Keeper) ClassesByAdmin(ctx context.Context, req *core.QueryClassesByAdminRequest) (*core.QueryClassesByAdminResponse, error) {
	admin, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	pg, err := ormutil.GogoPageReqToPulsarPageReq(req.Pagination)
	if err != nil {
		return nil, err
	}

	it, err := k.stateStore.ClassInfoTable().List(ctx, api.ClassInfoAdminIndexKey{}.WithAdmin(admin), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}

	classes := make([]*core.ClassDetails, 0)
	for it.Next() {
		class, err := it.Value()
		if err != nil {
			return nil, err
		}

		info := core.ClassDetails{
			Id:               class.Id,
			Admin:            admin.String(),
			Metadata:         class.Metadata,
			CreditTypeAbbrev: class.CreditTypeAbbrev,
		}

		classes = append(classes, &info)
	}

	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, err
	}

	return &core.QueryClassesByAdminResponse{Classes: classes, Pagination: pr}, nil
}
