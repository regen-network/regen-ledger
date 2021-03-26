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
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group"
)

func TestTallyVotesInvariant(t *testing.T) {
	var s serverImpl
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
	s.proposalTable = proposalTableBuilder.Build()

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

	prevProposal := &group.Proposal{
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
	}

	curProposal := &group.Proposal{
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
	}

	_, err = s.proposalTable.Create(prevCtx, prevProposal)
	if err != nil {
		fmt.Println(err)
		panic("create proposal")
	}

	_, err = s.proposalTable.Create(curCtx, curProposal)
	if err != nil {
		fmt.Println(err)
		panic("create proposal")
	}

	msg, broken := tallyVotesInvariant(prevProposal, curProposal)
	fmt.Println(msg, broken)
	if broken == true {
		panic("Invariant broken")
	}
}
