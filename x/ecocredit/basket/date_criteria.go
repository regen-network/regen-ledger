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
	if x := d.GetMinStartDate(); x != nil {
		return &basketv1.DateCriteria{MinStartDate: &timestamppb.Timestamp{Seconds: x.Seconds, Nanos: x.Nanos}}
	} else if x := d.GetStartDateWindow(); x != nil {
		return &basketv1.DateCriteria{StartDateWindow: &durationpb.Duration{Seconds: x.Seconds, Nanos: x.Nanos}}
	}
	return nil
}
