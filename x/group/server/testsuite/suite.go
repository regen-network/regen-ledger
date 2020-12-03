package testsuite

import (
	"context"
	"strings"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/testutil/server"
	"github.com/regen-network/regen-ledger/testutil/testdata"
	groupserver "github.com/regen-network/regen-ledger/x/group/server"
	"github.com/regen-network/regen-ledger/x/group/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory server.FixtureFactory
	fixture        server.Fixture

	ctx              context.Context
	sdkCtx           sdk.Context
	msgClient        types.MsgClient
	addr1            sdk.AccAddress
	addr2            sdk.AccAddress
	groupAccountAddr sdk.AccAddress
	groupID          types.GroupID

	groupKeeper groupserver.Keeper
	bankKeeper  bankkeeper.Keeper
	router      sdk.Router

	blockTime time.Time
}

func NewIntegrationTestSuite(
	fixtureFactory server.FixtureFactory, groupKeeper groupserver.Keeper,
	bankKeeper bankkeeper.Keeper, router sdk.Router) *IntegrationTestSuite {
	return &IntegrationTestSuite{
		fixtureFactory: fixtureFactory,
		groupKeeper:    groupKeeper,
		bankKeeper:     bankKeeper,
		router:         router,
	}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.fixture = s.fixtureFactory.Setup()
	s.ctx = s.fixture.Context()

	s.blockTime = time.Now().UTC()

	sdkCtx := sdk.UnwrapSDKContext(s.ctx).WithBlockTime(s.blockTime)
	s.sdkCtx = sdkCtx
	s.ctx = sdk.WrapSDKContext(sdkCtx)

	s.groupKeeper.SetParams(sdkCtx, types.DefaultParams())

	totalSupply := banktypes.NewSupply(sdk.NewCoins(sdk.NewInt64Coin("test", 400000000)))
	s.bankKeeper.SetSupply(sdkCtx, totalSupply)
	s.bankKeeper.SetParams(sdkCtx, banktypes.DefaultParams())

	s.msgClient = types.NewMsgClient(s.fixture.TxConn())
	if len(s.fixture.Signers()) < 2 {
		s.FailNow("expected at least 2 signers, got %d", s.fixture.Signers())
	}
	s.addr1 = s.fixture.Signers()[0]
	s.addr2 = s.fixture.Signers()[1]

	members := []types.Member{
		{Address: s.addr2, Power: "1"},
	}
	groupID, err := s.groupKeeper.CreateGroup(s.sdkCtx, s.addr1, members, "test")
	s.Require().NoError(err)
	s.groupID = groupID

	policy := types.NewThresholdDecisionPolicy(
		"1",
		gogotypes.Duration{Seconds: 1},
	)
	accountAddr, err := s.groupKeeper.CreateGroupAccount(s.sdkCtx, s.addr1, groupID, policy, "test")
	s.Require().NoError(err)
	s.groupAccountAddr = accountAddr

	s.Require().NoError(s.bankKeeper.SetBalances(s.sdkCtx, accountAddr, sdk.Coins{sdk.NewInt64Coin("test", 10000)}))
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.fixture.Teardown()
}

