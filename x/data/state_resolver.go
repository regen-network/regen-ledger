package data

import (
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate performs basic validation of the Resolver state type
func (m *Resolver) Validate() error {
	if m.Id == 0 {
		return ErrParseFailure.Wrap("id cannot be empty")
	}

	if m.Url == "" {
		return ErrParseFailure.Wrap("url cannot be empty")
	}

	if _, err := url.ParseRequestURI(m.Url); err != nil {
		return ErrParseFailure.Wrap("url: invalid url format")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(m.Manager).String()); err != nil {
		return ErrParseFailure.Wrapf("manager: %s", err)
	}

	return nil
}
