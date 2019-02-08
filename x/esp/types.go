package esp

import "gitlab.com/regen-network/regen-ledger/x/agent"

//type SchemaType int
//
//const (
//	JSONSchema SchemaType = 1
//)

type ESPVersionSpec struct {
	//SchemaType SchemaType
	//Schema string
	Curator agent.AgentID `json:"curator"`
	Name string `json:"name"`
	Version string `json:"version"`
	Verifiers []agent.AgentID `json:"verifiers"`
	// TODO
	// * compute functions run by oracles vs agent verifiers
	// * commission for verifying ESP that goes to curator - simple WASM function
	//   taking hectares and dates as input
	// * can we include cost of verification or is that in another subsystem?
}

type ESPResult struct {
	Curator agent.AgentID `json:"curator"`
	Name string `json:"name"`
	Version string `json:"version"`
	Verifier agent.AgentID `json:"verifier"`
	GeoID []byte `json:"geo_id,omitempty"`
	Data []byte `json:"data"`
	// TODO link this to data module for either on or off-chain result storage
}
