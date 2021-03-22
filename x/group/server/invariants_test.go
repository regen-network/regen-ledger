package server

import (
	"math"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/group"
	// "github.com/regen-network/regen-ledger/types"
)

func TestTallyVotesInvariant(t *testing.T) {
	ff := server.NewFixtureFactory(t, 6)

	sdkCtx, _ := ff.Setup().Context().(types.Context)
	sdkCtx1, _ := sdkCtx.CacheContext()
	sdkCtx1 = sdkCtx1.WithBlockHeight(10)
	// var ctx sdk.Context
	// ctx := ff.Setup().Context()
	// blockHeight := ctx.BlockHeight()
	var proposalTable orm.AutoUInt64Table
	// ctx1 := types.Context{Context: sdkCtx1}
	proposalIterator, err := proposalTable.PrefixScan(sdkCtx1, 1, math.MaxUint64)
	if err != nil {
		panic(err)
	}
	var test require.TestingT
	var curProposals *group.Proposal
	_, err = orm.ReadAll(proposalIterator, &curProposals)
	require.NoError(test, err, &curProposals)
	// invar := Invar{
	// 	sdkCtx:        ctx, // types.Context{Context: ctx},
	// 	proposalTable: proposalTable,
	// }
	// curYesCount, err := curProposals.VoteState.GetYesCount()
	// curNoCount, err := curProposals.VoteState.GetNoCount()
	// curAbstainCount, err := curProposals.VoteState.GetAbstainCount()
	// curVetoCount, err := curProposals.VoteState.GetVetoCount()

	_, _, addr := testdata.KeyTestPubAddr()
	groupAddr := addr.String()

	_, _, addr = testdata.KeyTestPubAddr()
	memberAddr := addr.String()

	specs := map[string]struct {
		src    *group.MsgCreateProposalRequest
		expErr bool
	}{
		"all good with minimum fields set": {
			src: &group.MsgCreateProposalRequest{
				GroupAccount: groupAddr,
				Proposers:    []string{memberAddr},
			},
		},
		"group account required": {
			src: &group.MsgCreateProposalRequest{
				Proposers: []string{memberAddr},
			},
			expErr: true,
		},
		"proposers required": {
			src: &group.MsgCreateProposalRequest{
				GroupAccount: groupAddr,
			},
			expErr: true,
		},
		"valid proposer address required": {
			src: &group.MsgCreateProposalRequest{
				GroupAccount: groupAddr,
				Proposers:    []string{"invalid-member-address"},
			},
			expErr: true,
		},
		"no duplicate proposers": {
			src: &group.MsgCreateProposalRequest{
				GroupAccount: groupAddr,
				Proposers:    []string{memberAddr, memberAddr},
			},
			expErr: true,
		},
		"empty proposer address not allowed": {
			src: &group.MsgCreateProposalRequest{
				GroupAccount: groupAddr,
				Proposers:    []string{memberAddr, ""},
			},
			expErr: true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := spec.src.ValidateBasic()
			if spec.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

	specs2 := map[string]struct {
		src    *group.MsgVoteRequest
		expErr bool
	}{
		"all good with minimum fields set": {
			src: &group.MsgVoteRequest{
				ProposalId: 1,
				Choice:     group.Choice_CHOICE_YES,
				Voter:      memberAddr,
			},
		},
		"proposal required": {
			src: &group.MsgVoteRequest{
				Choice: group.Choice_CHOICE_YES,
				Voter:  memberAddr,
			},
			expErr: true,
		},
		"choice required": {
			src: &group.MsgVoteRequest{
				ProposalId: 1,
				Voter:      memberAddr,
			},
			expErr: true,
		},
		"valid choice required": {
			src: &group.MsgVoteRequest{
				ProposalId: 1,
				Choice:     5,
				Voter:      memberAddr,
			},
			expErr: true,
		},
		"voter required": {
			src: &group.MsgVoteRequest{
				ProposalId: 1,
				Choice:     group.Choice_CHOICE_YES,
			},
			expErr: true,
		},
		"valid voter address required": {
			src: &group.MsgVoteRequest{
				ProposalId: 1,
				Choice:     group.Choice_CHOICE_YES,
				Voter:      "invalid-member-address",
			},
			expErr: true,
		},
		"empty voters address not allowed": {
			src: &group.MsgVoteRequest{
				ProposalId: 1,
				Choice:     group.Choice_CHOICE_YES,
				Voter:      "",
			},
			expErr: true,
		},
	}
	for msg, spec := range specs2 {
		t.Run(msg, func(t *testing.T) {
			err := spec.src.ValidateBasic()
			if spec.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			// ch := spec.src.Choice
			// a := group.Choice_value
			//group.Choice_value(spec.src.Choice)

			// vote := spec.src.GetChoice().String()

		})
	}
	// tallyVotesInvariant(invar)

}