func (s *IntegrationTestSuite) TestCreateGroup() {
	members := []types.Member{{
		Address: sdk.AccAddress([]byte("one--member--address")),
		Power:   "1",
		Comment: "first",
	}, {
		Address: sdk.AccAddress([]byte("other-member-address")),
		Power:   "2",
		Comment: "second",
	}}

	specs := map[string]struct {
		req    *types.MsgCreateGroupRequest
		expErr bool
	}{
		"all good": {
			req: &types.MsgCreateGroupRequest{
				Admin:   []byte("valid--admin-address"),
				Members: members,
				Comment: "test",
			},
		},
		"group comment too long": {
			req: &types.MsgCreateGroupRequest{
				Admin:   []byte("valid--admin-address"),
				Members: members,
				Comment: strings.Repeat("a", 256),
			},
			expErr: true,
		},
		"member comment too long": {
			req: &types.MsgCreateGroupRequest{
				Admin: []byte("valid--admin-address"),
				Members: []types.Member{{
					Address: []byte("valid-member-address"),
					Power:   "1",
					Comment: strings.Repeat("a", 256),
				}},
				Comment: "test",
			},
			expErr: true,
		},
	}
	var seq uint32 = 1
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			res, err := s.msgClient.CreateGroup(s.ctx, spec.req)
			if spec.expErr {
				s.Require().Error(err)
				s.Require().False(s.groupKeeper.HasGroup(s.sdkCtx, types.GroupID(seq+1).Bytes()))
				return
			}
			s.Require().NoError(err)
			id := res.GroupId

			seq++
			s.Assert().Equal(types.GroupID(seq), id)

			// then all data persisted
			loadedGroup, err := s.groupKeeper.GetGroup(s.sdkCtx, id)
			s.Require().NoError(err)
			s.Assert().Equal(sdk.AccAddress([]byte(spec.req.Admin)), loadedGroup.Admin)
			s.Assert().Equal(spec.req.Comment, loadedGroup.Comment)
			s.Assert().Equal(id, loadedGroup.GroupId)
			s.Assert().Equal(uint64(1), loadedGroup.Version)

			// and members are stored as well
			it, err := s.groupKeeper.GetGroupMembers(s.sdkCtx, id)
			s.Require().NoError(err)
			var loadedMembers []types.GroupMember
			_, err = orm.ReadAll(it, &loadedMembers)
			s.Require().NoError(err)
			s.Assert().Equal(len(members), len(loadedMembers))
			for i := range loadedMembers {
				s.Assert().Equal(members[i].Comment, loadedMembers[i].Comment)
				s.Assert().Equal(members[i].Address, loadedMembers[i].Member)
				s.Assert().Equal(members[i].Power, loadedMembers[i].Weight)
				s.Assert().Equal(id, loadedMembers[i].GroupId)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateGroupAdmin() {
	members := []types.Member{{
		Address: sdk.AccAddress([]byte("valid-member-address")),
		Power:   "1",
		Comment: "first member",
	}}
	oldAdmin := []byte("my-old-admin-address")
	groupID, err := s.groupKeeper.CreateGroup(s.sdkCtx, oldAdmin, members, "test")
	s.Require().NoError(err)

	specs := map[string]struct {
		req       *types.MsgUpdateGroupAdminRequest
		expStored types.GroupMetadata
		expErr    bool
	}{
		"with correct admin": {
			req: &types.MsgUpdateGroupAdminRequest{
				GroupId:  groupID,
				Admin:    oldAdmin,
				NewAdmin: []byte("my-new-admin-address"),
			},
			expStored: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       []byte("my-new-admin-address"),
				Comment:     "test",
				TotalWeight: "1",
				Version:     2,
			},
		},
		"with wrong admin": {
			req: &types.MsgUpdateGroupAdminRequest{
				GroupId:  groupID,
				Admin:    []byte("unknown-address"),
				NewAdmin: []byte("my-new-admin-address"),
			},
			expErr: true,
			expStored: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
		},
		"with unknown groupID": {
			req: &types.MsgUpdateGroupAdminRequest{
				GroupId:  999,
				Admin:    oldAdmin,
				NewAdmin: []byte("my-new-admin-address"),
			},
			expErr: true,
			expStored: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			_, err := s.msgClient.UpdateGroupAdmin(s.ctx, spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			// then
			loaded, err := s.groupKeeper.GetGroup(s.sdkCtx, groupID)
			s.Require().NoError(err)
			s.Assert().Equal(spec.expStored, loaded)
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateGroupComment() {
	oldComment := "first"
	members := []types.Member{{
		Address: sdk.AccAddress([]byte("valid-member-address")),
		Power:   "1",
		Comment: oldComment,
	}}

	oldAdmin := []byte("my-old-admin-address")
	groupID, err := s.groupKeeper.CreateGroup(s.sdkCtx, oldAdmin, members, "test")
	s.Require().NoError(err)

	specs := map[string]struct {
		req       *types.MsgUpdateGroupCommentRequest
		expErr    bool
		expStored types.GroupMetadata
	}{
		"with correct admin": {
			req: &types.MsgUpdateGroupCommentRequest{
				GroupId: groupID,
				Admin:   oldAdmin,
				Comment: "new comment",
			},
			expStored: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Comment:     "new comment",
				TotalWeight: "1",
				Version:     2,
			},
		},
		"with wrong admin": {
			req: &types.MsgUpdateGroupCommentRequest{
				GroupId: groupID,
				Admin:   []byte("unknown-address"),
				Comment: "new comment",
			},
			expErr: true,
			expStored: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
		},
		"with unknown groupid": {
			req: &types.MsgUpdateGroupCommentRequest{
				GroupId: 999,
				Admin:   []byte("unknown-address"),
				Comment: "new comment",
			},
			expErr: true,
			expStored: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			_, err := s.msgClient.UpdateGroupComment(s.ctx, spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			// then
			loaded, err := s.groupKeeper.GetGroup(s.sdkCtx, groupID)
			s.Require().NoError(err)
			s.Assert().Equal(spec.expStored, loaded)
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateGroupMembers() {
	members := []types.Member{{
		Address: sdk.AccAddress([]byte("valid-member-address")),
		Power:   "1",
		Comment: "first",
	}}

	myAdmin := []byte("valid--admin-address")
	groupID, err := s.groupKeeper.CreateGroup(s.sdkCtx, myAdmin, members, "test")
	s.Require().NoError(err)

	specs := map[string]struct {
		req        *types.MsgUpdateGroupMembersRequest
		expErr     bool
		expGroup   types.GroupMetadata
		expMembers []types.GroupMember
	}{
		"add new member": {
			req: &types.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []types.Member{{
					Address: sdk.AccAddress([]byte("other-member-address")),
					Power:   "2",
					Comment: "second",
				}},
			},
			expGroup: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "3",
				Version:     2,
			},
			expMembers: []types.GroupMember{
				{
					Member:  sdk.AccAddress([]byte("other-member-address")),
					GroupId: groupID,
					Weight:  "2",
					Comment: "second",
				},
				{
					Member:  sdk.AccAddress([]byte("valid-member-address")),
					GroupId: groupID,
					Weight:  "1",
					Comment: "first",
				},
			},
		},
		"update member": {
			req: &types.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []types.Member{{
					Address: sdk.AccAddress([]byte("valid-member-address")),
					Power:   "2",
					Comment: "updated",
				}},
			},
			expGroup: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "2",
				Version:     2,
			},
			expMembers: []types.GroupMember{
				{
					Member:  sdk.AccAddress([]byte("valid-member-address")),
					GroupId: groupID,
					Weight:  "2",
					Comment: "updated",
				},
			},
		},
		"update member with same data": {
			req: &types.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []types.Member{{
					Address: sdk.AccAddress([]byte("valid-member-address")),
					Power:   "1",
					Comment: "first",
				}},
			},
			expGroup: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     2,
			},
			expMembers: []types.GroupMember{
				{
					Member:  sdk.AccAddress([]byte("valid-member-address")),
					GroupId: groupID,
					Weight:  "1",
					Comment: "first",
				},
			},
		},
		"replace member": {
			req: &types.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []types.Member{
					{
						Address: sdk.AccAddress([]byte("valid-member-address")),
						Power:   "0",
						Comment: "good bye",
					},
					{
						Address: s.addr1,
						Power:   "1",
						Comment: "welcome",
					},
				},
			},
			expGroup: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     2,
			},
			expMembers: []types.GroupMember{{
				Member:  s.addr1,
				GroupId: groupID,
				Weight:  "1",
				Comment: "welcome",
			}},
		},
		"remove existing member": {
			req: &types.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []types.Member{{
					Address: sdk.AccAddress([]byte("valid-member-address")),
					Power:   "0",
					Comment: "good bye",
				}},
			},
			expGroup: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "0",
				Version:     2,
			},
			expMembers: []types.GroupMember{},
		},
		"remove unknown member": {
			req: &types.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []types.Member{{
					Address: sdk.AccAddress([]byte("unknown-member-addrs")),
					Power:   "0",
					Comment: "good bye",
				}},
			},
			expErr: true,
			expGroup: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
			expMembers: []types.GroupMember{{
				Member:  sdk.AccAddress([]byte("valid-member-address")),
				GroupId: groupID,
				Weight:  "1",
				Comment: "first",
			}},
		},
		"with wrong admin": {
			req: &types.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   []byte("unknown-address"),
				MemberUpdates: []types.Member{{
					Address: sdk.AccAddress([]byte("other-member-address")),
					Power:   "2",
					Comment: "second",
				}},
			},
			expErr: true,
			expGroup: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
			expMembers: []types.GroupMember{{
				Member:  sdk.AccAddress([]byte("valid-member-address")),
				GroupId: groupID,
				Weight:  "1",
				Comment: "first",
			}},
		},
		"with unknown groupID": {
			req: &types.MsgUpdateGroupMembersRequest{
				GroupId: 999,
				Admin:   myAdmin,
				MemberUpdates: []types.Member{{
					Address: sdk.AccAddress([]byte("other-member-address")),
					Power:   "2",
					Comment: "second",
				}},
			},
			expErr: true,
			expGroup: types.GroupMetadata{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
			expMembers: []types.GroupMember{{
				Member:  sdk.AccAddress([]byte("valid-member-address")),
				GroupId: groupID,
				Weight:  "1",
				Comment: "first",
			}},
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			ctx, _ := s.sdkCtx.CacheContext()
			_, err := s.msgClient.UpdateGroupMembers(sdk.WrapSDKContext(ctx), spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			// then
			loaded, err := s.groupKeeper.GetGroup(ctx, groupID)
			s.Require().NoError(err)
			s.Assert().Equal(spec.expGroup, loaded)

			// and members persisted
			it, err := s.groupKeeper.GetGroupMembers(ctx, groupID)
			s.Require().NoError(err)
			var loadedMembers []types.GroupMember
			_, err = orm.ReadAll(it, &loadedMembers)
			s.Require().NoError(err)
			s.Assert().Equal(spec.expMembers, loadedMembers)
		})
	}
}

