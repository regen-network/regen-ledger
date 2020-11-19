package server

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/regen-network/regen-ledger/app"
	"github.com/regen-network/regen-ledger/orm"

	// "github.com/regen-network/regen-ledger/x/group/server"
	"github.com/regen-network/regen-ledger/x/group/testutil"
	"github.com/regen-network/regen-ledger/x/group/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

// TODO: Setup test suite
func TestCreateGroup(t *testing.T) {
	encodingConfig := app.MakeEncodingConfig()

	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, types.DefaultParamspace)

	groupKey := sdk.NewKVStoreKey(types.StoreKey)
	k := NewGroupKeeper(groupKey, paramSpace, baseapp.NewRouter())
	ctx := testutil.NewContext(pKey, pTKey, groupKey)
	defaultParams := types.DefaultParams()
	paramSpace.SetParamSet(ctx, &defaultParams)

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
		t.Run(msg, func(t *testing.T) {
			id, err := k.CreateGroup(ctx, spec.srcAdmin, spec.srcMembers, spec.srcComment)
			if spec.expErr {
				require.Error(t, err)
				require.False(t, k.HasGroup(ctx, types.GroupID(seq+1).Bytes()))
				return
			}
			require.NoError(t, err)

			seq++
			assert.Equal(t, types.GroupID(seq), id)

			// then all data persisted
			loadedGroup, err := k.GetGroup(ctx, id)
			require.NoError(t, err)
			assert.Equal(t, sdk.AccAddress([]byte(spec.srcAdmin)), loadedGroup.Admin)
			assert.Equal(t, spec.srcComment, loadedGroup.Comment)
			assert.Equal(t, id, loadedGroup.Group)
			assert.Equal(t, uint64(1), loadedGroup.Version)

			// and members are stored as well
			it, err := k.GetGroupMembersByGroup(ctx, id)
			require.NoError(t, err)
			var loadedMembers []types.GroupMember
			_, err = orm.ReadAll(it, &loadedMembers)
			require.NoError(t, err)
			assert.Equal(t, len(members), len(loadedMembers))
			for i := range loadedMembers {
				assert.Equal(t, members[i].Comment, loadedMembers[i].Comment)
				assert.Equal(t, members[i].Address, loadedMembers[i].Member)
				assert.Equal(t, members[i].Power, loadedMembers[i].Weight)
				assert.Equal(t, id, loadedMembers[i].Group)
			}
		})
	}
}

func TestCreateGroupAccount(t *testing.T) {
	encodingConfig := app.MakeEncodingConfig()

	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, types.DefaultParamspace)

	groupKey := sdk.NewKVStoreKey(types.StoreKey)
	k := NewGroupKeeper(groupKey, paramSpace, baseapp.NewRouter())
	ctx := testutil.NewContext(pKey, pTKey, groupKey)
	defaultParams := types.DefaultParams()
	paramSpace.SetParamSet(ctx, &defaultParams)

	myGroupID, err := k.CreateGroup(ctx, []byte("valid--admin-address"), nil, "test")
	require.NoError(t, err)

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
		t.Run(msg, func(t *testing.T) {
			addr, err := k.CreateGroupAccount(ctx, spec.srcAdmin, spec.srcGroupID, spec.srcPolicy, spec.srcComment)
			if spec.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// then all data persisted
			groupAccount, err := k.GetGroupAccount(ctx, addr)
			require.NoError(t, err)
			assert.Equal(t, addr, groupAccount.GroupAccount)
			assert.Equal(t, myGroupID, groupAccount.Group)
			assert.Equal(t, sdk.AccAddress([]byte(spec.srcAdmin)), groupAccount.Admin)
			assert.Equal(t, spec.srcComment, groupAccount.Comment)
			assert.Equal(t, uint64(1), groupAccount.Version)
			// TODO Fix (ORM should unpack Any's properly)
			// assert.Equal(t, &spec.srcPolicy, groupAccount.GetDecisionPolicy())
		})
	}
}

