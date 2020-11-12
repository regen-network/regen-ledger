package server

// import (
// 	"math"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/cosmos/cosmos-sdk/baseapp"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
// 	"github.com/gogo/protobuf/types"
// 	"github.com/regen-network/regen-ledger/app"
// 	"github.com/regen-network/regen-ledger/orm"
// 	testdatagroup "github.com/regen-network/regen-ledger/testutil/testdata/group"
// 	"github.com/regen-network/regen-ledger/x/group"
// 	"github.com/regen-network/regen-ledger/x/group/testutil"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// // TODO: Setup test suite
// func TestCreateGroup(t *testing.T) {
// 	encodingConfig := app.MakeEncodingConfig()

// 	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
// 	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, group.DefaultParamspace)

// 	groupKey := sdk.NewKVStoreKey(group.StoreKey)
// 	k := group.NewGroupKeeper(groupKey, paramSpace, baseapp.NewRouter(), &testutil.MockProposalI{})
// 	ctx := testutil.NewContext(pKey, pTKey, groupKey)
// 	defaultParams := group.DefaultParams()
// 	paramSpace.SetParamSet(ctx, &defaultParams)

// 	members := []group.Member{{
// 		Address: sdk.AccAddress([]byte("one--member--address")),
// 		Power:   sdk.NewDec(1),
// 		Comment: "first",
// 	}, {
// 		Address: sdk.AccAddress([]byte("other-member-address")),
// 		Power:   sdk.NewDec(2),
// 		Comment: "second",
// 	}}
// 	specs := map[string]struct {
// 		srcAdmin   sdk.AccAddress
// 		srcMembers []group.Member
// 		srcComment string
// 		expErr     bool
// 	}{
// 		"all good": {
// 			srcAdmin:   []byte("valid--admin-address"),
// 			srcMembers: members,
// 			srcComment: "test",
// 		},
// 		"group comment too long": {
// 			srcAdmin:   []byte("valid--admin-address"),
// 			srcMembers: members,
// 			srcComment: strings.Repeat("a", 256),
// 			expErr:     true,
// 		},
// 		"member comment too long": {
// 			srcAdmin: []byte("valid--admin-address"),
// 			srcMembers: []group.Member{{
// 				Address: []byte("valid-member-address"),
// 				Power:   sdk.OneDec(),
// 				Comment: strings.Repeat("a", 256),
// 			}},
// 			srcComment: "test",
// 			expErr:     true,
// 		},
// 	}
// 	var seq uint32
// 	for msg, spec := range specs {
// 		t.Run(msg, func(t *testing.T) {
// 			id, err := k.CreateGroup(ctx, spec.srcAdmin, spec.srcMembers, spec.srcComment)
// 			if spec.expErr {
// 				require.Error(t, err)
// 				require.False(t, k.HasGroup(ctx, group.ID(seq+1).Bytes()))
// 				return
// 			}
// 			require.NoError(t, err)

// 			seq++
// 			assert.Equal(t, group.ID(seq), id)

// 			// then all data persisted
// 			loadedGroup, err := k.GetGroup(ctx, id)
// 			require.NoError(t, err)
// 			assert.Equal(t, sdk.AccAddress([]byte(spec.srcAdmin)), loadedGroup.Admin)
// 			assert.Equal(t, spec.srcComment, loadedGroup.Comment)
// 			assert.Equal(t, id, loadedGroup.GroupId)
// 			assert.Equal(t, uint64(1), loadedGroup.Version)

// 			// and members are stored as well
// 			it, err := k.GetGroupMembersByGroup(ctx, id)
// 			require.NoError(t, err)
// 			var loadedMembers []group.GroupMember
// 			_, err = orm.ReadAll(it, &loadedMembers)
// 			require.NoError(t, err)
// 			assert.Equal(t, len(members), len(loadedMembers))
// 			for i := range loadedMembers {
// 				assert.Equal(t, members[i].Comment, loadedMembers[i].Comment)
// 				assert.Equal(t, members[i].Address, loadedMembers[i].Member)
// 				assert.Equal(t, members[i].Power, loadedMembers[i].Weight)
// 				assert.Equal(t, id, loadedMembers[i].GroupId)
// 			}
// 		})
// 	}
// }

