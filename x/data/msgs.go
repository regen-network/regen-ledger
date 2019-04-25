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

// MsgTrackRawData tracks off-chain raw data with the specified hash and optional URL.
type MsgTrackRawData struct {
	// Sha256Hash is the SHA-256 hash of the data.
	Sha256Hash []byte
	// Url is an optional Url from which the data can be retrieved. It can be omitted
	// if for some reason, there is a desire to timestamp data without providing any
	// reachable URL for the time-being. It is also not an error to submit MsgTrackRawData
	// multiple times with different URL's but the same hash, each new URL will be stored.
	Url string
	// Signer is the message signer.
	Signer sdk.AccAddress
}

// NewMsgStoreGraph creates a MsgStoreGraph
func NewMsgStoreGraph(hash []byte, data []byte, signer sdk.AccAddress) MsgStoreGraph {
	return MsgStoreGraph{Hash: hash, Data: data, Signer: signer}
}

// Route returns the Msg route
func (msg MsgStoreGraph) Route() string { return "data" }

// Type returns the Msg type
func (msg MsgStoreGraph) Type() string { return "data.store" }

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

// Route implements the Msg interface.
func (msg MsgTrackRawData) Route() string { return "data" }

// Type implements the Msg interface.
func (msg MsgTrackRawData) Type() string { return "data.track-raw" }

// ValidateBasic implements the Msg interface.
func (msg MsgTrackRawData) ValidateBasic() sdk.Error {
	if len(msg.Sha256Hash) == 0 {
		return sdk.ErrUnknownRequest("Sha256Hash cannot be empty")
	}
	return nil
}

// GetSignBytes implements the Msg interface.
func (msg MsgTrackRawData) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners implements the Msg interface.
func (msg MsgTrackRawData) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
