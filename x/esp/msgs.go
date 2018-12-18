package esp

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/regen-network/regen-ledger/x/agent"
)

type MsgRegisterESPVersion struct {
	Curator agent.AgentId
	Name string
	Version string
	Spec ESPVersionSpec
	Signers []sdk.AccAddress
}

type MsgReportESPResult struct {
	Curator agent.AgentId
	Name string
	Version string
	Verifier agent.AgentId
	Result ESPResult
	Signers []sdk.AccAddress
}

func (msg MsgRegisterESPVersion) Route() string { return "esp" }

func (msg MsgRegisterESPVersion) Type() string { return "register" }

func (msg MsgRegisterESPVersion) ValidateBasic() sdk.Error {
	return nil
}

func (msg MsgRegisterESPVersion) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgRegisterESPVersion) GetSigners() []sdk.AccAddress {
	return msg.Signers
}

func (msg MsgReportESPResult) Route() string { return "esp" }

func (msg MsgReportESPResult) Type() string { return "report_result" }

func (msg MsgReportESPResult) ValidateBasic() sdk.Error {
	// TODO validate schema

	return nil
}

func (msg MsgReportESPResult) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgReportESPResult) GetSigners() []sdk.AccAddress {
	return msg.Signers
}
