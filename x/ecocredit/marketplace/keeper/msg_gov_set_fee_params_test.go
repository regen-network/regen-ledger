package keeper

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/google/go-cmp/cmp"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"

	marketplacev1 "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/types/v1"
)

type govSetFeeParams struct {
	*baseSuite
	err error
	msg *types.MsgGovSetFeeParams
}

func TestGovSetFeeParams(t *testing.T) {
	gocuke.NewRunner(t, &govSetFeeParams{}).
		Path("./features/msg_gov_set_fee_params.feature").
		Step(`^fee\s+params\s+\x60([^\x60]*)\x60$`, (*govSetFeeParams).FeeParamsInline).
		Run()
}

func (s *govSetFeeParams) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t, 2)
	s.msg = &types.MsgGovSetFeeParams{}
}

func (s *govSetFeeParams) Authority(a string) {
	s.msg.Authority = a
}

func (s *govSetFeeParams) FeeParams(a gocuke.DocString) {
	if a.Content != "" {
		s.msg.Fees = &types.FeeParams{}
		require.NoError(s.t, jsonpb.UnmarshalString(a.Content, s.msg.Fees))
	}
}

func (s *govSetFeeParams) FeeParamsInline(a string) {
	if a != "" {
		s.msg.Fees = &types.FeeParams{}
		require.NoError(s.t, jsonpb.UnmarshalString(a, s.msg.Fees))
	}
}

func (s *govSetFeeParams) TheMessage(a gocuke.DocString) {
	s.msg = &types.MsgGovSetFeeParams{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *govSetFeeParams) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *govSetFeeParams) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *govSetFeeParams) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *govSetFeeParams) AuthorityIsSetToTheKeeperAuthority() {
	s.msg.Authority = s.k.authority.String()
}

func (s *govSetFeeParams) AuthorityIsSetTo(a string) {
	s.msg.Authority = a
}

func (s *govSetFeeParams) FeeParamsAreSet() {
	_, s.err = s.k.GovSetFeeParams(s.ctx, s.msg)
}

func (s *govSetFeeParams) ExpectErrorContains(a string) {
	if s.err != nil {
		require.ErrorContains(s.t, s.err, a)
	} else {
		require.NoError(s.t, s.err)
	}
}

func (s *govSetFeeParams) ExpectFeeParams(a gocuke.DocString) {
	var expected marketplacev1.FeeParams
	require.NoError(s.t, jsonpb.UnmarshalString(a.Content, &expected))

	actual, err := s.k.stateStore.FeeParamsTable().Get(s.ctx)
	require.NoError(s.t, err)

	if diff := cmp.Diff(&expected, actual, protocmp.Transform()); diff != "" {
		require.Fail(s.t, "unexpected fee params", diff)
	}
}
