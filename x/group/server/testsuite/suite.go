package testsuite

import (
	"context"
	"fmt"
	"strings"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/regen-network/regen-ledger/testutil/server"
	"github.com/regen-network/regen-ledger/testutil/testdata"
	"github.com/regen-network/regen-ledger/x/group"
	groupserver "github.com/regen-network/regen-ledger/x/group/server"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory server.FixtureFactory
	fixture        server.Fixture

	ctx              context.Context
	sdkCtx           sdk.Context
	msgClient        group.MsgClient
	queryClient      group.QueryClient
	addr1            sdk.AccAddress
	addr2            sdk.AccAddress
	groupAccountAddr sdk.AccAddress
	groupID          group.ID

	groupSubspace paramstypes.Subspace
	bankKeeper    bankkeeper.Keeper
	router        sdk.Router

	blockTime time.Time
}

func NewIntegrationTestSuite(
	fixtureFactory server.FixtureFactory, groupSubspace paramstypes.Subspace,
	bankKeeper bankkeeper.Keeper, router sdk.Router) *IntegrationTestSuite {
	return &IntegrationTestSuite{
		fixtureFactory: fixtureFactory,
		groupSubspace:  groupSubspace,
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

	groupParams := group.DefaultParams()
	if !s.groupSubspace.HasKeyTable() {
		s.groupSubspace = s.groupSubspace.WithKeyTable(paramstypes.NewKeyTable().RegisterParamSet(&group.Params{}))
	}
	s.groupSubspace.SetParamSet(sdkCtx, &groupParams)

	totalSupply := banktypes.NewSupply(sdk.NewCoins(sdk.NewInt64Coin("test", 400000000)))
	s.bankKeeper.SetSupply(sdkCtx, totalSupply)
	s.bankKeeper.SetParams(sdkCtx, banktypes.DefaultParams())

	s.msgClient = group.NewMsgClient(s.fixture.TxConn())
	s.queryClient = group.NewQueryClient(s.fixture.QueryConn())

	if len(s.fixture.Signers()) < 2 {
		s.FailNow("expected at least 2 signers, got %d", s.fixture.Signers())
	}
	s.addr1 = s.fixture.Signers()[0]
	s.addr2 = s.fixture.Signers()[1]

	members := []group.Member{
		{Address: s.addr2, Power: "1"},
	}
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:   s.addr1,
		Members: members,
		Comment: "test",
	})
	s.Require().NoError(err)
	s.groupID = groupRes.GroupId

	policy := group.NewThresholdDecisionPolicy(
		"1",
		gogotypes.Duration{Seconds: 1},
	)
	accountReq := &group.MsgCreateGroupAccountRequest{
		Admin:   s.addr1,
		GroupId: s.groupID,
		Comment: "test",
	}
	err = accountReq.SetDecisionPolicy(policy)
	s.Require().NoError(err)
	accountRes, err := s.msgClient.CreateGroupAccount(s.ctx, accountReq)
	s.Require().NoError(err)
	s.groupAccountAddr = accountRes.GroupAccount

	s.Require().NoError(s.bankKeeper.SetBalances(s.sdkCtx, s.groupAccountAddr, sdk.Coins{sdk.NewInt64Coin("test", 10000)}))
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.fixture.Teardown()
}

