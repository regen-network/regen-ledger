package ordermatch

import (
	"testing"

	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestMatchDates(t *testing.T) {
	type args struct {
		batch    *ecocreditv1beta1.BatchInfo
		minStart *timestamppb.Timestamp
		maxEnd   *timestamppb.Timestamp
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchDates(tt.args.batch, tt.args.minStart, tt.args.maxEnd); got != tt.want {
				t.Errorf("matchDates() = %v, want %v", got, tt.want)
			}
		})
	}
}
