package basket

import (
	"fmt"

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

// Validate validates DateCriteria, ensuring only one field is set,
// minimum start date is after 1900-01-01, start date window is at
// least one day, and years in the past is at least one year
func (d *DateCriteria) Validate() error {
	if d == nil {
		return nil
	}

	minStartDate := d.GetMinStartDate()
	startDateWindow := d.GetStartDateWindow()
	yearsInThePast := d.GetYearsInThePast()

	if (minStartDate != nil && startDateWindow != nil) ||
		(startDateWindow != nil && yearsInThePast != 0) ||
		(minStartDate != nil && yearsInThePast != 0) {
		return fmt.Errorf("only one of min_start_date, start_date_window, or years_in_the_past must be set")
	}

	if minStartDate != nil {
		if minStartDate.Seconds < -2208992400 {
			return fmt.Errorf("min_start_date must be after 1900-01-01")
		}
	} else if startDateWindow != nil {
		if startDateWindow.Seconds < 24*3600 {
			return fmt.Errorf("start_date_window must be at least 1 day")
		}
	}

	return nil
}
