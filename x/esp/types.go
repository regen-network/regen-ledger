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
}

type ESPResult struct {
	Curator agent.AgentID `json:"curator"`
	Name string `json:"name"`
	Version string `json:"version"`
	Verifier agent.AgentID `json:"verifier"`
	// TODO maybe use geo keeper to save space with large polygons
	PolygonEWKB []byte `json:"polygon_ewkb,omitempty"`
	Data []byte `json:"data"`
}
