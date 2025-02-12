package app

import (
	simappparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
)

// MakeEncodingConfig returns the application's encoding configuration with all
// types and interfaces registered.
func MakeEncodingConfig() simappparams.EncodingConfig {
	encodingConfig := testutil.MakeTestEncodingConfig()
	// std.RegisterLegacyAminoCodec(encodingConfig.Amino) // removed because it gets called on init of codec/legacy
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return simappparams.EncodingConfig{
		InterfaceRegistry: encodingConfig.InterfaceRegistry,
		Codec:             encodingConfig.Codec,
		TxConfig:          encodingConfig.TxConfig,
		Amino:             encodingConfig.Amino,
	}
}
