package esp

import "gitlab.com/regen-network/regen-ledger/x/agent"

type SchemaType int

const (
	JSONSchema SchemaType = 1
)

type ESPVersionSpec struct {
	//SchemaType SchemaType
	//Schema string
	Verifiers []agent.AgentID `json:"verifiers"`
}

type ESPResult struct {
	// TODO use geo keeper to save space
	PolygonEWKB []byte `json:"polygon_ewkb,omitempty"`
	Data []byte `json:"data"`
}
