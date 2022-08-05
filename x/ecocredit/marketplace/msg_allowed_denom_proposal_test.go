package marketplace

import (
	"testing"
)

func TestAllowDenomProposal_ValidateBasic(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Denom       *AllowedDenom
	}
	validDenom := &AllowedDenom{
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
				Title:       "foo",
				Description: "bar",
				Denom:       validDenom,
			},
			wantErr: false,
		},
		{
			name: "no title",
			fields: fields{
				Description: "foo",
				Denom:       validDenom,
			},
			wantErr: true,
		},
		{
			name: "no desc",
			fields: fields{
				Title: "foo",
				Denom: validDenom,
			},
			wantErr: true,
		},
		{
			name: "no allowed denom",
			fields: fields{
				Title:       "foo",
				Description: "bar",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AllowDenomProposal{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Denom:       tt.fields.Denom,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
