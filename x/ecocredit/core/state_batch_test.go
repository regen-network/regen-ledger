package core

import (
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"
)

func TestBatchValidate(t *testing.T) {
	tests := []struct {
		name   string
		batch  Batch
		errMsg string
	}{
		{
			name: "valid",
			batch: Batch{
				Denom:      "C01-001-20200101-20210101-001",
				ProjectKey: 1,
				StartDate:  &types.Timestamp{Seconds: 1},
				EndDate:    &types.Timestamp{Seconds: 2},
				Issuer:     []byte("BTZfSbi0JKqguZ/tIAPUIhdAa7Y="),
			},
		},
		{
			name:   "empty denom",
			batch:  Batch{},
			errMsg: "batch denom cannot be empty: parse error",
		},
		{
			name: "invalid denom",
			batch: Batch{
				Denom: "foo",
			},
			errMsg: "invalid batch denom: expected format A00-000-00000000-00000000-000: parse error",
		},
		{
			name: "empty project key",
			batch: Batch{
				Denom: "C01-001-20200101-20210101-001",
			},
			errMsg: "project key cannot be zero: invalid request",
		},
		{
			name: "empty start date",
			batch: Batch{
				Denom:      "C01-001-20200101-20210101-001",
				ProjectKey: 1,
			},
			errMsg: "must provide a start date for the credit batch: invalid request",
		},
		{
			name: "empty end date",
			batch: Batch{
				Denom:      "C01-001-20200101-20210101-001",
				ProjectKey: 1,
				StartDate:  &types.Timestamp{Seconds: 1},
			},
			errMsg: "must provide an end date for the credit batch: invalid request",
		},
		{
			name: "start date equal to end date",
			batch: Batch{
				Denom:      "C01-001-20200101-20210101-001",
				ProjectKey: 1,
				StartDate:  &types.Timestamp{Seconds: 1},
				EndDate:    &types.Timestamp{Seconds: 1},
			},
			errMsg: "the batch end date (1970-01-01T00:00:01Z) must be the same as or after the batch start date (1970-01-01T00:00:01Z): invalid request",
		},
		{
			name: "start date after end date",
			batch: Batch{
				Denom:      "C01-001-20200101-20210101-001",
				ProjectKey: 1,
				StartDate:  &types.Timestamp{Seconds: 2},
				EndDate:    &types.Timestamp{Seconds: 1},
			},
			errMsg: "the batch end date (1970-01-01T00:00:01Z) must be the same as or after the batch start date (1970-01-01T00:00:02Z): invalid request",
		},
		{
			name: "empty issuer",
			batch: Batch{
				Denom:      "C01-001-20200101-20210101-001",
				ProjectKey: 1,
				StartDate:  &types.Timestamp{Seconds: 1},
				EndDate:    &types.Timestamp{Seconds: 2},
			},
			errMsg: "issuer: empty address string is not allowed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.batch.Validate()
			if tt.errMsg != "" {
				require.EqualError(t, err, tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
