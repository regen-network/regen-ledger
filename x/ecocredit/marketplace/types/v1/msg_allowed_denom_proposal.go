package v1

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

var _ govtypes.Content = &AllowDenomProposal{}

const (
	AllowDenomProposalType = "AllowDenomProposal"
)

func init() {
	govtypes.RegisterProposalType(AllowDenomProposalType)
}

func (m AllowDenomProposal) ProposalRoute() string { return "ecocredit" }

func (m AllowDenomProposal) ProposalType() string { return AllowDenomProposalType }

func (m AllowDenomProposal) ValidateBasic() error {
	if m.Denom == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("denom cannot be empty")
	}
	if err := m.Denom.Validate(); err != nil {
		return err
	}
	return govtypes.ValidateAbstract(&m)
}
