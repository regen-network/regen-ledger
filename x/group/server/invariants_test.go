package server

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/types/known/timestamppb"

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
	proposalTableBuilder := orm.NewAutoUInt64TableBuilder(ProposalTablePrefix, ProposalTableSeqPrefix, key, &group.Proposal{}, cdc)
	proposalTable := proposalTableBuilder.Build()

	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	curBlockTime, err := gogotypes.TimestampProto(curCtx.BlockTime())
	require.NoError(t, err)
	prevBlockTime, err := gogotypes.TimestampProto(prevCtx.BlockTime())
	require.NoError(t, err)

	specs := map[string]struct {
		prevReq []*group.Proposal
		curReq  []*group.Proposal
		expErr  bool
	}{
		"invariant not broken": {
			prevReq: []*group.Proposal{
				{
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
			},

			curReq: []*group.Proposal{
				{
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
		},
		"current block yes vote count must be greater than previous block yes vote count": {
			prevReq: []*group.Proposal{
				{
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
			},
			curReq: []*group.Proposal{
				{
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
			},
			expErr: true,
		},
		"current block no vote count must be greater than previous block no vote count": {
			prevReq: []*group.Proposal{
				{
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
			},
			curReq: []*group.Proposal{
				{
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
			},
			expErr: true,
		},
		"current block abstain vote count must be greater than previous block abstain vote count": {
			prevReq: []*group.Proposal{
				{
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
			},
			curReq: []*group.Proposal{
				{
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
			},
			expErr: true,
		},
		"current block veto vote count must be greater than previous block veto vote count": {
			prevReq: []*group.Proposal{
				{
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
			},
			curReq: []*group.Proposal{
				{
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
			},
			expErr: true,
		},
	}

	for _, spec := range specs {

		prevProposals := spec.prevReq
		curProposals := spec.curReq

		cachePrevCtx, _ := prevCtx.CacheContext()
		cacheCurCtx, _ := curCtx.CacheContext()

		for i := 0; i < len(prevProposals) && i < len(curProposals); i++ {
			_, err = proposalTable.Create(cachePrevCtx, prevProposals[i])
			require.NoError(t, err)
			_, err = proposalTable.Create(cacheCurCtx, curProposals[i])
			require.NoError(t, err)
		}
		_, broken, _ := tallyVotesInvariant(cacheCurCtx, cachePrevCtx, proposalTable)
		require.Equal(t, spec.expErr, broken)
	}
}

func TestGroupTotalWeightInvariant(t *testing.T) {
	curCtx, cdc, key := getCtxCodecKey(t)

	// Group Table
	groupTableBuilder := orm.NewTableBuilder(GroupTablePrefix, key, &group.GroupInfo{}, orm.FixLengthIndexKeys(orm.EncodedSeqLength), cdc)
	groupTable := groupTableBuilder.Build()

	// Group Member Table
	groupMemberTableBuilder := orm.NewPrimaryKeyTableBuilder(GroupMemberTablePrefix, key, &group.GroupMember{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	groupMemberByGroupIndex := orm.NewUInt64Index(groupMemberTableBuilder, GroupMemberByGroupIndexPrefix, func(val interface{}) ([]uint64, error) {
		group := val.(*group.GroupMember).GroupId
		return []uint64{group}, nil
	})
	groupMemberTable := groupMemberTableBuilder.Build()

	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	specs := map[string]struct {
		groupReq   []*group.GroupInfo
		membersReq []*group.GroupMember
		expErr     bool
	}{
		"invariant not broken": {
			groupReq: []*group.GroupInfo{
				{
					GroupId:     1,
					Admin:       addr1.String(),
					Version:     1,
					TotalWeight: "3",
				},
			},
			membersReq: []*group.GroupMember{
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
			expErr: false,
		},

		"group's TotalWeight must be equal to sum of its members weight ": {
			groupReq: []*group.GroupInfo{
				{
					GroupId:     1,
					Admin:       addr1.String(),
					Version:     1,
					TotalWeight: "3",
				},
			},
			membersReq: []*group.GroupMember{
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
			expErr: true,
		},
	}

	for _, spec := range specs {
		cacheCurCtx, _ := curCtx.CacheContext()
		groupReq := spec.groupReq
		members := spec.membersReq

		for i := 0; i < len(groupReq); i++ {
			err := groupTable.Create(cacheCurCtx, group.ID(groupReq[i].GroupId).Bytes(), groupReq[i])
			require.NoError(t, err)
		}

		for i := 0; i < len(members); i++ {
			err := groupMemberTable.Create(cacheCurCtx, members[i])
			require.NoError(t, err)
		}

		_, broken, _ := groupTotalWeightInvariant(cacheCurCtx, groupTable, groupMemberByGroupIndex)
		require.Equal(t, spec.expErr, broken)
	}
}

func TestProposalTallyInvariant(t *testing.T) {
	curCtx, cdc, key := getCtxCodecKey(t)

	// Group Table
	groupTableBuilder := orm.NewTableBuilder(GroupTablePrefix, key, &group.GroupInfo{}, orm.FixLengthIndexKeys(orm.EncodedSeqLength), cdc)
	groupTable := groupTableBuilder.Build()

	// Group Account Table
	_ = orm.NewSequence(key, GroupAccountTableSeqPrefix)
	groupAccountTableBuilder := orm.NewPrimaryKeyTableBuilder(GroupAccountTablePrefix, key, &group.GroupAccountInfo{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	groupAccountByGroupIndex := orm.NewUInt64Index(groupAccountTableBuilder, GroupAccountByGroupIndexPrefix, func(value interface{}) ([]uint64, error) {
		group := value.(*group.GroupAccountInfo).GroupId
		return []uint64{group}, nil
	})
	groupAccountTable := groupAccountTableBuilder.Build()

	// Group Member Table
	groupMemberTableBuilder := orm.NewPrimaryKeyTableBuilder(GroupMemberTablePrefix, key, &group.GroupMember{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	groupMemberByGroupIndex := orm.NewUInt64Index(groupMemberTableBuilder, GroupMemberByGroupIndexPrefix, func(val interface{}) ([]uint64, error) {
		group := val.(*group.GroupMember).GroupId
		return []uint64{group}, nil
	})
	groupMemberTable := groupMemberTableBuilder.Build()

	// Proposal Table
	proposalTableBuilder := orm.NewAutoUInt64TableBuilder(ProposalTablePrefix, ProposalTableSeqPrefix, key, &group.Proposal{}, cdc)
	proposalByGroupAccountIndex := orm.NewIndex(proposalTableBuilder, ProposalByGroupAccountIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		account := value.(*group.Proposal).Address
		addr, err := sdk.AccAddressFromBech32(account)
		if err != nil {
			return nil, err
		}
		return []orm.RowID{addr.Bytes()}, nil
	})
	proposalTable := proposalTableBuilder.Build()

	// Vote Table
	voteTableBuilder := orm.NewPrimaryKeyTableBuilder(VoteTablePrefix, key, &group.Vote{}, orm.Max255DynamicLengthIndexKeyCodec{}, cdc)
	// voteByProposalIndex := orm.NewUInt64Index(voteTableBuilder, VoteByProposalIndexPrefix, func(value interface{}) ([]uint64, error) {
	// 	return []uint64{value.(*group.Vote).ProposalId}, nil
	// })
	// voteByVoterIndex := orm.NewIndex(voteTableBuilder, VoteByVoterIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
	// 	addr, err := sdk.AccAddressFromBech32(value.(*group.Vote).Voter)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return []orm.RowID{addr.Bytes()}, nil
	// })
	voteTable := voteTableBuilder.Build()

	_, _, adminAddr := testdata.KeyTestPubAddr()
	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()
	// _, _, addr3 := testdata.KeyTestPubAddr()

	curBlockTime, err := gogotypes.TimestampProto(curCtx.BlockTime())
	require.NoError(t, err)

	specs := map[string]struct {
		groupReq    []*group.GroupInfo
		groupAccReq []*group.GroupAccountInfo
		policy      group.DecisionPolicy
		membersReq  []*group.GroupMember
		proposalReq []*group.Proposal
		voteReq     []*group.Vote
		expErr      bool
	}{
		"invariant not broken": {
			groupReq: []*group.GroupInfo{
				{
					GroupId:     1,
					Admin:       adminAddr.String(),
					Version:     1,
					TotalWeight: "3",
				},
			},
			groupAccReq: []*group.GroupAccountInfo{
				{
					Address: addr1.String(),
					GroupId: 1,
					Admin:   adminAddr.String(),
					Version: 1,
				},
				{
					Address: addr2.String(),
					GroupId: 1,
					Admin:   adminAddr.String(),
					Version: 1,
				},
			},
			policy: group.NewThresholdDecisionPolicy(
				"1",
				gogotypes.Duration{Seconds: 1},
			),
			membersReq: []*group.GroupMember{
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
						Weight:  "1",
					},
				},
			},
			proposalReq: []*group.Proposal{
				{
					ProposalId:          1,
					Address:             addr1.String(),
					Proposers:           []string{addr1.String()},
					SubmittedAt:         *curBlockTime,
					GroupVersion:        1,
					GroupAccountVersion: 1,
					Status:              group.ProposalStatusSubmitted,
					Result:              group.ProposalResultUnfinalized,
					VoteState:           group.Tally{YesCount: "2", NoCount: "1", AbstainCount: "0", VetoCount: "0"},
					Timeout:             gogotypes.Timestamp{Seconds: 600},
					ExecutorResult:      group.ProposalExecutorResultNotRun,
				},
			},
			voteReq: []*group.Vote{
				{
					ProposalId: 1,
					Voter:      addr1.String(),
					Choice:     group.Choice_CHOICE_YES,
					SubmittedAt: gogotypes.Timestamp{
						Seconds: timestamppb.Now().Seconds,
						Nanos:   timestamppb.Now().Nanos,
					},
				},
				{
					ProposalId: 1,
					Voter:      addr2.String(),
					Choice:     group.Choice_CHOICE_NO,
					SubmittedAt: gogotypes.Timestamp{
						Seconds: timestamppb.Now().Seconds,
						Nanos:   timestamppb.Now().Nanos,
					},
				},
			},
			expErr: false,
		},
	}

	for _, spec := range specs {
		cacheCurCtx, _ := curCtx.CacheContext()
		groupReq := spec.groupReq
		proposals := spec.proposalReq
		members := spec.membersReq
		votes := spec.voteReq
		groupAcc := spec.groupAccReq

		for i := 0; i < len(groupReq); i++ {
			err := groupTable.Create(cacheCurCtx, group.ID(groupReq[i].GroupId).Bytes(), groupReq[i])
			require.NoError(t, err)
		}

		for i := 0; i < len(groupAcc); i++ {
			err := groupAcc[i].SetDecisionPolicy(spec.policy)
			require.NoError(t, err)
		}

		for i := 0; i < len(groupAcc); i++ {
			err = groupAccountTable.Create(cacheCurCtx, groupAcc[i])
			require.NoError(t, err)
		}

		for i := 0; i < len(members); i++ {
			err = groupMemberTable.Create(cacheCurCtx, members[i])
			require.NoError(t, err)
		}

		for i := 0; i < len(proposals); i++ {
			_, err = proposalTable.Create(cacheCurCtx, proposals[i])
			require.NoError(t, err)
		}

		for i := 0; i < len(votes); i++ {
			err = voteTable.Create(cacheCurCtx, votes[i])
			require.NoError(t, err)
		}

		// _, broken, _ := proposalTallyInvariant(cacheCurCtx, proposalTable, groupAccountByGroupIndex, groupMemberByGroupIndex, voteByProposalIndex, proposalByGroupAccountIndex, groupMemberTable, groupAccountTable, groupTable, voteByVoterIndex)
		_, broken, _ := proposalTallyInvariant(cacheCurCtx, groupAccountByGroupIndex, groupMemberByGroupIndex, proposalByGroupAccountIndex, groupTable)
		require.Equal(t, spec.expErr, broken)

	}
}
