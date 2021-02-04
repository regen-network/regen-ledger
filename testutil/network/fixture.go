package network

import (
	"context"
	"fmt"
	"testing"

	"github.com/spf13/pflag"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/regen-network/regen-ledger/testutil/server"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
)

type fixtureFactory struct {
	t      *testing.T
	config Config
}

func NewFixtureFactory(t *testing.T, config Config) *fixtureFactory {
	return &fixtureFactory{t: t, config: config}
}

var _ server.FixtureFactory = fixtureFactory{}

func (f fixtureFactory) Setup(setupHooks ...func(cdc *codec.ProtoCodec, app *baseapp.BaseApp)) server.Fixture {
	if len(setupHooks) != 0 {
		f.t.Fatal("setup hooks are not supported by ")
	}

	network := New(f.t, f.config)
	_, err := network.WaitForHeight(1)
	require.NoError(f.t, err)
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
	req, ok := args.(sdk.MsgRequest)
	if !ok {
		return fmt.Errorf("expected %T, got %T", (*sdk.MsgRequest)(nil), args)
	}

	signers := req.GetSigners()
	if len(signers) != 1 {
		return fmt.Errorf("execpted exactly 1 signer, got %+v", signers)
	}

	signer := signers[0]
	for _, val := range c.network.Validators {
		if val.Address.Equals(signer) {
			clientCtx := val.ClientCtx
			msg := &sdk.ServiceMsg{
				MethodName: method,
				Request:    req,
			}

			clientCtx = clientCtx.
				WithFromAddress(signer).
				WithFromName(val.Moniker).
				WithSkipConfirmation(true).
				WithBroadcastMode("block")

			txf := tx.NewFactoryCLI(clientCtx, &pflag.FlagSet{})
			err := tx.BroadcastTx(clientCtx, txf, msg)
			if err != nil {
				return err
			}
		}
	}
	return fmt.Errorf("do not know how to sign transactions with address %s", signer)
}
