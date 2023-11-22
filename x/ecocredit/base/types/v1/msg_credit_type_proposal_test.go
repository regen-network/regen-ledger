package v1

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestCreditTypeProposal_ValidateBasic(t *testing.T) {
	validCreditType := &CreditType{"C", "carbon", "carbon ton", 6}
	type fields struct {
		Title       string
		Description string
		CreditType  *CreditType
	}
	tests := []struct {
		name   string
		fields fields
		errMsg string
	}{
		{
			name: "valid",
			fields: fields{
				Title:       "hello",
				Description: "world",
				CreditType:  validCreditType,
			},
		},
		{
			name: "invalid: nil credit type",
			fields: fields{
				Title:       "hi",
				Description: "hello",
				CreditType:  nil,
			},
			errMsg: "credit type cannot be nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &CreditTypeProposal{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				CreditType:  tt.fields.CreditType,
			}
			err := m.ValidateBasic()
			if len(tt.errMsg) != 0 {
				assert.ErrorContains(t, err, tt.errMsg)
			} else {
				assert.NilError(t, err)
			}
		})
	}
}
