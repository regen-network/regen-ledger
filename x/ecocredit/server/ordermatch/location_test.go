package ordermatch

import (
	"testing"

	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
)

func TestMatchLocations(t *testing.T) {
	type args struct {
		project   *ecocreditv1beta1.ProjectInfo
		locations []string
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
			if got := matchLocations(tt.args.project, tt.args.locations); got != tt.want {
				t.Errorf("matchLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}
