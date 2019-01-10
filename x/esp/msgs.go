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

func (msg ActionRegisterESPVersion) Type() string { return "esp.register_version" }

func (msg ActionRegisterESPVersion) ValidateBasic() sdk.Error {
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if len(msg.Version) == 0 {
		return sdk.ErrUnknownRequest("Version cannot be empty")
	}
	if len(msg.Spec.Verifiers) == 0 {
		return sdk.ErrUnknownRequest("Verifiers cannot be empty")
	}
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

func (msg ActionReportESPResult) Type() string { return "esp.report_result" }

func (msg ActionReportESPResult) ValidateBasic() sdk.Error {
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if len(msg.Version) == 0 {
		return sdk.ErrUnknownRequest("Version cannot be empty")
	}
	if len(msg.Result.Data) == 0 {
		return sdk.ErrUnknownRequest("Result data cannot be empty")
	}
	return nil
}

func (msg ActionReportESPResult) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}
