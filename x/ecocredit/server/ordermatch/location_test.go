package ordermatch

import (
	"testing"
)

func TestMatchLocations(t *testing.T) {
	tests := []struct {
		name     string
		location string
		filters  []string
		want     bool
	}{
		{
			"empty filters",
			"US-NY 01000",
			nil,
			true,
		},
		{
			"country filter",
			"US-NY 01000",
			[]string{"FR", "US"},
			true,
		},
		{
			"country filter false",
			"US-NY 01000",
			[]string{"FR", "DE"},
			false,
		},
		{
			"country, sub-div filter",
			"US-NY 01000",
			[]string{"FR", "US-NY"},
			true,
		},
		{
			"country, sub-div filter false",
			"US-NY 01000",
			[]string{"FR", "US-CA"},
			false,
		},
		{
			"country, sub-div, postal filter",
			"US-NY 01000",
			[]string{"FR", "US-NY 12345 ", "US-NY 01000"},
			true,
		},
		{
			"country, sub-div, postal filter false",
			"US-NY 01000",
			[]string{"FR", "US-NY 12345 "},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchLocation(tt.location, tt.filters); got != tt.want {
				t.Errorf("matchLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}
