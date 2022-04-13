package basket

import (
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types"
)

// ToApi converts to pulsar based data structure
func (d *DateCriteria) ToApi() *api.DateCriteria {
	if d == nil {
		return nil
	}
	if x := d.GetMinStartDate(); x != nil {
		return &api.DateCriteria{MinStartDate: types.GogoToProtobufTimestamp(x)}
	} else if x := d.GetStartDateWindow(); x != nil {
		return &api.DateCriteria{StartDateWindow: types.GogoToProtobufDuration(x)}
	}
	return nil
}