// func TestCreateGroupAccount(t *testing.T) {
// 	encodingConfig := app.MakeEncodingConfig()

// 	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
// 	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, group.DefaultParamspace)

// 	groupKey := sdk.NewKVStoreKey(group.StoreKey)
// 	k := group.NewGroupKeeper(groupKey, paramSpace, baseapp.NewRouter(), &testutil.MockProposalI{})
// 	ctx := testutil.NewContext(pKey, pTKey, groupKey)
// 	defaultParams := group.DefaultParams()
// 	paramSpace.SetParamSet(ctx, &defaultParams)

// 	myGroupID, err := k.CreateGroup(ctx, []byte("valid--admin-address"), nil, "test")
// 	require.NoError(t, err)

// 	specs := map[string]struct {
// 		srcAdmin   sdk.AccAddress
// 		srcGroupID group.ID
// 		srcPolicy  group.DecisionPolicy
// 		srcComment string
// 		expErr     bool
// 	}{
// 		"all good": {
// 			srcAdmin:   []byte("valid--admin-address"),
// 			srcComment: "test",
// 			srcGroupID: myGroupID,
// 			srcPolicy: group.NewThresholdDecisionPolicy(
// 				sdk.OneDec(),
// 				types.Duration{Seconds: 1},
// 			),
// 		},
// 		"decision policy threshold > total group weight": {
// 			srcAdmin:   []byte("valid--admin-address"),
// 			srcComment: "test",
// 			srcGroupID: myGroupID,
// 			srcPolicy: group.NewThresholdDecisionPolicy(
// 				sdk.NewDec(math.MaxInt64),
// 				types.Duration{Seconds: 1},
// 			),
// 		},
// 		"group id does not exists": {
// 			srcAdmin:   []byte("valid--admin-address"),
// 			srcComment: "test",
// 			srcGroupID: 9999,
// 			srcPolicy: group.NewThresholdDecisionPolicy(
// 				sdk.OneDec(),
// 				types.Duration{Seconds: 1},
// 			),
// 			expErr: true,
// 		},
// 		"admin not group admin": {
// 			srcAdmin:   []byte("other--admin-address"),
// 			srcComment: "test",
// 			srcGroupID: myGroupID,
// 			srcPolicy: group.NewThresholdDecisionPolicy(
// 				sdk.OneDec(),
// 				types.Duration{Seconds: 1},
// 			),
// 			expErr: true,
// 		},
// 		"comment too long": {
// 			srcAdmin:   []byte("valid--admin-address"),
// 			srcComment: strings.Repeat("a", 256),
// 			srcGroupID: myGroupID,
// 			srcPolicy: group.NewThresholdDecisionPolicy(
// 				sdk.OneDec(),
// 				types.Duration{Seconds: 1},
// 			),
// 			expErr: true,
// 		},
// 	}
// 	for msg, spec := range specs {
// 		t.Run(msg, func(t *testing.T) {
// 			addr, err := k.CreateGroupAccount(ctx, spec.srcAdmin, spec.srcGroupID, spec.srcPolicy, spec.srcComment)
// 			if spec.expErr {
// 				require.Error(t, err)
// 				return
// 			}
// 			require.NoError(t, err)

// 			// then all data persisted
// 			groupAccount, err := k.GetGroupAccount(ctx, addr)
// 			require.NoError(t, err)
// 			assert.Equal(t, addr, groupAccount.GroupAccount)
// 			assert.Equal(t, myGroupID, groupAccount.GroupId)
// 			assert.Equal(t, sdk.AccAddress([]byte(spec.srcAdmin)), groupAccount.Admin)
// 			assert.Equal(t, spec.srcComment, groupAccount.Comment)
// 			assert.Equal(t, uint64(1), groupAccount.Version)
// 			assert.Equal(t, &spec.srcPolicy, groupAccount.GetDecisionPolicy())
// 		})
// 	}
// }

// func TestCreateProposal(t *testing.T) {
// 	encodingConfig := app.MakeEncodingConfig()

// 	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
// 	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, group.DefaultParamspace)

// 	groupKey := sdk.NewKVStoreKey(group.StoreKey)
// 	k := group.NewGroupKeeper(groupKey, paramSpace, baseapp.NewRouter(), &testdatagroup.MyAppProposal{})
// 	blockTime := time.Now()
// 	ctx := testutil.NewContext(pKey, pTKey, groupKey).WithBlockTime(blockTime)
// 	defaultParams := group.DefaultParams()
// 	paramSpace.SetParamSet(ctx, &defaultParams)

