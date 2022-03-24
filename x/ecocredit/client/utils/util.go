package utils

import (
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ParseDate parses a date using the format yyyy-mm-dd.
func ParseDate(field string, date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, sdkerrors.ErrInvalidRequest.Wrapf("%s must have format yyyy-mm-dd, but received %v", field, date)
	}
	return t, nil
}

// ParseAndSetDate is as helper function which sets the time do the provided argument if
// the ParseDate was successful.
func ParseAndSetDate(dest **time.Time, field string, date string) error {
	t, err := ParseDate(field, date)
	if err == nil {
		*dest = &t
	}
	return err
}
