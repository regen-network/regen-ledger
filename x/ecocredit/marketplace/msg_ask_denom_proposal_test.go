package marketplace

import "testing"

func TestAskDenomProposal_ValidateBasic(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		AskDenom    *AskDenom
	}
	validAskDenom := &AskDenom{
		Denom:        "uregen",
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
				AskDenom:    validAskDenom,
			},
			wantErr: false,
		},
		{
			name: "no title",
			fields: fields{
				Description: "foo",
				AskDenom:    validAskDenom,
			},
			wantErr: true,
		},
		{
			name: "no desc",
			fields: fields{
				Title:    "foo",
				AskDenom: validAskDenom,
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
			m := &AskDenomProposal{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				AskDenom:    tt.fields.AskDenom,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
