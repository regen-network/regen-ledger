package server

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(group.ModuleName, "Tally-Votes", s.tallyVotesInvariant())
}

func (s serverImpl) AllInvariants() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return s.tallyVotesInvariant()(ctx)
	}
}

func (s serverImpl) tallyVotesInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var msg string
		var broken bool
		if ctx.BlockHeight()-1 < 0 {
			return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", "Not enough blocks to perform TallyVotesInvariant"), false
		}
		prevCtx := ctx.WithBlockHeight(ctx.BlockHeight() - 1)
		prevIt, err := s.proposalTable.PrefixScan(prevCtx, 1, math.MaxUint64)
		if err != nil {
			panic(err)
		}
		curIt, err := s.proposalTable.PrefixScan(ctx, 1, math.MaxUint64)
		if err != nil {
			panic(err)
		}

		var t require.TestingT
		var curProposals []*group.Proposal
		_, err = orm.ReadAll(curIt, &curProposals)
		require.NoError(t, err, &curProposals)

		var prevProposals []*group.Proposal
		_, err = orm.ReadAll(prevIt, &prevProposals)
		require.NoError(t, err, &curProposals)

		for i := 0; i < len(prevProposals) && i < len(curProposals); i++ {
			msg, broken = tallyVotesInvariant(prevProposals[i], curProposals[i])
		}
		return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", msg), broken
	}
}

func tallyVotesInvariant(prevProposal *group.Proposal, curProposal *group.Proposal) (string, bool) {

	var msg string
	var broken bool
	prevYesCount, err := prevProposal.VoteState.GetYesCount()
	if err != nil {
		panic(err)
	}
	curYesCount, err := curProposal.VoteState.GetYesCount()
	if err != nil {
		panic(err)
	}
	prevNoCount, err := prevProposal.VoteState.GetNoCount()
	if err != nil {
		panic(err)
	}
	curNoCount, err := curProposal.VoteState.GetNoCount()
	if err != nil {
		panic(err)
	}
	prevAbstainCount, err := prevProposal.VoteState.GetAbstainCount()
	if err != nil {
		panic(err)
	}
	curAbstainCount, err := curProposal.VoteState.GetAbstainCount()
	if err != nil {
		panic(err)
	}
	prevVetoCount, err := prevProposal.VoteState.GetVetoCount()
	if err != nil {
		panic(err)
	}
	curVetoCount, err := curProposal.VoteState.GetVetoCount()
	if err != nil {
		panic(err)
	}
	if (curYesCount.Cmp(prevYesCount) == -1) || (curNoCount.Cmp(prevNoCount) == -1) || (curAbstainCount.Cmp(prevAbstainCount) == -1) || (curVetoCount.Cmp(prevVetoCount) == -1) {
		broken = true
		msg += "vote tally sums must never have less than the block before\n"

	}
	return msg, broken
}
