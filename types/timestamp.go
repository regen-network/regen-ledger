package types

import (
	"fmt"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GogoToProtobufTimestamp converts a gogo timestamp to a protobuf timestamp.
func GogoToProtobufTimestamp(ts *gogotypes.Timestamp) *timestamppb.Timestamp {
	if ts == nil {
		return nil
	}
	return &timestamppb.Timestamp{
		Seconds: ts.Seconds,
		Nanos:   ts.Nanos,
	}
}

// ProtobufToGogoTimestamp converts a protobuf timestamp to a gogo timestamp.
func ProtobufToGogoTimestamp(ts *timestamppb.Timestamp) *gogotypes.Timestamp {
	if ts == nil {
		return nil
	}
	return &gogotypes.Timestamp{
		Seconds: ts.Seconds,
		Nanos:   ts.Nanos,
	}
}

// GogoToProtobufDuration converts a gogo duration to a protobuf duration.
func GogoToProtobufDuration(d *gogotypes.Duration) *durationpb.Duration {
	if d == nil {
		return nil
	}
	return &durationpb.Duration{
		Seconds: d.Seconds,
		Nanos:   d.Nanos,
	}
}

// ParseDate parses a date using the format yyyy-mm-dd.
func ParseDate(field string, date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, fmt.Errorf("%s must have format yyyy-mm-dd, but received %v", field, date)
	}
	return t, nil
}
