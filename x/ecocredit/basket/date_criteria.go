package basket

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
)

// ToApi converts to pulsar based data structure
func (d *DateCriteria) ToApi() *basketv1.DateCriteria {
	if d == nil {
		return nil
	}
	// NOTE(@aaronc): I would choose a marshal/unmarshal approach to
	// automatically handle any new cases that get added to the oneof.
	// This is faster but more brittle to changes. We have no performance reason
	// to do this because basket creation is very infrequent.
	if x := d.GetMinStartDate(); x != nil {
		return &basketv1.DateCriteria{Sum: &basketv1.DateCriteria_MinStartDate{
			&timestamppb.Timestamp{Seconds: x.Seconds, Nanos: x.Nanos}}}
	} else if x := d.GetStartDateWindow(); x != nil {
		return &basketv1.DateCriteria{Sum: &basketv1.DateCriteria_StartDateWindow{
			&durationpb.Duration{Seconds: x.Seconds, Nanos: x.Nanos}}}
	}
	return nil
}