// 	members := []group.Member{{
// 		Address: []byte("valid-member-address"),
// 		Power:   sdk.OneDec(),
// 	}}
// 	myGroupID, err := k.CreateGroup(ctx, []byte("valid--admin-address"), members, "test")
// 	require.NoError(t, err)

// 	policy := group.NewThresholdDecisionPolicy(
// 		sdk.OneDec(),
// 		types.Duration{Seconds: 1},
// 	)
// 	accountAddr, err := k.CreateGroupAccount(ctx, []byte("valid--admin-address"), myGroupID, policy, "test")
// 	require.NoError(t, err)

// 	policy = group.NewThresholdDecisionPolicy(
// 		sdk.NewDec(math.MaxInt64),
// 		types.Duration{Seconds: 1},
// 	)
// 	bigThresholdAddr, err := k.CreateGroupAccount(ctx, []byte("valid--admin-address"), myGroupID, policy, "test")
// 	require.NoError(t, err)

// 	specs := map[string]struct {
// 		srcAccount   sdk.AccAddress
// 		srcProposers []sdk.AccAddress
// 		srcMsgs      []sdk.Msg
// 		srcComment   string
// 		expErr       bool
// 	}{
// 		"all good with minimal fields set": {
// 			srcAccount:   accountAddr,
// 			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
// 		},
// 		"all good with good msg payload": {
// 			srcAccount:   accountAddr,
// 			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcMsgs:      []sdk.Msg{&testdatagroup.MsgAlwaysSucceed{}, &testdatagroup.MsgAlwaysFail{}},
// 		},
// 		"invalid payload should be rejected": {
// 			srcAccount:   accountAddr,
// 			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcMsgs:      []sdk.Msg{&testdatagroup.MsgAlwaysSucceed{}},
// 			srcComment:   "payload not a pointer",
// 			expErr:       true,
// 		},
// 		"comment too long": {
// 			srcAccount:   accountAddr,
// 			srcComment:   strings.Repeat("a", 256),
// 			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
// 			expErr:       true,
// 		},
// 		"group account required": {
// 			srcComment:   "test",
// 			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
// 			expErr:       true,
// 		},
// 		"existing group account required": {
// 			srcAccount:   []byte("non-existing-account"),
// 			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
// 			expErr:       true,
// 		},
// 		"impossible case: decision policy threshold > total group weight": {
// 			srcAccount:   bigThresholdAddr,
// 			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
// 			expErr:       true,
// 		},
// 		"only group members can create a proposal": {
// 			srcAccount:   accountAddr,
// 			srcProposers: []sdk.AccAddress{[]byte("non--member-address")},
// 			expErr:       true,
// 		},
// 		"all proposers must be in group": {
// 			srcAccount:   accountAddr,
// 			srcProposers: []sdk.AccAddress{[]byte("valid-member-address"), []byte("non--member-address")},
// 			expErr:       true,
// 		},
// 		"proposers must not be nil": {
// 			srcAccount:   accountAddr,
// 			srcProposers: []sdk.AccAddress{[]byte("valid-member-address"), nil},
// 			expErr:       true,
// 		},
// 		"admin that is not a group member can not create proposal": {
// 			srcAccount:   accountAddr,
// 			srcComment:   "test",
// 			srcProposers: []sdk.AccAddress{[]byte("valid--admin-address")},
// 			expErr:       true,
// 		},
// 		"reject msgs that are not authz by group account": {
// 			srcAccount:   accountAddr,
// 			srcComment:   "test",
// 			srcMsgs:      []sdk.Msg{&testdatagroup.MsgAuthenticate{Signers: []sdk.AccAddress{[]byte("not-group-acct-addrs")}}},
// 			srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
// 			expErr:       true,
// 		},
// 	}
// 	for msg, spec := range specs {
// 		t.Run(msg, func(t *testing.T) {
// 			id, err := k.CreateProposal(ctx, spec.srcAccount, spec.srcComment, spec.srcProposers, spec.srcMsgs)
// 			if spec.expErr {
// 				require.Error(t, err)
// 				return
// 			}
// 			require.NoError(t, err)

