package group

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
)

func TestConditionUnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		json          string
		wantErr       *errors.Error
		wantCondition Condition
	}{
		"default decoding": {
			json:          `"foo/bar/636f6e646974696f6e64617461"`,
			wantCondition: NewCondition("foo", "bar", []byte("conditiondata")),
		},
		"invalid condition format": {
			json:    `"foo/636f6e646974696f6e64617461"`,
			wantErr: ErrInvalid,
		},
		"invalid condition data": {
			json:    `"foo/bar/zzzzz"`,
			wantErr: ErrInvalid,
		},
		"zero address": {
			json:          `""`,
			wantCondition: nil,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			var got Condition
			err := json.Unmarshal([]byte(tc.json), &got)
			if !tc.wantErr.Is(err) {
				t.Fatalf("got error: %+v", err)
			}
			if err == nil && !got.Equals(tc.wantCondition) {
				t.Fatalf("expected %q but got condition: %q", tc.wantCondition, got)
			}
		})
	}
}

func TestConditionMarshalJSON(t *testing.T) {
	cases := map[string]struct {
		source   Condition
		wantJson string
	}{
		"cond encoding": {
			source:   NewCondition("foo", "bar", []byte("conditiondata")),
			wantJson: `"foo/bar/636F6E646974696F6E64617461"`,
		},
		"nil encoding": {
			source:   nil,
			wantJson: `""`,
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			got, err := json.Marshal(tc.source)
			assert.Nil(t, err)
			assert.Equal(t, tc.wantJson, string(got))
		})
	}
}
