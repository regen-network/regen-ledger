package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type govSetFeeParams struct {
	gocuke.TestingT
	msg *MsgGovSetFeeParams
	err error
}

func TestGovSetFeeParams(t *testing.T) {
	gocuke.NewRunner(t, &govSetFeeParams{}).
		Path("./features/msg_gov_set_fee_params.feature").
		Step("^fee\\s+params\\s+`([^`]*)`$", (*govSetFeeParams).FeeParams).
		Run()
}

func (s *govSetFeeParams) Before() {
	s.msg = &MsgGovSetFeeParams{}
}

func (s *govSetFeeParams) Authority(a string) {
	s.msg.Authority = a
}

func (s *govSetFeeParams) FeeParams(a string) {
	if a != "" {
		s.msg.Fees = &FeeParams{}
		require.NoError(s, jsonpb.UnmarshalString(a, s.msg.Fees))
	}
}

func (s *govSetFeeParams) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *govSetFeeParams) ExpectErrorTrue() {
	require.Error(s, s.err)
}

func (s *govSetFeeParams) ExpectErrorFalse() {
	require.NoError(s, s.err)
}
