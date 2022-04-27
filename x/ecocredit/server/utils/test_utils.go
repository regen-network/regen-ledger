package utils

import (
	"github.com/golang/mock/gomock"

	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
)

func ExpectParamGet[T any](obj *T, paramKeeper *mocks.MockParamKeeper, times int) {
	gmAny := gomock.Any()
	paramKeeper.EXPECT().Get(gmAny, gmAny, gmAny).Do(func(a, b interface{}, param *T) {
		*param = *obj
	}).Times(times)
}
