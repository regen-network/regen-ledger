package group_test

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/modules/incubator/group"
	"github.com/stretchr/testify/assert"
)

func TestConditionUnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		json          string
		wantErr       *errors.Error
		wantCondition group.Condition
	}{
		"default decoding": {
			json:          `"foo/bar/636f6e646974696f6e64617461"`,
			wantCondition: group.NewCondition("foo", "bar", []byte("conditiondata")),
		},
		"invalid condition format": {
			json:    `"foo/636f6e646974696f6e64617461"`,
			wantErr: group.ErrInvalid,
		},
		"invalid condition data": {
			json:    `"foo/bar/zzzzz"`,
			wantErr: group.ErrInvalid,
		},
		"zero address": {
			json:          `""`,
			wantCondition: nil,
		},
	}

	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			var got group.Condition
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
		source   group.Condition
		wantJson string
	}{
		"cond encoding": {
			source:   group.NewCondition("foo", "bar", []byte("conditiondata")),
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