func (s *IntegrationTestSuite) TestCreateGroupAccount() {
	myGroupID, err := s.groupKeeper.CreateGroup(s.sdkCtx, []byte("valid--admin-address"), nil, "test")
	s.Require().NoError(err)

	specs := map[string]struct {
		req    *types.MsgCreateGroupAccountRequest
		policy types.DecisionPolicy
		expErr bool
	}{
		"all good": {
			req: &types.MsgCreateGroupAccountRequest{
				Admin:   []byte("valid--admin-address"),
				Comment: "test",
				GroupId: myGroupID,
			},
			policy: types.NewThresholdDecisionPolicy(
				"1",
				gogotypes.Duration{Seconds: 1},
			),
		},
		"decision policy threshold > total group weight": {
			req: &types.MsgCreateGroupAccountRequest{
				Admin:   []byte("valid--admin-address"),
				Comment: "test",
				GroupId: myGroupID,
			},
			policy: types.NewThresholdDecisionPolicy(
				"10",
				gogotypes.Duration{Seconds: 1},
			),
		},
		"group id does not exists": {
			req: &types.MsgCreateGroupAccountRequest{
				Admin:   []byte("valid--admin-address"),
				Comment: "test",
				GroupId: 9999,
			},
			policy: types.NewThresholdDecisionPolicy(
				"1",
				gogotypes.Duration{Seconds: 1},
			),
			expErr: true,
		},
		"admin not group admin": {
			req: &types.MsgCreateGroupAccountRequest{
				Admin:   []byte("other--admin-address"),
				Comment: "test",
				GroupId: myGroupID,
			},
			policy: types.NewThresholdDecisionPolicy(
				"1",
				gogotypes.Duration{Seconds: 1},
			),
			expErr: true,
		},
		"comment too long": {
			req: &types.MsgCreateGroupAccountRequest{
				Admin:   []byte("valid--admin-address"),
				Comment: strings.Repeat("a", 256),
				GroupId: myGroupID,
			},
			policy: types.NewThresholdDecisionPolicy(
				"1",
				gogotypes.Duration{Seconds: 1},
			),
			expErr: true,
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			err := spec.req.SetDecisionPolicy(spec.policy)
			s.Require().NoError(err)

			res, err := s.msgClient.CreateGroupAccount(s.ctx, spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			addr := res.GroupAccount

			// then all data persisted
			groupAccount, err := s.groupKeeper.GetGroupAccount(s.sdkCtx, addr)
			s.Require().NoError(err)
			s.Assert().Equal(addr, groupAccount.GroupAccount)
			s.Assert().Equal(myGroupID, groupAccount.GroupId)
			s.Assert().Equal(sdk.AccAddress([]byte(spec.req.Admin)), groupAccount.Admin)
			s.Assert().Equal(spec.req.Comment, groupAccount.Comment)
			s.Assert().Equal(uint64(1), groupAccount.Version)
			s.Assert().Equal(spec.policy.(*types.ThresholdDecisionPolicy), groupAccount.GetDecisionPolicy())
		})
	}
}

