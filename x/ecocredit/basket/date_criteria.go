package basket

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
)

// ToApi converts to pulsar based data structure
func (d *DateCriteria) ToApi() *api.DateCriteria {
	if d == nil {
		return nil
	}
	if x := d.GetMinStartDate(); x != nil {
		return &api.DateCriteria{MinStartDate: &timestamppb.Timestamp{Seconds: x.Seconds, Nanos: x.Nanos}}
	} else if x := d.GetStartDateWindow(); x != nil {
		return &api.DateCriteria{StartDateWindow: &durationpb.Duration{Seconds: x.Seconds, Nanos: x.Nanos}}
	}
	return nil
}
