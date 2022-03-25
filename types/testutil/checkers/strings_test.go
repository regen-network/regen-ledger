package structvalid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrMaxLen(t *testing.T) {
	require := require.New(t)
	require.Nil(StrMaxLen("abc", "", 0, nil))
	require.Nil(StrMaxLen("abc", "a", 1, nil))
	require.Nil(StrMaxLen("abc", "a", 10, nil))

	errs := StrMaxLen("abc1", "xyz1", 2, nil)
	require.Len(errs, 1)

	errs = StrMaxLen("abc2", "xyz2", 4, errs)
	require.Len(errs, 1)

	errs = StrMaxLen("abc3", "xyz3", 3, errs)
	require.Len(errs, 2)

	errStr := ErrsToError(errs).Error()
	require.Contains(errStr, "abc1")
	require.Contains(errStr, "abc3")
	require.NotContains(errStr, "abc2")
}
