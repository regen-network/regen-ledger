package testsuite

import (
	"bytes"
	"context"
	"sort"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/suite"

	"github.com/regen-network/regen-ledger/types"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/group"
	groupserver "github.com/regen-network/regen-ledger/x/group/server"
	"github.com/regen-network/regen-ledger/x/group/testdata"
)

type IntegrationTestSuite struct {
	suite.Suite

	fixtureFactory *servermodule.FixtureFactory
	fixture        testutil.Fixture

	ctx              context.Context
	sdkCtx           sdk.Context
	genesisCtx       types.Context
	msgClient        group.MsgClient
	queryClient      group.QueryClient
	addr1            sdk.AccAddress
	addr2            sdk.AccAddress
	addr3            sdk.AccAddress
	addr4            sdk.AccAddress
	addr5            sdk.AccAddress
	addr6            sdk.AccAddress
	groupAccountAddr sdk.AccAddress
	groupID          uint64

	accountKeeper authkeeper.AccountKeeper
	bankKeeper    bankkeeper.Keeper
	mintKeeper    mintkeeper.Keeper

	blockTime time.Time
}

func NewIntegrationTestSuite(fixtureFactory *servermodule.FixtureFactory, accountKeeper authkeeper.AccountKeeper, bankKeeper bankkeeper.BaseKeeper, mintKeeper mintkeeper.Keeper) *IntegrationTestSuite {
	return &IntegrationTestSuite{
		fixtureFactory: fixtureFactory,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
		mintKeeper:     mintKeeper,
	}
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.fixture = s.fixtureFactory.Setup()

	s.blockTime = time.Now().UTC()

	// TODO clean up once types.Context merged upstream into sdk.Context
	sdkCtx := s.fixture.Context().(types.Context).WithBlockTime(s.blockTime)
	s.sdkCtx, _ = sdkCtx.CacheContext()
	s.ctx = types.Context{Context: s.sdkCtx}
	s.genesisCtx = types.Context{Context: sdkCtx}

	s.Require().NoError(s.bankKeeper.MintCoins(s.sdkCtx, minttypes.ModuleName, sdk.NewCoins(sdk.NewInt64Coin("test", 400000000))))

	s.bankKeeper.SetParams(sdkCtx, banktypes.DefaultParams())

	s.msgClient = group.NewMsgClient(s.fixture.TxConn())
	s.queryClient = group.NewQueryClient(s.fixture.QueryConn())

	s.Require().GreaterOrEqual(len(s.fixture.Signers()), 6)
	s.addr1 = s.fixture.Signers()[0]
	s.addr2 = s.fixture.Signers()[1]
	s.addr3 = s.fixture.Signers()[2]
	s.addr4 = s.fixture.Signers()[3]
	s.addr5 = s.fixture.Signers()[4]
	s.addr6 = s.fixture.Signers()[5]

	// Initial group, group account and balance setup
	members := []group.Member{
		{Address: s.addr2.String(), Weight: "1"},
	}
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:    s.addr1.String(),
		Members:  members,
		Metadata: nil,
	})
	s.Require().NoError(err)
	s.groupID = groupRes.GroupId

	policy := group.NewThresholdDecisionPolicy(
		"1",
		gogotypes.Duration{Seconds: 1},
	)
	accountReq := &group.MsgCreateGroupAccountRequest{
		Admin:    s.addr1.String(),
		GroupId:  s.groupID,
		Metadata: nil,
	}
	err = accountReq.SetDecisionPolicy(policy)
	s.Require().NoError(err)
	accountRes, err := s.msgClient.CreateGroupAccount(s.ctx, accountReq)
	s.Require().NoError(err)
	addr, err := sdk.AccAddressFromBech32(accountRes.Address)
	s.Require().NoError(err)
	s.groupAccountAddr = addr
	s.Require().NoError(fundAccount(s.bankKeeper, s.sdkCtx, s.groupAccountAddr, sdk.Coins{sdk.NewInt64Coin("test", 10000)}))
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.fixture.Teardown()
}

// TODO: https://github.com/cosmos/cosmos-sdk/issues/9346
func fundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}
	return bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}