func (s *IntegrationTestSuite) TestCreateGroup() {
	members := []group.Member{{
		Address: sdk.AccAddress([]byte("one--member--address")),
		Power:   "1",
		Comment: "first",
	}, {
		Address: sdk.AccAddress([]byte("other-member-address")),
		Power:   "2",
		Comment: "second",
	}}

	expGroups := []*group.GroupInfo{
		&group.GroupInfo{
			GroupId:     s.groupID,
			Version:     1,
			Admin:       s.addr1,
			TotalWeight: "1",
			Comment:     "test",
		},
		&group.GroupInfo{
			GroupId:     2,
			Version:     1,
			Admin:       s.addr1,
			TotalWeight: "3",
			Comment:     "test",
		},
	}

	specs := map[string]struct {
		req       *group.MsgCreateGroupRequest
		expErr    bool
		expGroups []*group.GroupInfo
	}{
		"all good": {
			req: &group.MsgCreateGroupRequest{
				Admin:   s.addr1,
				Members: members,
				Comment: "test",
			},
			expGroups: expGroups,
		},
		"group comment too long": {
			req: &group.MsgCreateGroupRequest{
				Admin:   s.addr1,
				Members: members,
				Comment: strings.Repeat("a", 256),
			},
			expErr: true,
		},
		"member comment too long": {
			req: &group.MsgCreateGroupRequest{
				Admin: s.addr1,
				Members: []group.Member{{
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
				_, err := s.queryClient.GroupInfo(s.ctx, &group.QueryGroupInfoRequest{GroupId: group.ID(seq + 1)})
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			id := res.GroupId

			seq++
			s.Assert().Equal(group.ID(seq), id)

			// then all data persisted
			loadedGroupRes, err := s.queryClient.GroupInfo(s.ctx, &group.QueryGroupInfoRequest{GroupId: id})
			s.Require().NoError(err)
			s.Assert().Equal(sdk.AccAddress([]byte(spec.req.Admin)), loadedGroupRes.Info.Admin)
			s.Assert().Equal(spec.req.Comment, loadedGroupRes.Info.Comment)
			s.Assert().Equal(id, loadedGroupRes.Info.GroupId)
			s.Assert().Equal(uint64(1), loadedGroupRes.Info.Version)

			// and members are stored as well
			membersRes, err := s.queryClient.GroupMembers(s.ctx, &group.QueryGroupMembersRequest{GroupId: id})
			s.Require().NoError(err)
			loadedMembers := membersRes.Members
			s.Require().Equal(len(members), len(loadedMembers))
			for i := range loadedMembers {
				s.Assert().Equal(members[i].Comment, loadedMembers[i].Comment)
				s.Assert().Equal(members[i].Address, loadedMembers[i].Member)
				s.Assert().Equal(members[i].Power, loadedMembers[i].Weight)
				s.Assert().Equal(id, loadedMembers[i].GroupId)
			}

			// query groups by admin
			groupsRes, err := s.queryClient.GroupsByAdmin(s.ctx, &group.QueryGroupsByAdminRequest{Admin: s.addr1})
			s.Require().NoError(err)
			loadedGroups := groupsRes.Groups
			s.Require().Equal(len(spec.expGroups), len(loadedGroups))
			for i := range loadedGroups {
				s.Assert().Equal(spec.expGroups[i].Comment, loadedGroups[i].Comment)
				s.Assert().Equal(spec.expGroups[i].Admin, loadedGroups[i].Admin)
				s.Assert().Equal(spec.expGroups[i].TotalWeight, loadedGroups[i].TotalWeight)
				s.Assert().Equal(spec.expGroups[i].GroupId, loadedGroups[i].GroupId)
				s.Assert().Equal(spec.expGroups[i].Version, loadedGroups[i].Version)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateGroupAdmin() {
	members := []group.Member{{
		Address: sdk.AccAddress([]byte("valid-member-address")),
		Power:   "1",
		Comment: "first member",
	}}
	oldAdmin := []byte("my-old-admin-address")
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:   oldAdmin,
		Members: members,
		Comment: "test",
	})
	s.Require().NoError(err)
	groupID := groupRes.GroupId

	specs := map[string]struct {
		req       *group.MsgUpdateGroupAdminRequest
		expStored *group.GroupInfo
		expErr    bool
	}{
		"with correct admin": {
			req: &group.MsgUpdateGroupAdminRequest{
				GroupId:  groupID,
				Admin:    oldAdmin,
				NewAdmin: []byte("my-new-admin-address"),
			},
			expStored: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       []byte("my-new-admin-address"),
				Comment:     "test",
				TotalWeight: "1",
				Version:     2,
			},
		},
		"with wrong admin": {
			req: &group.MsgUpdateGroupAdminRequest{
				GroupId:  groupID,
				Admin:    []byte("unknown-address"),
				NewAdmin: []byte("my-new-admin-address"),
			},
			expErr: true,
			expStored: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
		},
		"with unknown groupID": {
			req: &group.MsgUpdateGroupAdminRequest{
				GroupId:  999,
				Admin:    oldAdmin,
				NewAdmin: []byte("my-new-admin-address"),
			},
			expErr: true,
			expStored: &group.GroupInfo{
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
			res, err := s.queryClient.GroupInfo(s.ctx, &group.QueryGroupInfoRequest{GroupId: groupID})
			s.Require().NoError(err)
			s.Assert().Equal(spec.expStored, res.Info)
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateGroupComment() {
	oldComment := "first"
	members := []group.Member{{
		Address: sdk.AccAddress([]byte("valid-member-address")),
		Power:   "1",
		Comment: oldComment,
	}}

	oldAdmin := []byte("my-old-admin-address")
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:   oldAdmin,
		Members: members,
		Comment: "test",
	})
	s.Require().NoError(err)
	groupID := groupRes.GroupId

	specs := map[string]struct {
		req       *group.MsgUpdateGroupCommentRequest
		expErr    bool
		expStored *group.GroupInfo
	}{
		"with correct admin": {
			req: &group.MsgUpdateGroupCommentRequest{
				GroupId: groupID,
				Admin:   oldAdmin,
				Comment: "new comment",
			},
			expStored: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Comment:     "new comment",
				TotalWeight: "1",
				Version:     2,
			},
		},
		"with wrong admin": {
			req: &group.MsgUpdateGroupCommentRequest{
				GroupId: groupID,
				Admin:   []byte("unknown-address"),
				Comment: "new comment",
			},
			expErr: true,
			expStored: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
		},
		"with unknown groupid": {
			req: &group.MsgUpdateGroupCommentRequest{
				GroupId: 999,
				Admin:   []byte("unknown-address"),
				Comment: "new comment",
			},
			expErr: true,
			expStored: &group.GroupInfo{
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
			res, err := s.queryClient.GroupInfo(s.ctx, &group.QueryGroupInfoRequest{GroupId: groupID})
			s.Require().NoError(err)
			s.Assert().Equal(spec.expStored, res.Info)
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateGroupMembers() {
	members := []group.Member{{
		Address: sdk.AccAddress([]byte("valid-member-address")),
		Power:   "1",
		Comment: "first",
	}}

	myAdmin := []byte("valid--admin-address")
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:   myAdmin,
		Members: members,
		Comment: "test",
	})
	s.Require().NoError(err)
	groupID := groupRes.GroupId

	specs := map[string]struct {
		req        *group.MsgUpdateGroupMembersRequest
		expErr     bool
		expGroup   *group.GroupInfo
		expMembers []*group.GroupMember
	}{
		"add new member": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{{
					Address: sdk.AccAddress([]byte("other-member-address")),
					Power:   "2",
					Comment: "second",
				}},
			},
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "3",
				Version:     2,
			},
			expMembers: []*group.GroupMember{
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
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{{
					Address: sdk.AccAddress([]byte("valid-member-address")),
					Power:   "2",
					Comment: "updated",
				}},
			},
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "2",
				Version:     2,
			},
			expMembers: []*group.GroupMember{
				{
					Member:  sdk.AccAddress([]byte("valid-member-address")),
					GroupId: groupID,
					Weight:  "2",
					Comment: "updated",
				},
			},
		},
		"update member with same data": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{{
					Address: sdk.AccAddress([]byte("valid-member-address")),
					Power:   "1",
					Comment: "first",
				}},
			},
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     2,
			},
			expMembers: []*group.GroupMember{
				{
					Member:  sdk.AccAddress([]byte("valid-member-address")),
					GroupId: groupID,
					Weight:  "1",
					Comment: "first",
				},
			},
		},
		"replace member": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{
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
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     2,
			},
			expMembers: []*group.GroupMember{{
				Member:  s.addr1,
				GroupId: groupID,
				Weight:  "1",
				Comment: "welcome",
			}},
		},
		"remove existing member": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{{
					Address: sdk.AccAddress([]byte("valid-member-address")),
					Power:   "0",
					Comment: "good bye",
				}},
			},
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "0",
				Version:     2,
			},
			expMembers: []*group.GroupMember{},
		},
		"remove unknown member": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{{
					Address: sdk.AccAddress([]byte("unknown-member-addrs")),
					Power:   "0",
					Comment: "good bye",
				}},
			},
			expErr: true,
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
			expMembers: []*group.GroupMember{{
				Member:  sdk.AccAddress([]byte("valid-member-address")),
				GroupId: groupID,
				Weight:  "1",
				Comment: "first",
			}},
		},
		"with wrong admin": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   []byte("unknown-address"),
				MemberUpdates: []group.Member{{
					Address: sdk.AccAddress([]byte("other-member-address")),
					Power:   "2",
					Comment: "second",
				}},
			},
			expErr: true,
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
			expMembers: []*group.GroupMember{{
				Member:  sdk.AccAddress([]byte("valid-member-address")),
				GroupId: groupID,
				Weight:  "1",
				Comment: "first",
			}},
		},
		"with unknown groupID": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: 999,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{{
					Address: sdk.AccAddress([]byte("other-member-address")),
					Power:   "2",
					Comment: "second",
				}},
			},
			expErr: true,
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Comment:     "test",
				TotalWeight: "1",
				Version:     1,
			},
			expMembers: []*group.GroupMember{{
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
			sdkCtx, _ := s.sdkCtx.CacheContext()
			ctx := sdk.WrapSDKContext(sdkCtx)
			_, err := s.msgClient.UpdateGroupMembers(ctx, spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			// then
			res, err := s.queryClient.GroupInfo(ctx, &group.QueryGroupInfoRequest{GroupId: groupID})
			s.Require().NoError(err)
			s.Assert().Equal(spec.expGroup, res.Info)

			// and members persisted
			membersRes, err := s.queryClient.GroupMembers(ctx, &group.QueryGroupMembersRequest{GroupId: groupID})
			s.Require().NoError(err)
			loadedMembers := membersRes.Members
			s.Assert().Equal(spec.expMembers, loadedMembers)
		})
	}
}

