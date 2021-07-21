package server

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s serverImpl) getCreditType(ctx sdk.Context, creditTypeName string) (*ecocredit.CreditType, error) {
	var params ecocredit.Params
	s.paramSpace.GetParamSet(ctx, &params)
	creditTypes := params.CreditTypes

	for _, creditType := range creditTypes {
		if creditType.Type == creditTypeName {
			return creditType, nil
		}
	}
	return nil, fmt.Errorf("%s is not a valid credit type", creditTypeName)
}
