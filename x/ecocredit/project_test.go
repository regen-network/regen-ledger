package ecocredit

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProjectID(t *testing.T) {
	testCases := []struct {
		name      string
		projectID string
		expectErr bool
	}{
		{
			"valid project id",
			"A123",
			false,
		},
		{
			"invalid project id min length",
			"a",
			true,
		},
		{
			"invalid project id max length",
			"abcdef123456ghijklmnop789",
			true,
		},
		{
			"invalid project id special characters",
			"abcd@1",
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := validateProjectID(tc.projectID); tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