func (s *IntegrationTestSuite) TestCreateGroupAccount() {
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:   s.addr1,
		Members: nil,
		Comment: "test",
	})
	s.Require().NoError(err)
	myGroupID := groupRes.GroupId

	specs := map[string]struct {
		req    *group.MsgCreateGroupAccountRequest
		policy group.DecisionPolicy
		expErr bool
	}{
		"all good": {
			req: &group.MsgCreateGroupAccountRequest{
				Admin:   s.addr1,
				Comment: "test 1",
				GroupId: myGroupID,
			},
			policy: group.NewThresholdDecisionPolicy(
				"1",
				gogotypes.Duration{Seconds: 1},
			),
		},
		"decision policy threshold > total group weight": {
			req: &group.MsgCreateGroupAccountRequest{
				Admin:   s.addr1,
				Comment: "test 2",
				GroupId: myGroupID,
			},
			policy: group.NewThresholdDecisionPolicy(
				"10",
				gogotypes.Duration{Seconds: 1},
			),
		},
		"group id does not exists": {
			req: &group.MsgCreateGroupAccountRequest{
				Admin:   s.addr1,
				Comment: "test",
				GroupId: 9999,
			},
			policy: group.NewThresholdDecisionPolicy(
				"1",
				gogotypes.Duration{Seconds: 1},
			),
			expErr: true,
		},
		"admin not group admin": {
			req: &group.MsgCreateGroupAccountRequest{
				Admin:   []byte("other--admin-address"),
				Comment: "test",
				GroupId: myGroupID,
			},
			policy: group.NewThresholdDecisionPolicy(
				"1",
				gogotypes.Duration{Seconds: 1},
			),
			expErr: true,
		},
		"comment too long": {
			req: &group.MsgCreateGroupAccountRequest{
				Admin:   s.addr1,
				Comment: strings.Repeat("a", 256),
				GroupId: myGroupID,
			},
			policy: group.NewThresholdDecisionPolicy(
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
			groupAccountRes, err := s.queryClient.GroupAccountInfo(s.ctx, &group.QueryGroupAccountInfoRequest{GroupAccount: addr})
			s.Require().NoError(err)

			groupAccount := groupAccountRes.Info
			s.Assert().Equal(addr, groupAccount.GroupAccount)
			s.Assert().Equal(myGroupID, groupAccount.GroupId)
			s.Assert().Equal(sdk.AccAddress([]byte(spec.req.Admin)), groupAccount.Admin)
			s.Assert().Equal(spec.req.Comment, groupAccount.Comment)
			s.Assert().Equal(uint64(1), groupAccount.Version)
			s.Assert().Equal(spec.policy.(*group.ThresholdDecisionPolicy), groupAccount.GetDecisionPolicy())
		})
	}
}