// 			// then all data persisted
// 			proposal, err := k.GetProposal(ctx, id)
// 			require.NoError(t, err)

// 			base := proposal.GetBase()
// 			assert.Equal(t, accountAddr, base.GroupAccount)
// 			assert.Equal(t, spec.srcComment, base.Comment)
// 			assert.Equal(t, spec.srcProposers, base.Proposers)

// 			submittedAt, err := types.TimestampFromProto(&base.SubmittedAt)
// 			require.NoError(t, err)
// 			assert.Equal(t, blockTime.UTC(), submittedAt)

// 			assert.Equal(t, uint64(1), base.GroupVersion)
// 			assert.Equal(t, uint64(1), base.GroupAccountVersion)
// 			assert.Equal(t, group.ProposalStatusSubmitted, base.Status)
// 			assert.Equal(t, group.ProposalResultUndefined, base.Result)
// 			assert.Equal(t, group.Tally{
// 				YesCount:     sdk.ZeroDec(),
// 				NoCount:      sdk.ZeroDec(),
// 				AbstainCount: sdk.ZeroDec(),
// 				VetoCount:    sdk.ZeroDec(),
// 			}, base.VoteState)

// 			timeout, err := types.TimestampFromProto(&base.Timeout)
// 			require.NoError(t, err)
// 			assert.Equal(t, blockTime.Add(time.Second).UTC(), timeout)

// 			if spec.srcMsgs == nil { // then empty list is ok
// 				assert.Len(t, proposal.GetMsgs(), 0)
// 			} else {
// 				assert.Equal(t, spec.srcMsgs, proposal.GetMsgs())
// 			}
// 		})
// 	}
// }

// func TestVote(t *testing.T) {
// 	encodingConfig := app.MakeEncodingConfig()

// 	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
// 	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, group.DefaultParamspace)

// 	groupKey := sdk.NewKVStoreKey(group.StoreKey)
// 	k := group.NewGroupKeeper(groupKey, paramSpace, baseapp.NewRouter(), &testdatagroup.MyAppProposal{})
// 	blockTime := time.Now().UTC()
// 	parentCtx := testutil.NewContext(pKey, pTKey, groupKey).WithBlockTime(blockTime)
// 	defaultParams := group.DefaultParams()
// 	paramSpace.SetParamSet(parentCtx, &defaultParams)

// 	members := []group.Member{
// 		{Address: []byte("valid-member-address"), Power: sdk.OneDec()},
// 		{Address: []byte("power-member-address"), Power: sdk.NewDec(2)},
// 	}
// 	myGroupID, err := k.CreateGroup(parentCtx, []byte("valid--admin-address"), members, "test")
// 	require.NoError(t, err)

// 	policy := group.NewThresholdDecisionPolicy(
// 		sdk.NewDec(2),
// 		types.Duration{Seconds: 1},
// 	)
// 	accountAddr, err := k.CreateGroupAccount(parentCtx, []byte("valid--admin-address"), myGroupID, policy, "test")
// 	require.NoError(t, err)
// 	myProposalID, err := k.CreateProposal(parentCtx, accountAddr, "integration test", []sdk.AccAddress{[]byte("valid-member-address")}, nil)
// 	require.NoError(t, err)

