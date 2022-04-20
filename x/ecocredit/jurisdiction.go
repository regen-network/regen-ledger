package ecocredit

import (
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var reJurisdiction = regexp.MustCompile(`^([A-Z]{2})(?:-([A-Z0-9]{1,3})(?: ([a-zA-Z0-9 \-]{1,64}))?)?$`)

// ValidateJurisdiction checks that the country and region conform to ISO 3166 and
// the postal code is valid. This is a simple regex check and doesn't check that
// the country or subdivision codes actually exist. This is because the codes
// could change at short notice and we don't want to hardfork to keep up-to-date
// with that information.
func ValidateJurisdiction(jurisdiction string) error {
	matches := reJurisdiction.FindStringSubmatch(jurisdiction)
	if matches == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid jurisdiction: %s.\nJurisdiction should have format <country-code>[-<region-code>[ <postal-code>]].\n", jurisdiction)
	}

	return nil
}