func (s *IntegrationTestSuite) TestGroupAccountsByAdminOrGroup() {
	admin := s.addr2
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:   admin,
		Members: nil,
		Comment: "test",
	})
	s.Require().NoError(err)
	myGroupID := groupRes.GroupId

	policies := []group.DecisionPolicy{
		group.NewThresholdDecisionPolicy(
			"1",
			gogotypes.Duration{Seconds: 1},
		),
		group.NewThresholdDecisionPolicy(
			"10",
			gogotypes.Duration{Seconds: 1},
		),
	}

	count := 2
	addrs := make([]sdk.AccAddress, count)
	reqs := make([]*group.MsgCreateGroupAccountRequest, count)
	for i := range addrs {
		req := &group.MsgCreateGroupAccountRequest{
			Admin:   admin,
			Comment: fmt.Sprintf("test %d", i),
			GroupId: myGroupID,
		}
		err := req.SetDecisionPolicy(policies[i])
		s.Require().NoError(err)
		reqs[i] = req
		res, err := s.msgClient.CreateGroupAccount(s.ctx, req)
		s.Require().NoError(err)
		addrs[i] = res.GroupAccount
	}

	// query group account by group
	accountsByGroupRes, err := s.queryClient.GroupAccountsByGroup(s.ctx, &group.QueryGroupAccountsByGroupRequest{
		GroupId: myGroupID,
	})
	s.Require().NoError(err)
	accounts := accountsByGroupRes.GroupAccounts
	s.Require().Equal(len(accounts), count)
	for i := range accounts {
		s.Assert().Equal(addrs[i], accounts[i].GroupAccount)
		s.Assert().Equal(myGroupID, accounts[i].GroupId)
		s.Assert().Equal(admin, accounts[i].Admin)
		s.Assert().Equal(reqs[i].Comment, accounts[i].Comment)
		s.Assert().Equal(uint64(1), accounts[i].Version)
		s.Assert().Equal(policies[i].(*group.ThresholdDecisionPolicy), accounts[i].GetDecisionPolicy())
	}

	// query group account by admin
	accountsByAdminRes, err := s.queryClient.GroupAccountsByAdmin(s.ctx, &group.QueryGroupAccountsByAdminRequest{
		Admin: admin,
	})
	s.Require().NoError(err)
	accounts = accountsByAdminRes.GroupAccounts
	s.Require().Equal(len(accounts), count)
	for i := range accounts {
		s.Assert().Equal(addrs[i], accounts[i].GroupAccount)
		s.Assert().Equal(myGroupID, accounts[i].GroupId)
		s.Assert().Equal(admin, accounts[i].Admin)
		s.Assert().Equal(reqs[i].Comment, accounts[i].Comment)
		s.Assert().Equal(uint64(1), accounts[i].Version)
		s.Assert().Equal(policies[i].(*group.ThresholdDecisionPolicy), accounts[i].GetDecisionPolicy())
	}
}

