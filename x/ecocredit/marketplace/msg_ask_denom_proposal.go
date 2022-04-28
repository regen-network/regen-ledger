package marketplace

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ govtypes.Content = &AskDenomProposal{}

const (
	ProposalType = "AskDenomProposal"
)

func init() {
	govtypes.RegisterProposalType(ProposalType)
}

func (m *AskDenomProposal) ProposalRoute() string { return ecocredit.RouterKey }

func (m *AskDenomProposal) ProposalType() string { return ProposalType }

func (m *AskDenomProposal) ValidateBasic() error {
	if m.AskDenom == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("ask_denom cannot be nil")
	}
	if err := m.AskDenom.Validate(); err != nil {
		return err
	}
	return govtypes.ValidateAbstract(m)
}

func (m *AskDenomProposal) String() string {
	return fmt.Sprintf(`Credit Type Proposal:
  Title:       %s
  Description: %s
  Ask Denom: %v
`, m.Title, m.Description, m.AskDenom)
}
