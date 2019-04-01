package claim

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
)

// MsgSignClaim creates a transaction which indicates that the Signers of the
// indicated Claim are willing to "sign" the being, to the best of their
// knowledge true, optionally linking to some Evidence to support that
// assessment.
type MsgSignClaim struct {
	// Claim is the claim which is being signed. It must point to an on or
	// off chain graph stored or tracked by the data module.
	Claim types.DataAddress
	// Evidence is optional data which is being pointed to as evidence to the
	// Claim's veracity by the Signers of the Claim. It can point to on or
	// off chain graphs as well as raw data tracked by the data module.
	Evidence []types.DataAddress
	// Signers are the signers of this Claim
	Signers []sdk.AccAddress
}

// Implements Msg.
func (msg MsgSignClaim) Route() string { return "claim" }

// Implements Msg.
func (msg MsgSignClaim) Type() string { return "claim.sign" }

// Implements Msg.
func (msg MsgSignClaim) ValidateBasic() sdk.Error {
	if len(msg.Claim) == 0 {
		return sdk.ErrUnknownRequest("Claim cannot be empty")
	}
	if !types.IsGraphDataAddress(msg.Claim) {
		return sdk.ErrUnknownRequest("Claim must point to graph data")
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