// 	specs := map[string]struct {
// 		srcProposalID     group.ProposalId
// 		srcVoters         []sdk.AccAddress
// 		srcChoice         group.Choice
// 		srcComment        string
// 		srcCtx            sdk.Context
// 		doBefore          func(t *testing.T, ctx sdk.Context)
// 		expErr            bool
// 		expVoteState      group.Tally
// 		expProposalStatus group.ProposalBase_Status
// 		expResult         group.ProposalBase_Result
// 	}{
// 		"vote yes": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_YES,
// 			expVoteState: group.Tally{
// 				YesCount:     sdk.OneDec(),
// 				NoCount:      sdk.ZeroDec(),
// 				AbstainCount: sdk.ZeroDec(),
// 				VetoCount:    sdk.ZeroDec(),
// 			},
// 			expProposalStatus: group.ProposalStatusSubmitted,
// 			expResult:         group.ProposalResultUndefined,
// 		},
// 		"vote no": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_NO,
// 			expVoteState: group.Tally{
// 				YesCount:     sdk.ZeroDec(),
// 				NoCount:      sdk.OneDec(),
// 				AbstainCount: sdk.ZeroDec(),
// 				VetoCount:    sdk.ZeroDec(),
// 			},
// 			expProposalStatus: group.ProposalStatusSubmitted,
// 			expResult:         group.ProposalResultUndefined,
// 		},
// 		"vote abstain": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_ABSTAIN,
// 			expVoteState: group.Tally{
// 				YesCount:     sdk.ZeroDec(),
// 				NoCount:      sdk.ZeroDec(),
// 				AbstainCount: sdk.OneDec(),
// 				VetoCount:    sdk.ZeroDec(),
// 			},
// 			expProposalStatus: group.ProposalStatusSubmitted,
// 			expResult:         group.ProposalResultUndefined,
// 		},
// 		"vote veto": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_VETO,
// 			expVoteState: group.Tally{
// 				YesCount:     sdk.ZeroDec(),
// 				NoCount:      sdk.ZeroDec(),
// 				AbstainCount: sdk.ZeroDec(),
// 				VetoCount:    sdk.OneDec(),
// 			},
// 			expProposalStatus: group.ProposalStatusSubmitted,
// 			expResult:         group.ProposalResultUndefined,
// 		},
// 		"apply decision policy early": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("power-member-address")},
// 			srcChoice:     group.Choice_YES,
// 			expVoteState: group.Tally{
// 				YesCount:     sdk.NewDec(2),
// 				NoCount:      sdk.ZeroDec(),
// 				AbstainCount: sdk.ZeroDec(),
// 				VetoCount:    sdk.ZeroDec(),
// 			},
// 			expProposalStatus: group.ProposalStatusClosed,
// 			expResult:         group.ProposalResultAccepted,
// 		},
// 		"reject new votes when final decision is made already": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_YES,
// 			doBefore: func(t *testing.T, ctx sdk.Context) {
// 				require.NoError(t, k.Vote(ctx, myProposalID, []sdk.AccAddress{[]byte("power-member-address")}, group.Choice_VETO, ""))
// 			},
// 			expErr: true,
// 		},
// 		"comment too long": {
// 			srcProposalID: myProposalID,
// 			srcComment:    strings.Repeat("a", 256),
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_NO,
// 			expErr:        true,
// 		},
// 		"existing proposal required": {
// 			srcProposalID: 9999,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_NO,
// 			expErr:        true,
// 		},
// 		"empty choice": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			expErr:        true,
// 		},
// 		"invalid choice": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     5,
// 			expErr:        true,
// 		},
// 		"all voters must be in group": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address"), []byte("non--member-address")},
// 			srcChoice:     group.Choice_NO,
// 			expErr:        true,
// 		},
// 		"voters must not include nil": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address"), nil},
// 			srcChoice:     group.Choice_NO,
// 			expErr:        true,
// 		},
// 		"voters must not be nil": {
// 			srcProposalID: myProposalID,
// 			srcChoice:     group.Choice_NO,
// 			expErr:        true,
// 		},
// 		"voters must not be empty": {
// 			srcProposalID: myProposalID,
// 			srcChoice:     group.Choice_NO,
// 			srcVoters:     []sdk.AccAddress{},
// 			expErr:        true,
// 		},
// 		"admin that is not a group member can not vote": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid--admin-address")},
// 			srcChoice:     group.Choice_NO,
// 			expErr:        true,
// 		},
// 		"on timeout": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_NO,
// 			srcCtx:        parentCtx.WithBlockTime(blockTime.Add(time.Second)),
// 			expErr:        true,
// 		},
// 		"closed already": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_NO,
// 			doBefore: func(t *testing.T, ctx sdk.Context) {
// 				err := k.Vote(ctx, myProposalID, []sdk.AccAddress{[]byte("power-member-address")}, group.Choice_YES, "")
// 				require.NoError(t, err)
// 			},
// 			expErr: true,
// 		},
// 		"voted already": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_NO,
// 			doBefore: func(t *testing.T, ctx sdk.Context) {
// 				err := k.Vote(ctx, myProposalID, []sdk.AccAddress{[]byte("valid-member-address")}, group.Choice_YES, "")
// 				require.NoError(t, err)
// 			},
// 			expErr: true,
// 		},
// 		"with group modified": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_NO,
// 			doBefore: func(t *testing.T, ctx sdk.Context) {
// 				g, err := k.GetGroup(ctx, myGroupID)
// 				require.NoError(t, err)
// 				g.Comment = "modified"
// 				require.NoError(t, k.UpdateGroup(ctx, &g))
// 			},
// 			expErr: true,
// 		},
// 		"with policy modified": {
// 			srcProposalID: myProposalID,
// 			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
// 			srcChoice:     group.Choice_NO,
// 			doBefore: func(t *testing.T, ctx sdk.Context) {
// 				a, err := k.GetGroupAccount(ctx, accountAddr)
// 				require.NoError(t, err)
// 				a.Comment = "modified"
// 				require.NoError(t, k.UpdateGroupAccount(ctx, &a))
// 			},
// 			expErr: true,
// 		},
// 	}
// 	for msg, spec := range specs {
// 		t.Run(msg, func(t *testing.T) {
// 			ctx := parentCtx
// 			if !spec.srcCtx.IsZero() {
// 				ctx = spec.srcCtx
// 			}
// 			ctx, _ = ctx.CacheContext()

