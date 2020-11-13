package ecocredit

import (
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	reDigits = regexp.MustCompile(`\d+`)
)

// BatchID2ClassID validates id has a correct format end extracts class id from the batch denom.
func BatchID2ClassID(id string) (string, error) {
	ids := strings.Split(id, "/")
	if len(ids) != 2 || !reDigits.MatchString(ids[0]) || !reDigits.MatchString(ids[1]) {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"Invalid batch_denom (%s). Valid format: '<number>/<number>'", id)
	}
	return ids[0], nil
}
