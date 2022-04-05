package server

import (
	"testing"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group"
)

func getCtxCodecKey(t *testing.T) (sdk.Context, *codec.ProtoCodec, *sdk.KVStoreKey) {
	interfaceRegistry := types.NewInterfaceRegistry()
	group.RegisterTypes(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)
	key := sdk.NewKVStoreKey(group.ModuleName)
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	err := cms.LoadLatestVersion()
	require.NoError(t, err)
	curCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	curCtx = curCtx.WithBlockHeight(10)
	return curCtx, cdc, key
}

func TestTallyVotesInvariant(t *testing.T) {
	curCtx, cdc, key := getCtxCodecKey(t)
	prevCtx, _ := curCtx.CacheContext()
	prevCtx = prevCtx.WithBlockHeight(curCtx.BlockHeight() - 1)

	// Proposal Table
	proposalTableBuilder, err := orm.NewAutoUInt64TableBuilder(ProposalTablePrefix, ProposalTableSeqPrefix, key, &group.Proposal{}, cdc)
	require.NoError(t, err)
	proposalTable := proposalTableBuilder.Build()

	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	curBlockTime, err := gogotypes.TimestampProto(curCtx.BlockTime())
	require.NoError(t, err)
	prevBlockTime, err := gogotypes.TimestampProto(prevCtx.BlockTime())
	require.NoError(t, err)

	specs := map[string]struct {
		prevProposal *group.Proposal
		curProposal  *group.Proposal
		expBroken    bool
	}{
		"invariant not broken": {
			prevProposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr1.String(),
				Proposers:           []string{addr1.String()},
				SubmittedAt:         *prevBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "1", NoCount: "0", AbstainCount: "0", VetoCount: "0"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},

			curProposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr2.String(),
				Proposers:           []string{addr2.String()},
				SubmittedAt:         *curBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "2", NoCount: "0", AbstainCount: "0", VetoCount: "0"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
		},
		"current block yes vote count must be greater than previous block yes vote count": {
			prevProposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr1.String(),
				Proposers:           []string{addr1.String()},
				SubmittedAt:         *prevBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "2", NoCount: "0", AbstainCount: "0", VetoCount: "0"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
			curProposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr2.String(),
				Proposers:           []string{addr2.String()},
				SubmittedAt:         *curBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "1", NoCount: "0", AbstainCount: "0", VetoCount: "0"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
			expBroken: true,
		},
		"current block no vote count must be greater than previous block no vote count": {
			prevProposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr1.String(),
				Proposers:           []string{addr1.String()},
				SubmittedAt:         *prevBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "0", NoCount: "2", AbstainCount: "0", VetoCount: "0"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
			curProposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr2.String(),
				Proposers:           []string{addr2.String()},
				SubmittedAt:         *curBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "0", NoCount: "1", AbstainCount: "0", VetoCount: "0"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
			expBroken: true,
		},
		"current block abstain vote count must be greater than previous block abstain vote count": {
			prevProposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr1.String(),
				Proposers:           []string{addr1.String()},
				SubmittedAt:         *prevBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "0", NoCount: "0", AbstainCount: "2", VetoCount: "0"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
			curProposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr2.String(),
				Proposers:           []string{addr2.String()},
				SubmittedAt:         *curBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "0", NoCount: "0", AbstainCount: "1", VetoCount: "0"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
			expBroken: true,
		},
		"current block veto vote count must be greater than previous block veto vote count": {
			prevProposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr1.String(),
				Proposers:           []string{addr1.String()},
				SubmittedAt:         *prevBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "0", NoCount: "0", AbstainCount: "0", VetoCount: "2"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
			curProposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr2.String(),
				Proposers:           []string{addr2.String()},
				SubmittedAt:         *curBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "0", NoCount: "0", AbstainCount: "0", VetoCount: "1"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
			expBroken: true,
		},
	}

	for _, spec := range specs {

		prevProposal := spec.prevProposal
		curProposal := spec.curProposal

		cachePrevCtx, _ := prevCtx.CacheContext()
		cacheCurCtx, _ := curCtx.CacheContext()

		_, err = proposalTable.Create(cachePrevCtx, prevProposal)
		require.NoError(t, err)
		_, err = proposalTable.Create(cacheCurCtx, curProposal)
		require.NoError(t, err)

		_, broken := tallyVotesInvariant(cacheCurCtx, cachePrevCtx, proposalTable)
		require.Equal(t, spec.expBroken, broken)
	}
}

