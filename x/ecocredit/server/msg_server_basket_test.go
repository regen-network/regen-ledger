package server

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFilter(t *testing.T) {

	testCases := []struct {
		name     string
		filters  []*ecocredit.Filter
		batch    ecocredit.BatchInfo
		class    ecocredit.ClassInfo
		basket   ecocredit.Basket
		owner    string
		expErr   bool
		expDepth int
	}{
		{
			name: "depth 1 filter",
			filters: []*ecocredit.Filter{
				{Sum: &ecocredit.Filter_ClassId{ClassId: "foo"}},
			},
			batch: ecocredit.BatchInfo{
				ClassId:         "foo",
				BatchDenom:      "",
				Issuer:          "",
				TotalAmount:     "",
				Metadata:        nil,
				AmountCancelled: "",
				StartDate:       nil,
				EndDate:         nil,
				ProjectLocation: "",
			},
			expDepth: 0,
			expErr:   false,
		},
		{
			name: "depth 2 filter OR",
			filters: []*ecocredit.Filter{
				{
					Sum: &ecocredit.Filter_Or_{Or: &ecocredit.Filter_Or{Filters: []*ecocredit.Filter{
						{
							Sum: &ecocredit.Filter_ClassId{ClassId: "foo"},
						},
						{
							Sum: &ecocredit.Filter_CreditTypeName{CreditTypeName: "doo"},
						},
					},
					}},
				},
			},
			batch: ecocredit.BatchInfo{
				ClassId:         "foo",
				BatchDenom:      "",
				Issuer:          "",
				TotalAmount:     "",
				Metadata:        nil,
				AmountCancelled: "",
				StartDate:       nil,
				EndDate:         nil,
				ProjectLocation: "",
			},
			class: ecocredit.ClassInfo{
				ClassId:  "",
				Admin:    "",
				Issuers:  nil,
				Metadata: nil,
				CreditType: &ecocredit.CreditType{
					Name:         "doo",
					Abbreviation: "",
					Unit:         "",
					Precision:    0,
				},
				NumBatches: 0,
			},
			expDepth: 0,
			expErr:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// classInfo ecocredit.ClassInfo, batchInfo ecocredit.BatchInfo, basketInfo ecocredit.Basket, owner string)
			depth, err := checkFilters(tc.filters, tc.class, tc.batch, tc.basket, tc.owner)
			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expDepth, depth)
			}
		})
	}
}