// 			if spec.doBefore != nil {
// 				spec.doBefore(t, ctx)
// 			}
// 			err := k.Vote(ctx, spec.srcProposalID, spec.srcVoters, spec.srcChoice, spec.srcComment)
// 			if spec.expErr {
// 				require.Error(t, err)
// 				return
// 			}
// 			require.NoError(t, err)

// 			// and all votes are stored
// 			for _, voter := range spec.srcVoters {
// 				// then all data persisted
// 				loaded, err := k.GetVote(ctx, spec.srcProposalID, voter)
// 				require.NoError(t, err)
// 				assert.Equal(t, spec.srcProposalID, loaded.ProposalId)
// 				assert.Equal(t, voter, loaded.Voter)
// 				assert.Equal(t, spec.srcChoice, loaded.Choice)
// 				assert.Equal(t, spec.srcComment, loaded.Comment)
// 				submittedAt, err := types.TimestampFromProto(&loaded.SubmittedAt)
// 				require.NoError(t, err)
// 				assert.Equal(t, blockTime, submittedAt)
// 			}

// 			// and proposal is updated
// 			proposal, err := k.GetProposal(ctx, spec.srcProposalID)
// 			require.NoError(t, err)
// 			assert.Equal(t, spec.expVoteState, proposal.GetBase().VoteState)
// 			assert.Equal(t, spec.expResult, proposal.GetBase().Result)
// 			assert.Equal(t, spec.expProposalStatus, proposal.GetBase().Status)
// 		})
// 	}
// }

// func TestExecProposal(t *testing.T) {
// 	encodingConfig := app.MakeEncodingConfig()

// 	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
// 	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, group.DefaultParamspace)

// 	router := baseapp.NewRouter()
// 	groupKey := sdk.NewKVStoreKey(group.StoreKey)
// 	k := group.NewGroupKeeper(groupKey, paramSpace, router, &testdatagroup.MyAppProposal{})
// 	testdataKey := sdk.NewKVStoreKey(testdatagroup.ModuleName)
// 	testdataKeeper := testdatagroup.NewKeeper(testdataKey, k)
// 	router.AddRoute(sdk.NewRoute(testdatagroup.ModuleName, testdatagroup.NewHandler(testdataKeeper)))

// 	blockTime := time.Now().UTC()
// 	parentCtx := testutil.NewContext(pKey, pTKey, groupKey, testdataKey).WithBlockTime(blockTime)
// 	defaultParams := group.DefaultParams()
// 	paramSpace.SetParamSet(parentCtx, &defaultParams)

// 	members := []group.Member{
// 		{Address: []byte("valid-member-address"), Power: sdk.OneDec()},
// 	}
// 	myGroupID, err := k.CreateGroup(parentCtx, []byte("valid--admin-address"), members, "test")
// 	require.NoError(t, err)

// 	policy := group.NewThresholdDecisionPolicy(
// 		sdk.OneDec(),
// 		types.Duration{Seconds: 1},
// 	)
// 	accountAddr, err := k.CreateGroupAccount(parentCtx, []byte("valid--admin-address"), myGroupID, policy, "test")
// 	require.NoError(t, err)

