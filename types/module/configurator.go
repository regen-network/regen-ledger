package module

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gogo/protobuf/grpc"
)

// Configurator extends the cosmos sdk Configurator interface
// with BinaryMarshaler()
type Configurator interface {
	sdkmodule.Configurator

	BinaryMarshaler() codec.BinaryMarshaler
}

type configurator struct {
	msgServer   grpc.Server
	queryServer grpc.Server
	cdc         codec.BinaryMarshaler
}

// NewConfigurator returns a new Configurator instance
func NewConfigurator(msgServer grpc.Server, queryServer grpc.Server, cdc codec.BinaryMarshaler) Configurator {
	return configurator{msgServer: msgServer, queryServer: queryServer, cdc: cdc}
}

var _ Configurator = configurator{}

// MsgServer implements the Configurator.MsgServer method
func (c configurator) MsgServer() grpc.Server {
	return c.msgServer
}

// QueryServer implements the Configurator.QueryServer method
func (c configurator) QueryServer() grpc.Server {
	return c.queryServer
}

// BinaryMarshaler implements the Configurator.BinaryMarshaler method
func (c configurator) BinaryMarshaler() codec.BinaryMarshaler {
	return c.cdc
}
