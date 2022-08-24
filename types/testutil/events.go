package testutil

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"
)

// MatchEvent matches the values in a proto message to a sdk.Event.
func MatchEvent(expected proto.Message, emitted sdk.Event) error {
	msg, err := sdk.ParseTypedEvent(abci.Event(emitted))
	if err != nil {
		return err
	}
	equal := proto.Equal(expected, msg)
	if !equal {
		return fmt.Errorf("expected %s\ngot %s", expected.String(), msg.String())
	}
	return nil
}

// GetEvent searches through the sdk.Events and attempts to return the sdk.Event whose type matches
// the proto.MessageName of the passed msg.
func GetEvent(msg proto.Message, events []sdk.Event) (e sdk.Event, found bool) {
	eventName := proto.MessageName(msg)
	for _, e := range events {
		if eventName == e.Type {
			return e, true
		}
	}
	return e, false
}
