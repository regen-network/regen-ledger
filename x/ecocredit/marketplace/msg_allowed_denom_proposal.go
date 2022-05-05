package marketplace

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ govtypes.Content = &AllowedDenomProposal{}

const (
	AllowedDenomProposalType = "AllowedDenomProposal"
)

func init() {
	govtypes.RegisterProposalType(AllowedDenomProposalType)
	govtypes.RegisterProposalTypeCodec(&AllowedDenomProposal{}, "regen/AllowedDenomProposal")
}

func (m AllowedDenomProposal) ProposalRoute() string { return ecocredit.RouterKey }

func (m AllowedDenomProposal) ProposalType() string { return AllowedDenomProposalType }

func (m AllowedDenomProposal) ValidateBasic() error {
	if m.Denom == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("allowed denom cannot be nil")
	}
	if err := m.Denom.Validate(); err != nil {
		return err
	}
	return govtypes.ValidateAbstract(&m)
}

func (m AllowedDenomProposal) String() string {
	return fmt.Sprintf(`Credit Type Proposal:
  Title:       %s
  Description: %s
  Allowed Denom: %v
`, m.Title, m.Description, m.Denom)
}
