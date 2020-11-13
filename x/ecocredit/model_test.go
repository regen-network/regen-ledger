package ecocredit

import (
	"github.com/stretchr/testify/suite"
)

type ModelSuite struct {
	suite.Suite
}

func (s ModelSuite) TestBatchID2ClassID() {
	require := s.Require()
	testCases := []struct {
		denom, expected string
	}{
		{"", ""},
		{"12", ""},
		{"12/", ""},
		{"/", ""},
		{"/1", ""},
		{"a/1", ""},
		{"1/a", ""},
		{"12/0", "12"},
		{"01/01", "01"},
		{"123213123213213124123123199/1", "123213123213213124123123199"},
	}
	for _, tc := range testCases {
		id, err := BatchID2ClassID(tc.denom)
		if tc.expected == "" {
			require.Error(err)
		} else {
			require.NoError(err)
			require.Equal(id, tc.expected)
		}
	}
}
