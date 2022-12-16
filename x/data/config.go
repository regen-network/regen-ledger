package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func init() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
}

const (
	DefaultIRIPrefix  = "regen"
	DefaultIRIVersion = iriVersion0
)

// Config is a config struct used for initializing the data module to avoid using globals.
type Config struct {
	// IRIPrefix defines the IRI prefix to use (e.g regen).
	IRIPrefix string
	// IRIVersion defines the IRI version to use (e.g 0).
	IRIVersion byte
}

// DefaultConfig returns the default config for the data module.
func DefaultConfig() Config {
	return Config{
		IRIPrefix:  DefaultIRIPrefix,
		IRIVersion: DefaultIRIVersion,
	}
}
