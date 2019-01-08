package proposal

import sdk "github.com/cosmos/cosmos-sdk/types"

type Proposal struct {
	Proposer  sdk.AccAddress `json:"proposer"`
	Action    ProposalAction `json:"action"`
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

	//Handle(ctx sdk.Context, votes []sdk.AccAddress) sdk.Result
}

type ProposalHandler interface {
	CanHandle(action ProposalAction) bool
	Handle(ctx sdk.Context, action ProposalAction, voters []sdk.AccAddress) sdk.Result
}
