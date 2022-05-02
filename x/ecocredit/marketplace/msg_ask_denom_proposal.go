package marketplace

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ govtypes.Content = &AskDenomProposal{}

const (
	AskDenomProposalType = "AskDenomProposal"
)

func init() {
	govtypes.RegisterProposalType(AskDenomProposalType)
}

func (m AskDenomProposal) ProposalRoute() string { return ecocredit.RouterKey }

func (m AskDenomProposal) ProposalType() string { return AskDenomProposalType }

func (m AskDenomProposal) ValidateBasic() error {
	if m.AllowedDenom == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("allowed_denom cannot be nil")
	}
	if err := m.AllowedDenom.Validate(); err != nil {
		return err
	}
	return govtypes.ValidateAbstract(&m)
}

func (m AskDenomProposal) String() string {
	return fmt.Sprintf(`Credit Type Proposal:
  Title:       %s
  Description: %s
  Allowed Denom: %v
`, m.Title, m.Description, m.AllowedDenom)
}
