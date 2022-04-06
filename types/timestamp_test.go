package types_test

import (
	"reflect"
	"testing"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/regen-network/regen-ledger/types"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
			want: &durationpb.Duration{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := types.GogoToProtobufDuration(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GogoToProtobufDuration() = %v, want %v", got, tt.want)
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
			want: &timestamppb.Timestamp{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := types.GogoToProtobufTimestamp(tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GogoToProtobufTimestamp() = %v, want %v", got, tt.want)
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
			want: &gogotypes.Timestamp{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := types.ProtobufToGogoTimestamp(tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProtobufToGogoTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
