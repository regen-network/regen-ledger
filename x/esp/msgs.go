package esp

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/regen-network/regen-ledger/x/agent"
)

type ActionRegisterESPVersion struct {
	Curator agent.AgentID `json:"curator"`
	Name string `json:"name"`
	Version string `json:"version"`
	Spec ESPVersionSpec `json:"spec"`
}

type ActionReportESPResult struct {
	Curator agent.AgentID `json:"curator"`
	Name string `json:"name"`
	Version string `json:"version"`
	Verifier agent.AgentID `json:"verifier"`
	Result ESPResult `json:"result"`
}

func NewActionRegisterESPVersion(curator agent.AgentID, name string, version string, spec ESPVersionSpec) ActionRegisterESPVersion {
	return ActionRegisterESPVersion{
		Curator:curator,
		Name:name,
		Version:version,
		Spec:spec,
	}

}

func (msg ActionRegisterESPVersion) Route() string { return "esp" }

func (msg ActionRegisterESPVersion) Type() string { return "register" }

func (msg ActionRegisterESPVersion) ValidateBasic() sdk.Error {
	return nil
}

func (msg ActionRegisterESPVersion) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg ActionReportESPResult) Route() string { return "esp" }

func (msg ActionReportESPResult) Type() string { return "report_result" }

func (msg ActionReportESPResult) ValidateBasic() sdk.Error {
	// TODO validate schema

	return nil
}

func (msg ActionReportESPResult) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}
