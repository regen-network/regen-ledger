package v1

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type feeParams struct {
	gocuke.TestingT
	params *FeeParams
	err    error
}

func TestFeeParams(t *testing.T) {
	gocuke.NewRunner(t, &feeParams{}).Path("./features/state_fee_params.feature").Run()
}

func (s *feeParams) Before() {
	s.params = &FeeParams{}
}

func (s *feeParams) BuyerPercentageFee(a string) {
	s.params.BuyerPercentageFee = a
}

func (s *feeParams) SellerPercentageFee(a string) {
	s.params.SellerPercentageFee = a
}

func (s *feeParams) IValidateTheFeeParams() {
	s.err = s.params.Validate()
}

func (s *feeParams) ExpectErrorToBeTrue() {
	require.Error(s, s.err)
}

func (s *feeParams) ExpectErrorToBeFalse() {
	require.NoError(s, s.err)
}
