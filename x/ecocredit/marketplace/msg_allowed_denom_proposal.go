package marketplace

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ govv1beta1.Content = &AllowDenomProposal{}

const (
	AllowDenomProposalType = "AllowDenomProposal"
)

func init() {
	govv1beta1.RegisterProposalType(AllowDenomProposalType)
}

func (m AllowDenomProposal) ProposalRoute() string { return ecocredit.RouterKey }

func (m AllowDenomProposal) ProposalType() string { return AllowDenomProposalType }

func (m AllowDenomProposal) ValidateBasic() error {
	if m.Denom == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("denom cannot be empty")
	}
	if err := m.Denom.Validate(); err != nil {
		return err
	}
	return govv1beta1.ValidateAbstract(&m)
}

func (m AllowDenomProposal) String() string {
	return fmt.Sprintf(`Credit Type Proposal:
  Title:       %s
  Description: %s
  Allowed Denom: %v
`, m.Title, m.Description, m.Denom)
}
