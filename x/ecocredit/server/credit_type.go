package server

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s serverImpl) getCreditType(ctx sdk.Context, creditTypeName string) (*ecocredit.CreditType, error) {
	creditTypes := s.getAllCreditTypes(ctx)
	creditTypeName = ecocredit.NormalizeCreditTypeName(creditTypeName)
	for _, creditType := range creditTypes {
		if creditType.Name == creditTypeName {
			return creditType, nil
		}
	}
	return nil, fmt.Errorf("%s is not a valid credit type", creditTypeName)
}

func (s serverImpl) getAllCreditTypes(ctx sdk.Context) []*ecocredit.CreditType {
	var params ecocredit.Params
	s.paramSpace.GetParamSet(ctx, &params)
	return params.CreditTypes
}
