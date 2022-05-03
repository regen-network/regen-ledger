package marketplace

import (
	"testing"
)

func TestAskDenomProposal_ValidateBasic(t *testing.T) {
	type fields struct {
		Title        string
		Description  string
		AllowedDenom *AllowedDenom
	}
	validAskDenom := &AllowedDenom{
		BankDenom:    "uregen",
		DisplayDenom: "regen",
		Exponent:     18,
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				Title:        "foo",
				Description:  "bar",
				AllowedDenom: validAskDenom,
			},
			wantErr: false,
		},
		{
			name: "no title",
			fields: fields{
				Description:  "foo",
				AllowedDenom: validAskDenom,
			},
			wantErr: true,
		},
		{
			name: "no desc",
			fields: fields{
				Title:        "foo",
				AllowedDenom: validAskDenom,
			},
			wantErr: true,
		},
		{
			name: "no ask denom",
			fields: fields{
				Title:       "foo",
				Description: "bar",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AllowedDenomProposal{
				Title:        tt.fields.Title,
				Description:  tt.fields.Description,
				AllowedDenom: tt.fields.AllowedDenom,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
