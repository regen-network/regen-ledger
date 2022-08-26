package types

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckDuplicateKey(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		expErr    bool
		expErrMsg string
	}{
		{
			name:      "invalid json",
			input:     "{foo:bar}",
			expErr:    true,
			expErrMsg: "invalid character",
		},
		{
			name:      "invalid json duplicate keys",
			input:     `{"foo": "bar", "foo": "bar"}`,
			expErr:    true,
			expErrMsg: "duplicate key foo",
		},
		{
			name:      "invalid json duplicate nested keys",
			input:     `{"foo": "bar", "baz": [{"foo": "bar", "foo": "baz"}]}`,
			expErr:    true,
			expErrMsg: "duplicate key foo",
		},
		{
			name:  "valid json",
			input: `{"foo": "bar", "baz": "foo"}`,
		},
		{
			name:  "valid json nested",
			input: `{"foo": "bar", "baz": [{"foo": "bar", "baz": "foo"}]}`,
		},
	}

	for _, tc := range testCases {
		input, expErr, expErrMsg := tc.input, tc.expErr, tc.expErrMsg
		t.Run(tc.name, func(t *testing.T) {
			err := CheckDuplicateKey(json.NewDecoder(strings.NewReader(input)), nil)
			if expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), expErrMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
