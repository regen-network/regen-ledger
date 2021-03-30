package server

import (
	"fmt"
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

func TestTallyVotesInvariant(t *testing.T) {
	interfaceRegistry := types.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	key := sdk.NewKVStoreKey(group.ModuleName)
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}
	curCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	curCtx = curCtx.WithBlockHeight(10)
	prevCtx := curCtx.WithBlockHeight(curCtx.BlockHeight() - 1)

	// Proposal Table
	proposalTableBuilder := orm.NewAutoUInt64TableBuilder(ProposalTablePrefix, ProposalTableSeqPrefix, key, &group.Proposal{}, cdc)
	proposalTable := proposalTableBuilder.Build()

	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	curBlockTime, err := gogotypes.TimestampProto(curCtx.BlockTime())
	if err != nil {
		fmt.Println("block time conversion")
		panic(err)
	}
	prevBlockTime, err := gogotypes.TimestampProto(prevCtx.BlockTime())
	if err != nil {
		fmt.Println("block time conversion")
		panic(err)
	}

	specs := map[string]struct {
		prevReq *group.Proposal
		curReq  *group.Proposal
		expErr  bool
	}{
		"invariant not broken": {
			prevReq: &group.Proposal{
				GroupAccount:        addr1.String(),
				Proposers:           []string{addr1.String()},
				SubmittedAt:         *prevBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Result:              group.ProposalResultUnfinalized,
				Status:              group.ProposalStatusSubmitted,
				ExecutorResult:      group.ProposalExecutorResultNotRun,
				Timeout: gogotypes.Timestamp{
					Seconds: 600,
				},
				VoteState: group.Tally{
					YesCount:     "1",
					NoCount:      "0",
					AbstainCount: "0",
					VetoCount:    "0",
				},
			},
			curReq: &group.Proposal{
				GroupAccount:        addr2.String(),
				Proposers:           []string{addr2.String()},
				SubmittedAt:         *curBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Result:              group.ProposalResultUnfinalized,
				Status:              group.ProposalStatusSubmitted,
				ExecutorResult:      group.ProposalExecutorResultNotRun,
				Timeout: gogotypes.Timestamp{
					Seconds: 600,
				},
				VoteState: group.Tally{
					YesCount:     "2",
					NoCount:      "0",
					AbstainCount: "0",
					VetoCount:    "0",
				},
			},
		},
		"invariant broken": {
			prevReq: &group.Proposal{
				GroupAccount:        addr1.String(),
				Proposers:           []string{addr1.String()},
				SubmittedAt:         *prevBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Result:              group.ProposalResultUnfinalized,
				Status:              group.ProposalStatusSubmitted,
				ExecutorResult:      group.ProposalExecutorResultNotRun,
				Timeout: gogotypes.Timestamp{
					Seconds: 600,
				},
				VoteState: group.Tally{
					YesCount:     "2",
					NoCount:      "0",
					AbstainCount: "0",
					VetoCount:    "0",
				},
			},
			curReq: &group.Proposal{
				GroupAccount:        addr2.String(),
				Proposers:           []string{addr2.String()},
				SubmittedAt:         *curBlockTime,
				GroupVersion:        1,
				GroupAccountVersion: 1,
				Result:              group.ProposalResultUnfinalized,
				Status:              group.ProposalStatusSubmitted,
				ExecutorResult:      group.ProposalExecutorResultNotRun,
				Timeout: gogotypes.Timestamp{
					Seconds: 600,
				},
				VoteState: group.Tally{
					YesCount:     "1",
					NoCount:      "0",
					AbstainCount: "0",
					VetoCount:    "0",
				},
			},
			expErr: true,
		},
	}
	for _, spec := range specs {
		spec := spec
		prevProposal := spec.prevReq
		curProposal := spec.curReq

		_, err = proposalTable.Create(prevCtx, prevProposal)
		if err != nil {
			fmt.Println(err)
			panic("create proposal")
		}

		_, err = proposalTable.Create(curCtx, curProposal)
		if err != nil {
			fmt.Println(err)
			panic("create proposal")
		}

		var test require.TestingT
		_, broken := tallyVotesInvariant(prevProposal, curProposal)
		require.Equal(test, spec.expErr, broken)
	}
}
