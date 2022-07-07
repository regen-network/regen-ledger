package core

import (
	"fmt"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreditType_Validate(t *testing.T) {
	type fields struct {
		Abbreviation string
		Name         string
		Unit         string
		Precision    uint32
	}
	tests := []struct {
		name   string
		fields fields
		errMsg string
	}{
		{
			name: "valid",
			fields: fields{
				Abbreviation: "C",
				Name:         "carbon",
				Unit:         "ton",
				Precision:    6,
			},
		},
		{
			name: "invalid: no abbreviation",
			fields: fields{
				Abbreviation: "",
			},
			errMsg: "credit type abbreviation must be 1-3 uppercase latin letters",
		},
		{
			name: "invalid: abbreviation too long",
			fields: fields{
				Abbreviation: "CARB",
			},
			errMsg: "credit type abbreviation must be 1-3 uppercase latin letters",
		},
		{
			name: "invalid no name",
			fields: fields{
				Abbreviation: "C",
				Name:         "",
			},
			errMsg: "name cannot be empty",
		},
		{
			name: "invalid name too long name",
			fields: fields{
				Abbreviation: "C",
				Name:         strings.Repeat("x", maxCreditTypeNameLength+1),
			},
			errMsg: fmt.Sprintf("credit type name cannot exceed %d characters", maxCreditTypeNameLength),
		},
		{
			name: "invalid no unit",
			fields: fields{
				Abbreviation: "C",
				Name:         "carbon",
			},
			errMsg: "unit cannot be empty",
		},
		{
			name: "invalid precision",
			fields: fields{
				Abbreviation: "C",
				Name:         "carbon",
				Unit:         "ton",
				Precision:    3,
			},
			errMsg: "credit type precision is currently locked to 6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := CreditType{
				Abbreviation: tt.fields.Abbreviation,
				Name:         tt.fields.Name,
				Unit:         tt.fields.Unit,
				Precision:    tt.fields.Precision,
			}
			err := m.Validate()
			if len(tt.errMsg) != 0 {
				assert.ErrorContains(t, err, tt.errMsg)
			} else {
				assert.NilError(t, err)
			}
		})
	}
}
