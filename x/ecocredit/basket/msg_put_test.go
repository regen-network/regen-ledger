package basket

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestMsgPut_ValidateBasic(t *testing.T) {
	t.Parallel()

	_, _, addr := testdata.KeyTestPubAddr()
	t1, t2 := time.Now(), time.Now()
	denom, err := core.FormatDenom("C02", 1, &t1, &t2)
	require.NoError(t, err)

	type fields struct {
		Owner       string
		BasketDenom string
		Credits     []*BasketCredit
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: "COOL",
				Credits:     []*BasketCredit{{BatchDenom: denom, Amount: "100.5302"}},
			},
		},
		{
			name: "bad addr",
			fields: fields{
				Owner:       "oops!",
				BasketDenom: "COOL",
				Credits:     []*BasketCredit{{BatchDenom: denom, Amount: "100.5302"}},
			},
			wantErr: true,
		},
		{
			name: "no credits",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: "COOL",
			},
			wantErr: true,
		},
		{
			name: "bad batch denom",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: "COOL",
				Credits:     []*BasketCredit{{BatchDenom: "bad bad not good!", Amount: "100.5302"}},
			},
			wantErr: true,
		},
		{
			name: "bad amount",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: "COOL",
				Credits:     []*BasketCredit{{BatchDenom: denom, Amount: "100.52.302.35.2"}},
			},
			wantErr: true,
		},
		{
			name: "zero amount",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: "COOL",
				Credits:     []*BasketCredit{{BatchDenom: denom, Amount: "0"}},
			},
			wantErr: true,
		},
		{
			name: "negative amount",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: "COOL",
				Credits:     []*BasketCredit{{BatchDenom: denom, Amount: "-50.329"}},
			},
			wantErr: true,
		},
		{
			name: "bad basket denom",
			fields: fields{
				Owner:       addr.String(),
				BasketDenom: "CO:OL",
				Credits:     []*BasketCredit{{BatchDenom: denom, Amount: "100"}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := MsgPut{
				Owner:       tt.fields.Owner,
				BasketDenom: tt.fields.BasketDenom,
				Credits:     tt.fields.Credits,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