func (s *IntegrationTestSuite) TestCreateProposal() {
	members := []types.Member{{
		Address: []byte("valid-member-address"),
		Power:   "1",
	}}
	myGroupID, err := s.groupKeeper.CreateGroup(s.sdkCtx, []byte("valid--admin-address"), members, "test")
	s.Require().NoError(err)

	policy := types.NewThresholdDecisionPolicy(
		"1",
		gogotypes.Duration{Seconds: 1},
	)
	accountAddr, err := s.groupKeeper.CreateGroupAccount(s.sdkCtx, []byte("valid--admin-address"), myGroupID, policy, "test")
	s.Require().NoError(err)

	policy = types.NewThresholdDecisionPolicy(
		"100",
		gogotypes.Duration{Seconds: 1},
	)
	bigThresholdAddr, err := s.groupKeeper.CreateGroupAccount(s.sdkCtx, []byte("valid--admin-address"), myGroupID, policy, "test")
	s.Require().NoError(err)

	specs := map[string]struct {
		req    *types.MsgCreateProposalRequest
		msgs   []sdk.Msg
		expErr bool
	}{
		"all good with minimal fields set": {
			req: &types.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
			},
		},
		"all good with good msg payload": {
			req: &types.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
			},
			msgs: []sdk.Msg{&banktypes.MsgSend{
				FromAddress: accountAddr.String(),
				ToAddress:   s.addr2.String(),
				Amount:      sdk.Coins{sdk.NewInt64Coin("token", 100)},
			}},
		},
		"comment too long": {
			req: &types.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Comment:      strings.Repeat("a", 256),
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"group account required": {
			req: &types.MsgCreateProposalRequest{
				Comment:   "test",
				Proposers: []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"existing group account required": {
			req: &types.MsgCreateProposalRequest{
				GroupAccount: []byte("non-existing-account"),
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"impossible case: decision policy threshold > total group weight": {
			req: &types.MsgCreateProposalRequest{
				GroupAccount: bigThresholdAddr,
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"only group members can create a proposal": {
			req: &types.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Proposers:    []sdk.AccAddress{[]byte("non--member-address")},
			},
			expErr: true,
		},
		"all proposers must be in group": {
			req: &types.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address"), []byte("non--member-address")},
			},
			expErr: true,
		},
		"proposers must not be nil": {
			req: &types.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address"), nil},
			},
			expErr: true,
		},
		"admin that is not a group member can not create proposal": {
			req: &types.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Comment:      "test",
				Proposers:    []sdk.AccAddress{[]byte("valid--admin-address")},
			},
			expErr: true,
		},
		"reject msgs that are not authz by group account": {
			req: &types.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Comment:      "test",
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
			},
			msgs:   []sdk.Msg{&testdata.MsgAuthenticated{Signers: []sdk.AccAddress{[]byte("not-group-acct-addrs")}}},
			expErr: true,
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			err := spec.req.SetMsgs(spec.msgs)
			s.Require().NoError(err)

			res, err := s.msgClient.CreateProposal(s.ctx, spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			id := res.ProposalId

			// then all data persisted
			proposal, err := s.groupKeeper.GetProposal(s.sdkCtx, id)
			s.Require().NoError(err)

			s.Assert().Equal(accountAddr, proposal.GroupAccount)
			s.Assert().Equal(spec.req.Comment, proposal.Comment)
			s.Assert().Equal(spec.req.Proposers, proposal.Proposers)

			submittedAt, err := gogotypes.TimestampFromProto(&proposal.SubmittedAt)
			s.Require().NoError(err)
			s.Assert().Equal(s.blockTime, submittedAt)

			s.Assert().Equal(uint64(1), proposal.GroupVersion)
			s.Assert().Equal(uint64(1), proposal.GroupAccountVersion)
			s.Assert().Equal(types.ProposalStatusSubmitted, proposal.Status)
			s.Assert().Equal(types.ProposalResultUnfinalized, proposal.Result)
			s.Assert().Equal(types.Tally{
				YesCount:     "0",
				NoCount:      "0",
				AbstainCount: "0",
				VetoCount:    "0",
			}, proposal.VoteState)

			timeout, err := gogotypes.TimestampFromProto(&proposal.Timeout)
			s.Require().NoError(err)
			s.Assert().Equal(s.blockTime.Add(time.Second), timeout)

			if spec.msgs == nil { // then empty list is ok
				s.Assert().Len(proposal.GetMsgs(), 0)
			} else {
				s.Assert().Equal(spec.msgs, proposal.GetMsgs())
			}
		})
	}
}

