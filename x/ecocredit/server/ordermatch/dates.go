package ordermatch

import (
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func matchDates(batch *ecocreditv1beta1.BatchInfo, minStart, maxEnd *timestamppb.Timestamp) bool {
	// TODO: use cosmos-proto timepb support functions
	if minStart != nil && batch.StartDate.AsTime().Before(minStart.AsTime()) {
		return false
	}

	if maxEnd != nil && batch.EndDate.AsTime().After(maxEnd.AsTime()) {
		return false
	}

	return true
}