// 	specs := map[string]struct {
// 		srcBlockTime      time.Time
// 		setupProposal     func(t *testing.T, ctx sdk.Context) group.ProposalId
// 		expErr            bool
// 		expProposalStatus group.ProposalBase_Status
// 		expProposalResult group.ProposalBase_Result
// 		expExecutorResult group.ProposalBase_ExecutorResult
// 		expPayloadCounter uint64
// 	}{
// 		"proposal executed when accepted": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgIncCounter{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, group.Choice_YES, ""))
// 				return myProposalID
// 			},
// 			expProposalStatus: group.ProposalStatusClosed,
// 			expProposalResult: group.ProposalResultAccepted,
// 			expExecutorResult: group.ProposalExecutorResultSuccess,
// 			expPayloadCounter: 1,
// 		},
// 		"proposal with multiple messages executed when accepted": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgIncCounter{}, &testdatagroup.MsgIncCounter{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, group.Choice_YES, ""))
// 				return myProposalID
// 			},
// 			expProposalStatus: group.ProposalStatusClosed,
// 			expProposalResult: group.ProposalResultAccepted,
// 			expExecutorResult: group.ProposalExecutorResultSuccess,
// 			expPayloadCounter: 2,
// 		},
// 		"proposal not executed when rejected": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, group.Choice_NO, ""))
// 				return myProposalID
// 			},
// 			expProposalStatus: group.ProposalStatusClosed,
// 			expProposalResult: group.ProposalResultRejected,
// 			expExecutorResult: group.ProposalExecutorResultNotRun,
// 		},
// 		"open proposal must not fail": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				return myProposalID
// 			},
// 			expProposalStatus: group.ProposalStatusSubmitted,
// 			expProposalResult: group.ProposalResultUndefined,
// 			expExecutorResult: group.ProposalExecutorResultNotRun,
// 		},
// 		"existing proposal required": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				return 9999
// 			},
// 			expErr: true,
// 		},
// 		"Decision policy also applied on timeout": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, group.Choice_NO, ""))
// 				return myProposalID
// 			},
// 			srcBlockTime:      blockTime.Add(time.Second),
// 			expProposalStatus: group.ProposalStatusClosed,
// 			expProposalResult: group.ProposalResultRejected,
// 			expExecutorResult: group.ProposalExecutorResultNotRun,
// 		},
// 		"Decision policy also applied after timeout": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, group.Choice_NO, ""))
// 				return myProposalID
// 			},
// 			srcBlockTime:      blockTime.Add(time.Second).Add(time.Millisecond),
// 			expProposalStatus: group.ProposalStatusClosed,
// 			expProposalResult: group.ProposalResultRejected,
// 			expExecutorResult: group.ProposalExecutorResultNotRun,
// 		},
// 		"with group modified before tally": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				// then modify group
// 				g, err := k.GetGroup(ctx, myGroupID)
// 				require.NoError(t, err)
// 				g.Comment = "modified"
// 				require.NoError(t, k.UpdateGroup(ctx, &g))
// 				return myProposalID
// 			},
// 			expProposalStatus: group.ProposalStatusAborted,
// 			expProposalResult: group.ProposalResultUndefined,
// 			expExecutorResult: group.ProposalExecutorResultNotRun,
// 		},
// 		"with group account modified before tally": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				// then modify group account
// 				a, err := k.GetGroupAccount(ctx, accountAddr)
// 				require.NoError(t, err)
// 				a.Comment = "modified"
// 				require.NoError(t, k.UpdateGroupAccount(ctx, &a))
// 				return myProposalID
// 			},
// 			expProposalStatus: group.ProposalStatusAborted,
// 			expProposalResult: group.ProposalResultUndefined,
// 			expExecutorResult: group.ProposalExecutorResultNotRun,
// 		},
// 		"with group modified after tally": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, group.Choice_YES, ""))
// 				// then modify group after tally on vote
// 				g, err := k.GetGroup(ctx, myGroupID)
// 				require.NoError(t, err)
// 				g.Comment = "modified"
// 				require.NoError(t, k.UpdateGroup(ctx, &g))
// 				return myProposalID
// 			},
// 			expProposalStatus: group.ProposalStatusClosed,
// 			expProposalResult: group.ProposalResultAccepted,
// 			expExecutorResult: group.ProposalExecutorResultFailure,
// 		},
// 		"with group account modified after tally": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				// then modify group account
// 				a, err := k.GetGroupAccount(ctx, accountAddr)
// 				require.NoError(t, err)
// 				a.Comment = "modified"
// 				require.NoError(t, k.UpdateGroupAccount(ctx, &a))
// 				return myProposalID
// 			},
// 			expProposalStatus: group.ProposalStatusAborted,
// 			expProposalResult: group.ProposalResultUndefined,
// 			expExecutorResult: group.ProposalExecutorResultNotRun,
// 		},
// 		"prevent double execution when successful": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgIncCounter{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, group.Choice_YES, ""))
// 				require.NoError(t, k.ExecProposal(ctx, myProposalID))
// 				return myProposalID
// 			},
// 			expPayloadCounter: 1,
// 			expProposalStatus: group.ProposalStatusClosed,
// 			expProposalResult: group.ProposalResultAccepted,
// 			expExecutorResult: group.ProposalExecutorResultSuccess,
// 		},
// 		"rollback all msg updates on failure": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgIncCounter{}, &testdatagroup.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, group.Choice_YES, ""))
// 				return myProposalID
// 			},
// 			expProposalStatus: group.ProposalStatusClosed,
// 			expProposalResult: group.ProposalResultAccepted,
// 			expExecutorResult: group.ProposalExecutorResultFailure,
// 		},
// 		"executable when failed before": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) group.ProposalId {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatagroup.MsgConditional{ExpectedCounter: 1}, &testdatagroup.MsgIncCounter{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, group.Choice_YES, ""))
// 				require.NoError(t, k.ExecProposal(ctx, myProposalID))
// 				testdataKeeper.IncCounter(ctx)
// 				return myProposalID
// 			},
// 			expPayloadCounter: 2,
// 			expProposalStatus: group.ProposalStatusClosed,
// 			expProposalResult: group.ProposalResultAccepted,
// 			expExecutorResult: group.ProposalExecutorResultSuccess,
// 		},
// 	}
// 	for msg, spec := range specs {
// 		t.Run(msg, func(t *testing.T) {
// 			ctx, _ := parentCtx.CacheContext()
// 			proposalID := spec.setupProposal(t, ctx)

