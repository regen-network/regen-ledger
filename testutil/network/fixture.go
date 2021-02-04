package network

import (
	"context"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/regen-network/regen-ledger/testutil/server"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
)

type fixtureFactory struct {
	config Config
	t      *testing.T
}

var _ server.FixtureFactory = fixtureFactory{}

func (f fixtureFactory) Setup(setupHooks ...func(cdc *codec.ProtoCodec, app *baseapp.BaseApp)) server.Fixture {
	if len(setupHooks) != 0 {
		f.t.Fatal("setup hooks are not supported by ")
	}

	network := New(f.t, f.config)
	return fixture{network: network}
}

type fixture struct {
	network *Network
}

var _ server.Fixture = fixture{}

func (f fixture) Context() context.Context {
	return context.Background()
}

func (f fixture) TxConn() grpc.ClientConnInterface {
	return txConn{fixture: f}
}

func (f fixture) QueryConn() grpc.ClientConnInterface {
	return f.network.Validators[0].ClientCtx
}

func (f fixture) Signers() []sdk.AccAddress {
	var addrs []sdk.AccAddress
	validators := f.network.Validators

	for _, val := range validators {
		addrs = append(addrs, val.Address)
	}

	return addrs
}

func (f fixture) Teardown() {
	f.network.Cleanup()
}

type txConn struct {
	fixture
}

func (c txConn) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("not supported")
}

var _ grpc.ClientConnInterface = txConn{}

func (c txConn) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	clientCtx := c.network.Validators[0].ClientCtx
	return tx.GenerateOrBroadcastTxCLI(clientCtx, nil)
}
