package server

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/util"
	"strings"
)

func (s serverImpl) getCreditType(ctx sdk.Context, creditTypeName string) (*ecocredit.CreditType, error) {
	creditTypes := s.getAllCreditTypes(ctx)
	creditTypeName = strings.ToLower(creditTypeName)
	creditTypeName = util.FastRemoveWhitespace(creditTypeName)
	for _, creditType := range creditTypes {
		if creditType.Type == creditTypeName {
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
