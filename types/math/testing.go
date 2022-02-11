package math

import (
	"fmt"

	"github.com/golang/mock/gomock"
)

// MatchEq is a gomock.Matcher which compares a decimal value to an
// expected value in gomock calls.
func MatchEq(dec Dec) gomock.Matcher {
	return &decEq{dec: dec}
}

type decEq struct {
	dec Dec
	msg string
}

func (d *decEq) Matches(x interface{}) bool {
	d.msg = ""
	y, ok := x.(Dec)
	if !ok {
		return false
	}

	if y.Cmp(d.dec) != 0 {
		d.msg = fmt.Sprintf("%s != %s", d.dec, y)
		return false
	} else {
		return true
	}
}

func (d *decEq) String() string {
	return d.msg
}