func (s *IntegrationTestSuite) TestVote() {
	members := []types.Member{
		{Address: []byte("valid-member-address"), Power: "1"},
		{Address: []byte("power-member-address"), Power: "2"},
	}
	myGroupID, err := s.groupKeeper.CreateGroup(s.sdkCtx, []byte("valid--admin-address"), members, "test")
	s.Require().NoError(err)

	policy := types.NewThresholdDecisionPolicy(
		"2",
		gogotypes.Duration{Seconds: 1},
	)
	accountAddr, err := s.groupKeeper.CreateGroupAccount(s.sdkCtx, []byte("valid--admin-address"), myGroupID, policy, "test")
	s.Require().NoError(err)
	myProposalID, err := s.groupKeeper.CreateProposal(s.sdkCtx, accountAddr, "integration test", []sdk.AccAddress{[]byte("valid-member-address")}, nil)
	s.Require().NoError(err)

	specs := map[string]struct {
		req               *types.MsgVoteRequest
		srcCtx            sdk.Context
		doBefore          func(ctx sdk.Context)
		expErr            bool
		expVoteState      types.Tally
		expProposalStatus types.Proposal_Status
		expResult         types.Proposal_Result
	}{
		"vote yes": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_YES,
			},
			expVoteState: types.Tally{
				YesCount:     "1",
				NoCount:      "0",
				AbstainCount: "0",
				VetoCount:    "0",
			},
			expProposalStatus: types.ProposalStatusSubmitted,
			expResult:         types.ProposalResultUnfinalized,
		},
		"vote no": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_NO,
			},
			expVoteState: types.Tally{
				YesCount:     "0",
				NoCount:      "1",
				AbstainCount: "0",
				VetoCount:    "0",
			},
			expProposalStatus: types.ProposalStatusSubmitted,
			expResult:         types.ProposalResultUnfinalized,
		},
		"vote abstain": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_ABSTAIN,
			},
			expVoteState: types.Tally{
				YesCount:     "0",
				NoCount:      "0",
				AbstainCount: "1",
				VetoCount:    "0",
			},
			expProposalStatus: types.ProposalStatusSubmitted,
			expResult:         types.ProposalResultUnfinalized,
		},
		"vote veto": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_VETO,
			},
			expVoteState: types.Tally{
				YesCount:     "0",
				NoCount:      "0",
				AbstainCount: "0",
				VetoCount:    "1",
			},
			expProposalStatus: types.ProposalStatusSubmitted,
			expResult:         types.ProposalResultUnfinalized,
		},
		"apply decision policy early": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("power-member-address")},
				Choice:     types.Choice_CHOICE_YES,
			},
			expVoteState: types.Tally{
				YesCount:     "2",
				NoCount:      "0",
				AbstainCount: "0",
				VetoCount:    "0",
			},
			expProposalStatus: types.ProposalStatusClosed,
			expResult:         types.ProposalResultAccepted,
		},
		"reject new votes when final decision is made already": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_YES,
			},
			doBefore: func(ctx sdk.Context) {
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, []sdk.AccAddress{[]byte("power-member-address")}, types.Choice_CHOICE_VETO, ""))
			},
			expErr: true,
		},
		"comment too long": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Comment:    strings.Repeat("a", 256),
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"existing proposal required": {
			req: &types.MsgVoteRequest{
				ProposalId: 999,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"empty choice": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"invalid choice": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     5,
			},
			expErr: true,
		},
		"all voters must be in group": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address"), []byte("non--member-address")},
				Choice:     types.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"voters must not include nil": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address"), nil},
				Choice:     types.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"voters must not be nil": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Choice:     types.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"voters must not be empty": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Choice:     types.Choice_CHOICE_NO,
				Voters:     []sdk.AccAddress{},
			},
			expErr: true,
		},
		"admin that is not a group member can not vote": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid--admin-address")},
				Choice:     types.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"on timeout": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_NO,
			},
			srcCtx: s.sdkCtx.WithBlockTime(s.blockTime.Add(time.Second)),
			expErr: true,
		},
		"closed already": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_NO,
			},
			doBefore: func(ctx sdk.Context) {
				err := s.groupKeeper.Vote(ctx, myProposalID, []sdk.AccAddress{[]byte("power-member-address")}, types.Choice_CHOICE_YES, "")
				s.Require().NoError(err)
			},
			expErr: true,
		},
		"voted already": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_NO,
			},
			doBefore: func(ctx sdk.Context) {
				err := s.groupKeeper.Vote(ctx, myProposalID, []sdk.AccAddress{[]byte("valid-member-address")}, types.Choice_CHOICE_YES, "")
				s.Require().NoError(err)
			},
			expErr: true,
		},
		"with group modified": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_NO,
			},
			doBefore: func(ctx sdk.Context) {
				g, err := s.groupKeeper.GetGroup(ctx, myGroupID)
				s.Require().NoError(err)
				g.Comment = "group modified"
				s.Require().NoError(s.groupKeeper.UpdateGroup(ctx, &g))
			},
			expErr: true,
		},
		"with policy modified": {
			req: &types.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     types.Choice_CHOICE_NO,
			},
			doBefore: func(ctx sdk.Context) {
				a, err := s.groupKeeper.GetGroupAccount(ctx, accountAddr)
				s.Require().NoError(err)
				a.Comment = "policy modified"
				s.Require().NoError(s.groupKeeper.UpdateGroupAccount(ctx, &a))
			},
			expErr: true,
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			ctx := s.sdkCtx
			if !spec.srcCtx.IsZero() {
				ctx = spec.srcCtx
			}
			ctx, _ = ctx.CacheContext()

			if spec.doBefore != nil {
				spec.doBefore(ctx)
			}
			_, err := s.msgClient.Vote(sdk.WrapSDKContext(ctx), spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			// and all votes are stored
			for _, voter := range spec.req.Voters {
				// then all data persisted
				loaded, err := s.groupKeeper.GetVote(ctx, spec.req.ProposalId, voter)
				s.Require().NoError(err)
				s.Assert().Equal(spec.req.ProposalId, loaded.ProposalId)
				s.Assert().Equal(voter, loaded.Voter)
				s.Assert().Equal(spec.req.Choice, loaded.Choice)
				s.Assert().Equal(spec.req.Comment, loaded.Comment)
				submittedAt, err := gogotypes.TimestampFromProto(&loaded.SubmittedAt)
				s.Require().NoError(err)
				s.Assert().Equal(s.blockTime, submittedAt)
			}

			// and proposal is updated
			proposal, err := s.groupKeeper.GetProposal(ctx, spec.req.ProposalId)
			s.Require().NoError(err)
			s.Assert().Equal(spec.expVoteState, proposal.VoteState)
			s.Assert().Equal(spec.expResult, proposal.Result)
			s.Assert().Equal(spec.expProposalStatus, proposal.Status)
		})
	}
}

