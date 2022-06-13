package core

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestOriginTx_Validate(t *testing.T) {
	testCases := []struct {
		name string
		msg  OriginTx
		err  error
	}{
		{
			name: "valid",
			msg: OriginTx{
				Id:     "0x123",
				Source: "0x123",
			},
		},
		{
			name: "no id",
			msg:  OriginTx{},
			err:  fmt.Errorf("invalid OriginTx: no id"),
		},
		{
			name: "no source",
			msg: OriginTx{
				Id: "0x123",
			},
			err: fmt.Errorf("invalid OriginTx: no source"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.Validate()
			if tc.err != nil {
				assert.ErrorContains(t, err, tc.err.Error())
			} else {
				assert.NilError(t, err)
			}
		})
	}
}