// 			if !spec.srcBlockTime.IsZero() {
// 				ctx = ctx.WithBlockTime(spec.srcBlockTime)
// 			}
// 			err := k.ExecProposal(ctx, proposalID)
// 			if spec.expErr {
// 				require.Error(t, err)
// 				return
// 			}
// 			require.NoError(t, err)

// 			// and proposal is updated
// 			proposal, err := k.GetProposal(ctx, proposalID)
// 			require.NoError(t, err)
// 			exp := group.ProposalBase_Result_name[int32(spec.expProposalResult)]
// 			got := group.ProposalBase_Result_name[int32(proposal.GetBase().Result)]
// 			assert.Equal(t, exp, got)

// 			exp = group.ProposalBase_Status_name[int32(spec.expProposalStatus)]
// 			got = group.ProposalBase_Status_name[int32(proposal.GetBase().Status)]
// 			assert.Equal(t, exp, got)

// 			exp = group.ProposalBase_ExecutorResult_name[int32(spec.expExecutorResult)]
// 			got = group.ProposalBase_ExecutorResult_name[int32(proposal.GetBase().ExecutorResult)]
// 			assert.Equal(t, exp, got)

// 			// and proposal messages executed
// 			assert.Equal(t, spec.expPayloadCounter, testdataKeeper.GetCounter(ctx), "counter")
// 		})
// 	}
// }

// func TestLoadParam(t *testing.T) {
// 	encodingConfig := app.MakeEncodingConfig()

// 	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
// 	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, group.DefaultParamspace)

// 	groupKey := sdk.NewKVStoreKey(group.StoreKey)
// 	k := group.NewGroupKeeper(groupKey, paramSpace, baseapp.NewRouter(), &testutil.MockProposalI{})

// 	ctx := testutil.NewContext(pKey, pTKey, groupKey)

// 	myParams := group.Params{MaxCommentLength: 1}
// 	paramSpace.SetParamSet(ctx, &myParams)

// 	assert.Equal(t, myParams, k.GetParams(ctx))
// }