func TestGroupTotalWeightInvariant(t *testing.T) {
	curCtx, cdc, key := getCtxCodecKey(t)

	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	// Group Table
	groupTableBuilder, err := orm.NewAutoUInt64TableBuilder(GroupTablePrefix, GroupTableSeqPrefix, key, &group.GroupInfo{}, cdc)
	require.NoError(t, err)
	groupTable := groupTableBuilder.Build()

	groupInfo := &group.GroupInfo{
		GroupId:     1,
		Admin:       addr1.String(),
		Version:     1,
		TotalWeight: "3",
	}
	rowID, err := groupTable.Create(curCtx, groupInfo)
	require.NoError(t, err)
	require.Equal(t, uint64(1), rowID)

	// Group Member Table
	groupMemberTableBuilder, err := orm.NewPrimaryKeyTableBuilder(GroupMemberTablePrefix, key, &group.GroupMember{}, cdc)
	require.NoError(t, err)
	groupMemberByGroupIndex, err := orm.NewIndex(groupMemberTableBuilder, GroupMemberByGroupIndexPrefix, func(val interface{}) ([]interface{}, error) {
		group := val.(*group.GroupMember).GroupId
		return []interface{}{group}, nil
	}, group.GroupMember{}.GroupId)
	require.NoError(t, err)
	groupMemberTable := groupMemberTableBuilder.Build()

	specs := map[string]struct {
		groupMembers []*group.GroupMember
		expBroken    bool
	}{
		"invariant not broken": {
			groupMembers: []*group.GroupMember{
				{
					GroupId: 1,
					Member: &group.Member{
						Address: addr1.String(),
						Weight:  "1",
					},
				},
				{
					GroupId: 1,
					Member: &group.Member{
						Address: addr2.String(),
						Weight:  "2",
					},
				},
			},
			expBroken: false,
		},

		"group's TotalWeight must be equal to sum of its members weight ": {
			groupMembers: []*group.GroupMember{
				{
					GroupId: 1,
					Member: &group.Member{
						Address: addr1.String(),
						Weight:  "2",
					},
				},
				{
					GroupId: 1,
					Member: &group.Member{
						Address: addr2.String(),
						Weight:  "2",
					},
				},
			},
			expBroken: true,
		},
	}

	for _, spec := range specs {
		cacheCurCtx, _ := curCtx.CacheContext()
		groupMembers := spec.groupMembers

		for i := 0; i < len(groupMembers); i++ {
			err := groupMemberTable.Create(cacheCurCtx, groupMembers[i])
			require.NoError(t, err)
		}

		_, broken := groupTotalWeightInvariant(cacheCurCtx, groupTable, groupMemberByGroupIndex)
		require.Equal(t, spec.expBroken, broken)
	}
}

