package testsuite

import (
	"context"
	"math"
	"strings"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/testutil/server"
	groupserver "github.com/regen-network/regen-ledger/x/group/server"
	"github.com/regen-network/regen-ledger/x/group/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory server.FixtureFactory
	fixture        server.Fixture

	ctx       context.Context
	sdkCtx    sdk.Context
	msgClient types.MsgClient
	addr1     sdk.AccAddress
	addr2     sdk.AccAddress

	groupKeeper *groupserver.Keeper
	bankKeeper  bankkeeper.Keeper

	blockTime time.Time
}

func NewIntegrationTestSuite(
	fixtureFactory server.FixtureFactory, groupKeeper *groupserver.Keeper,
	bankKeeper bankkeeper.Keeper) *IntegrationTestSuite {
	return &IntegrationTestSuite{
		fixtureFactory: fixtureFactory,
		groupKeeper:    groupKeeper,
		bankKeeper:     bankKeeper,
	}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.fixture = s.fixtureFactory.Setup()
	s.ctx = s.fixture.Context()

	s.blockTime = time.Now().UTC()

	sdkContext := sdk.UnwrapSDKContext(s.ctx).WithBlockTime(s.blockTime)
	s.groupKeeper.SetParams(sdkContext, types.DefaultParams())
	s.sdkCtx = sdkContext

	s.msgClient = types.NewMsgClient(s.fixture.TxConn())
	if len(s.fixture.Signers()) < 2 {
		s.FailNow("expected at least 2 signers, got %d", s.fixture.Signers())
	}
	s.addr1 = s.fixture.Signers()[0]
	s.addr2 = s.fixture.Signers()[1]
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.fixture.Teardown()
}