func TestCreateProposal(t *testing.T) {
	encodingConfig := app.MakeEncodingConfig()

	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, types.DefaultParamspace)

	groupKey := sdk.NewKVStoreKey(types.StoreKey)
	k := NewGroupKeeper(groupKey, paramSpace, baseapp.NewRouter())
	blockTime := time.Now()
	ctx := testutil.NewContext(pKey, pTKey, groupKey).WithBlockTime(blockTime)
	defaultParams := types.DefaultParams()
	paramSpace.SetParamSet(ctx, &defaultParams)

	members := []types.Member{{
		Address: []byte("valid-member-address"),
		Power:   sdk.OneDec(),
	}}
	myGroupID, err := k.CreateGroup(ctx, []byte("valid--admin-address"), members, "test")
	require.NoError(t, err)

	policy := types.NewThresholdDecisionPolicy(
		sdk.OneDec(),
		gogotypes.Duration{Seconds: 1},
	)
	accountAddr, err := k.CreateGroupAccount(ctx, []byte("valid--admin-address"), myGroupID, policy, "test")
	require.NoError(t, err)

	policy = types.NewThresholdDecisionPolicy(
		sdk.NewDec(math.MaxInt64),
		gogotypes.Duration{Seconds: 1},
	)
	bigThresholdAddr, err := k.CreateGroupAccount(ctx, []byte("valid--admin-address"), myGroupID, policy, "test")
	require.NoError(t, err)

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
			srcMsgs:      []sdk.Msg{testdata.NewServiceMsgCreateDog(&testdata.MsgCreateDog{})},
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
		// "reject msgs that are not authz by group account": {
		// 	srcAccount:   accountAddr,
		// 	srcComment:   "test",
		// 	srcMsgs:      []sdk.Msg{&banktypes.MsgSend{FromAddress: "regen1yq8lgssgxlx9smjhes6ryjasmqmd3ts2559g0t"}},
		// 	srcProposers: []sdk.AccAddress{[]byte("valid-member-address")},
		// 	expErr:       true,
		// },
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			id, err := k.CreateProposal(ctx, spec.srcAccount, spec.srcComment, spec.srcProposers, spec.srcMsgs)
			if spec.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// then all data persisted
			proposal, err := k.GetProposal(ctx, id)
			require.NoError(t, err)

			assert.Equal(t, accountAddr, proposal.GroupAccount)
			assert.Equal(t, spec.srcComment, proposal.Comment)
			assert.Equal(t, spec.srcProposers, proposal.Proposers)

			submittedAt, err := gogotypes.TimestampFromProto(&proposal.SubmittedAt)
			require.NoError(t, err)
			assert.Equal(t, blockTime.UTC(), submittedAt)

			assert.Equal(t, uint64(1), proposal.GroupVersion)
			assert.Equal(t, uint64(1), proposal.GroupAccountVersion)
			assert.Equal(t, types.ProposalStatusSubmitted, proposal.Status)
			assert.Equal(t, types.ProposalResultUndefined, proposal.Result)
			assert.Equal(t, types.Tally{
				YesCount:     sdk.ZeroDec(),
				NoCount:      sdk.ZeroDec(),
				AbstainCount: sdk.ZeroDec(),
				VetoCount:    sdk.ZeroDec(),
			}, proposal.VoteState)

			timeout, err := gogotypes.TimestampFromProto(&proposal.Timeout)
			require.NoError(t, err)
			assert.Equal(t, blockTime.Add(time.Second).UTC(), timeout)

			if spec.srcMsgs == nil { // then empty list is ok
				assert.Len(t, proposal.GetMsgs(), 0)
			} else {
				assert.Equal(t, spec.srcMsgs, proposal.GetMsgs())
			}
		})
	}
}

