package ordermatch

import (
	"testing"
	"time"

	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestMatchDates(t *testing.T) {
	batch := &ecocreditv1beta1.BatchInfo{
		StartDate: timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
		EndDate:   timestamppb.New(time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)),
	}
	tests := []struct {
		name     string
		minStart *timestamppb.Timestamp
		maxEnd   *timestamppb.Timestamp
		want     bool
	}{
		{
			"no filters",
			nil,
			nil,
			true,
		},
		{
			"start",
			timestamppb.New(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
			nil,
			true,
		},
		{
			"end",
			nil,
			timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			true,
		},
		{
			"start and end",
			timestamppb.New(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
			timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			true,
		},
		{
			"start false",
			timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			nil,
			false,
		},
		{
			"end false",
			nil,
			timestamppb.New(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
			false,
		},
		{
			"start and end false",
			timestamppb.New(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)),
			timestamppb.New(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchDates(batch, tt.minStart, tt.maxEnd); got != tt.want {
				t.Errorf("matchDates() = %v, want %v", got, tt.want)
			}
		})
	}
}
