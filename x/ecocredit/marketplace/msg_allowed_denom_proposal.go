package marketplace

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ govtypes.Content = &AllowDenomProposal{}

const (
	AllowDenomProposalType = "AllowDenomProposal"
)

func init() {
	govtypes.RegisterProposalType(AllowDenomProposalType)
	govtypes.RegisterProposalTypeCodec(&AllowDenomProposal{}, "regen/AllowDenomProposal")
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
	return govtypes.ValidateAbstract(&m)
}

func (m AllowDenomProposal) String() string {
	return fmt.Sprintf(`Credit Type Proposal:
  Title:       %s
  Description: %s
  Allowed Denom: %v
`, m.Title, m.Description, m.Denom)
}