func (s *IntegrationTestSuite) TestDoExecuteMsgs() {
	msgSend := &banktypes.MsgSend{
		FromAddress: s.groupAccountAddr.String(),
		ToAddress:   s.addr2.String(),
		Amount:      sdk.Coins{sdk.NewInt64Coin("test", 100)},
	}

	unauthzMsgSend := &banktypes.MsgSend{
		FromAddress: s.addr1.String(),
		ToAddress:   s.addr2.String(),
		Amount:      sdk.Coins{sdk.NewInt64Coin("test", 100)},
	}

	specs := map[string]struct {
		srcMsgs    []sdk.Msg
		srcHandler sdk.Handler
		expErr     bool
	}{
		"all good": {
			srcMsgs: []sdk.Msg{msgSend},
		},
		"not authz by group account": {
			srcMsgs: []sdk.Msg{unauthzMsgSend},
			expErr:  true,
		},
		"mixed group account msgs": {
			srcMsgs: []sdk.Msg{
				msgSend,
				unauthzMsgSend,
			},
			expErr: true,
		},
		"no handler": {
			srcMsgs: []sdk.Msg{&testdata.MsgAuthenticated{Signers: []sdk.AccAddress{s.groupAccountAddr}}},
			expErr:  true,
		},
		"not panic on nil result": {
			srcMsgs: []sdk.Msg{&testdata.MsgAuthenticated{Signers: []sdk.AccAddress{s.groupAccountAddr}}},
			srcHandler: func(ctx sdk.Context, msg sdk.Msg) (result *sdk.Result, err error) {
				return nil, nil
			},
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			ctx, _ := s.sdkCtx.CacheContext()

			var router sdk.Router
			if spec.srcHandler != nil {
				router = baseapp.NewRouter().AddRoute(sdk.NewRoute("MsgAuthenticated", spec.srcHandler))
			} else {
				router = s.router
			}
			_, err := groupserver.DoExecuteMsgs(ctx, router, s.groupAccountAddr, spec.srcMsgs)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
		})
	}
}