func TestVote(t *testing.T) {
	encodingConfig := app.MakeEncodingConfig()

	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, types.DefaultParamspace)

	groupKey := sdk.NewKVStoreKey(types.StoreKey)
	k := NewGroupKeeper(groupKey, paramSpace, baseapp.NewRouter())
	blockTime := time.Now().UTC()
	parentCtx := testutil.NewContext(pKey, pTKey, groupKey).WithBlockTime(blockTime)
	defaultParams := types.DefaultParams()
	paramSpace.SetParamSet(parentCtx, &defaultParams)

	members := []types.Member{
		{Address: []byte("valid-member-address"), Power: sdk.OneDec()},
		{Address: []byte("power-member-address"), Power: sdk.NewDec(2)},
	}
	myGroupID, err := k.CreateGroup(parentCtx, []byte("valid--admin-address"), members, "test")
	require.NoError(t, err)

	policy := types.NewThresholdDecisionPolicy(
		sdk.NewDec(2),
		gogotypes.Duration{Seconds: 1},
	)
	accountAddr, err := k.CreateGroupAccount(parentCtx, []byte("valid--admin-address"), myGroupID, policy, "test")
	require.NoError(t, err)
	myProposalID, err := k.CreateProposal(parentCtx, accountAddr, "integration test", []sdk.AccAddress{[]byte("valid-member-address")}, nil)
	require.NoError(t, err)

	specs := map[string]struct {
		srcProposalID     types.ProposalID
		srcVoters         []sdk.AccAddress
		srcChoice         types.Choice
		srcComment        string
		srcCtx            sdk.Context
		doBefore          func(t *testing.T, ctx sdk.Context)
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
			doBefore: func(t *testing.T, ctx sdk.Context) {
				require.NoError(t, k.Vote(ctx, myProposalID, []sdk.AccAddress{[]byte("power-member-address")}, types.Choice_VETO, ""))
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
			srcCtx:        parentCtx.WithBlockTime(blockTime.Add(time.Second)),
			expErr:        true,
		},
		"closed already": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			doBefore: func(t *testing.T, ctx sdk.Context) {
				err := k.Vote(ctx, myProposalID, []sdk.AccAddress{[]byte("power-member-address")}, types.Choice_YES, "")
				require.NoError(t, err)
			},
			expErr: true,
		},
		"voted already": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			doBefore: func(t *testing.T, ctx sdk.Context) {
				err := k.Vote(ctx, myProposalID, []sdk.AccAddress{[]byte("valid-member-address")}, types.Choice_YES, "")
				require.NoError(t, err)
			},
			expErr: true,
		},
		"with group modified": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			doBefore: func(t *testing.T, ctx sdk.Context) {
				g, err := k.GetGroup(ctx, myGroupID)
				require.NoError(t, err)
				g.Comment = "modified"
				require.NoError(t, k.UpdateGroup(ctx, &g))
			},
			expErr: true,
		},
		"with policy modified": {
			srcProposalID: myProposalID,
			srcVoters:     []sdk.AccAddress{[]byte("valid-member-address")},
			srcChoice:     types.Choice_NO,
			doBefore: func(t *testing.T, ctx sdk.Context) {
				a, err := k.GetGroupAccount(ctx, accountAddr)
				require.NoError(t, err)
				a.Comment = "modified"
				require.NoError(t, k.UpdateGroupAccount(ctx, &a))
			},
			expErr: true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			ctx := parentCtx
			if !spec.srcCtx.IsZero() {
				ctx = spec.srcCtx
			}
			ctx, _ = ctx.CacheContext()

			if spec.doBefore != nil {
				spec.doBefore(t, ctx)
			}
			err := k.Vote(ctx, spec.srcProposalID, spec.srcVoters, spec.srcChoice, spec.srcComment)
			if spec.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// and all votes are stored
			for _, voter := range spec.srcVoters {
				// then all data persisted
				loaded, err := k.GetVote(ctx, spec.srcProposalID, voter)
				require.NoError(t, err)
				assert.Equal(t, spec.srcProposalID, loaded.Proposal)
				assert.Equal(t, voter, loaded.Voter)
				assert.Equal(t, spec.srcChoice, loaded.Choice)
				assert.Equal(t, spec.srcComment, loaded.Comment)
				submittedAt, err := gogotypes.TimestampFromProto(&loaded.SubmittedAt)
				require.NoError(t, err)
				assert.Equal(t, blockTime, submittedAt)
			}

			// and proposal is updated
			proposal, err := k.GetProposal(ctx, spec.srcProposalID)
			require.NoError(t, err)
			assert.Equal(t, spec.expVoteState, proposal.VoteState)
			assert.Equal(t, spec.expResult, proposal.Result)
			assert.Equal(t, spec.expProposalStatus, proposal.Status)
		})
	}
}

// func TestExecProposal(t *testing.T) {
// 	encodingConfig := app.MakeEncodingConfig()

// 	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
// 	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, types.DefaultParamspace)

// 	router := baseapp.NewRouter()
// 	groupKey := sdk.NewKVStoreKey(types.StoreKey)
// 	k := NewGroupKeeper(groupKey, paramSpace, router)
// 	testdataKey := sdk.NewKVStoreKey(testdatatypes.ModuleName)
// 	testdataKeeper := testdatatypes.NewKeeper(testdataKey, k)
// 	router.AddRoute(sdk.NewRoute(testdatatypes.ModuleName, testdatatypes.NewHandler(testdataKeeper)))

// 	blockTime := time.Now().UTC()
// 	parentCtx := testutil.NewContext(pKey, pTKey, groupKey, testdataKey).WithBlockTime(blockTime)
// 	defaultParams := types.DefaultParams()
// 	paramSpace.SetParamSet(parentCtx, &defaultParams)

