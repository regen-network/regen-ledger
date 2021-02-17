package module

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gogo/protobuf/grpc"
)

// Configurator extends the cosmos sdk Configurator interface
// with Marshaler()
type Configurator interface {
	sdkmodule.Configurator

	Marshaler() codec.Marshaler
}

type configurator struct {
	msgServer   grpc.Server
	queryServer grpc.Server
	cdc         codec.Marshaler
}

// NewConfigurator returns a new Configurator instance
func NewConfigurator(msgServer grpc.Server, queryServer grpc.Server, cdc codec.Marshaler) Configurator {
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

// Marshaler implements the Configurator.Marshaler method
func (c configurator) Marshaler() codec.Marshaler {
	return c.cdc
}
