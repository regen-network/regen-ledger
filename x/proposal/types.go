package proposal

import sdk "github.com/cosmos/cosmos-sdk/types"

type Proposal struct {
	Proposer  sdk.AccAddress   `json:"proposer"`
	Action    ProposalAction   `json:"action"`
	Approvers []sdk.AccAddress `json:"approvers,omitempty"`
}

type ProposalAction interface {
	// Return the message type.
	// Must be alphanumeric or empty.
	Route() string

	// Returns a human-readable string for the message, intended for utilization
	// within tags
	Type() string

	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	ValidateBasic() sdk.Error

	// Get the canonical byte representation of the Msg.
	GetSignBytes() []byte

	//HandleProposal(ctx sdk.Context, votes []sdk.AccAddress) sdk.Result
}

type ProposalHandler interface {
	// Returns a true or false value as to whether the handler can handle this type of
	// proposal and a result that either can contain Tags for indexing for valid proposals
	// or error messsages for invalidate proposals
	CheckProposal(ctx sdk.Context, action ProposalAction) (bool, sdk.Result)
	HandleProposal(ctx sdk.Context, action ProposalAction, voters []sdk.AccAddress) sdk.Result
}
