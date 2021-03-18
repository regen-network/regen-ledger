package server

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// isServiceMsg checks if a type URL corresponds to a service method name,
// i.e. /cosmos.bank.Msg/Send vs /cosmos.bank.MsgSend.
// Function copied from cosmos-sdk, once we add ADR 033 and group module to it,
// we'll want to merge them.
func isServiceMsg(typeURL string) bool {
	return strings.Count(typeURL, "/") >= 2
}

// SetMsgs takes a slice of sdk.Msg's and turn them into Any's.
// This is similar to what is in the cosmos-sdk tx builder
// and could eventually be merged in.
func SetMsgs(msgs []sdk.Msg) ([]*types.Any, error) {
	anys := make([]*types.Any, len(msgs))
	for i, msg := range msgs {
		var err error
		switch msg := msg.(type) {
		case sdk.ServiceMsg:
			anys[i], err = types.NewAnyWithCustomTypeURL(msg.Request, msg.MethodName)
		default:
			anys[i], err = types.NewAnyWithValue(msg)
		}
		if err != nil {
			return nil, err
		}
	}
	return anys, nil
}

// SetMsgs takes a slice of Any's and turn them into sdk.Msg's.
// This is similar to what is in the cosmos-sdk sdk.Tx
// and could eventually be merged in.
func GetMsgs(anys []*types.Any) []sdk.Msg {
	msgs := make([]sdk.Msg, len(anys))
	for i, any := range anys {
		var msg sdk.Msg
		if isServiceMsg(any.TypeUrl) {
			req := any.GetCachedValue()
			if req == nil {
				panic("Any cached value is nil. Transaction messages must be correctly packed Any values.")
			}
			msg = sdk.ServiceMsg{
				MethodName: any.TypeUrl,
				Request:    any.GetCachedValue().(sdk.MsgRequest),
			}
		} else {
			msg = any.GetCachedValue().(sdk.Msg)
		}
		msgs[i] = msg
	}
	return msgs
}

func UnpackInterfaces(unpacker types.AnyUnpacker, anys []*types.Any) error {
	for _, any := range anys {
		// If the any's typeUrl contains 2 slashes, then we unpack the any into
		// a ServiceMsg struct as per ADR-031.
		if isServiceMsg(any.TypeUrl) {
			var req sdk.MsgRequest
			err := unpacker.UnpackAny(any, &req)
			if err != nil {
				return err
			}
		} else {
			var msg sdk.Msg
			err := unpacker.UnpackAny(any, &msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
