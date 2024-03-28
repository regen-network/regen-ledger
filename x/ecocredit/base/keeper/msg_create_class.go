package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// CreateClass creates a new class of ecocredit
//
// The admin is charged a fee for creating the class. This is controlled by
// the global parameter CreditClassFee, which can be updated through the
// governance process.
func (k Keeper) CreateClass(goCtx context.Context, req *types.MsgCreateClass) (*types.MsgCreateClassResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	adminAddress, err := sdk.AccAddressFromBech32(req.Admin)
	if err != nil {
		return nil, err
	}

	if err := k.assertCanCreateClass(goCtx, adminAddress); err != nil {
		return nil, err
	}

	classFee, err := k.stateStore.ClassFeeTable().Get(goCtx)
	if err != nil {
		return nil, err
	}

	err = k.deductFee(goCtx, adminAddress, req.Fee, classFee.Fee)
	if err != nil {
		return nil, err
	}

	creditType, err := k.stateStore.CreditTypeTable().Get(goCtx, req.CreditTypeAbbrev)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not get credit type with abbreviation %s: %s", req.CreditTypeAbbrev, err)
	}

	// default the sequence to 1 for the `not found` case.
	// will get overwritten by the actual sequence if it exists.
	var seq uint64 = 1
	classSeq, err := k.stateStore.ClassSequenceTable().Get(goCtx, creditType.Abbreviation)
	if err != nil {
		if !ormerrors.IsNotFound(err) {
			return nil, err
		}
	} else {
		seq = classSeq.NextSequence
	}
	if err = k.stateStore.ClassSequenceTable().Save(goCtx, &api.ClassSequence{
		CreditTypeAbbrev: creditType.Abbreviation,
		NextSequence:     seq + 1,
	}); err != nil {
		return nil, err
	}

	classID := base.FormatClassID(creditType.Abbreviation, seq)

	key, err := k.stateStore.ClassTable().InsertReturningID(goCtx, &api.Class{
		Id:               classID,
		Admin:            adminAddress,
		Metadata:         req.Metadata,
		CreditTypeAbbrev: creditType.Abbreviation,
	})
	if err != nil {
		return nil, err
	}

	for _, issuer := range req.Issuers {
		issuer, err := sdk.AccAddressFromBech32(issuer)
		if err != nil {
			return nil, err
		}
		if err = k.stateStore.ClassIssuerTable().Insert(goCtx, &api.ClassIssuer{
			ClassKey: key,
			Issuer:   issuer,
		}); err != nil {
			return nil, err
		}

		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/MsgCreateClass issuer iteration")
	}

	err = sdkCtx.EventManager().EmitTypedEvent(&types.EventCreateClass{
		ClassId: classID,
	})
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateClassResponse{ClassId: classID}, nil
}

func (k Keeper) assertCanCreateClass(ctx context.Context, adminAddress sdk.AccAddress) error {
	allowListEnabled, err := k.stateStore.ClassCreatorAllowlistTable().Get(ctx)
	if err != nil {
		return err
	}

	if allowListEnabled.Enabled {
		_, err := k.stateStore.AllowedClassCreatorTable().Get(ctx, adminAddress)
		if err != nil {
			if ormerrors.NotFound.Is(err) {
				return sdkerrors.ErrUnauthorized.Wrapf("%s is not allowed to create credit classes", adminAddress.String())
			}
			return err
		}
	}
	return nil
}
