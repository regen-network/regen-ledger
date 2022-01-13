package ecocredit

import (
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var reLocation = regexp.MustCompile(`^([A-Z]{2})(?:-([A-Z0-9]{1,3})(?: ([a-zA-Z0-9 \-]{1,64}))?)?$`)

// ValidateLocation checks that the country and region conform to ISO 3166 and
// the postal code is valid. This is a simple regex check and doesn't check that
// the country or subdivision codes actually exist. This is because the codes
// could change at short notice and we don't want to hardfork to keep up-to-date
// with that information.
func ValidateLocation(location string) error {
	matches := reLocation.FindStringSubmatch(location)
	if matches == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid location: %s.\nLocation should have format <country-code>[-<region-code>[ <postal-code>]].\n", location)
	}

	return nil
}