func (s *IntegrationTestSuite) TestCreateGroup() {
	members := []types.Member{{
		Address: sdk.AccAddress([]byte("one--member--address")),
		Power:   sdk.NewDec(1),
		Comment: "first",
	}, {
		Address: sdk.AccAddress([]byte("other-member-address")),
		Power:   sdk.NewDec(2),
		Comment: "second",
	}}
	specs := map[string]struct {
		srcAdmin   sdk.AccAddress
		srcMembers []types.Member
		srcComment string
		expErr     bool
	}{
		"all good": {
			srcAdmin:   []byte("valid--admin-address"),
			srcMembers: members,
			srcComment: "test",
		},
		"group comment too long": {
			srcAdmin:   []byte("valid--admin-address"),
			srcMembers: members,
			srcComment: strings.Repeat("a", 256),
			expErr:     true,
		},
		"member comment too long": {
			srcAdmin: []byte("valid--admin-address"),
			srcMembers: []types.Member{{
				Address: []byte("valid-member-address"),
				Power:   sdk.OneDec(),
				Comment: strings.Repeat("a", 256),
			}},
			srcComment: "test",
			expErr:     true,
		},
	}
	var seq uint32
	for msg, spec := range specs {
		s.Run(msg, func() {
			id, err := s.groupKeeper.CreateGroup(s.sdkCtx, spec.srcAdmin, spec.srcMembers, spec.srcComment)
			if spec.expErr {
				s.Require().Error(err)
				s.Require().False(s.groupKeeper.HasGroup(s.sdkCtx, types.GroupID(seq+1).Bytes()))
				return
			}
			s.Require().NoError(err)

			seq++
			s.Assert().Equal(types.GroupID(seq), id)

			// then all data persisted
			loadedGroup, err := s.groupKeeper.GetGroup(s.sdkCtx, id)
			s.Require().NoError(err)
			s.Assert().Equal(sdk.AccAddress([]byte(spec.srcAdmin)), loadedGroup.Admin)
			s.Assert().Equal(spec.srcComment, loadedGroup.Comment)
			s.Assert().Equal(id, loadedGroup.Group)
			s.Assert().Equal(uint64(1), loadedGroup.Version)

			// and members are stored as well
			it, err := s.groupKeeper.GetGroupMembersByGroup(s.sdkCtx, id)
			s.Require().NoError(err)
			var loadedMembers []types.GroupMember
			_, err = orm.ReadAll(it, &loadedMembers)
			s.Require().NoError(err)
			s.Assert().Equal(len(members), len(loadedMembers))
			for i := range loadedMembers {
				s.Assert().Equal(members[i].Comment, loadedMembers[i].Comment)
				s.Assert().Equal(members[i].Address, loadedMembers[i].Member)
				s.Assert().Equal(members[i].Power, loadedMembers[i].Weight)
				s.Assert().Equal(id, loadedMembers[i].Group)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCreateGroupAccount() {
	myGroupID, err := s.groupKeeper.CreateGroup(s.sdkCtx, []byte("valid--admin-address"), nil, "test")
	s.Require().NoError(err)

	specs := map[string]struct {
		srcAdmin   sdk.AccAddress
		srcGroupID types.GroupID
		srcPolicy  types.DecisionPolicy
		srcComment string
		expErr     bool
	}{
		"all good": {
			srcAdmin:   []byte("valid--admin-address"),
			srcComment: "test",
			srcGroupID: myGroupID,
			srcPolicy: types.NewThresholdDecisionPolicy(
				sdk.OneDec(),
				gogotypes.Duration{Seconds: 1},
			),
		},
		"decision policy threshold > total group weight": {
			srcAdmin:   []byte("valid--admin-address"),
			srcComment: "test",
			srcGroupID: myGroupID,
			srcPolicy: types.NewThresholdDecisionPolicy(
				sdk.NewDec(math.MaxInt64),
				gogotypes.Duration{Seconds: 1},
			),
		},
		"group id does not exists": {
			srcAdmin:   []byte("valid--admin-address"),
			srcComment: "test",
			srcGroupID: 9999,
			srcPolicy: types.NewThresholdDecisionPolicy(
				sdk.OneDec(),
				gogotypes.Duration{Seconds: 1},
			),
			expErr: true,
		},
		"admin not group admin": {
			srcAdmin:   []byte("other--admin-address"),
			srcComment: "test",
			srcGroupID: myGroupID,
			srcPolicy: types.NewThresholdDecisionPolicy(
				sdk.OneDec(),
				gogotypes.Duration{Seconds: 1},
			),
			expErr: true,
		},
		"comment too long": {
			srcAdmin:   []byte("valid--admin-address"),
			srcComment: strings.Repeat("a", 256),
			srcGroupID: myGroupID,
			srcPolicy: types.NewThresholdDecisionPolicy(
				sdk.OneDec(),
				gogotypes.Duration{Seconds: 1},
			),
			expErr: true,
		},
	}
	for msg, spec := range specs {
		s.Run(msg, func() {
			addr, err := s.groupKeeper.CreateGroupAccount(s.sdkCtx, spec.srcAdmin, spec.srcGroupID, spec.srcPolicy, spec.srcComment)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			// then all data persisted
			groupAccount, err := s.groupKeeper.GetGroupAccount(s.sdkCtx, addr)
			s.Require().NoError(err)
			s.Assert().Equal(addr, groupAccount.GroupAccount)
			s.Assert().Equal(myGroupID, groupAccount.Group)
			s.Assert().Equal(sdk.AccAddress([]byte(spec.srcAdmin)), groupAccount.Admin)
			s.Assert().Equal(spec.srcComment, groupAccount.Comment)
			s.Assert().Equal(uint64(1), groupAccount.Version)
			// TODO Fix (ORM should unpack Any's properly)
			s.Assert().Equal(&spec.srcPolicy, groupAccount.GetDecisionPolicy())
		})
	}
}

func (s *IntegrationTestSuite) TestCreateProposal() {
	members := []types.Member{{
		Address: []byte("valid-member-address"),
		Power:   sdk.OneDec(),
	}}
	myGroupID, err := s.groupKeeper.CreateGroup(s.sdkCtx, []byte("valid--admin-address"), members, "test")
	s.Require().NoError(err)

	policy := types.NewThresholdDecisionPolicy(
		sdk.OneDec(),
		gogotypes.Duration{Seconds: 1},
	)
	accountAddr, err := s.groupKeeper.CreateGroupAccount(s.sdkCtx, []byte("valid--admin-address"), myGroupID, policy, "test")
	s.Require().NoError(err)

	policy = types.NewThresholdDecisionPolicy(
		sdk.NewDec(math.MaxInt64),
		gogotypes.Duration{Seconds: 1},
	)
	bigThresholdAddr, err := s.groupKeeper.CreateGroupAccount(s.sdkCtx, []byte("valid--admin-address"), myGroupID, policy, "test")
	s.Require().NoError(err)

	specs := map[string]struct {
		srcAccount   sdk.AccAddress
		srcProposers []sdk.AccAddress
		srcMsgs      []sdk.Msg
		srcComment   string
		expErr       bool
	}{
		"all good with minimal fields set": {
			srcAccount:   accountAddr,
			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
		},
		"all good with good msg payload": {
			srcAccount:   accountAddr,
			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
			srcMsgs:      []sdk.Msg{sdktestdata.NewServiceMsgCreateDog(&sdktestdata.MsgCreateDog{})},
		},
		"comment too long": {
			srcAccount:   accountAddr,
			srcComment:   strings.Repeat("a", 256),
			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
			expErr:       true,
		},
		"group account required": {
			srcComment:   "test",
			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
			expErr:       true,
		},
		"existing group account required": {
			srcAccount:   []byte("non-existing-account"),
			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
			expErr:       true,
		},
		"impossible case: decision policy threshold > total group weight": {
			srcAccount:   bigThresholdAddr,
			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
			expErr:       true,
		},
		"only group members can create a proposal": {
			srcAccount:   accountAddr,
			srcProposers: []sdk.AccAddress{[]byte("non--member-address")},
			expErr:       true,
		},
		"all proposers must be in group": {
			srcAccount:   accountAddr,
			srcProposers: []sdk.AccAddress{[]byte("valid-member-address"), []byte("non--member-address")},
			expErr:       true,
		},
		"proposers must not be nil": {
			srcAccount:   accountAddr,
			srcProposers: []sdk.AccAddress{[]byte("valid-member-address"), nil},
			expErr:       true,
		},
		"admin that is not a group member can not create proposal": {
			srcAccount:   accountAddr,
			srcComment:   "test",
			srcProposers: []sdk.AccAddress{[]byte("valid--admin-address")},
			expErr:       true,
		},
		"reject msgs that are not authz by group account": {
			srcAccount:   accountAddr,
			srcComment:   "test",
			srcMsgs:      []sdk.Msg{&testdata.MsgAuthenticated{Signers: []sdk.AccAddress{[]byte("not-group-acct-addrs")}}},
			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
			expErr:       true,
		},
	}
	for msg, spec := range specs {
		s.Run(msg, func() {
			id, err := s.groupKeeper.CreateProposal(s.sdkCtx, spec.srcAccount, spec.srcComment, spec.srcProposers, spec.srcMsgs)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			// then all data persisted
			proposal, err := s.groupKeeper.GetProposal(s.sdkCtx, id)
			s.Require().NoError(err)

			s.Assert().Equal(accountAddr, proposal.GroupAccount)
			s.Assert().Equal(spec.srcComment, proposal.Comment)
			s.Assert().Equal(spec.srcProposers, proposal.Proposers)

			submittedAt, err := gogotypes.TimestampFromProto(&proposal.SubmittedAt)
			s.Require().NoError(err)
			s.Assert().Equal(s.blockTime, submittedAt)

			s.Assert().Equal(uint64(1), proposal.GroupVersion)
			s.Assert().Equal(uint64(1), proposal.GroupAccountVersion)
			s.Assert().Equal(types.ProposalStatusSubmitted, proposal.Status)
			s.Assert().Equal(types.ProposalResultUndefined, proposal.Result)
			s.Assert().Equal(types.Tally{
				YesCount:     sdk.ZeroDec(),
				NoCount:      sdk.ZeroDec(),
				AbstainCount: sdk.ZeroDec(),
				VetoCount:    sdk.ZeroDec(),
			}, proposal.VoteState)

			timeout, err := gogotypes.TimestampFromProto(&proposal.Timeout)
			s.Require().NoError(err)
			s.Assert().Equal(s.blockTime.Add(time.Second), timeout)

			if spec.srcMsgs == nil { // then empty list is ok
				s.Assert().Len(proposal.GetMsgs(), 0)
			} else {
				s.Assert().Equal(spec.srcMsgs, proposal.GetMsgs())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestVote() {
	members := []types.Member{
		{Address: []byte("valid-member-address"), Power: sdk.OneDec()},
		{Address: []byte("power-member-address"), Power: sdk.NewDec(2)},
	}
	myGroupID, err := s.groupKeeper.CreateGroup(s.sdkCtx, []byte("valid--admin-address"), members, "test")
	s.Require().NoError(err)

	policy := types.NewThresholdDecisionPolicy(
		sdk.NewDec(2),
		gogotypes.Duration{Seconds: 1},
	)
	accountAddr, err := s.groupKeeper.CreateGroupAccount(s.sdkCtx, []byte("valid--admin-address"), myGroupID, policy, "test")
	s.Require().NoError(err)
	myProposalID, err := s.groupKeeper.CreateProposal(s.sdkCtx, accountAddr, "integration test", []sdk.AccAddress{[]byte("valid-member-address")}, nil)
	s.Require().NoError(err)

	specs := map[string]struct {
		srcProposalID     types.ProposalID
		srcVoters         []sdk.AccAddress
		srcChoice         types.Choice
		srcComment        string
		srcCtx            sdk.Context
		doBefore          func(ctx sdk.Context)
		expErr            bool
		expVoteState      types.Tally
		expProposalStatus types.Proposal_Status
		expResult         types.Proposal_Result
	}{
		"vote yes": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_YES,
			expVoteState: types.Tally{
				YesCount:     sdk.OneDec(),
				NoCount:      sdk.ZeroDec(),
				AbstainCount: sdk.ZeroDec(),
				VetoCount:    sdk.ZeroDec(),
			},
			expProposalStatus: types.ProposalStatusSubmitted,
			expResult:         types.ProposalResultUndefined,
		},
		"vote no": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			expVoteState: types.Tally{
				YesCount:     sdk.ZeroDec(),
				NoCount:      sdk.OneDec(),
				AbstainCount: sdk.ZeroDec(),
				VetoCount:    sdk.ZeroDec(),
			},
			expProposalStatus: types.ProposalStatusSubmitted,
			expResult:         types.ProposalResultUndefined,
		},
		"vote abstain": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_ABSTAIN,
			expVoteState: types.Tally{
				YesCount:     sdk.ZeroDec(),
				NoCount:      sdk.ZeroDec(),
				AbstainCount: sdk.OneDec(),
				VetoCount:    sdk.ZeroDec(),
			},
			expProposalStatus: types.ProposalStatusSubmitted,
			expResult:         types.ProposalResultUndefined,
		},
		"vote veto": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_VETO,
			expVoteState: types.Tally{
				YesCount:     sdk.ZeroDec(),
				NoCount:      sdk.ZeroDec(),
				AbstainCount: sdk.ZeroDec(),
				VetoCount:    sdk.OneDec(),
			},
			expProposalStatus: types.ProposalStatusSubmitted,
			expResult:         types.ProposalResultUndefined,
		},
		"apply decision policy early": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("power-member-address")},
			srcChoice:     types.Choice_YES,
			expVoteState: types.Tally{
				YesCount:     sdk.NewDec(2),
				NoCount:      sdk.ZeroDec(),
				AbstainCount: sdk.ZeroDec(),
				VetoCount:    sdk.ZeroDec(),
			},
			expProposalStatus: types.ProposalStatusClosed,
			expResult:         types.ProposalResultAccepted,
		},
		"reject new votes when final decision is made already": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_YES,
			doBefore: func(ctx sdk.Context) {
				s.Require().NoError(s.groupKeeper.Vote(s.sdkCtx, myProposalID, []sdk.AccAddress{[]byte("power-member-address")}, types.Choice_VETO, ""))
			},
			expErr: true,
		},
		"comment too long": {
			srcProposalID: myProposalID,
			srcComment:    strings.Repeat("a", 256),
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			expErr:        true,
		},
		"existing proposal required": {
			srcProposalID: 9999,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			expErr:        true,
		},
		"empty choice": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			expErr:        true,
		},
		"invalid choice": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     5,
			expErr:        true,
		},
		"all voters must be in group": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address"), []byte("non--member-address")},
			srcChoice:     types.Choice_NO,
			expErr:        true,
		},
		"voters must not include nil": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address"), nil},
			srcChoice:     types.Choice_NO,
			expErr:        true,
		},
		"voters must not be nil": {
			srcProposalID: myProposalID,
			srcChoice:     types.Choice_NO,
			expErr:        true,
		},
		"voters must not be empty": {
			srcProposalID: myProposalID,
			srcChoice:     types.Choice_NO,
			srcVoters:     []sdk.AccAddress{},
			expErr:        true,
		},
		"admin that is not a group member can not vote": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid--admin-address")},
			srcChoice:     types.Choice_NO,
			expErr:        true,
		},
		"on timeout": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			srcCtx:        s.sdkCtx.WithBlockTime(s.blockTime.Add(time.Second)),
			expErr:        true,
		},
		"closed already": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			doBefore: func(ctx sdk.Context) {
				err := s.groupKeeper.Vote(s.sdkCtx, myProposalID, []sdk.AccAddress{[]byte("power-member-address")}, types.Choice_YES, "")
				s.Require().NoError(err)
			},
			expErr: true,
		},
		"voted already": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			doBefore: func(ctx sdk.Context) {
				err := s.groupKeeper.Vote(s.sdkCtx, myProposalID, []sdk.AccAddress{[]byte("valid-member-address")}, types.Choice_YES, "")
				s.Require().NoError(err)
			},
			expErr: true,
		},
		"with group modified": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			doBefore: func(ctx sdk.Context) {
				g, err := s.groupKeeper.GetGroup(s.sdkCtx, myGroupID)
				s.Require().NoError(err)
				g.Comment = "modified"
				s.Require().NoError(s.groupKeeper.UpdateGroup(s.sdkCtx, &g))
			},
			expErr: true,
		},
		"with policy modified": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			doBefore: func(ctx sdk.Context) {
				a, err := s.groupKeeper.GetGroupAccount(s.sdkCtx, accountAddr)
				s.Require().NoError(err)
				a.Comment = "modified"
				s.Require().NoError(s.groupKeeper.UpdateGroupAccount(s.sdkCtx, &a))
			},
			expErr: true,
		},
	}
	for msg, spec := range specs {
		s.Run(msg, func() {
			ctx := s.sdkCtx
			if !spec.srcCtx.IsZero() {
				ctx = spec.srcCtx
			}
			ctx, _ = ctx.CacheContext()

			if spec.doBefore != nil {
				spec.doBefore(s.sdkCtx)
			}
			err := s.groupKeeper.Vote(s.sdkCtx, spec.srcProposalID, spec.srcVoters, spec.srcChoice, spec.srcComment)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			// and all votes are stored
			for _, voter := range spec.srcVoters {
				// then all data persisted
				loaded, err := s.groupKeeper.GetVote(s.sdkCtx, spec.srcProposalID, voter)
				s.Require().NoError(err)
				s.Assert().Equal(spec.srcProposalID, loaded.Proposal)
				s.Assert().Equal(voter, loaded.Voter)
				s.Assert().Equal(spec.srcChoice, loaded.Choice)
				s.Assert().Equal(spec.srcComment, loaded.Comment)
				submittedAt, err := gogotypes.TimestampFromProto(&loaded.SubmittedAt)
				s.Require().NoError(err)
				s.Assert().Equal(s.blockTime, submittedAt)
			}

			// and proposal is updated
			proposal, err := s.groupKeeper.GetProposal(s.sdkCtx, spec.srcProposalID)
			s.Require().NoError(err)
			s.Assert().Equal(spec.expVoteState, proposal.VoteState)
			s.Assert().Equal(spec.expResult, proposal.Result)
			s.Assert().Equal(spec.expProposalStatus, proposal.Status)
		})
	}
}