func (s *IntegrationTestSuite) TestCreateProposal() {
	members := []group.Member{{
		Address: []byte("valid-member-address"),
		Power:   "1",
	}}
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:   []byte("valid--admin-address"),
		Members: members,
		Comment: "test",
	})
	s.Require().NoError(err)
	myGroupID := groupRes.GroupId

	policy := group.NewThresholdDecisionPolicy(
		"1",
		gogotypes.Duration{Seconds: 1},
	)
	accountReq := &group.MsgCreateGroupAccountRequest{
		Admin:   []byte("valid--admin-address"),
		GroupId: myGroupID,
		Comment: "test",
	}
	err = accountReq.SetDecisionPolicy(policy)
	s.Require().NoError(err)
	accountRes, err := s.msgClient.CreateGroupAccount(s.ctx, accountReq)
	s.Require().NoError(err)
	accountAddr := accountRes.GroupAccount

	policy = group.NewThresholdDecisionPolicy(
		"100",
		gogotypes.Duration{Seconds: 1},
	)
	err = accountReq.SetDecisionPolicy(policy)
	s.Require().NoError(err)
	bigThresholdRes, err := s.msgClient.CreateGroupAccount(s.ctx, accountReq)
	s.Require().NoError(err)
	bigThresholdAddr := bigThresholdRes.GroupAccount

	specs := map[string]struct {
		req    *group.MsgCreateProposalRequest
		msgs   []sdk.Msg
		expErr bool
	}{
		"all good with minimal fields set": {
			req: &group.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
			},
		},
		"all good with good msg payload": {
			req: &group.MsgCreateProposalRequest{
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
			req: &group.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Comment:      strings.Repeat("a", 256),
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"group account required": {
			req: &group.MsgCreateProposalRequest{
				Comment:   "test",
				Proposers: []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"existing group account required": {
			req: &group.MsgCreateProposalRequest{
				GroupAccount: []byte("non-existing-account"),
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"impossible case: decision policy threshold > total group weight": {
			req: &group.MsgCreateProposalRequest{
				GroupAccount: bigThresholdAddr,
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"only group members can create a proposal": {
			req: &group.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Proposers:    []sdk.AccAddress{[]byte("non--member-address")},
			},
			expErr: true,
		},
		"all proposers must be in group": {
			req: &group.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address"), []byte("non--member-address")},
			},
			expErr: true,
		},
		"proposers must not be nil": {
			req: &group.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address"), nil},
			},
			expErr: true,
		},
		"admin that is not a group member can not create proposal": {
			req: &group.MsgCreateProposalRequest{
				GroupAccount: accountAddr,
				Comment:      "test",
				Proposers:    []sdk.AccAddress{[]byte("valid--admin-address")},
			},
			expErr: true,
		},
		"reject msgs that are not authz by group account": {
			req: &group.MsgCreateProposalRequest{
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
			proposalRes, err := s.queryClient.Proposal(s.ctx, &group.QueryProposalRequest{ProposalId: id})
			s.Require().NoError(err)
			proposal := proposalRes.Proposal

			s.Assert().Equal(accountAddr, proposal.GroupAccount)
			s.Assert().Equal(spec.req.Comment, proposal.Comment)
			s.Assert().Equal(spec.req.Proposers, proposal.Proposers)

			submittedAt, err := gogotypes.TimestampFromProto(&proposal.SubmittedAt)
			s.Require().NoError(err)
			s.Assert().Equal(s.blockTime, submittedAt)

			s.Assert().Equal(uint64(1), proposal.GroupVersion)
			s.Assert().Equal(uint64(1), proposal.GroupAccountVersion)
			s.Assert().Equal(group.ProposalStatusSubmitted, proposal.Status)
			s.Assert().Equal(group.ProposalResultUnfinalized, proposal.Result)
			s.Assert().Equal(group.Tally{
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
	members := []group.Member{
		{Address: []byte("valid-member-address"), Power: "1"},
		{Address: []byte("power-member-address"), Power: "2"},
	}
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:   []byte("valid--admin-address"),
		Members: members,
		Comment: "test",
	})
	s.Require().NoError(err)
	myGroupID := groupRes.GroupId

	policy := group.NewThresholdDecisionPolicy(
		"2",
		gogotypes.Duration{Seconds: 1},
	)
	accountReq := &group.MsgCreateGroupAccountRequest{
		Admin:   []byte("valid--admin-address"),
		GroupId: myGroupID,
		Comment: "test",
	}
	err = accountReq.SetDecisionPolicy(policy)
	s.Require().NoError(err)
	accountRes, err := s.msgClient.CreateGroupAccount(s.ctx, accountReq)
	s.Require().NoError(err)
	accountAddr := accountRes.GroupAccount

	req := &group.MsgCreateProposalRequest{
		GroupAccount: accountAddr,
		Comment:      "integration test",
		Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
		Msgs:         nil,
	}
	proposalRes, err := s.msgClient.CreateProposal(s.ctx, req)
	s.Require().NoError(err)
	myProposalID := proposalRes.ProposalId

	// proposals by group account
	proposalsRes, err := s.queryClient.ProposalsByGroupAccount(s.ctx, &group.QueryProposalsByGroupAccountRequest{
		GroupAccount: accountAddr,
	})
	s.Require().NoError(err)
	proposals := proposalsRes.Proposals
	s.Require().Equal(len(proposals), 1)
	s.Assert().Equal(req.GroupAccount, proposals[0].GroupAccount)
	s.Assert().Equal(req.Comment, proposals[0].Comment)
	s.Assert().Equal(req.Proposers, proposals[0].Proposers)

	submittedAt, err := gogotypes.TimestampFromProto(&proposals[0].SubmittedAt)
	s.Require().NoError(err)
	s.Assert().Equal(s.blockTime, submittedAt)

	s.Assert().Equal(uint64(1), proposals[0].GroupVersion)
	s.Assert().Equal(uint64(1), proposals[0].GroupAccountVersion)
	s.Assert().Equal(group.ProposalStatusSubmitted, proposals[0].Status)
	s.Assert().Equal(group.ProposalResultUnfinalized, proposals[0].Result)
	s.Assert().Equal(group.Tally{
		YesCount:     "0",
		NoCount:      "0",
		AbstainCount: "0",
		VetoCount:    "0",
	}, proposals[0].VoteState)

	specs := map[string]struct {
		req               *group.MsgVoteRequest
		srcCtx            sdk.Context
		doBefore          func(ctx context.Context)
		expErr            bool
		expVoteState      group.Tally
		expProposalStatus group.Proposal_Status
		expResult         group.Proposal_Result
	}{
		"vote yes": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     group.Choice_CHOICE_YES,
			},
			expVoteState: group.Tally{
				YesCount:     "1",
				NoCount:      "0",
				AbstainCount: "0",
				VetoCount:    "0",
			},
			expProposalStatus: group.ProposalStatusSubmitted,
			expResult:         group.ProposalResultUnfinalized,
		},
		"vote no": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     group.Choice_CHOICE_NO,
			},
			expVoteState: group.Tally{
				YesCount:     "0",
				NoCount:      "1",
				AbstainCount: "0",
				VetoCount:    "0",
			},
			expProposalStatus: group.ProposalStatusSubmitted,
			expResult:         group.ProposalResultUnfinalized,
		},
		"vote abstain": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     group.Choice_CHOICE_ABSTAIN,
			},
			expVoteState: group.Tally{
				YesCount:     "0",
				NoCount:      "0",
				AbstainCount: "1",
				VetoCount:    "0",
			},
			expProposalStatus: group.ProposalStatusSubmitted,
			expResult:         group.ProposalResultUnfinalized,
		},
		"vote veto": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     group.Choice_CHOICE_VETO,
			},
			expVoteState: group.Tally{
				YesCount:     "0",
				NoCount:      "0",
				AbstainCount: "0",
				VetoCount:    "1",
			},
			expProposalStatus: group.ProposalStatusSubmitted,
			expResult:         group.ProposalResultUnfinalized,
		},
		"apply decision policy early": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("power-member-address")},
				Choice:     group.Choice_CHOICE_YES,
			},
			expVoteState: group.Tally{
				YesCount:     "2",
				NoCount:      "0",
				AbstainCount: "0",
				VetoCount:    "0",
			},
			expProposalStatus: group.ProposalStatusClosed,
			expResult:         group.ProposalResultAccepted,
		},
		"reject new votes when final decision is made already": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     group.Choice_CHOICE_YES,
			},
			doBefore: func(ctx context.Context) {
				_, err := s.msgClient.Vote(ctx, &group.MsgVoteRequest{
					ProposalId: myProposalID,
					Voters:     []sdk.AccAddress{[]byte("power-member-address")},
					Choice:     group.Choice_CHOICE_VETO,
				})
				s.Require().NoError(err)
			},
			expErr: true,
		},
		"comment too long": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Comment:    strings.Repeat("a", 256),
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     group.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"existing proposal required": {
			req: &group.MsgVoteRequest{
				ProposalId: 999,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     group.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"empty choice": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"invalid choice": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     5,
			},
			expErr: true,
		},
		"all voters must be in group": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address"), []byte("non--member-address")},
				Choice:     group.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"voters must not include nil": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address"), nil},
				Choice:     group.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"voters must not be nil": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Choice:     group.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"voters must not be empty": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Choice:     group.Choice_CHOICE_NO,
				Voters:     []sdk.AccAddress{},
			},
			expErr: true,
		},
		"admin that is not a group member can not vote": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid--admin-address")},
				Choice:     group.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"on timeout": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     group.Choice_CHOICE_NO,
			},
			srcCtx: s.sdkCtx.WithBlockTime(s.blockTime.Add(time.Second)),
			expErr: true,
		},
		"closed already": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     group.Choice_CHOICE_NO,
			},
			doBefore: func(ctx context.Context) {
				_, err := s.msgClient.Vote(ctx, &group.MsgVoteRequest{
					ProposalId: myProposalID,
					Voters:     []sdk.AccAddress{[]byte("power-member-address")},
					Choice:     group.Choice_CHOICE_YES,
				})
				s.Require().NoError(err)
			},
			expErr: true,
		},
		"voted already": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     group.Choice_CHOICE_NO,
			},
			doBefore: func(ctx context.Context) {
				_, err := s.msgClient.Vote(ctx, &group.MsgVoteRequest{
					ProposalId: myProposalID,
					Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
					Choice:     group.Choice_CHOICE_YES,
				})
				s.Require().NoError(err)
			},
			expErr: true,
		},
		"with group modified": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
				Choice:     group.Choice_CHOICE_NO,
			},
			doBefore: func(ctx context.Context) {
				_, err = s.msgClient.UpdateGroupComment(ctx, &group.MsgUpdateGroupCommentRequest{
					GroupId: myGroupID,
					Admin:   []byte("valid--admin-address"),
					Comment: "group modified",
				})
				s.Require().NoError(err)
			},
			expErr: true,
		},
		// TODO Need to implement group account updates
		// "with policy modified": {
		// 	req: &group.MsgVoteRequest{
		// 		ProposalId: myProposalID,
		// 		Voters:     []sdk.AccAddress{[]byte("valid-member-address")},
		// 		Choice:     group.Choice_CHOICE_NO,
		// 	},
		// 	doBefore: func(ctx context.Context) {
		// 		a, err := s.groupKeeper.GetGroupAccount(ctx, accountAddr)
		// 		s.Require().NoError(err)
		// 		a.Comment = "policy modified"
		// 		s.Require().NoError(s.groupKeeper.UpdateGroupAccount(ctx, &a))
		// 	},
		// 	expErr: true,
		// },
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			sdkCtx := s.sdkCtx
			if !spec.srcCtx.IsZero() {
				sdkCtx = spec.srcCtx
			}
			sdkCtx, _ = sdkCtx.CacheContext()
			ctx := sdk.WrapSDKContext(sdkCtx)

			if spec.doBefore != nil {
				spec.doBefore(ctx)
			}
			_, err := s.msgClient.Vote(ctx, spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			s.Require().NoError(err)
			// and all votes are stored
			for _, voter := range spec.req.Voters {
				// then all data persisted
				res, err := s.queryClient.VoteByProposalVoter(ctx, &group.QueryVoteByProposalVoterRequest{
					ProposalId: spec.req.ProposalId,
					Voter:      voter,
				})
				s.Require().NoError(err)
				loaded := res.Vote
				s.Assert().Equal(spec.req.ProposalId, loaded.ProposalId)
				s.Assert().Equal(voter, loaded.Voter)
				s.Assert().Equal(spec.req.Choice, loaded.Choice)
				s.Assert().Equal(spec.req.Comment, loaded.Comment)
				submittedAt, err := gogotypes.TimestampFromProto(&loaded.SubmittedAt)
				s.Require().NoError(err)
				s.Assert().Equal(s.blockTime, submittedAt)
			}

			// query votes by proposal
			votesRes, err := s.queryClient.VotesByProposal(ctx, &group.QueryVotesByProposalRequest{
				ProposalId: spec.req.ProposalId,
			})
			s.Require().NoError(err)
			votes := votesRes.Votes
			s.Require().Equal(len(spec.req.Voters), len(votes))
			for i, vote := range votes {
				s.Assert().Equal(spec.req.ProposalId, vote.ProposalId)
				s.Assert().Equal(spec.req.Voters[i], vote.Voter)
				s.Assert().Equal(spec.req.Choice, vote.Choice)
				s.Assert().Equal(spec.req.Comment, vote.Comment)
				submittedAt, err := gogotypes.TimestampFromProto(&vote.SubmittedAt)
				s.Require().NoError(err)
				s.Assert().Equal(s.blockTime, submittedAt)
			}

			// query votes by voter
			for _, voter := range spec.req.Voters {
				// then all data persisted
				res, err := s.queryClient.VotesByVoter(ctx, &group.QueryVotesByVoterRequest{
					Voter: voter,
				})
				s.Require().NoError(err)
				votes := res.Votes
				s.Require().Equal(1, len(votes))
				s.Assert().Equal(spec.req.ProposalId, votes[0].ProposalId)
				s.Assert().Equal(voter, votes[0].Voter)
				s.Assert().Equal(spec.req.Choice, votes[0].Choice)
				s.Assert().Equal(spec.req.Comment, votes[0].Comment)
				submittedAt, err := gogotypes.TimestampFromProto(&votes[0].SubmittedAt)
				s.Require().NoError(err)
				s.Assert().Equal(s.blockTime, submittedAt)
			}

			// and proposal is updated
			proposalRes, err := s.queryClient.Proposal(ctx, &group.QueryProposalRequest{
				ProposalId: spec.req.ProposalId,
			})
			s.Require().NoError(err)
			proposal := proposalRes.Proposal
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
		setupProposal     func(ctx context.Context) group.ProposalID
		expErr            bool
		expProposalStatus group.Proposal_Status
		expProposalResult group.Proposal_Result
		expExecutorResult group.Proposal_ExecutorResult
		expFromBalances   sdk.Coins
		expToBalances     sdk.Coins
	}{
		"proposal executed when accepted": {
			setupProposal: func(ctx context.Context) group.ProposalID {
				msgs := []sdk.Msg{msgSend}
				return createProposalAndVote(ctx, s, msgs, proposers, group.Choice_CHOICE_YES)
			},
			expProposalStatus: group.ProposalStatusClosed,
			expProposalResult: group.ProposalResultAccepted,
			expExecutorResult: group.ProposalExecutorResultSuccess,
			expFromBalances:   sdk.Coins{sdk.NewInt64Coin("test", 9900)},
			expToBalances:     sdk.Coins{sdk.NewInt64Coin("test", 100)},
		},
		"proposal with multiple messages executed when accepted": {
			setupProposal: func(ctx context.Context) group.ProposalID {
				msgs := []sdk.Msg{msgSend, msgSend}
				return createProposalAndVote(ctx, s, msgs, proposers, group.Choice_CHOICE_YES)
			},
			expProposalStatus: group.ProposalStatusClosed,
			expProposalResult: group.ProposalResultAccepted,
			expExecutorResult: group.ProposalExecutorResultSuccess,
			expFromBalances:   sdk.Coins{sdk.NewInt64Coin("test", 9800)},
			expToBalances:     sdk.Coins{sdk.NewInt64Coin("test", 200)},
		},
		"proposal not executed when rejected": {
			setupProposal: func(ctx context.Context) group.ProposalID {
				msgs := []sdk.Msg{msgSend}
				return createProposalAndVote(ctx, s, msgs, proposers, group.Choice_CHOICE_NO)
			},
			expProposalStatus: group.ProposalStatusClosed,
			expProposalResult: group.ProposalResultRejected,
			expExecutorResult: group.ProposalExecutorResultNotRun,
		},
		"open proposal must not fail": {
			setupProposal: func(ctx context.Context) group.ProposalID {
				return createProposal(ctx, s, []sdk.Msg{msgSend}, proposers)
			},
			expProposalStatus: group.ProposalStatusSubmitted,
			expProposalResult: group.ProposalResultUnfinalized,
			expExecutorResult: group.ProposalExecutorResultNotRun,
		},
		"existing proposal required": {
			setupProposal: func(ctx context.Context) group.ProposalID {
				return 9999
			},
			expErr: true,
		},
		"Decision policy also applied on timeout": {
			setupProposal: func(ctx context.Context) group.ProposalID {
				msgs := []sdk.Msg{msgSend}
				return createProposalAndVote(ctx, s, msgs, proposers, group.Choice_CHOICE_NO)
			},
			srcBlockTime:      s.blockTime.Add(time.Second),
			expProposalStatus: group.ProposalStatusClosed,
			expProposalResult: group.ProposalResultRejected,
			expExecutorResult: group.ProposalExecutorResultNotRun,
		},
		"Decision policy also applied after timeout": {
			setupProposal: func(ctx context.Context) group.ProposalID {
				msgs := []sdk.Msg{msgSend}
				return createProposalAndVote(ctx, s, msgs, proposers, group.Choice_CHOICE_NO)
			},
			srcBlockTime:      s.blockTime.Add(time.Second).Add(time.Millisecond),
			expProposalStatus: group.ProposalStatusClosed,
			expProposalResult: group.ProposalResultRejected,
			expExecutorResult: group.ProposalExecutorResultNotRun,
		},
		"with group modified before tally": {
			setupProposal: func(ctx context.Context) group.ProposalID {
				myProposalID := createProposal(ctx, s, []sdk.Msg{msgSend}, proposers)

				// then modify group
				_, err := s.msgClient.UpdateGroupComment(ctx, &group.MsgUpdateGroupCommentRequest{
					Admin:   s.addr1,
					GroupId: s.groupID,
					Comment: "group modified before tally",
				})
				s.Require().NoError(err)
				return myProposalID
			},
			expProposalStatus: group.ProposalStatusAborted,
			expProposalResult: group.ProposalResultUnfinalized,
			expExecutorResult: group.ProposalExecutorResultNotRun,
		},
		// TODO Need to implement group account update
		// "with group account modified before tally": {
		// 	setupProposal: func(ctx context.Context) group.ProposalID {
		// 		myProposalID := createProposal(ctx, s, []sdk.Msg{msgSend}, proposers)

		// 		// then modify group account
		// 		a, err := s.groupKeeper.GetGroupAccount(ctx, s.groupAccountAddr)
		// 		s.Require().NoError(err)
		// 		a.Comment = "group account modified before tally"
		// 		s.Require().NoError(s.groupKeeper.UpdateGroupAccount(ctx, &a))
		// 		return myProposalID
		// 	},
		// 	expProposalStatus: group.ProposalStatusAborted,
		// 	expProposalResult: group.ProposalResultUnfinalized,
		// 	expExecutorResult: group.ProposalExecutorResultNotRun,
		// },
		"prevent double execution when successful": {
			setupProposal: func(ctx context.Context) group.ProposalID {
				myProposalID := createProposalAndVote(ctx, s, []sdk.Msg{msgSend}, proposers, group.Choice_CHOICE_YES)

				_, err := s.msgClient.Exec(ctx, &group.MsgExecRequest{ProposalId: myProposalID})
				s.Require().NoError(err)
				return myProposalID
			},
			expProposalStatus: group.ProposalStatusClosed,
			expProposalResult: group.ProposalResultAccepted,
			expExecutorResult: group.ProposalExecutorResultSuccess,
			expFromBalances:   sdk.Coins{sdk.NewInt64Coin("test", 9900)},
			expToBalances:     sdk.Coins{sdk.NewInt64Coin("test", 100)},
		},
		"rollback all msg updates on failure": {
			setupProposal: func(ctx context.Context) group.ProposalID {
				msgs := []sdk.Msg{
					msgSend, &banktypes.MsgSend{
						FromAddress: s.groupAccountAddr.String(),
						ToAddress:   s.addr2.String(),
						Amount:      sdk.Coins{sdk.NewInt64Coin("test", 10001)}},
				}
				return createProposalAndVote(ctx, s, msgs, proposers, group.Choice_CHOICE_YES)
			},
			expProposalStatus: group.ProposalStatusClosed,
			expProposalResult: group.ProposalResultAccepted,
			expExecutorResult: group.ProposalExecutorResultFailure,
		},
		"executable when failed before": {
			setupProposal: func(ctx context.Context) group.ProposalID {
				msgs := []sdk.Msg{
					&banktypes.MsgSend{
						FromAddress: s.groupAccountAddr.String(),
						ToAddress:   s.addr2.String(),
						Amount:      sdk.Coins{sdk.NewInt64Coin("test", 10001)}},
				}
				myProposalID := createProposalAndVote(ctx, s, msgs, proposers, group.Choice_CHOICE_YES)

				_, err := s.msgClient.Exec(ctx, &group.MsgExecRequest{ProposalId: myProposalID})
				s.Require().NoError(err)
				s.Require().NoError(s.bankKeeper.SetBalances(sdk.UnwrapSDKContext(ctx), s.groupAccountAddr, sdk.Coins{sdk.NewInt64Coin("test", 10002)}))
				return myProposalID
			},
			expProposalStatus: group.ProposalStatusClosed,
			expProposalResult: group.ProposalResultAccepted,
			expExecutorResult: group.ProposalExecutorResultSuccess,
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			sdkCtx, _ := s.sdkCtx.CacheContext()
			ctx := sdk.WrapSDKContext(sdkCtx)

			proposalID := spec.setupProposal(ctx)

			if !spec.srcBlockTime.IsZero() {
				sdkCtx = sdkCtx.WithBlockTime(spec.srcBlockTime)
				ctx = sdk.WrapSDKContext(sdkCtx)
			}

			_, err := s.msgClient.Exec(ctx, &group.MsgExecRequest{ProposalId: proposalID})
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			// and proposal is updated
			res, err := s.queryClient.Proposal(ctx, &group.QueryProposalRequest{ProposalId: proposalID})
			s.Require().NoError(err)
			proposal := res.Proposal

			exp := group.Proposal_Result_name[int32(spec.expProposalResult)]
			got := group.Proposal_Result_name[int32(proposal.Result)]
			s.Assert().Equal(exp, got)

			exp = group.Proposal_Status_name[int32(spec.expProposalStatus)]
			got = group.Proposal_Status_name[int32(proposal.Status)]
			s.Assert().Equal(exp, got)

			exp = group.Proposal_ExecutorResult_name[int32(spec.expExecutorResult)]
			got = group.Proposal_ExecutorResult_name[int32(proposal.ExecutorResult)]
			s.Assert().Equal(exp, got)

			if spec.expFromBalances != nil {
				fromBalances := s.bankKeeper.GetAllBalances(sdkCtx, s.groupAccountAddr)
				s.Require().Equal(spec.expFromBalances, fromBalances)
			}
			if spec.expToBalances != nil {
				toBalances := s.bankKeeper.GetAllBalances(sdkCtx, s.addr2)
				s.Require().Equal(spec.expToBalances, toBalances)
			}
		})
	}
}

func createProposal(
	ctx context.Context, s *IntegrationTestSuite, msgs []sdk.Msg,
	proposers []sdk.AccAddress) group.ProposalID {
	proposalReq := &group.MsgCreateProposalRequest{
		GroupAccount: s.groupAccountAddr,
		Proposers:    proposers,
		Comment:      "test",
	}
	err := proposalReq.SetMsgs(msgs)
	s.Require().NoError(err)

	proposalRes, err := s.msgClient.CreateProposal(ctx, proposalReq)
	s.Require().NoError(err)
	return proposalRes.ProposalId
}

func createProposalAndVote(
	ctx context.Context, s *IntegrationTestSuite, msgs []sdk.Msg,
	proposers []sdk.AccAddress, choice group.Choice) group.ProposalID {
	myProposalID := createProposal(ctx, s, msgs, proposers)

	_, err := s.msgClient.Vote(ctx, &group.MsgVoteRequest{
		ProposalId: myProposalID,
		Voters:     proposers,
		Choice:     choice,
	})
	s.Require().NoError(err)
	return myProposalID
}
