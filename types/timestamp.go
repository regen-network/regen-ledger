package types

import (
	gogotypes "github.com/gogo/protobuf/types"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GogoToProtobufTimestamp converts a gogo timestamp to a protobuf timestamp.
func GogoToProtobufTimestamp(ts *gogotypes.Timestamp) *timestamppb.Timestamp {
	if ts == nil {
		return &timestamppb.Timestamp{}
	}
	return &timestamppb.Timestamp{
		Seconds: ts.Seconds,
		Nanos:   ts.Nanos,
	}
}

// ProtobufToGogoTimestamp converts a protobuf timestamp to a gogo timestamp.
func ProtobufToGogoTimestamp(ts *timestamppb.Timestamp) *gogotypes.Timestamp {
	if ts == nil {
		return &gogotypes.Timestamp{}
	}
	return &gogotypes.Timestamp{
		Seconds: ts.Seconds,
		Nanos:   ts.Nanos,
	}
}

// GogoToProtobufDuration converts a gogo duration to a protobuf duration.
func GogoToProtobufDuration(d *gogotypes.Duration) *durationpb.Duration {
	if d == nil {
		return &durationpb.Duration{}
	}
	return &durationpb.Duration{
		Seconds: d.Seconds,
		Nanos:   d.Nanos,
	}
}
