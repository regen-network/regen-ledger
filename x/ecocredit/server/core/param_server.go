package core

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

func (s serverImpl) NewCreditType(ctx context.Context, request *v1beta1.MsgNewCreditTypeRequest) (*v1beta1.MsgNewCreditTypeResponse, error) {
	if err := s.calledByGovernance(request.RootAddress); err != nil {
		return nil, err
	}

	for _, ct := range request.CreditTypes {
		if err := s.stateStore.CreditTypeStore().Insert(ctx, &ecocreditv1beta1.CreditType{
			Name:         ct.Name,
			Abbreviation: ct.Abbreviation,
			Unit:         ct.Unit,
			Precision:    ct.Precision,
		}); err != nil {
			return nil, err
		}
	}

	return &v1beta1.MsgNewCreditTypeResponse{}, nil
}

func (s serverImpl) ToggleAllowList(ctx context.Context, request *v1beta1.MsgToggleAllowListRequest) (*v1beta1.MsgToggleAllowListResponse, error) {
	if err := s.calledByGovernance(request.RootAddress); err != nil {
		return nil, err
	}
	err := s.stateStore.AllowlistEnabledStore().Save(ctx, &ecocreditv1beta1.AllowlistEnabled{Enabled: request.Toggle})
	return &v1beta1.MsgToggleAllowListResponse{}, err
}

func (s serverImpl) calledByGovernance(addr string) error {
	rootAddress := s.accountKeeper.GetModuleAddress(govtypes.ModuleName).String()

	if addr != rootAddress {
		return sdkerrors.ErrUnauthorized.Wrapf("root address must be governance module address, got: %s, expected: %s", addr, rootAddress)
	}

	return nil
}