// 	members := []types.Member{
// 		{Address: []byte("valid-member-address"), Power: sdk.OneDec()},
// 	}
// 	myGroupID, err := k.CreateGroup(parentCtx, []byte("valid--admin-address"), members, "test")
// 	require.NoError(t, err)

// 	policy := types.NewThresholdDecisionPolicy(
// 		sdk.OneDec(),
// 		gogotypes.Duration{Seconds: 1},
// 	)
// 	accountAddr, err := k.CreateGroupAccount(parentCtx, []byte("valid--admin-address"), myGroupID, policy, "test")
// 	require.NoError(t, err)

// 	specs := map[string]struct {
// 		srcBlockTime      time.Time
// 		setupProposal     func(t *testing.T, ctx sdk.Context) types.Proposal
// 		expErr            bool
// 		expProposalStatus types.Proposal_Status
// 		expProposalResult types.Proposal_Result
// 		expExecutorResult types.Proposal_ExecutorResult
// 		expPayloadCounter uint64
// 	}{
// 		"proposal executed when accepted": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgIncCounter{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, types.Choice_YES, ""))
// 				return myProposalID
// 			},
// 			expProposalStatus: types.ProposalStatusClosed,
// 			expProposalResult: types.ProposalResultAccepted,
// 			expExecutorResult: types.ProposalExecutorResultSuccess,
// 			expPayloadCounter: 1,
// 		},
// 		"proposal with multiple messages executed when accepted": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgIncCounter{}, &testdatatypes.MsgIncCounter{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, types.Choice_YES, ""))
// 				return myProposalID
// 			},
// 			expProposalStatus: types.ProposalStatusClosed,
// 			expProposalResult: types.ProposalResultAccepted,
// 			expExecutorResult: types.ProposalExecutorResultSuccess,
// 			expPayloadCounter: 2,
// 		},
// 		"proposal not executed when rejected": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, types.Choice_NO, ""))
// 				return myProposalID
// 			},
// 			expProposalStatus: types.ProposalStatusClosed,
// 			expProposalResult: types.ProposalResultRejected,
// 			expExecutorResult: types.ProposalExecutorResultNotRun,
// 		},
// 		"open proposal must not fail": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				return myProposalID
// 			},
// 			expProposalStatus: types.ProposalStatusSubmitted,
// 			expProposalResult: types.ProposalResultUndefined,
// 			expExecutorResult: types.ProposalExecutorResultNotRun,
// 		},
// 		"existing proposal required": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				return 9999
// 			},
// 			expErr: true,
// 		},
// 		"Decision policy also applied on timeout": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, types.Choice_NO, ""))
// 				return myProposalID
// 			},
// 			srcBlockTime:      blockTime.Add(time.Second),
// 			expProposalStatus: types.ProposalStatusClosed,
// 			expProposalResult: types.ProposalResultRejected,
// 			expExecutorResult: types.ProposalExecutorResultNotRun,
// 		},
// 		"Decision policy also applied after timeout": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, types.Choice_NO, ""))
// 				return myProposalID
// 			},
// 			srcBlockTime:      blockTime.Add(time.Second).Add(time.Millisecond),
// 			expProposalStatus: types.ProposalStatusClosed,
// 			expProposalResult: types.ProposalResultRejected,
// 			expExecutorResult: types.ProposalExecutorResultNotRun,
// 		},
// 		"with group modified before tally": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				// then modify group
// 				g, err := k.GetGroup(ctx, myGroupID)
// 				require.NoError(t, err)
// 				g.Comment = "modified"
// 				require.NoError(t, k.UpdateGroup(ctx, &g))
// 				return myProposalID
// 			},
// 			expProposalStatus: types.ProposalStatusAborted,
// 			expProposalResult: types.ProposalResultUndefined,
// 			expExecutorResult: types.ProposalExecutorResultNotRun,
// 		},
// 		"with group account modified before tally": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				// then modify group account
// 				a, err := k.GetGroupAccount(ctx, accountAddr)
// 				require.NoError(t, err)
// 				a.Comment = "modified"
// 				require.NoError(t, k.UpdateGroupAccount(ctx, &a))
// 				return myProposalID
// 			},
// 			expProposalStatus: types.ProposalStatusAborted,
// 			expProposalResult: types.ProposalResultUndefined,
// 			expExecutorResult: types.ProposalExecutorResultNotRun,
// 		},
// 		"with group modified after tally": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, types.Choice_YES, ""))
// 				// then modify group after tally on vote
// 				g, err := k.GetGroup(ctx, myGroupID)
// 				require.NoError(t, err)
// 				g.Comment = "modified"
// 				require.NoError(t, k.UpdateGroup(ctx, &g))
// 				return myProposalID
// 			},
// 			expProposalStatus: types.ProposalStatusClosed,
// 			expProposalResult: types.ProposalResultAccepted,
// 			expExecutorResult: types.ProposalExecutorResultFailure,
// 		},
// 		"with group account modified after tally": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				// then modify group account
// 				a, err := k.GetGroupAccount(ctx, accountAddr)
// 				require.NoError(t, err)
// 				a.Comment = "modified"
// 				require.NoError(t, k.UpdateGroupAccount(ctx, &a))
// 				return myProposalID
// 			},
// 			expProposalStatus: types.ProposalStatusAborted,
// 			expProposalResult: types.ProposalResultUndefined,
// 			expExecutorResult: types.ProposalExecutorResultNotRun,
// 		},
// 		"prevent double execution when successful": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgIncCounter{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, types.Choice_YES, ""))
// 				require.NoError(t, k.ExecProposal(ctx, myProposalID))
// 				return myProposalID
// 			},
// 			expPayloadCounter: 1,
// 			expProposalStatus: types.ProposalStatusClosed,
// 			expProposalResult: types.ProposalResultAccepted,
// 			expExecutorResult: types.ProposalExecutorResultSuccess,
// 		},
// 		"rollback all msg updates on failure": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgIncCounter{}, &testdatatypes.MsgAlwaysFail{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, types.Choice_YES, ""))
// 				return myProposalID
// 			},
// 			expProposalStatus: types.ProposalStatusClosed,
// 			expProposalResult: types.ProposalResultAccepted,
// 			expExecutorResult: types.ProposalExecutorResultFailure,
// 		},
// 		"executable when failed before": {
// 			setupProposal: func(t *testing.T, ctx sdk.Context) types.Proposal {
// 				member := []sdk.AccAddress{[]byte("valid-member-address")}
// 				myProposalID, err := k.CreateProposal(ctx, accountAddr, "test", member, []sdk.Msg{
// 					&testdatatypes.MsgConditional{ExpectedCounter: 1}, &testdatatypes.MsgIncCounter{},
// 				})
// 				require.NoError(t, err)
// 				require.NoError(t, k.Vote(ctx, myProposalID, member, types.Choice_YES, ""))
// 				require.NoError(t, k.ExecProposal(ctx, myProposalID))
// 				testdataKeeper.IncCounter(ctx)
// 				return myProposalID
// 			},
// 			expPayloadCounter: 2,
// 			expProposalStatus: types.ProposalStatusClosed,
// 			expProposalResult: types.ProposalResultAccepted,
// 			expExecutorResult: types.ProposalExecutorResultSuccess,
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
// 			exp := types.Proposal_Result_name[int32(spec.expProposalResult)]
// 			got := types.Proposal_Result_name[int32(proposal.GetBase().Result)]
// 			assert.Equal(t, exp, got)

