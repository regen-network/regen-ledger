package server

import (
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pariz/gountries"
)

// validateLocation checks that the country and region conform to ISO 3166 and
// the postal code is valid
func validateLocation(location string) error {
	countryCode, regionCode, postalCode, err := parseLocation(location)
	if err != nil {
		return err
	}

	iso3166query := gountries.New()

	country, err := iso3166query.FindCountryByAlpha(countryCode)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	if regionCode != "" {
		_, err = country.FindSubdivisionByCode(regionCode)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}

	if err := validatePostalCode(postalCode); err != nil {
		return err
	}

	return nil
}

func parseLocation(location string) (string, string, string, error) {
	strings := strings.Split(location, "-")
	switch numStrings := len(strings); numStrings {
	case 1:
		return strings[0], "", "", nil
	case 2:
		return strings[0], strings[1], "", nil
	case 3:
		return strings[0], strings[1], strings[2], nil
	default:
		return "", "", "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Location should have format <country-code>[-<region-code>[-<postal-code>]]")
	}
}

// validatePostalCode currently checks that the postal code is not longer than
// 10 characters, as currently the longest postal codes in the world are 10
// characters
func validatePostalCode(postalCode string) error {
	if len(postalCode) > 10 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The postal code must not be longer than 10 characters")
	}

	return nil
}