func (s *IntegrationTestSuite) TestExecProposal() {
	msgSend := &banktypes.MsgSend{
		FromAddress: s.groupAccountAddr.String(),
		ToAddress:   s.addr2.String(),
		Amount:      sdk.Coins{sdk.NewInt64Coin("test", 100)},
	}
	proposers := []sdk.AccAddress{s.addr2}

	specs := map[string]struct {
		srcBlockTime      time.Time
		setupProposal     func(ctx sdk.Context) types.ProposalID
		expErr            bool
		expProposalStatus types.Proposal_Status
		expProposalResult types.Proposal_Result
		expExecutorResult types.Proposal_ExecutorResult
		expFromBalances   sdk.Coins
		expToBalances     sdk.Coins
	}{
		"proposal executed when accepted": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, s.groupAccountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_CHOICE_YES, ""))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultAccepted,
			expExecutorResult: types.ProposalExecutorResultSuccess,
			expFromBalances:   sdk.Coins{sdk.NewInt64Coin("test", 9900)},
			expToBalances:     sdk.Coins{sdk.NewInt64Coin("test", 100)},
		},
		"proposal with multiple messages executed when accepted": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, s.groupAccountAddr, "test", proposers, []sdk.Msg{
					msgSend, msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_CHOICE_YES, ""))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultAccepted,
			expExecutorResult: types.ProposalExecutorResultSuccess,
			expFromBalances:   sdk.Coins{sdk.NewInt64Coin("test", 9800)},
			expToBalances:     sdk.Coins{sdk.NewInt64Coin("test", 200)},
		},
		"proposal not executed when rejected": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, s.groupAccountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_CHOICE_NO, ""))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultRejected,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"open proposal must not fail": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, s.groupAccountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusSubmitted,
			expProposalResult: types.ProposalResultUnfinalized,
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
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, s.groupAccountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_CHOICE_NO, ""))
				return myProposalID
			},
			srcBlockTime:      s.blockTime.Add(time.Second),
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultRejected,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"Decision policy also applied after timeout": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, s.groupAccountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_CHOICE_NO, ""))
				return myProposalID
			},
			srcBlockTime:      s.blockTime.Add(time.Second).Add(time.Millisecond),
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultRejected,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"with group modified before tally": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, s.groupAccountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				// then modify group
				g, err := s.groupKeeper.GetGroup(ctx, s.groupID)
				s.Require().NoError(err)
				g.Comment = "group modified before tally"
				s.Require().NoError(s.groupKeeper.UpdateGroup(ctx, &g))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusAborted,
			expProposalResult: types.ProposalResultUnfinalized,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"with group account modified before tally": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, s.groupAccountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				// then modify group account
				a, err := s.groupKeeper.GetGroupAccount(ctx, s.groupAccountAddr)
				s.Require().NoError(err)
				a.Comment = "group account modified before tally"
				s.Require().NoError(s.groupKeeper.UpdateGroupAccount(ctx, &a))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusAborted,
			expProposalResult: types.ProposalResultUnfinalized,
			expExecutorResult: types.ProposalExecutorResultNotRun,
		},
		"prevent double execution when successful": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, s.groupAccountAddr, "test", proposers, []sdk.Msg{
					msgSend,
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_CHOICE_YES, ""))
				s.Require().NoError(s.groupKeeper.ExecProposal(ctx, myProposalID))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultAccepted,
			expExecutorResult: types.ProposalExecutorResultSuccess,
			expFromBalances:   sdk.Coins{sdk.NewInt64Coin("test", 9900)},
			expToBalances:     sdk.Coins{sdk.NewInt64Coin("test", 100)},
		},
		"rollback all msg updates on failure": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, s.groupAccountAddr, "test", proposers, []sdk.Msg{
					msgSend, &banktypes.MsgSend{
						FromAddress: s.groupAccountAddr.String(),
						ToAddress:   s.addr2.String(),
						Amount:      sdk.Coins{sdk.NewInt64Coin("test", 10001)}},
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_CHOICE_YES, ""))
				return myProposalID
			},
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultAccepted,
			expExecutorResult: types.ProposalExecutorResultFailure,
		},
		"executable when failed before": {
			setupProposal: func(ctx sdk.Context) types.ProposalID {
				myProposalID, err := s.groupKeeper.CreateProposal(ctx, s.groupAccountAddr, "test", proposers, []sdk.Msg{
					&banktypes.MsgSend{
						FromAddress: s.groupAccountAddr.String(),
						ToAddress:   s.addr2.String(),
						Amount:      sdk.Coins{sdk.NewInt64Coin("test", 10001)}},
				})
				s.Require().NoError(err)
				s.Require().NoError(s.groupKeeper.Vote(ctx, myProposalID, proposers, types.Choice_CHOICE_YES, ""))
				s.Require().NoError(s.groupKeeper.ExecProposal(ctx, myProposalID))
				s.Require().NoError(s.bankKeeper.SetBalances(ctx, s.groupAccountAddr, sdk.Coins{sdk.NewInt64Coin("test", 10002)}))

				return myProposalID
			},
			expProposalStatus: types.ProposalStatusClosed,
			expProposalResult: types.ProposalResultAccepted,
			expExecutorResult: types.ProposalExecutorResultSuccess,
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			ctx, _ := s.sdkCtx.CacheContext()
			proposalID := spec.setupProposal(ctx)

			if !spec.srcBlockTime.IsZero() {
				ctx = ctx.WithBlockTime(spec.srcBlockTime)
			}
			_, err := s.msgClient.Exec(sdk.WrapSDKContext(ctx), &types.MsgExecRequest{ProposalId: proposalID})
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

			if spec.expFromBalances != nil {
				fromBalances := s.bankKeeper.GetAllBalances(ctx, s.groupAccountAddr)
				s.Require().Equal(spec.expFromBalances, fromBalances)
			}
			if spec.expToBalances != nil {
				toBalances := s.bankKeeper.GetAllBalances(ctx, s.addr2)
				s.Require().Equal(spec.expToBalances, toBalances)
			}
		})
	}
}
