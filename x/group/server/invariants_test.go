package server

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
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

type invar struct {
	s   serverImpl
	cdc codec.Marshaler
}

func TestTallyVotesInvariant(t *testing.T) {
	var in invar
	key := sdk.NewKVStoreKey(group.ModuleName)
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
	ctx = ctx.WithBlockHeight(10)

	// Proposal Table
	proposalTableBuilder := orm.NewAutoUInt64TableBuilder(ProposalTablePrefix, ProposalTableSeqPrefix, key, &group.Proposal{}, in.cdc)
	in.s.proposalByGroupAccountIndex = orm.NewIndex(proposalTableBuilder, ProposalByGroupAccountIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		account := value.(*group.Proposal).GroupAccount
		addr, err := sdk.AccAddressFromBech32(account)
		if err != nil {
			return nil, err
		}
		return []orm.RowID{addr.Bytes()}, nil
	})
	in.s.proposalByProposerIndex = orm.NewIndex(proposalTableBuilder, ProposalByProposerIndexPrefix, func(value interface{}) ([]orm.RowID, error) {
		proposers := value.(*group.Proposal).Proposers
		r := make([]orm.RowID, len(proposers))
		for i := range proposers {
			addr, err := sdk.AccAddressFromBech32(proposers[i])
			if err != nil {
				return nil, err
			}
			r[i] = addr.Bytes()
		}
		return r, nil
	})
	in.s.proposalTable = proposalTableBuilder.Build()

	_, _, addr1 := testdata.KeyTestPubAddr()
	// addr1 := sdk.AccAddress("foo_________________")

	blockTime, err := gogotypes.TimestampProto(ctx.BlockTime())
	if err != nil {
		fmt.Println("block time conversion")
		panic(err)
	}

	m1 := &group.Proposal{
		GroupAccount:        addr1.String(),
		Proposers:           []string{addr1.String()},
		SubmittedAt:         *blockTime,
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

	_, err = in.s.proposalTable.Create(ctx, m1)
	if err != nil {
		fmt.Println(err)
		panic("create proposal")
	}

	fmt.Println("It worked")
	panic("")
}
