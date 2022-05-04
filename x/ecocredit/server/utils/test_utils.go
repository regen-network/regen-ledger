package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/mocks"
)

// ExpectParamGet is a helper function that sets up an expected mock call for the provided type.
// Once we switch to Go 1.18+ we can switch this impl to be generic:
// func ExpectParamGet[T any](obj *T, paramKeeper *mocks.MockParamKeeper, times int) {
//	gmAny := gomock.Any()
//	paramKeeper.EXPECT().Get(gmAny, gmAny, gmAny).Do(func(_, _ any, param *T) {
//		*param = *obj
//	}).AnyTimes()
// }
func ExpectParamGet(obj interface{}, paramKeeper *mocks.MockParamKeeper, key []byte, times int) {
	gmAny := gomock.Any()
	switch obj.(type) {
	case *[]string:
		s := obj.(*[]string)
		var expected []string
		paramKeeper.EXPECT().Get(gmAny, key, &expected).Do(func(_ sdk.Context, _ []byte, param *[]string) {
			*param = *s
		}).Times(times)
	case *sdk.Coins:
		coins := obj.(*sdk.Coins)
		var expected sdk.Coins
		paramKeeper.EXPECT().Get(gmAny, key, &expected).Do(func(_ sdk.Context, _ []byte, param *sdk.Coins) {
			*param = *coins
		}).Times(times)
	case *bool:
		b := obj.(*bool)
		var expected bool
		paramKeeper.EXPECT().Get(gmAny, key, &expected).Do(func(_ sdk.Context, _ []byte, param *bool) {
			*param = *b
		}).Times(times)
	case *[]*core.AskDenom:
		askDenoms := obj.(*[]*core.AskDenom)
		var expected []*core.AskDenom
		paramKeeper.EXPECT().Get(gmAny, key, &expected).Do(func(_ sdk.Context, _ []byte, param *[]*core.AskDenom) {
			*param = *askDenoms
		}).Times(times)
	}
}
