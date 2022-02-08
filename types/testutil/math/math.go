package mathtestutil

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/regen-network/regen-ledger/types/math"
)

func MatchDecFromString(x string) gomock.Matcher {
	dec, err := math.NewDecFromString(x)
	if err != nil {
		panic(err)
	}
	return DecMatcher{dec}
}

func MatchDecFromInt64(x int64) gomock.Matcher {
	return DecMatcher{math.NewDecFromInt64(x)}
}

type DecMatcher struct {
	math.Dec
}

func (d DecMatcher) Matches(x interface{}) bool {
	return x.(math.Dec).Equal(d.Dec)
}

var _ gomock.Matcher = DecMatcher{}

func MatchInt(x int64) gomock.Matcher {
	return IntMatcher{sdk.NewInt(x)}
}

type IntMatcher struct {
	sdk.Int
}

func (i IntMatcher) Matches(x interface{}) bool {
	return x.(sdk.Int).Equal(i.Int)
}

var _ gomock.Matcher = IntMatcher{}
