package claim

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
)

// MsgSignClaim creates a transaction which indicates that the Signers of the
// claim are stating that to the best of their knowledge, the Content of the
// claim is true, optionally linking to some Evidence that supports their
// assessment.
type MsgSignClaim struct {
	// Content is the content of the claim which is being signed
	// - the statement about what is being claimed as true. It must point
	// to an on or off chain graph stored or tracked by the data module.
	Content types.DataAddress
	// Evidence is optional data which is being pointed to as evidence to the
	// Content's veracity by the Signers of the Content. It can point to on or
	// off chain graphs as well as raw data tracked by the data module.
	Evidence []types.DataAddress
	// Signers are the signers of this claim. By signing this claim they
	// are asserting to the best of their knowledge the Content is true
	Signers []sdk.AccAddress
}

// Implements Msg.
func (msg MsgSignClaim) Route() string { return "claim" }

// Implements Msg.
func (msg MsgSignClaim) Type() string { return "claim.sign" }

// Implements Msg.
func (msg MsgSignClaim) ValidateBasic() sdk.Error {
	if len(msg.Content) == 0 {
		return sdk.ErrUnknownRequest("Content cannot be empty")
	}
	if !types.IsGraphDataAddress(msg.Content) {
		return sdk.ErrUnknownRequest("Content must point to graph data")
	}
	return nil
}

// Implements Msg.
func (msg MsgSignClaim) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// Implements Msg.
func (msg MsgSignClaim) GetSigners() []sdk.AccAddress {
	return msg.Signers
}
