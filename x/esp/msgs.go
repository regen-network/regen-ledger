package esp

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ActionRegisterESPVersion struct {
	ESPVersionSpec `json:"spec"`
}

type ActionReportESPResult struct {
	ESPResult `json:"result"`
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
	if len(msg.Verifiers) == 0 {
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
	if len(msg.Data) == 0 {
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
