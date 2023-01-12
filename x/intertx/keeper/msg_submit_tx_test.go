package keeper

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"

	v1 "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

func TestSubmitTx(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	owner := s.addrs[0]
	receiver := s.addrs[1]
	sendAmt := sdk.NewCoins(sdk.NewInt64Coin("uregen", 10))
	connectionID := "ch-5"

	msgSend := banktypes.MsgSend{
		FromAddress: owner.String(),
		ToAddress:   receiver.String(),
		Amount:      sendAmt,
	}
	anyMsg, err := types.NewAnyWithValue(&msgSend)
	assert.NilError(t, err)

	msg := v1.MsgSubmitTx{
		Owner:        owner.String(),
		ConnectionId: connectionID,
		Msg:          anyMsg,
	}

	portID, err := icatypes.NewControllerPortID(msg.Owner)
	assert.NilError(t, err)
	channelID := "ch-1"
	serializedTx, err := icatypes.SerializeCosmosTx(s.cdc, []sdk.Msg{&msgSend})
	assert.NilError(t, err)
	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: serializedTx,
	}
	blockTime := time.Unix(35235, 30)
	s.sdkCtx.WithBlockTime(blockTime)
	timeOut := s.sdkCtx.BlockTime().Add(time.Minute).UnixNano()
	capability := &capabilitytypes.Capability{Index: 32}
	gomock.InOrder(
		s.ica.EXPECT().
			GetActiveChannelID(s.sdkCtx, msg.ConnectionId, portID).
			Return(channelID, true).
			Times(1),
		s.cap.EXPECT().
			GetCapability(s.sdkCtx, host.ChannelCapabilityPath(portID, channelID)).
			Return(capability, true).
			Times(1),
		s.ica.EXPECT().
			SendTx(s.sdkCtx, capability, msg.ConnectionId, portID, packetData, uint64(timeOut)).
			Return(uint64(0), nil).
			Times(1),
	)

	res, err := s.k.SubmitTx(s.ctx, &msg)
	assert.NilError(t, err)
	assert.Check(t, res != nil)
}
