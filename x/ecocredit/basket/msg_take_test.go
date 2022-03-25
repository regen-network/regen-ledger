package basket

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

func TestMsgTakeValidateBasic(t *testing.T) {
	t.Parallel()

	_, _, addr := testdata.KeyTestPubAddr()

	type fields struct {
		Owner              string
		BasketDenom        string
		Amount             string
		RetirementLocation string
		RetireOnTake       bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid message",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "BCT",
				Amount:             "1234",
				RetirementLocation: "US-WA",
				RetireOnTake:       true,
			},
		},
		{
			name: "valid message - do not retire",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "BCT",
				Amount:             "1234",
				RetirementLocation: "",
				RetireOnTake:       false,
			},
		},
		{
			name: "invalid owner address - empty",
			fields: fields{
				Owner:              "",
				BasketDenom:        "BCT",
				Amount:             "1234",
				RetirementLocation: "US-WA",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
		{
			name: "invalid owner address - format",
			fields: fields{
				Owner:              "foo",
				BasketDenom:        "BCT",
				Amount:             "1234",
				RetirementLocation: "US-WA",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
		{
			name: "invalid basket denom - empty",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "",
				Amount:             "1234",
				RetirementLocation: "US-WA",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
		{
			name: "invalid basket denom - format",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "foo!bar",
				Amount:             "1234",
				RetirementLocation: "US-WA",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
		{
			name: "invalid amount - empty",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "BCT",
				Amount:             "",
				RetirementLocation: "US-WA",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
		{
			name: "invalid amount - empty",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "BCT",
				Amount:             "",
				RetirementLocation: "US-WA",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
		{
			name: "invalid amount - not integer",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "BCT",
				Amount:             "12.34",
				RetirementLocation: "US-WA",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
		{
			name: "invalid amount - format",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "BCT",
				Amount:             "12.34.56",
				RetirementLocation: "US-WA",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
		{
			name: "invalid amount - zero",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "BCT",
				Amount:             "0",
				RetirementLocation: "US-WA",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
		{
			name: "invalid amount - negative",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "BCT",
				Amount:             "-1234",
				RetirementLocation: "US-WA",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
		{
			name: "invalid retirement location - empty",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "BCT",
				Amount:             "1234",
				RetirementLocation: "",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
		{
			name: "invalid retirement location - format",
			fields: fields{
				Owner:              addr.String(),
				BasketDenom:        "BCT",
				Amount:             "1234",
				RetirementLocation: "foo-bar",
				RetireOnTake:       true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := MsgTake{
				Owner:              tt.fields.Owner,
				BasketDenom:        tt.fields.BasketDenom,
				Amount:             tt.fields.Amount,
				RetirementLocation: tt.fields.RetirementLocation,
				RetireOnTake:       tt.fields.RetireOnTake,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
