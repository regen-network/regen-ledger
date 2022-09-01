package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// AddCreditType adds a new credit type to the network.
func (k Keeper) AddCreditType(ctx context.Context, req *types.MsgAddCreditType) (*types.MsgAddCreditTypeResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	if err := k.stateStore.CreditTypeTable().Insert(ctx, &api.CreditType{
		Abbreviation: req.CreditType.Abbreviation,
		Name:         req.CreditType.Name,
		Unit:         req.CreditType.Unit,
		Precision:    req.CreditType.Precision,
	}); err != nil {
		if ormerrors.AlreadyExists.Is(err) {
			return nil, sdkerrors.ErrConflict.Wrapf("credit type abbreviation %s already exists", req.CreditType.Abbreviation)
		} else if ormerrors.UniqueKeyViolation.Is(err) {
			return nil, sdkerrors.ErrConflict.Wrapf("credit type with %s name already exists", req.CreditType.Name)
		}

		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not add credit type: %s", err.Error())
	}

	return &types.MsgAddCreditTypeResponse{}, nil
}
