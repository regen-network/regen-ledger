package data

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgStoreGraph stores a graph in binary format with the specified hash
type MsgStoreGraph struct {
	Hash   []byte         `json:"hash"`
	Data   []byte         `json:"data"`
	Signer sdk.AccAddress `json:"signer"`
}

// NewMsgStoreGraph creates a MsgStoreGraph
func NewMsgStoreGraph(hash []byte, data []byte, signer sdk.AccAddress) MsgStoreGraph {
	return MsgStoreGraph{Hash: hash, Data: data, Signer: signer}
}

// Route returns the Msg route
func (msg MsgStoreGraph) Route() string { return "data" }

// Type returns the Msg type
func (msg MsgStoreGraph) Type() string { return "store_data" }

// ValidateBasic performs basic validation
func (msg MsgStoreGraph) ValidateBasic() sdk.Error {
	if len(msg.Hash) == 0 {
		return sdk.ErrUnknownRequest("Hash cannot be empty")
	}
	if len(msg.Data) == 0 {
		return sdk.ErrUnknownRequest("Data cannot be empty")
	}
	return nil
}

// GetSignBytes gets bytes to sign over
func (msg MsgStoreGraph) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners returns the Msg signers
func (msg MsgStoreGraph) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
