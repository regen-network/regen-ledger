package math

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BenchmarkSdkIntTrim(b *testing.B) {
	s := "12345678901234567890.12345678901234567890"
	d, err := NewDecFromString(s)
	if err != nil {
		b.Error("can't convert test number")
	}

	b.Run("exp", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d.SdkIntTrim()
		}
	})

	b.Run("quo-integer", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sdkIntTrimQuo(d)
		}
	})

	b.Run("string", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			sdkIntTrimNaive(d)
		}
	})

}

func sdkIntTrimQuo(d Dec) sdk.Int {
	d, err := d.QuoInteger(NewDecFromInt64(1))
	if err != nil {
		panic(err)
	}

	i, err := d.BigInt()
	if err != nil {
		panic(err)
	}
	return sdk.NewIntFromBigInt(i)
}

func sdkIntTrimNaive(d Dec) sdk.Int {
	d, err := d.QuoInteger(NewDecFromInt64(1))
	if err != nil {
		panic(err)
	}

	s := d.String()
	i, ok := sdk.NewIntFromString(s)
	if !ok {
		panic("can't convert from string")
	}
	return i
}
