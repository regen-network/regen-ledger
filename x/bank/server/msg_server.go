package server

import (
	"bytes"
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/util/storehelpers"
	"github.com/regen-network/regen-ledger/x/bank"
	"github.com/regen-network/regen-ledger/x/bank/math"
)

func (s serverImpl) CreateDenom(goCtx context.Context, req *bank.MsgCreateDenomRequest) (*bank.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	namespace := req.DenomNamespace
	denom := fmt.Sprintf("%s/%s", namespace, req.DenomName)
	store := ctx.KVStore(s.key)
	minterKey := DenomAdminKey(denom)
	if store.Has(minterKey) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "denom %s already exists", denom)
	}

	namespaceAdmin, found := s.denomNamespaceAdmins[namespace]
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "denom namespace %s has no authorized admin", namespace)
	}

	admin, err := sdk.AccAddressFromBech32(req.AdminAddress)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(admin, namespaceAdmin.Address()) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not admin for denom namespace %s", admin, namespace)
	}

	minter, err := sdk.AccAddressFromBech32(req.MinterAddress)
	if err != nil {
		return nil, err
	}

	store.Set(minterKey, minter)

	return &bank.MsgCreateDenomResponse{Denom: denom}, nil
}

func isAdmin(store sdk.KVStore, denom string, addr sdk.AccAddress) bool {
	admin := store.Get(DenomAdminKey(denom))
	return bytes.Equal(addr, admin)
}

func (s serverImpl) Mint(goCtx context.Context, req *bank.MsgMintRequest) (*bank.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.key)
	minter, err := sdk.AccAddressFromBech32(req.MinterAddress)
	if err != nil {
		return nil, err
	}

	for _, issuance := range req.Issuance {
		recipient, err := sdk.AccAddressFromBech32(issuance.Recipient)
		if err != nil {
			return nil, err
		}

		for _, coin := range issuance.Coins {
			denom := coin.Denom

			if !isAdmin(store, denom, minter) {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to mint coins for denom %s", minter, denom)
			}

			amount, err := math.ParseNonNegativeDecimal(coin.Amount)
			if err != nil {
				return nil, err
			}

			maxDecPlacesKey := MaxDecimalPlacesKey(denom)
			decPlaces := math.NumDecimalPlaces(amount)
			maxDecimalPlaces, err := storehelpers.GetUint32(store, maxDecPlacesKey)
			if err != nil {
				return nil, err
			}

			if decPlaces > maxDecimalPlaces {
				err = storehelpers.SetUInt32(store, maxDecPlacesKey, decPlaces)
				if err != nil {
					return nil, err
				}
			}

			// add balance
			err = storehelpers.GetAddAndSetDecimal(store, BalanceKey(recipient, denom), amount)
			if err != nil {
				return nil, err
			}

			// add supply
			err = storehelpers.GetAddAndSetDecimal(store, SupplyKey(denom), amount)
			if err != nil {
				return nil, err
			}
		}
	}

	return &bank.MsgMintResponse{}, nil
}

func (s serverImpl) Send(goCtx context.Context, req *bank.MsgSendRequest) (*bank.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.key)
	from, err := sdk.AccAddressFromBech32(req.FromAddress)
	if err != nil {
		return nil, err
	}

	to, err := sdk.AccAddressFromBech32(req.ToAddress)
	if err != nil {
		return nil, err
	}

	for _, coin := range req.Amount {
		denom := coin.Denom

		maxDecimalPlaces, err := storehelpers.GetUint32(store, MaxDecimalPlacesKey(denom))
		if err != nil {
			return nil, err
		}

		amount, err := math.ParseNonNegativeFixedDecimal(coin.Amount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		// subtract balance
		err = storehelpers.GetSubAndSetDecimal(store, BalanceKey(from, denom), amount)
		if err != nil {
			return nil, err
		}

		// add balance
		err = storehelpers.GetAddAndSetDecimal(store, BalanceKey(to, denom), amount)
		if err != nil {
			return nil, err
		}
	}

	return &bank.MsgSendResponse{}, nil
}

func (s serverImpl) Burn(goCtx context.Context, req *bank.MsgBurnRequest) (*bank.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.key)
	burner, err := sdk.AccAddressFromBech32(req.BurnerAddress)
	if err != nil {
		return nil, err
	}

	for _, coin := range req.Coins {
		denom := coin.Denom

		maxDecimalPlaces, err := storehelpers.GetUint32(store, MaxDecimalPlacesKey(denom))
		if err != nil {
			return nil, err
		}

		amount, err := math.ParseNonNegativeFixedDecimal(coin.Amount, maxDecimalPlaces)
		if err != nil {
			return nil, err
		}

		// subtract balance
		err = storehelpers.GetSubAndSetDecimal(store, BalanceKey(burner, denom), amount)
		if err != nil {
			return nil, err
		}

		// subtract supply
		err = storehelpers.GetSubAndSetDecimal(store, SupplyKey(denom), amount)
		if err != nil {
			return nil, err
		}
	}

	return &bank.MsgBurnResponse{}, nil
}

func (s serverImpl) SetPrecision(goCtx context.Context, req *bank.MsgSetPrecisionRequest) (*bank.MsgSetPrecisionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.key)

	minter, err := sdk.AccAddressFromBech32(req.Minter)
	if err != nil {
		return nil, err
	}

	denom := req.Denom
	allowedMinter := store.Get(DenomAdminKey(denom))
	if !bytes.Equal(minter, allowedMinter) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to mint coins for denom %s", minter, denom)
	}

	key := MaxDecimalPlacesKey(denom)
	x, err := storehelpers.GetUint32(store, key)
	if err != nil {
		return nil, err
	}

	if req.MaxDecimalPlaces <= x {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("maximum decimal can only be increased, it is currently %d, and %d was requested", x, req.MaxDecimalPlaces))
	}

	err = storehelpers.SetUInt32(store, key, req.MaxDecimalPlaces)
	if err != nil {
		return nil, err
	}

	return &bank.MsgSetPrecisionResponse{}, nil
}
