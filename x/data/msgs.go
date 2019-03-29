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

func NewMsgStoreGraph(hash []byte, data []byte, signer sdk.AccAddress) MsgStoreGraph {
	return MsgStoreGraph{Hash: hash, Data: data, Signer: signer}
}

func (msg MsgStoreGraph) Route() string { return "data" }

func (msg MsgStoreGraph) Type() string { return "store_data" }

func (msg MsgStoreGraph) ValidateBasic() sdk.Error {
	if len(msg.Hash) == 0 {
		return sdk.ErrUnknownRequest("Hash cannot be empty")
	}
	if len(msg.Data) == 0 {
		return sdk.ErrUnknownRequest("Data cannot be empty")
	}
	return nil
}

func (msg MsgStoreGraph) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgStoreGraph) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}
