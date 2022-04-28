package utils

import (
	"github.com/golang/mock/gomock"

	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
)

// ExpectParamGet is a helper function that sets up an expected mock call for the provided type.
func ExpectParamGet[T any](obj *T, paramKeeper *mocks.MockParamKeeper, times int) {
	gmAny := gomock.Any()
	paramKeeper.EXPECT().Get(gmAny, gmAny, gmAny).Do(func(_, _ any, param *T) {
		*param = *obj
	}).Times(times)
}
