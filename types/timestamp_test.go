package types_test

import (
	"reflect"
	"testing"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/regen-network/regen-ledger/types"
)

func TestGogoToProtobufDuration(t *testing.T) {
	type args struct {
		d *gogotypes.Duration
	}
	tests := []struct {
		name string
		args args
		want *durationpb.Duration
	}{
		{
			name: "valid",
			args: args{d: &gogotypes.Duration{Seconds: 45, Nanos: 4}},
			want: &durationpb.Duration{Seconds: 45, Nanos: 4},
		},
		{
			name: "nil case",
			args: args{d: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		args, want := tt.args, tt.want
		t.Run(tt.name, func(t *testing.T) {
			if got := types.GogoToProtobufDuration(args.d); !reflect.DeepEqual(got, want) {
				t.Errorf("GogoToProtobufDuration() = %v, want %v", got, want)
			}
		})
	}
}

func TestGogoToProtobufTimestamp(t *testing.T) {
	type args struct {
		ts *gogotypes.Timestamp
	}
	tests := []struct {
		name string
		args args
		want *timestamppb.Timestamp
	}{
		{
			name: "valid",
			args: args{ts: &gogotypes.Timestamp{Seconds: 45, Nanos: 3}},
			want: &timestamppb.Timestamp{Seconds: 45, Nanos: 3},
		},
		{
			name: "nil",
			args: args{ts: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		args, want := tt.args, tt.want
		t.Run(tt.name, func(t *testing.T) {
			if got := types.GogoToProtobufTimestamp(args.ts); !reflect.DeepEqual(got, want) {
				t.Errorf("GogoToProtobufTimestamp() = %v, want %v", got, want)
			}
		})
	}
}

func TestProtobufToGogoTimestamp(t *testing.T) {
	type args struct {
		ts *timestamppb.Timestamp
	}
	tests := []struct {
		name string
		args args
		want *gogotypes.Timestamp
	}{
		{
			name: "valid",
			args: args{ts: &timestamppb.Timestamp{Seconds: 45, Nanos: 3}},
			want: &gogotypes.Timestamp{Seconds: 45, Nanos: 3},
		},
		{
			name: "nil",
			args: args{ts: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		args, want := tt.args, tt.want
		t.Run(tt.name, func(t *testing.T) {
			if got := types.ProtobufToGogoTimestamp(args.ts); !reflect.DeepEqual(got, want) {
				t.Errorf("ProtobufToGogoTimestamp() = %v, want %v", got, want)
			}
		})
	}
}

func TestParseDate(t *testing.T) {
	tcs := []struct {
		name   string
		date   string
		hasErr bool
	}{
		{"good", "2022-01-20", false},
		{"bad", "01-2021-20", true},
	}
	for _, tc := range tcs {
		date, hasErr := tc.date, tc.hasErr
		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)
			tm, err := types.ParseDate(date, date)
			if hasErr {
				require.Error(err)
				require.Equal(time.Time{}, tm)
			} else {
				require.NoError(err)
				require.NotNil(tm)
			}
		})
	}
}
