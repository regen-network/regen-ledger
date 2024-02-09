package testutil

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

// GetEvent serches an event in sdk.Events matching message proto name in the reverse order.
// Returns event and true if the event was found and false otherwise.
func GetEvent(msg proto.Message, events []sdk.Event) (sdk.Event, bool) {
	eventName := proto.MessageName(msg)
	for i := len(events) - 1; i >= 0; i-- {
		if eventName == events[i].Type {
			return events[i], true
		}
	}
	return sdk.Event{}, false
}
