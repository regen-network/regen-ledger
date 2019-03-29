package esp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
)

//type SchemaType int
//
//const (
//	JSONSchema SchemaType = 1
//)

type ESPVersionSpec struct {
	//SchemaType SchemaType
	//Schema string
	Curator   sdk.AccAddress   `json:"curator"`
	Name      string           `json:"name"`
	Version   string           `json:"version"`
	Verifiers []sdk.AccAddress `json:"verifiers"`
	// TODO
	// * compute functions run by oracles vs group verifiers
	// * commission for verifying ESP that goes to curator - simple WASM function
	//   taking hectares and dates as input
	// * can we include cost of verification or is that in another subsystem?
}

type ESPResult struct {
	Curator  sdk.AccAddress   `json:"curator"`
	Name     string           `json:"name"`
	Version  string           `json:"version"`
	Verifier sdk.AccAddress   `json:"verifier"`
	GeoID    types.GeoAddress `json:"geo_id,omitempty"`
	Data     []byte           `json:"data"`
	// TODO link this to data module for either on or off-chain result storage
}