// 			exp = types.Proposal_Status_name[int32(spec.expProposalStatus)]
// 			got = types.Proposal_Status_name[int32(proposal.GetBase().Status)]
// 			assert.Equal(t, exp, got)

// 			exp = types.Proposal_ExecutorResult_name[int32(spec.expExecutorResult)]
// 			got = types.Proposal_ExecutorResult_name[int32(proposal.GetBase().ExecutorResult)]
// 			assert.Equal(t, exp, got)

// 			// and proposal messages executed
// 			assert.Equal(t, spec.expPayloadCounter, testdataKeeper.GetCounter(ctx), "counter")
// 		})
// 	}
// }

func TestLoadParam(t *testing.T) {
	encodingConfig := app.MakeEncodingConfig()

	pKey, pTKey := sdk.NewKVStoreKey(paramstypes.StoreKey), sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	paramSpace := paramstypes.NewSubspace(encodingConfig.Marshaler, encodingConfig.Amino, pKey, pTKey, types.DefaultParamspace)

	groupKey := sdk.NewKVStoreKey(types.StoreKey)
	k := NewGroupKeeper(groupKey, paramSpace, baseapp.NewRouter())

	ctx := testutil.NewContext(pKey, pTKey, groupKey)

	myParams := types.Params{MaxCommentLength: 1}
	paramSpace.SetParamSet(ctx, &myParams)

	assert.Equal(t, myParams, k.GetParams(ctx))
}
