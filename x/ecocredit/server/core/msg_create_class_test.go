package core

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type createClassSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	aliceBalance     sdk.Coin
	params           core.Params
	creditTypeAbbrev string
	classId          string
	res              *core.MsgCreateClassResponse
	err              error
}

func TestCreateClass(t *testing.T) {
	gocuke.NewRunner(t, &createClassSuite{}).Path("./features/msg_create_class.feature").Run()
}

func (s *createClassSuite) Before(t gocuke.TestingT) {
	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")

	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.creditTypeAbbrev = "C"
	s.classId = "C01"
}

func (s *createClassSuite) ACreditType() {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *createClassSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
	})
	require.NoError(s.t, err)
}

func (s *createClassSuite) AllowlistEnabled(a string) {
	allowlistEnabled, err := strconv.ParseBool(a)
	require.NoError(s.t, err)

	s.params.AllowlistEnabled = allowlistEnabled
}

func (s *createClassSuite) AliceIsAnApprovedCreditClassCreator() {
	s.params.AllowedClassCreators = append(s.params.AllowedClassCreators, s.alice.String())
}

func (s *createClassSuite) AllowedCreditClassFee(a string) {
	creditClassFee, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s.t, err)

	s.params.CreditClassFee = creditClassFee
}

func (s *createClassSuite) AliceHasATokenBalance(a string) {
	balance, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.aliceBalance = balance
}

func (s *createClassSuite) AClassSequenceWithCreditTypeAndNextSequence(a string, b string) {
	nextSequence, err := strconv.ParseUint(b, 10, 64)
	require.NoError(s.t, err)

	err = s.stateStore.ClassSequenceTable().Insert(s.ctx, &api.ClassSequence{
		CreditTypeAbbrev: a,
		NextSequence:     nextSequence,
	})
	require.NoError(s.t, err)
}

func (s *createClassSuite) AliceAttemptsToCreateACreditClass() {
	s.createClassExpectCalls()

	s.res, s.err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createClassSuite) AliceAttemptsToCreateACreditClassWithCreditType(a string) {
	s.createClassExpectCalls()

	s.res, s.err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		CreditTypeAbbrev: a,
	})
}

func (s *createClassSuite) AliceAttemptsToCreateACreditClassWithIssuers(a gocuke.DocString) {
	var issuers []string

	err := json.Unmarshal([]byte(a.Content), &issuers)
	require.NoError(s.t, err)

	s.createClassExpectCalls()

	s.res, s.err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Issuers:          issuers,
	})
}

func (s *createClassSuite) AliceAttemptsToCreateACreditClassWithProperties(a gocuke.DocString) {
	var msg *core.MsgCreateClass

	err := json.Unmarshal([]byte(a.Content), &msg)
	require.NoError(s.t, err)

	s.createClassExpectCalls()

	s.res, s.err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		CreditTypeAbbrev: msg.CreditTypeAbbrev,
		Metadata:         msg.Metadata,
	})
}

func (s *createClassSuite) AliceAttemptsToCreateACreditClassWithFee(a string) {
	fee, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.createClassExpectCalls()

	s.res, s.err = s.k.CreateClass(s.ctx, &core.MsgCreateClass{
		Admin:            s.alice.String(),
		CreditTypeAbbrev: s.creditTypeAbbrev,
		Fee:              &fee,
	})
}

func (s *createClassSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createClassSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *createClassSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *createClassSuite) ExpectClassSequenceWithCreditTypeAndNextSequence(a string, b string) {
	nextSequence, err := strconv.ParseUint(b, 10, 64)
	require.NoError(s.t, err)

	classSequence, err := s.stateStore.ClassSequenceTable().Get(s.ctx, a)
	require.NoError(s.t, err)

	require.Equal(s.t, nextSequence, classSequence.NextSequence)
}

func (s *createClassSuite) ExpectClassIssuers(a gocuke.DocString) {
	var issuers []string

	err := json.Unmarshal([]byte(a.Content), &issuers)
	require.NoError(s.t, err)

	class, err := s.stateStore.ClassTable().GetById(s.ctx, s.classId)
	require.NoError(s.t, err)

	for _, issuer := range issuers {
		account, err := sdk.AccAddressFromBech32(issuer)
		require.NoError(s.t, err)

		exists, err := s.stateStore.ClassIssuerTable().Has(s.ctx, class.Key, account)
		require.NoError(s.t, err)
		require.True(s.t, exists)
	}
}

func (s *createClassSuite) ExpectClassProperties(a gocuke.DocString) {
	var expected *api.Class

	err := json.Unmarshal([]byte(a.Content), &expected)
	require.NoError(s.t, err)

	class, err := s.stateStore.ClassTable().GetById(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.CreditTypeAbbrev, class.CreditTypeAbbrev)
	require.Equal(s.t, expected.Metadata, class.Metadata)
}

func (s *createClassSuite) ExpectTheResponse(a gocuke.DocString) {
	var res *core.MsgCreateClassResponse

	err := json.Unmarshal([]byte(a.Content), &res)
	require.NoError(s.t, err)

	require.Equal(s.t, res, s.res)
}

func (s *createClassSuite) createClassExpectCalls() {
	var allowlistEnabled bool
	var allowedClassCreators []string
	var creditClassFee sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyAllowlistEnabled, &allowlistEnabled).
		Do(func(ctx sdk.Context, key []byte, allowlistEnabled *bool) {
			*allowlistEnabled = s.params.AllowlistEnabled
		}).
		AnyTimes() // not expected on failed attempt

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyAllowedClassCreators, &allowedClassCreators).
		Do(func(ctx sdk.Context, key []byte, allowedClassCreators *[]string) {
			*allowedClassCreators = s.params.AllowedClassCreators
		}).
		AnyTimes() // not expected on failed attempt

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyCreditClassFee, &creditClassFee).
		Do(func(ctx sdk.Context, key []byte, creditClassFee *sdk.Coins) {
			*creditClassFee = s.params.CreditClassFee
		}).
		AnyTimes() // not expected on failed attempt

	var expectedFee sdk.Coin
	var expectedFees sdk.Coins

	if len(s.params.CreditClassFee) == 1 {
		expectedFee = s.params.CreditClassFee[0]
		expectedFees = sdk.Coins{expectedFee}
	}

	if len(s.params.CreditClassFee) == 2 {
		expectedFee = s.params.CreditClassFee[1]
		expectedFees = sdk.Coins{expectedFee}
	}

	s.bankKeeper.EXPECT().
		GetBalance(s.sdkCtx, s.alice, expectedFee.Denom).
		Return(s.aliceBalance).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromAccountToModule(s.sdkCtx, s.alice, ecocredit.ModuleName, expectedFees).
		Do(func(sdk.Context, sdk.AccAddress, string, sdk.Coins) {
			if s.params.CreditClassFee != nil {
				// simulate token balance update unavailable with mocks
				s.aliceBalance = s.aliceBalance.Sub(expectedFee)
			}
		}).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		BurnCoins(s.sdkCtx, ecocredit.ModuleName, expectedFees).
		Return(nil).
		AnyTimes() // not expected on failed attempt
}