func (s *IntegrationTestSuite) TestCreateGroup() {
	members := []group.Member{{
		Address:  s.addr5.String(),
		Weight:   "1",
		Metadata: nil,
	}, {
		Address:  s.addr6.String(),
		Weight:   "2",
		Metadata: nil,
	}}

	expGroups := []*group.GroupInfo{
		{
			GroupId:     s.groupID,
			Version:     1,
			Admin:       s.addr1.String(),
			TotalWeight: "1",
			Metadata:    nil,
		},
		{
			GroupId:     2,
			Version:     1,
			Admin:       s.addr1.String(),
			TotalWeight: "3",
			Metadata:    nil,
		},
	}

	specs := map[string]struct {
		req       *group.MsgCreateGroupRequest
		expErr    bool
		expGroups []*group.GroupInfo
	}{
		"all good": {
			req: &group.MsgCreateGroupRequest{
				Admin:    s.addr1.String(),
				Members:  members,
				Metadata: nil,
			},
			expGroups: expGroups,
		},
		"group metadata too long": {
			req: &group.MsgCreateGroupRequest{
				Admin:    s.addr1.String(),
				Members:  members,
				Metadata: bytes.Repeat([]byte{1}, 256),
			},
			expErr: true,
		},
		"member metadata too long": {
			req: &group.MsgCreateGroupRequest{
				Admin: s.addr1.String(),
				Members: []group.Member{{
					Address:  s.addr3.String(),
					Weight:   "1",
					Metadata: bytes.Repeat([]byte{1}, 256),
				}},
				Metadata: nil,
			},
			expErr: true,
		},
		"zero member weight": {
			req: &group.MsgCreateGroupRequest{
				Admin: s.addr1.String(),
				Members: []group.Member{{
					Address:  s.addr3.String(),
					Weight:   "0",
					Metadata: nil,
				}},
				Metadata: nil,
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
				_, err := s.queryClient.GroupInfo(s.ctx, &group.QueryGroupInfoRequest{GroupId: uint64(seq + 1)})
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			id := res.GroupId

			seq++
			s.Assert().Equal(uint64(seq), id)

			// then all data persisted
			loadedGroupRes, err := s.queryClient.GroupInfo(s.ctx, &group.QueryGroupInfoRequest{GroupId: id})
			s.Require().NoError(err)
			s.Assert().Equal(spec.req.Admin, loadedGroupRes.Info.Admin)
			s.Assert().Equal(spec.req.Metadata, loadedGroupRes.Info.Metadata)
			s.Assert().Equal(id, loadedGroupRes.Info.GroupId)
			s.Assert().Equal(uint64(1), loadedGroupRes.Info.Version)

			// and members are stored as well
			membersRes, err := s.queryClient.GroupMembers(s.ctx, &group.QueryGroupMembersRequest{GroupId: id})
			s.Require().NoError(err)
			loadedMembers := membersRes.Members
			s.Require().Equal(len(members), len(loadedMembers))
			// we reorder members by address to be able to compare them
			sort.Slice(members, func(i, j int) bool { return members[i].Address < members[j].Address })
			for i := range loadedMembers {
				s.Assert().Equal(members[i].Metadata, loadedMembers[i].Member.Metadata)
				s.Assert().Equal(members[i].Address, loadedMembers[i].Member.Address)
				s.Assert().Equal(members[i].Weight, loadedMembers[i].Member.Weight)
				s.Assert().Equal(id, loadedMembers[i].GroupId)
			}

			// query groups by admin
			groupsRes, err := s.queryClient.GroupsByAdmin(s.ctx, &group.QueryGroupsByAdminRequest{Admin: s.addr1.String()})
			s.Require().NoError(err)
			loadedGroups := groupsRes.Groups
			s.Require().Equal(len(spec.expGroups), len(loadedGroups))
			for i := range loadedGroups {
				s.Assert().Equal(spec.expGroups[i].Metadata, loadedGroups[i].Metadata)
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
		Address:  s.addr1.String(),
		Weight:   "1",
		Metadata: nil,
	}}
	oldAdmin := s.addr2.String()
	newAdmin := s.addr3.String()
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:    oldAdmin,
		Members:  members,
		Metadata: nil,
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
				NewAdmin: newAdmin,
			},
			expStored: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       newAdmin,
				Metadata:    nil,
				TotalWeight: "1",
				Version:     2,
			},
		},
		"with wrong admin": {
			req: &group.MsgUpdateGroupAdminRequest{
				GroupId:  groupID,
				Admin:    s.addr4.String(),
				NewAdmin: newAdmin,
			},
			expErr: true,
			expStored: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Metadata:    nil,
				TotalWeight: "1",
				Version:     1,
			},
		},
		"with unknown groupID": {
			req: &group.MsgUpdateGroupAdminRequest{
				GroupId:  999,
				Admin:    oldAdmin,
				NewAdmin: newAdmin,
			},
			expErr: true,
			expStored: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Metadata:    nil,
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

func (s *IntegrationTestSuite) TestUpdateGroupMetadata() {
	oldAdmin := s.addr1.String()
	groupID := s.groupID

	specs := map[string]struct {
		req       *group.MsgUpdateGroupMetadataRequest
		expErr    bool
		expStored *group.GroupInfo
	}{
		"with correct admin": {
			req: &group.MsgUpdateGroupMetadataRequest{
				GroupId:  groupID,
				Admin:    oldAdmin,
				Metadata: []byte{1, 2, 3},
			},
			expStored: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Metadata:    []byte{1, 2, 3},
				TotalWeight: "1",
				Version:     2,
			},
		},
		"with wrong admin": {
			req: &group.MsgUpdateGroupMetadataRequest{
				GroupId:  groupID,
				Admin:    s.addr3.String(),
				Metadata: []byte{1, 2, 3},
			},
			expErr: true,
			expStored: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Metadata:    nil,
				TotalWeight: "1",
				Version:     1,
			},
		},
		"with unknown groupid": {
			req: &group.MsgUpdateGroupMetadataRequest{
				GroupId:  999,
				Admin:    oldAdmin,
				Metadata: []byte{1, 2, 3},
			},
			expErr: true,
			expStored: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       oldAdmin,
				Metadata:    nil,
				TotalWeight: "1",
				Version:     1,
			},
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			sdkCtx, _ := s.sdkCtx.CacheContext()
			ctx := types.Context{Context: sdkCtx}
			_, err := s.msgClient.UpdateGroupMetadata(ctx, spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)

			// then
			res, err := s.queryClient.GroupInfo(ctx, &group.QueryGroupInfoRequest{GroupId: groupID})
			s.Require().NoError(err)
			s.Assert().Equal(spec.expStored, res.Info)
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateGroupMembers() {
	member1 := s.addr5.String()
	member2 := s.addr6.String()
	members := []group.Member{{
		Address:  member1,
		Weight:   "1",
		Metadata: nil,
	}}

	myAdmin := s.addr4.String()
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:    myAdmin,
		Members:  members,
		Metadata: nil,
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
					Address:  member2,
					Weight:   "2",
					Metadata: nil,
				}},
			},
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Metadata:    nil,
				TotalWeight: "3",
				Version:     2,
			},
			expMembers: []*group.GroupMember{
				{
					Member: &group.Member{
						Address:  member2,
						Weight:   "2",
						Metadata: nil,
					},
					GroupId: groupID,
				},
				{
					Member: &group.Member{
						Address:  member1,
						Weight:   "1",
						Metadata: nil,
					},
					GroupId: groupID,
				},
			},
		},
		"update member": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{{
					Address:  member1,
					Weight:   "2",
					Metadata: []byte{1, 2, 3},
				}},
			},
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Metadata:    nil,
				TotalWeight: "2",
				Version:     2,
			},
			expMembers: []*group.GroupMember{
				{
					GroupId: groupID,
					Member: &group.Member{
						Address:  member1,
						Weight:   "2",
						Metadata: []byte{1, 2, 3},
					},
				},
			},
		},
		"update member with same data": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{{
					Address: member1,
					Weight:  "1",
				}},
			},
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Metadata:    nil,
				TotalWeight: "1",
				Version:     2,
			},
			expMembers: []*group.GroupMember{
				{
					GroupId: groupID,
					Member: &group.Member{
						Address: member1,
						Weight:  "1",
					},
				},
			},
		},
		"replace member": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{
					{
						Address:  member1,
						Weight:   "0",
						Metadata: nil,
					},
					{
						Address:  member2,
						Weight:   "1",
						Metadata: nil,
					},
				},
			},
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Metadata:    nil,
				TotalWeight: "1",
				Version:     2,
			},
			expMembers: []*group.GroupMember{{
				GroupId: groupID,
				Member: &group.Member{
					Address:  member2,
					Weight:   "1",
					Metadata: nil,
				},
			}},
		},
		"remove existing member": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{{
					Address:  member1,
					Weight:   "0",
					Metadata: nil,
				}},
			},
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Metadata:    nil,
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
					Address:  s.addr4.String(),
					Weight:   "0",
					Metadata: nil,
				}},
			},
			expErr: true,
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Metadata:    nil,
				TotalWeight: "1",
				Version:     1,
			},
			expMembers: []*group.GroupMember{{
				GroupId: groupID,
				Member: &group.Member{
					Address:  member1,
					Weight:   "1",
					Metadata: nil,
				},
			}},
		},
		"with wrong admin": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: groupID,
				Admin:   s.addr3.String(),
				MemberUpdates: []group.Member{{
					Address:  member1,
					Weight:   "2",
					Metadata: nil,
				}},
			},
			expErr: true,
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Metadata:    nil,
				TotalWeight: "1",
				Version:     1,
			},
			expMembers: []*group.GroupMember{{
				GroupId: groupID,
				Member: &group.Member{
					Address: member1,
					Weight:  "1",
				},
			}},
		},
		"with unknown groupID": {
			req: &group.MsgUpdateGroupMembersRequest{
				GroupId: 999,
				Admin:   myAdmin,
				MemberUpdates: []group.Member{{
					Address:  member1,
					Weight:   "2",
					Metadata: nil,
				}},
			},
			expErr: true,
			expGroup: &group.GroupInfo{
				GroupId:     groupID,
				Admin:       myAdmin,
				Metadata:    nil,
				TotalWeight: "1",
				Version:     1,
			},
			expMembers: []*group.GroupMember{{
				GroupId: groupID,
				Member: &group.Member{
					Address: member1,
					Weight:  "1",
				},
			}},
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			sdkCtx, _ := s.sdkCtx.CacheContext()
			ctx := types.Context{Context: sdkCtx}
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
			s.Require().Equal(len(spec.expMembers), len(loadedMembers))
			// we reorder group members by address to be able to compare them
			sort.Slice(spec.expMembers, func(i, j int) bool {
				return spec.expMembers[i].Member.Address < spec.expMembers[j].Member.Address
			})
			for i := range loadedMembers {
				s.Assert().Equal(spec.expMembers[i].Member.Metadata, loadedMembers[i].Member.Metadata)
				s.Assert().Equal(spec.expMembers[i].Member.Address, loadedMembers[i].Member.Address)
				s.Assert().Equal(spec.expMembers[i].Member.Weight, loadedMembers[i].Member.Weight)
				s.Assert().Equal(spec.expMembers[i].GroupId, loadedMembers[i].GroupId)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCreateGroupAccount() {
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:    s.addr1.String(),
		Members:  nil,
		Metadata: nil,
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
				Admin:    s.addr1.String(),
				Metadata: nil,
				GroupId:  myGroupID,
			},
			policy: group.NewThresholdDecisionPolicy(
				"1",
				gogotypes.Duration{Seconds: 1},
			),
		},
		"decision policy threshold > total group weight": {
			req: &group.MsgCreateGroupAccountRequest{
				Admin:    s.addr1.String(),
				Metadata: nil,
				GroupId:  myGroupID,
			},
			policy: group.NewThresholdDecisionPolicy(
				"10",
				gogotypes.Duration{Seconds: 1},
			),
		},
		"group id does not exists": {
			req: &group.MsgCreateGroupAccountRequest{
				Admin:    s.addr1.String(),
				Metadata: nil,
				GroupId:  9999,
			},
			policy: group.NewThresholdDecisionPolicy(
				"1",
				gogotypes.Duration{Seconds: 1},
			),
			expErr: true,
		},
		"admin not group admin": {
			req: &group.MsgCreateGroupAccountRequest{
				Admin:    s.addr4.String(),
				Metadata: nil,
				GroupId:  myGroupID,
			},
			policy: group.NewThresholdDecisionPolicy(
				"1",
				gogotypes.Duration{Seconds: 1},
			),
			expErr: true,
		},
		"metadata too long": {
			req: &group.MsgCreateGroupAccountRequest{
				Admin:    s.addr1.String(),
				Metadata: []byte(strings.Repeat("a", 256)),
				GroupId:  myGroupID,
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
			addr := res.Address

			// then all data persisted
			groupAccountRes, err := s.queryClient.GroupAccountInfo(s.ctx, &group.QueryGroupAccountInfoRequest{Address: addr})
			s.Require().NoError(err)

			groupAccount := groupAccountRes.Info
			s.Assert().Equal(addr, groupAccount.Address)
			s.Assert().Equal(myGroupID, groupAccount.GroupId)
			s.Assert().Equal(spec.req.Admin, groupAccount.Admin)
			s.Assert().Equal(spec.req.Metadata, groupAccount.Metadata)
			s.Assert().Equal(uint64(1), groupAccount.Version)
			s.Assert().Equal(spec.policy.(*group.ThresholdDecisionPolicy), groupAccount.GetDecisionPolicy())
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateGroupAccountAdmin() {
	admin, newAdmin := s.addr1, s.addr2
	groupAccountAddr, myGroupID, policy := createGroupAndGroupAccount(admin, s)

	specs := map[string]struct {
		req             *group.MsgUpdateGroupAccountAdminRequest
		expGroupAccount *group.GroupAccountInfo
		expErr          bool
	}{
		"with wrong admin": {
			req: &group.MsgUpdateGroupAccountAdminRequest{
				Admin:    s.addr5.String(),
				Address:  groupAccountAddr,
				NewAdmin: newAdmin.String(),
			},
			expGroupAccount: &group.GroupAccountInfo{
				Admin:          admin.String(),
				Address:        groupAccountAddr,
				GroupId:        myGroupID,
				Metadata:       nil,
				Version:        2,
				DecisionPolicy: nil,
			},
			expErr: true,
		},
		"with wrong group account": {
			req: &group.MsgUpdateGroupAccountAdminRequest{
				Admin:    admin.String(),
				Address:  s.addr5.String(),
				NewAdmin: newAdmin.String(),
			},
			expGroupAccount: &group.GroupAccountInfo{
				Admin:          admin.String(),
				Address:        groupAccountAddr,
				GroupId:        myGroupID,
				Metadata:       nil,
				Version:        2,
				DecisionPolicy: nil,
			},
			expErr: true,
		},
		"correct data": {
			req: &group.MsgUpdateGroupAccountAdminRequest{
				Admin:    admin.String(),
				Address:  groupAccountAddr,
				NewAdmin: newAdmin.String(),
			},
			expGroupAccount: &group.GroupAccountInfo{
				Admin:          newAdmin.String(),
				Address:        groupAccountAddr,
				GroupId:        myGroupID,
				Metadata:       nil,
				Version:        2,
				DecisionPolicy: nil,
			},
			expErr: false,
		},
	}

	for msg, spec := range specs {
		spec := spec
		err := spec.expGroupAccount.SetDecisionPolicy(policy)
		s.Require().NoError(err)

		s.Run(msg, func() {
			_, err := s.msgClient.UpdateGroupAccountAdmin(s.ctx, spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			res, err := s.queryClient.GroupAccountInfo(s.ctx, &group.QueryGroupAccountInfoRequest{
				Address: groupAccountAddr,
			})
			s.Require().NoError(err)
			s.Assert().Equal(spec.expGroupAccount, res.Info)
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateGroupAccountMetadata() {
	admin := s.addr1
	groupAccountAddr, myGroupID, policy := createGroupAndGroupAccount(admin, s)

	specs := map[string]struct {
		req             *group.MsgUpdateGroupAccountMetadataRequest
		expGroupAccount *group.GroupAccountInfo
		expErr          bool
	}{
		"with wrong admin": {
			req: &group.MsgUpdateGroupAccountMetadataRequest{
				Admin:    s.addr5.String(),
				Address:  groupAccountAddr,
				Metadata: []byte("hello"),
			},
			expGroupAccount: &group.GroupAccountInfo{},
			expErr:          true,
		},
		"with wrong group account": {
			req: &group.MsgUpdateGroupAccountMetadataRequest{
				Admin:    admin.String(),
				Address:  s.addr5.String(),
				Metadata: []byte("hello"),
			},
			expGroupAccount: &group.GroupAccountInfo{},
			expErr:          true,
		},
		"with comment too long": {
			req: &group.MsgUpdateGroupAccountMetadataRequest{
				Admin:    admin.String(),
				Address:  s.addr5.String(),
				Metadata: []byte(strings.Repeat("a", 256)),
			},
			expGroupAccount: &group.GroupAccountInfo{},
			expErr:          true,
		},
		"correct data": {
			req: &group.MsgUpdateGroupAccountMetadataRequest{
				Admin:    admin.String(),
				Address:  groupAccountAddr,
				Metadata: []byte("hello"),
			},
			expGroupAccount: &group.GroupAccountInfo{
				Admin:          admin.String(),
				Address:        groupAccountAddr,
				GroupId:        myGroupID,
				Metadata:       []byte("hello"),
				Version:        2,
				DecisionPolicy: nil,
			},
			expErr: false,
		},
	}

	for msg, spec := range specs {
		spec := spec
		err := spec.expGroupAccount.SetDecisionPolicy(policy)
		s.Require().NoError(err)

		s.Run(msg, func() {
			_, err := s.msgClient.UpdateGroupAccountMetadata(s.ctx, spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			res, err := s.queryClient.GroupAccountInfo(s.ctx, &group.QueryGroupAccountInfoRequest{
				Address: groupAccountAddr,
			})
			s.Require().NoError(err)
			s.Assert().Equal(spec.expGroupAccount, res.Info)
		})
	}
}

func (s *IntegrationTestSuite) TestUpdateGroupAccountDecisionPolicy() {
	admin := s.addr1
	groupAccountAddr, myGroupID, policy := createGroupAndGroupAccount(admin, s)

	specs := map[string]struct {
		req             *group.MsgUpdateGroupAccountDecisionPolicyRequest
		policy          group.DecisionPolicy
		expGroupAccount *group.GroupAccountInfo
		expErr          bool
	}{
		"with wrong admin": {
			req: &group.MsgUpdateGroupAccountDecisionPolicyRequest{
				Admin:   s.addr5.String(),
				Address: groupAccountAddr,
			},
			policy:          policy,
			expGroupAccount: &group.GroupAccountInfo{},
			expErr:          true,
		},
		"with wrong group account": {
			req: &group.MsgUpdateGroupAccountDecisionPolicyRequest{
				Admin:   admin.String(),
				Address: s.addr5.String(),
			},
			policy:          policy,
			expGroupAccount: &group.GroupAccountInfo{},
			expErr:          true,
		},
		"correct data": {
			req: &group.MsgUpdateGroupAccountDecisionPolicyRequest{
				Admin:   admin.String(),
				Address: groupAccountAddr,
			},
			policy: group.NewThresholdDecisionPolicy(
				"2",
				gogotypes.Duration{Seconds: 2},
			),
			expGroupAccount: &group.GroupAccountInfo{
				Admin:          admin.String(),
				Address:        groupAccountAddr,
				GroupId:        myGroupID,
				Metadata:       nil,
				Version:        2,
				DecisionPolicy: nil,
			},
			expErr: false,
		},
	}

	for msg, spec := range specs {
		spec := spec
		err := spec.expGroupAccount.SetDecisionPolicy(spec.policy)
		s.Require().NoError(err)

		err = spec.req.SetDecisionPolicy(spec.policy)
		s.Require().NoError(err)

		s.Run(msg, func() {
			_, err := s.msgClient.UpdateGroupAccountDecisionPolicy(s.ctx, spec.req)
			if spec.expErr {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
			res, err := s.queryClient.GroupAccountInfo(s.ctx, &group.QueryGroupAccountInfoRequest{
				Address: groupAccountAddr,
			})
			s.Require().NoError(err)
			s.Assert().Equal(spec.expGroupAccount, res.Info)
		})
	}
}

func (s *IntegrationTestSuite) TestGroupAccountsByAdminOrGroup() {
	admin := s.addr2
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:    admin.String(),
		Members:  nil,
		Metadata: nil,
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
	expectAccs := make([]*group.GroupAccountInfo, count)
	for i := range expectAccs {
		req := &group.MsgCreateGroupAccountRequest{
			Admin:    admin.String(),
			Metadata: nil,
			GroupId:  myGroupID,
		}
		err := req.SetDecisionPolicy(policies[i])
		s.Require().NoError(err)
		res, err := s.msgClient.CreateGroupAccount(s.ctx, req)
		s.Require().NoError(err)

		expectAcc := &group.GroupAccountInfo{
			Address:  res.Address,
			Admin:    admin.String(),
			Metadata: nil,
			GroupId:  myGroupID,
			Version:  uint64(1),
		}
		err = expectAcc.SetDecisionPolicy(policies[i])
		s.Require().NoError(err)
		expectAccs[i] = expectAcc
	}
	sort.Slice(expectAccs, func(i, j int) bool { return expectAccs[i].Address < expectAccs[j].Address })

	// query group account by group
	accountsByGroupRes, err := s.queryClient.GroupAccountsByGroup(s.ctx, &group.QueryGroupAccountsByGroupRequest{
		GroupId: myGroupID,
	})
	s.Require().NoError(err)
	accounts := accountsByGroupRes.GroupAccounts
	s.Require().Equal(len(accounts), count)
	// we reorder accounts by address to be able to compare them
	sort.Slice(accounts, func(i, j int) bool { return accounts[i].Address < accounts[j].Address })
	for i := range accounts {
		s.Assert().Equal(accounts[i].Address, expectAccs[i].Address)
		s.Assert().Equal(accounts[i].GroupId, expectAccs[i].GroupId)
		s.Assert().Equal(accounts[i].Admin, expectAccs[i].Admin)
		s.Assert().Equal(accounts[i].Metadata, expectAccs[i].Metadata)
		s.Assert().Equal(accounts[i].Version, expectAccs[i].Version)
		s.Assert().Equal(accounts[i].GetDecisionPolicy(), expectAccs[i].GetDecisionPolicy())
	}

	// query group account by admin
	accountsByAdminRes, err := s.queryClient.GroupAccountsByAdmin(s.ctx, &group.QueryGroupAccountsByAdminRequest{
		Admin: admin.String(),
	})
	s.Require().NoError(err)
	accounts = accountsByAdminRes.GroupAccounts
	s.Require().Equal(len(accounts), count)
	// we reorder accounts by address to be able to compare them
	sort.Slice(accounts, func(i, j int) bool { return accounts[i].Address < accounts[j].Address })
	for i := range accounts {
		s.Assert().Equal(accounts[i].Address, expectAccs[i].Address)
		s.Assert().Equal(accounts[i].GroupId, expectAccs[i].GroupId)
		s.Assert().Equal(accounts[i].Admin, expectAccs[i].Admin)
		s.Assert().Equal(accounts[i].Metadata, expectAccs[i].Metadata)
		s.Assert().Equal(accounts[i].Version, expectAccs[i].Version)
		s.Assert().Equal(accounts[i].GetDecisionPolicy(), expectAccs[i].GetDecisionPolicy())
	}
}

func (s *IntegrationTestSuite) TestCreateProposal() {
	myGroupID := s.groupID

	accountReq := &group.MsgCreateGroupAccountRequest{
		Admin:    s.addr1.String(),
		GroupId:  myGroupID,
		Metadata: nil,
	}
	accountAddr := s.groupAccountAddr

	policy := group.NewThresholdDecisionPolicy(
		"100",
		gogotypes.Duration{Seconds: 1},
	)
	err := accountReq.SetDecisionPolicy(policy)
	s.Require().NoError(err)
	bigThresholdRes, err := s.msgClient.CreateGroupAccount(s.ctx, accountReq)
	s.Require().NoError(err)
	bigThresholdAddr := bigThresholdRes.Address

	specs := map[string]struct {
		req    *group.MsgCreateProposalRequest
		msgs   []sdk.Msg
		expErr bool
	}{
		"all good with minimal fields set": {
			req: &group.MsgCreateProposalRequest{
				Address:   accountAddr.String(),
				Proposers: []string{s.addr2.String()},
			},
		},
		"all good with good msg payload": {
			req: &group.MsgCreateProposalRequest{
				Address:   accountAddr.String(),
				Proposers: []string{s.addr2.String()},
			},
			msgs: []sdk.Msg{&banktypes.MsgSend{
				FromAddress: accountAddr.String(),
				ToAddress:   s.addr2.String(),
				Amount:      sdk.Coins{sdk.NewInt64Coin("token", 100)},
			}},
		},
		"metadata too long": {
			req: &group.MsgCreateProposalRequest{
				Address:   accountAddr.String(),
				Metadata:  bytes.Repeat([]byte{1}, 256),
				Proposers: []string{s.addr2.String()},
			},
			expErr: true,
		},
		"group account required": {
			req: &group.MsgCreateProposalRequest{
				Metadata:  nil,
				Proposers: []string{s.addr2.String()},
			},
			expErr: true,
		},
		"existing group account required": {
			req: &group.MsgCreateProposalRequest{
				Address:   s.addr1.String(),
				Proposers: []string{s.addr2.String()},
			},
			expErr: true,
		},
		"impossible case: decision policy threshold > total group weight": {
			req: &group.MsgCreateProposalRequest{
				Address:   bigThresholdAddr,
				Proposers: []string{s.addr2.String()},
			},
			expErr: true,
		},
		"only group members can create a proposal": {
			req: &group.MsgCreateProposalRequest{
				Address:   accountAddr.String(),
				Proposers: []string{s.addr3.String()},
			},
			expErr: true,
		},
		"all proposers must be in group": {
			req: &group.MsgCreateProposalRequest{
				Address:   accountAddr.String(),
				Proposers: []string{s.addr2.String(), s.addr4.String()},
			},
			expErr: true,
		},
		"proposers must not be empty": {
			req: &group.MsgCreateProposalRequest{
				Address:   accountAddr.String(),
				Proposers: []string{s.addr2.String(), ""},
			},
			expErr: true,
		},
		"admin that is not a group member can not create proposal": {
			req: &group.MsgCreateProposalRequest{
				Address:   accountAddr.String(),
				Metadata:  nil,
				Proposers: []string{s.addr1.String()},
			},
			expErr: true,
		},
		"reject msgs that are not authz by group account": {
			req: &group.MsgCreateProposalRequest{
				Address:   accountAddr.String(),
				Metadata:  nil,
				Proposers: []string{s.addr2.String()},
			},
			msgs:   []sdk.Msg{&testdata.MsgAuthenticated{Signers: []sdk.AccAddress{s.addr1}}},
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

			s.Assert().Equal(accountAddr.String(), proposal.Address)
			s.Assert().Equal(spec.req.Metadata, proposal.Metadata)
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
		{Address: s.addr2.String(), Weight: "1"},
		{Address: s.addr3.String(), Weight: "2"},
	}
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:    s.addr1.String(),
		Members:  members,
		Metadata: nil,
	})
	s.Require().NoError(err)
	myGroupID := groupRes.GroupId

	policy := group.NewThresholdDecisionPolicy(
		"2",
		gogotypes.Duration{Seconds: 1},
	)
	accountReq := &group.MsgCreateGroupAccountRequest{
		Admin:    s.addr1.String(),
		GroupId:  myGroupID,
		Metadata: nil,
	}
	err = accountReq.SetDecisionPolicy(policy)
	s.Require().NoError(err)
	accountRes, err := s.msgClient.CreateGroupAccount(s.ctx, accountReq)
	s.Require().NoError(err)
	accountAddr := accountRes.Address
	groupAccount, err := sdk.AccAddressFromBech32(accountAddr)
	s.Require().NoError(err)
	s.Require().NotNil(groupAccount)

	req := &group.MsgCreateProposalRequest{
		Address:   accountAddr,
		Metadata:  nil,
		Proposers: []string{s.addr2.String()},
		Msgs:      nil,
	}
	proposalRes, err := s.msgClient.CreateProposal(s.ctx, req)
	s.Require().NoError(err)
	myProposalID := proposalRes.ProposalId

	// proposals by group account
	proposalsRes, err := s.queryClient.ProposalsByGroupAccount(s.ctx, &group.QueryProposalsByGroupAccountRequest{
		Address: accountAddr,
	})
	s.Require().NoError(err)
	proposals := proposalsRes.Proposals
	s.Require().Equal(len(proposals), 1)
	s.Assert().Equal(req.Address, proposals[0].Address)
	s.Assert().Equal(req.Metadata, proposals[0].Metadata)
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
				Voter:      s.addr2.String(),
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
				Voter:      s.addr2.String(),
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
				Voter:      s.addr2.String(),
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
				Voter:      s.addr2.String(),
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
				Voter:      s.addr3.String(),
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
				Voter:      s.addr2.String(),
				Choice:     group.Choice_CHOICE_YES,
			},
			doBefore: func(ctx context.Context) {
				_, err := s.msgClient.Vote(ctx, &group.MsgVoteRequest{
					ProposalId: myProposalID,
					Voter:      s.addr3.String(),
					Choice:     group.Choice_CHOICE_VETO,
				})
				s.Require().NoError(err)
			},
			expErr: true,
		},
		"metadata too long": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voter:      s.addr2.String(),
				Metadata:   bytes.Repeat([]byte{1}, 256),
				Choice:     group.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"existing proposal required": {
			req: &group.MsgVoteRequest{
				ProposalId: 999,
				Voter:      s.addr2.String(),
				Choice:     group.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"empty choice": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voter:      s.addr2.String(),
			},
			expErr: true,
		},
		"invalid choice": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voter:      s.addr2.String(),
				Choice:     5,
			},
			expErr: true,
		},
		"voter must be in group": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voter:      s.addr4.String(),
				Choice:     group.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"voter must not be empty": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voter:      "",
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
		"admin that is not a group member can not vote": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voter:      s.addr1.String(),
				Choice:     group.Choice_CHOICE_NO,
			},
			expErr: true,
		},
		"on timeout": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voter:      s.addr2.String(),
				Choice:     group.Choice_CHOICE_NO,
			},
			srcCtx: s.sdkCtx.WithBlockTime(s.blockTime.Add(time.Second)),
			expErr: true,
		},
		"closed already": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voter:      s.addr2.String(),
				Choice:     group.Choice_CHOICE_NO,
			},
			doBefore: func(ctx context.Context) {
				_, err := s.msgClient.Vote(ctx, &group.MsgVoteRequest{
					ProposalId: myProposalID,
					Voter:      s.addr3.String(),
					Choice:     group.Choice_CHOICE_YES,
				})
				s.Require().NoError(err)
			},
			expErr: true,
		},
		"voted already": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voter:      s.addr2.String(),
				Choice:     group.Choice_CHOICE_NO,
			},
			doBefore: func(ctx context.Context) {
				_, err := s.msgClient.Vote(ctx, &group.MsgVoteRequest{
					ProposalId: myProposalID,
					Voter:      s.addr2.String(),
					Choice:     group.Choice_CHOICE_YES,
				})
				s.Require().NoError(err)
			},
			expErr: true,
		},
		"with group modified": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voter:      s.addr2.String(),
				Choice:     group.Choice_CHOICE_NO,
			},
			doBefore: func(ctx context.Context) {
				_, err = s.msgClient.UpdateGroupMetadata(ctx, &group.MsgUpdateGroupMetadataRequest{
					GroupId:  myGroupID,
					Admin:    s.addr1.String(),
					Metadata: []byte{1, 2, 3},
				})
				s.Require().NoError(err)
			},
			expErr: true,
		},
		"with policy modified": {
			req: &group.MsgVoteRequest{
				ProposalId: myProposalID,
				Voter:      s.addr2.String(),
				Choice:     group.Choice_CHOICE_NO,
			},
			doBefore: func(ctx context.Context) {
				m, err := group.NewMsgUpdateGroupAccountDecisionPolicyRequest(
					s.addr1,
					groupAccount,
					&group.ThresholdDecisionPolicy{
						Threshold: "1",
						Timeout:   gogotypes.Duration{Seconds: 1},
					},
				)
				s.Require().NoError(err)

				_, err = s.msgClient.UpdateGroupAccountDecisionPolicy(ctx, m)
				s.Require().NoError(err)
			},
			expErr: true,
		},
	}
	for msg, spec := range specs {
		spec := spec
		s.Run(msg, func() {
			sdkCtx := s.sdkCtx
			if !spec.srcCtx.IsZero() {
				sdkCtx = spec.srcCtx
			}
			sdkCtx, _ = sdkCtx.CacheContext()
			ctx := types.Context{Context: sdkCtx}

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
			// vote is stored and all data persisted
			res, err := s.queryClient.VoteByProposalVoter(ctx, &group.QueryVoteByProposalVoterRequest{
				ProposalId: spec.req.ProposalId,
				Voter:      spec.req.Voter,
			})
			s.Require().NoError(err)
			loaded := res.Vote
			s.Assert().Equal(spec.req.ProposalId, loaded.ProposalId)
			s.Assert().Equal(spec.req.Voter, loaded.Voter)
			s.Assert().Equal(spec.req.Choice, loaded.Choice)
			s.Assert().Equal(spec.req.Metadata, loaded.Metadata)
			submittedAt, err := gogotypes.TimestampFromProto(&loaded.SubmittedAt)
			s.Require().NoError(err)
			s.Assert().Equal(s.blockTime, submittedAt)

			// query votes by proposal
			votesByProposalRes, err := s.queryClient.VotesByProposal(ctx, &group.QueryVotesByProposalRequest{
				ProposalId: spec.req.ProposalId,
			})
			s.Require().NoError(err)
			votesByProposal := votesByProposalRes.Votes
			s.Require().Equal(1, len(votesByProposal))
			vote := votesByProposal[0]
			s.Assert().Equal(spec.req.ProposalId, vote.ProposalId)
			s.Assert().Equal(spec.req.Voter, vote.Voter)
			s.Assert().Equal(spec.req.Choice, vote.Choice)
			s.Assert().Equal(spec.req.Metadata, vote.Metadata)
			submittedAt, err = gogotypes.TimestampFromProto(&vote.SubmittedAt)
			s.Require().NoError(err)
			s.Assert().Equal(s.blockTime, submittedAt)

			// query votes by voter
			voter := spec.req.Voter
			votesByVoterRes, err := s.queryClient.VotesByVoter(ctx, &group.QueryVotesByVoterRequest{
				Voter: voter,
			})
			s.Require().NoError(err)
			votesByVoter := votesByVoterRes.Votes
			s.Require().Equal(1, len(votesByVoter))
			s.Assert().Equal(spec.req.ProposalId, votesByVoter[0].ProposalId)
			s.Assert().Equal(voter, votesByVoter[0].Voter)
			s.Assert().Equal(spec.req.Choice, votesByVoter[0].Choice)
			s.Assert().Equal(spec.req.Metadata, votesByVoter[0].Metadata)
			submittedAt, err = gogotypes.TimestampFromProto(&votesByVoter[0].SubmittedAt)
			s.Require().NoError(err)
			s.Assert().Equal(s.blockTime, submittedAt)

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
				router = baseapp.NewRouter().AddRoute(sdk.NewRoute(banktypes.ModuleName, bank.NewHandler(s.bankKeeper)))
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
	proposers := []string{s.addr2.String()}

	specs := map[string]struct {
		srcBlockTime      time.Time
		setupProposal     func(ctx context.Context) uint64
		expErr            bool
		expProposalStatus group.Proposal_Status
		expProposalResult group.Proposal_Result
		expExecutorResult group.Proposal_ExecutorResult
		expFromBalances   sdk.Coins
		expToBalances     sdk.Coins
	}{
		"proposal executed when accepted": {
			setupProposal: func(ctx context.Context) uint64 {
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
			setupProposal: func(ctx context.Context) uint64 {
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
			setupProposal: func(ctx context.Context) uint64 {
				msgs := []sdk.Msg{msgSend}
				return createProposalAndVote(ctx, s, msgs, proposers, group.Choice_CHOICE_NO)
			},
			expProposalStatus: group.ProposalStatusClosed,
			expProposalResult: group.ProposalResultRejected,
			expExecutorResult: group.ProposalExecutorResultNotRun,
		},
		"open proposal must not fail": {
			setupProposal: func(ctx context.Context) uint64 {
				return createProposal(ctx, s, []sdk.Msg{msgSend}, proposers)
			},
			expProposalStatus: group.ProposalStatusSubmitted,
			expProposalResult: group.ProposalResultUnfinalized,
			expExecutorResult: group.ProposalExecutorResultNotRun,
		},
		"existing proposal required": {
			setupProposal: func(ctx context.Context) uint64 {
				return 9999
			},
			expErr: true,
		},
		"Decision policy also applied on timeout": {
			setupProposal: func(ctx context.Context) uint64 {
				msgs := []sdk.Msg{msgSend}
				return createProposalAndVote(ctx, s, msgs, proposers, group.Choice_CHOICE_NO)
			},
			srcBlockTime:      s.blockTime.Add(time.Second),
			expProposalStatus: group.ProposalStatusClosed,
			expProposalResult: group.ProposalResultRejected,
			expExecutorResult: group.ProposalExecutorResultNotRun,
		},
		"Decision policy also applied after timeout": {
			setupProposal: func(ctx context.Context) uint64 {
				msgs := []sdk.Msg{msgSend}
				return createProposalAndVote(ctx, s, msgs, proposers, group.Choice_CHOICE_NO)
			},
			srcBlockTime:      s.blockTime.Add(time.Second).Add(time.Millisecond),
			expProposalStatus: group.ProposalStatusClosed,
			expProposalResult: group.ProposalResultRejected,
			expExecutorResult: group.ProposalExecutorResultNotRun,
		},
		"with group modified before tally": {
			setupProposal: func(ctx context.Context) uint64 {
				myProposalID := createProposal(ctx, s, []sdk.Msg{msgSend}, proposers)

				// then modify group
				_, err := s.msgClient.UpdateGroupMetadata(ctx, &group.MsgUpdateGroupMetadataRequest{
					Admin:    s.addr1.String(),
					GroupId:  s.groupID,
					Metadata: []byte{1, 2, 3},
				})
				s.Require().NoError(err)
				return myProposalID
			},
			expProposalStatus: group.ProposalStatusAborted,
			expProposalResult: group.ProposalResultUnfinalized,
			expExecutorResult: group.ProposalExecutorResultNotRun,
		},
		"with group account modified before tally": {
			setupProposal: func(ctx context.Context) uint64 {
				myProposalID := createProposal(ctx, s, []sdk.Msg{msgSend}, proposers)
				_, err := s.msgClient.UpdateGroupAccountMetadata(ctx, &group.MsgUpdateGroupAccountMetadataRequest{
					Admin:    s.addr1.String(),
					Address:  s.groupAccountAddr.String(),
					Metadata: []byte("group account modified before tally"),
				})
				s.Require().NoError(err)
				return myProposalID
			},
			expProposalStatus: group.ProposalStatusAborted,
			expProposalResult: group.ProposalResultUnfinalized,
			expExecutorResult: group.ProposalExecutorResultNotRun,
		},
		"prevent double execution when successful": {
			setupProposal: func(ctx context.Context) uint64 {
				myProposalID := createProposalAndVote(ctx, s, []sdk.Msg{msgSend}, proposers, group.Choice_CHOICE_YES)

				_, err := s.msgClient.Exec(ctx, &group.MsgExecRequest{Signer: s.addr1.String(), ProposalId: myProposalID})
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
			setupProposal: func(ctx context.Context) uint64 {
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
			setupProposal: func(ctx context.Context) uint64 {
				msgs := []sdk.Msg{
					&banktypes.MsgSend{
						FromAddress: s.groupAccountAddr.String(),
						ToAddress:   s.addr2.String(),
						Amount:      sdk.Coins{sdk.NewInt64Coin("test", 10001)}},
				}
				myProposalID := createProposalAndVote(ctx, s, msgs, proposers, group.Choice_CHOICE_YES)

				_, err := s.msgClient.Exec(ctx, &group.MsgExecRequest{Signer: s.addr1.String(), ProposalId: myProposalID})
				s.Require().NoError(err)
				s.Require().NoError(fundAccount(s.bankKeeper, ctx.(types.Context).Context, s.groupAccountAddr, sdk.Coins{sdk.NewInt64Coin("test", 10002)}))

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
			ctx := types.Context{Context: sdkCtx}

			proposalID := spec.setupProposal(ctx)

			if !spec.srcBlockTime.IsZero() {
				sdkCtx = sdkCtx.WithBlockTime(spec.srcBlockTime)
				ctx = types.Context{Context: sdkCtx}
			}

			_, err := s.msgClient.Exec(ctx, &group.MsgExecRequest{Signer: s.addr1.String(), ProposalId: proposalID})
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
	proposers []string) uint64 {
	proposalReq := &group.MsgCreateProposalRequest{
		Address:   s.groupAccountAddr.String(),
		Proposers: proposers,
		Metadata:  nil,
	}
	err := proposalReq.SetMsgs(msgs)
	s.Require().NoError(err)

	proposalRes, err := s.msgClient.CreateProposal(ctx, proposalReq)
	s.Require().NoError(err)
	return proposalRes.ProposalId
}

func createProposalAndVote(
	ctx context.Context, s *IntegrationTestSuite, msgs []sdk.Msg,
	proposers []string, choice group.Choice) uint64 {
	s.Require().Greater(len(proposers), 0)
	myProposalID := createProposal(ctx, s, msgs, proposers)

	_, err := s.msgClient.Vote(ctx, &group.MsgVoteRequest{
		ProposalId: myProposalID,
		Voter:      proposers[0],
		Choice:     choice,
	})
	s.Require().NoError(err)
	return myProposalID
}

func createGroupAndGroupAccount(
	admin sdk.AccAddress,
	s *IntegrationTestSuite,
) (string, uint64, group.DecisionPolicy) {
	groupRes, err := s.msgClient.CreateGroup(s.ctx, &group.MsgCreateGroupRequest{
		Admin:    admin.String(),
		Members:  nil,
		Metadata: nil,
	})
	s.Require().NoError(err)

	myGroupID := groupRes.GroupId
	groupAccount := &group.MsgCreateGroupAccountRequest{
		Admin:    admin.String(),
		GroupId:  myGroupID,
		Metadata: nil,
	}

	policy := group.NewThresholdDecisionPolicy(
		"1",
		gogotypes.Duration{Seconds: 1},
	)
	err = groupAccount.SetDecisionPolicy(policy)
	s.Require().NoError(err)

	groupAccountRes, err := s.msgClient.CreateGroupAccount(s.ctx, groupAccount)
	s.Require().NoError(err)

	return groupAccountRes.Address, myGroupID, policy
}
