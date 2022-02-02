package orderbook

import ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"

func matchLocation(project *ecocreditv1beta1.ProjectInfo, filter string) bool {
	target := project.ProjectLocation

	n := len(filter)

	// if the target is shorter than the filter than we know we don't have a match
	if len(target) < n {
		return false
	}

	// if the filter length is less than 2 we should match anything (the only filters less than 2 should be totally empty)
	if n < 2 {
		return true
	}

	// check if country matches
	if target[:2] != filter[:2] {
		return false
	}

	panic("TODO")
}