func TestTallyVotesSumInvariant(t *testing.T) {
	curCtx, cdc, key := getCtxCodecKey(t)

	// Group Table
	groupTableBuilder, err := orm.NewAutoUInt64TableBuilder(GroupTablePrefix, GroupTableSeqPrefix, key, &group.GroupInfo{}, cdc)
	require.NoError(t, err)
	groupTable := groupTableBuilder.Build()

	// Group Account Table
	groupAccountTableBuilder, err := orm.NewPrimaryKeyTableBuilder(GroupAccountTablePrefix, key, &group.GroupAccountInfo{}, cdc)
	require.NoError(t, err)
	groupAccountTable := groupAccountTableBuilder.Build()

	// Group Member Table
	groupMemberTableBuilder, err := orm.NewPrimaryKeyTableBuilder(GroupMemberTablePrefix, key, &group.GroupMember{}, cdc)
	require.NoError(t, err)
	groupMemberTable := groupMemberTableBuilder.Build()

	// Proposal Table
	proposalTableBuilder, err := orm.NewAutoUInt64TableBuilder(ProposalTablePrefix, ProposalTableSeqPrefix, key, &group.Proposal{}, cdc)
	require.NoError(t, err)
	proposalTable := proposalTableBuilder.Build()

	// Vote Table
	voteTableBuilder, err := orm.NewPrimaryKeyTableBuilder(VoteTablePrefix, key, &group.Vote{}, cdc)
	require.NoError(t, err)
	voteByProposalIndex, err := orm.NewIndex(voteTableBuilder, VoteByProposalIndexPrefix, func(value interface{}) ([]interface{}, error) {
		return []interface{}{value.(*group.Vote).ProposalId}, nil
	}, group.Vote{}.ProposalId)
	require.NoError(t, err)
	voteTable := voteTableBuilder.Build()

	_, _, adminAddr := testdata.KeyTestPubAddr()
	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	curBlockTime, err := gogotypes.TimestampProto(curCtx.BlockTime())
	require.NoError(t, err)

	specs := map[string]struct {
		groupsInfo   *group.GroupInfo
		groupAcc     *group.GroupAccountInfo
		groupMembers []*group.GroupMember
		proposal     *group.Proposal
		votes        []*group.Vote
		expBroken    bool
	}{
		"invariant not broken": {
			groupsInfo: &group.GroupInfo{
				GroupId:     1,
				Admin:       adminAddr.String(),
				Version:     1,
				TotalWeight: "7",
			},
			groupAcc: &group.GroupAccountInfo{
				Address:       addr1.String(),
				GroupId:       1,
				Admin:         adminAddr.String(),
				Version:       1,
				DerivationKey: []byte("derivation-key"),
			},
			groupMembers: []*group.GroupMember{
				{
					GroupId: 1,
					Member: &group.Member{
						Address: addr1.String(),
						Weight:  "4",
					},
				},
				{
					GroupId: 1,
					Member: &group.Member{
						Address: addr2.String(),
						Weight:  "3",
					},
				},
			},
			proposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr1.String(),
				Proposers:           []string{addr1.String()},
				SubmittedAt:         *curBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "4", NoCount: "3", AbstainCount: "0", VetoCount: "0"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
			votes: []*group.Vote{
				{
					ProposalId:  1,
					Voter:       addr1.String(),
					Choice:      group.Choice_CHOICE_YES,
					SubmittedAt: *gogotypes.TimestampNow(),
				},
				{
					ProposalId:  1,
					Voter:       addr2.String(),
					Choice:      group.Choice_CHOICE_NO,
					SubmittedAt: *gogotypes.TimestampNow(),
				},
			},
			expBroken: false,
		},
		"proposal tally must correspond to the sum of vote weights": {
			groupsInfo: &group.GroupInfo{
				GroupId:     1,
				Admin:       adminAddr.String(),
				Version:     1,
				TotalWeight: "5",
			},
			groupAcc: &group.GroupAccountInfo{
				Address:       addr1.String(),
				GroupId:       1,
				Admin:         adminAddr.String(),
				Version:       1,
				DerivationKey: []byte("derivation-key"),
			},
			groupMembers: []*group.GroupMember{
				{
					GroupId: 1,
					Member: &group.Member{
						Address: addr1.String(),
						Weight:  "2",
					},
				},
				{
					GroupId: 1,
					Member: &group.Member{
						Address: addr2.String(),
						Weight:  "3",
					},
				},
			},
			proposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr1.String(),
				Proposers:           []string{addr1.String()},
				SubmittedAt:         *curBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "6", NoCount: "0", AbstainCount: "0", VetoCount: "0"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
			votes: []*group.Vote{
				{
					ProposalId:  1,
					Voter:       addr1.String(),
					Choice:      group.Choice_CHOICE_YES,
					SubmittedAt: *gogotypes.TimestampNow(),
				},
				{
					ProposalId:  1,
					Voter:       addr2.String(),
					Choice:      group.Choice_CHOICE_YES,
					SubmittedAt: *gogotypes.TimestampNow(),
				},
			},
			expBroken: true,
		},
		"proposal VoteState must correspond to the vote choice": {
			groupsInfo: &group.GroupInfo{
				GroupId:     1,
				Admin:       adminAddr.String(),
				Version:     1,
				TotalWeight: "7",
			},
			groupAcc: &group.GroupAccountInfo{
				Address:       addr1.String(),
				GroupId:       1,
				Admin:         adminAddr.String(),
				Version:       1,
				DerivationKey: []byte("derivation-key"),
			},
			groupMembers: []*group.GroupMember{
				{
					GroupId: 1,
					Member: &group.Member{
						Address: addr1.String(),
						Weight:  "4",
					},
				},
				{
					GroupId: 1,
					Member: &group.Member{
						Address: addr2.String(),
						Weight:  "3",
					},
				},
			},
			proposal: &group.Proposal{
				ProposalId:          1,
				Address:             addr1.String(),
				Proposers:           []string{addr1.String()},
				SubmittedAt:         *curBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Status:              group.ProposalStatusSubmitted,
				Result:              group.ProposalResultUnfinalized,
				VoteState:           group.Tally{YesCount: "4", NoCount: "3", AbstainCount: "0", VetoCount: "0"},
				Timeout:             gogotypes.Timestamp{Seconds: 600},
				ExecutorResult:      group.ProposalExecutorResultNotRun,
			},
			votes: []*group.Vote{
				{
					ProposalId:  1,
					Voter:       addr1.String(),
					Choice:      group.Choice_CHOICE_YES,
					SubmittedAt: *gogotypes.TimestampNow(),
				},
				{
					ProposalId:  1,
					Voter:       addr2.String(),
					Choice:      group.Choice_CHOICE_ABSTAIN,
					SubmittedAt: *gogotypes.TimestampNow(),
				},
			},
			expBroken: true,
		},
	}

	for _, spec := range specs {
		cacheCurCtx, _ := curCtx.CacheContext()
		groupsInfo := spec.groupsInfo
		proposal := spec.proposal
		groupAcc := spec.groupAcc
		groupMembers := spec.groupMembers
		votes := spec.votes

		groupID, err := groupTable.Create(cacheCurCtx, groupsInfo)
		require.NoError(t, err)
		require.Equal(t, groupsInfo.GroupId, groupID)

		err = groupAcc.SetDecisionPolicy(group.NewThresholdDecisionPolicy("1", gogotypes.Duration{Seconds: 1}))
		require.NoError(t, err)
		err = groupAccountTable.Create(cacheCurCtx, groupAcc)
		require.NoError(t, err)

		for i := 0; i < len(groupMembers); i++ {
			err = groupMemberTable.Create(cacheCurCtx, groupMembers[i])
			require.NoError(t, err)
		}

		_, err = proposalTable.Create(cacheCurCtx, proposal)
		require.NoError(t, err)

		for i := 0; i < len(votes); i++ {
			err = voteTable.Create(cacheCurCtx, votes[i])
			require.NoError(t, err)
		}

		_, broken := tallyVotesSumInvariant(cacheCurCtx, groupTable, proposalTable, groupMemberTable, voteByProposalIndex, groupAccountTable)
		require.Equal(t, spec.expBroken, broken)
	}
}
