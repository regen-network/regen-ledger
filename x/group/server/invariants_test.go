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

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group"
)

func getCtxCodecKey(t *testing.T) (sdk.Context, *codec.ProtoCodec, *sdk.KVStoreKey) {
	interfaceRegistry := types.NewInterfaceRegistry()
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
					ProposalId:          0,
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
					ProposalId:          0,
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
					ProposalId:          0,
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
					ProposalId:          0,
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
					ProposalId:          0,
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
					ProposalId:          0,
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
					ProposalId:          0,
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
					ProposalId:          0,
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
					ProposalId:          0,
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
					ProposalId:          0,
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
