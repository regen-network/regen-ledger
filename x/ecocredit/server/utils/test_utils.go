package utils

import (
	"github.com/golang/mock/gomock"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
)

// ExpectParamGet is a helper function that sets up an expected mock call for the provided type.
func ExpectParamGet[T any](obj *T, paramKeeper *mocks.MockParamKeeper, key []byte, times int) {
	gmAny := gomock.Any()
	var expectedType T
	paramKeeper.EXPECT().Get(gmAny, key, &expectedType).Do(func(_ sdk.Context, _ []byte, param *T) {
		*param = *obj
	}).Times(times)
}
