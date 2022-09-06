package testutil

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	"github.com/cosmos/cosmos-sdk/x/nft"
)

func TestMatchEvent(t *testing.T) {
	event := nft.EventSend{
		ClassId:  "foo",
		Id:       "bar",
		Sender:   "baz",
		Receiver: "qux",
	}

	sdkEvent, err := sdk.TypedEventToEvent(&event)
	require.NoError(t, err)

	err = MatchEvent(&event, sdkEvent)
	require.NoError(t, err)

	event.Receiver = "fail"
	err = MatchEvent(&event, sdkEvent)
	require.Error(t, err)
}

func TestGetEvent(t *testing.T) {
	events := sdk.Events{}

	event := nft.EventSend{
		ClassId:  "foo",
		Id:       "bar",
		Sender:   "baz",
		Receiver: "qux",
	}
	event2 := group.EventCreateGroup{GroupId: 2}

	sdkEvent, err := sdk.TypedEventToEvent(&event)
	require.NoError(t, err)

	sdkEvent2, err := sdk.TypedEventToEvent(&event2)
	require.NoError(t, err)
	events = append(events, sdkEvent, sdkEvent2)

	gotEvent, found := GetEvent(&event, events)
	require.True(t, found)

	require.Equal(t, gotEvent, sdkEvent)

	gotEvent2, found := GetEvent(&event2, events)
	require.True(t, found)
	require.Equal(t, gotEvent2, sdkEvent2)

	notInEvents := group.EventSubmitProposal{}
	_, found = GetEvent(&notInEvents, events)
	require.False(t, found)
}
