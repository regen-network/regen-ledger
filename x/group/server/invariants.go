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
		msg, broken := tallyVotesInvariant(ctx, s.proposalTable)
		return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", msg), broken
	}
}

func tallyVotesInvariant(ctx sdk.Context, proposalTable orm.AutoUInt64Table) (string, bool) {

	var msg string
	var broken bool

	if ctx.BlockHeight()-1 < 0 {
		return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", "Not enough blocks to perform TallyVotesInvariant"), false
	}
	prevCtx := ctx.WithBlockHeight(ctx.BlockHeight() - 1)

	prevIt, err := proposalTable.PrefixScan(prevCtx, 1, math.MaxUint64)
	if err != nil {
		panic(err)
	}
	curIt, err := proposalTable.PrefixScan(ctx, 1, math.MaxUint64)
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
		prevYesCount, err := prevProposals[i].VoteState.GetYesCount()
		if err != nil {
			broken = true
			msg = err.Error()
		}
		curYesCount, err := curProposals[i].VoteState.GetYesCount()
		if err != nil {
			broken = true
			msg = err.Error()
		}
		prevNoCount, err := prevProposals[i].VoteState.GetNoCount()
		if err != nil {
			broken = true
			msg = err.Error()
		}
		curNoCount, err := curProposals[i].VoteState.GetNoCount()
		if err != nil {
			broken = true
			msg = err.Error()
		}
		prevAbstainCount, err := prevProposals[i].VoteState.GetAbstainCount()
		if err != nil {
			broken = true
			msg = err.Error()
		}
		curAbstainCount, err := curProposals[i].VoteState.GetAbstainCount()
		if err != nil {
			broken = true
			msg = err.Error()
		}
		prevVetoCount, err := prevProposals[i].VoteState.GetVetoCount()
		if err != nil {
			broken = true
			msg = err.Error()
		}
		curVetoCount, err := curProposals[i].VoteState.GetVetoCount()
		if err != nil {
			broken = true
			msg = err.Error()
		}
		if (curYesCount.Cmp(prevYesCount) == -1) || (curNoCount.Cmp(prevNoCount) == -1) || (curAbstainCount.Cmp(prevAbstainCount) == -1) || (curVetoCount.Cmp(prevVetoCount) == -1) {
			broken = true
			msg += "vote tally sums must never have less than the block before\n"
		}
	}
	return msg, broken
}