func (s *IntegrationTestSuite) TestExecProposal() {
	members := []types.Member{
		{Address: s.addr2, Power: sdk.OneDec()},
	}
	myGroupID, err := s.groupKeeper.CreateGroup(s.sdkCtx, s.addr1, members, "test")
	s.Require().NoError(err)

	policy := types.NewThresholdDecisionPolicy(
		sdk.OneDec(),
		gogotypes.Duration{Seconds: 1},
	)
	accountAddr, err := s.groupKeeper.CreateGroupAccount(s.sdkCtx, s.addr1, myGroupID, policy, "test")
	s.Require().NoError(err)

	msgSend := &banktypes.MsgSend{
		FromAddress: accountAddr.String(),
		ToAddress:   s.addr2.String(),
		Amount:      sdk.Coins{sdk.NewInt64Coin("token", 100)},
	}
	proposers := []sdk.AccAddress{s.addr2}

	specs := map[string]struct {
		srcBlockTime      time.Time
		setupProposal     func(ctx sdk.Context) types.ProposalID
		expErr            bool
		expProposalStatus types.Proposal_Status
		expProposalResult types.Proposal_Result
		expExecutorResult types.Proposal_ExecutorResult
		expPayloadCounter uint64
	}{
		"proposal executed when accepted": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_YES, ""))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultAccepted,
			expExecutorResult: types.ProposalExecutorResultSuccess,
			expPayloadCounter: 1,
		},
		"proposal with multiple messages executed when accepted": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend, msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_YES, ""))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultAccepted,
			expExecutorResult: types.ProposalExecutorResultSuccess,
			expPayloadCounter: 2,
		},
		"proposal not executed when rejected": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_NO, ""))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultRejected,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"open proposal must not fail": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusSubmitted,
			expProposalResult: types.ProposalResultUndefined,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"existing proposal required": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				return 9999
			},
			expErr: true,
		},
		"Decision policy also applied on timeout": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_NO, ""))
				return myProposalID
			},
			srcBlockTime:      s.blockTime.Add(time.Second),
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultRejected,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"Decision policy also applied after timeout": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_NO, ""))
				return myProposalID
			},
			srcBlockTime:      s.blockTime.Add(time.Second).Add(time.Millisecond),
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultRejected,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"with group modified before tally": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				// then modify group
				g, err := s.groupKeeper.GetGroup(ctx, myGroupID)
				s.Require().NoError(err)
				g.Comment = "modified"
				s.Require().NoError(s.groupKeeper.UpdateGroup(ctx, &g))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusAborted,
			expProposalResult: types.ProposalResultUndefined,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"with group account modified before tally": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				// then modify group account
				a, err := s.groupKeeper.GetGroupAccount(ctx, accountAddr)
				s.Require().NoError(err)
				a.Comment = "modified"
				s.Require().NoError(s.groupKeeper.UpdateGroupAccount(ctx, &a))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusAborted,
			expProposalResult: types.ProposalResultUndefined,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"with group modified after tally": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_YES, ""))
				// then modify group after tally on vote
				g, err := s.groupKeeper.GetGroup(ctx, myGroupID)
				s.Require().NoError(err)
				g.Comment = "modified"
				s.Require().NoError(s.groupKeeper.UpdateGroup(ctx, &g))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultAccepted,
			expExecutorResult: types.ProposalExecutorResultFailure,
		},
		"with group account modified after tally": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				// then modify group account
				a, err := s.groupKeeper.GetGroupAccount(ctx, accountAddr)
				s.Require().NoError(err)
				a.Comment = "modified"
				s.Require().NoError(s.groupKeeper.UpdateGroupAccount(ctx, &a))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusAborted,
			expProposalResult: types.ProposalResultUndefined,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"prevent double execution when successful": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_YES, ""))
				s.Require().NoError(s.groupKeeper.ExecProposal(ctx, myProposalID))
				return myProposalID
			},
			expPayloadCounter: 1,
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultAccepted,
			expExecutorResult: types.ProposalExecutorResultSuccess,
		},
		"rollback all msg updates on failure": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", proposers, []sdk.Msg{
					msgSend, &banktypes.MsgSend{},
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_YES, ""))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultAccepted,
			expExecutorResult: types.ProposalExecutorResultFailure,
		},
		// "executable when failed before": {
		// 	setupProposal: func(ctx sdk.Context) types.ProposalID {
		// 		member := []sdk.AccAddress{[]byte("valid-member-address")}
		// 		myProposalID, err := s.groupKeeper.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
		// 			&testdatatypes.MsgConditional{ExpectedCounter: 1}, &testdatatypes.MsgIncCounter{},
		// 		})
		// 		s.Require().NoError(err)
		// 		s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, member, types.Choice_YES, ""))
		// 		s.Require().NoError(s.groupKeeper.ExecProposal(ctx, myProposalID))
		// 		testdataKeeper.IncCounter(ctx)
		// 		return myProposalID
		// 	},
		// 	expPayloadCounter: 2,
		// 	expProposalStatus: types.ProposalStatusClosed,
		// 	expProposalResult: types.ProposalResultAccepted,
		// 	expExecutorResult: types.ProposalExecutorResultSuccess,
		// },
	}
	for msg, spec := range specs {
		s.Run(msg, func() {
			ctx, _ := s.sdkCtx.CacheContext()
			proposalID := spec.setupProposal(ctx)

			if !spec.srcBlockTime.IsZero() {
				ctx = ctx.WithBlockTime(spec.srcBlockTime)
			}
			err := s.groupKeeper.ExecProposal(ctx, proposalID)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			// and proposal is updated
			proposal, err := s.groupKeeper.GetProposal(ctx, proposalID)
			s.Require().NoError(err)
			exp := types.Proposal_Result_name[int32(spec.expProposalResult)]
			got := types.Proposal_Result_name[int32(proposal.Result)]
			s.Assert().Equal(exp, got)

			exp = types.Proposal_Status_name[int32(spec.expProposalStatus)]
			got = types.Proposal_Status_name[int32(proposal.Status)]
			s.Assert().Equal(exp, got)

			exp = types.Proposal_ExecutorResult_name[int32(spec.expExecutorResult)]
			got = types.Proposal_ExecutorResult_name[int32(proposal.ExecutorResult)]
			s.Assert().Equal(exp, got)

			// TODO verify proposal messages executed
			// s.Assert().Equal(spec.expPayloadCounter, testdataKeeper.GetCounter(ctx), "counter")
		})
	}
}
