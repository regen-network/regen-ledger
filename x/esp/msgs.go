package esp

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateESP struct {
	Curator []byte
	Name string
	Signers []sdk.AccAddress
}

type MsgRegisterESPVersion struct {
	Curator []byte
	Name string
	Version string
	Schema string
	SchemaType SchemaType
	Signers []sdk.AccAddress
}

type MsgReportESPResult struct {
	Curator []byte
	Name string
	Version string
	Verifier []byte
	Data []byte
	PolygonEWKB []byte
	Signers []sdk.AccAddress
}

type SchemaType int

const (
	JSONSchema SchemaType = 1
)

func NewMsgCreateESP(name string, org []byte, signers []sdk.AccAddress) MsgCreateESP {
	return MsgCreateESP{
		Name:name,
		Org:org,
		Signers:signers,
	}
}

func (msg MsgCreateESP) Route() string { return "esp" }

func (msg MsgCreateESP) Type() string { return "create_esp" }

func (msg MsgCreateESP) ValidateBasic() sdk.Error {
	return nil
}

func (msg MsgCreateESP) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgCreateESP) GetSigners() []sdk.AccAddress {
	return msg.Signers
}

func (msg MsgRegisterESPVersion) Route() string { return "esp" }

func (msg MsgRegisterESPVersion) Type() string { return "register_esp_version" }

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
