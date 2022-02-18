package basketsims

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"

	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func GetAndShuffleClasses(sdkCtx sdk.Context, r *rand.Rand, qryClient ecocredit.QueryClient) ([]*ecocredit.ClassInfo, error) {
	ctx := regentypes.Context{Context: sdkCtx}
	res, err := qryClient.Classes(ctx, &ecocredit.QueryClassesRequest{})
	if err != nil {
		return nil, err
	}

	classes := res.Classes
	if len(classes) <= 1 {
		return classes, nil
	}

	r.Shuffle(len(classes), func(i, j int) { classes[i], classes[j] = classes[j], classes[i] })
	return classes, nil
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func randomExponent(r *rand.Rand, precision uint32) uint32 {
	exponents := []uint32{0, 1, 2, 3, 6, 9, 12, 15, 18, 21, 24}
	for {
		x := exponents[r.Intn(len(exponents))]
		if x > precision {
			return x
		}
	}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}