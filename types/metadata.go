package types

import (
	"encoding/base64"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func DecodeMetadata(metadataStr string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(metadataStr)
	if err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("metadata is malformed, proper base64 string is required")
	}

	return b, nil
}
