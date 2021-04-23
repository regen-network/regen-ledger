package testdata

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/jsonpb"
)

var _ sdk.Msg = &MsgAuthenticated{}

func (m MsgAuthenticated) Route() string { return "MsgAuthenticated" }

func (m MsgAuthenticated) Type() string { return "Msg Authenticated" }

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgAuthenticated) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgAuthenticated) ValidateBasic() error {
	return nil
}
