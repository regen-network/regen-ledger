package server

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(group.ModuleName, "Tally-Votes", s.tallyVotesInvariant())
}

func (s serverImpl) tallyVotesInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if ctx.BlockHeight()-1 < 0 {
			return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", "Not enough blocks to perform TallyVotesInvariant"), false
		}
		prevCtx, _ := ctx.CacheContext()
		prevCtx = prevCtx.WithBlockHeight(ctx.BlockHeight() - 1)
		msg, broken, err := tallyVotesInvariant(ctx, prevCtx, s.proposalTable)
		if err != nil {
			panic(err)
		}
		return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", msg), broken
	}
}

func tallyVotesInvariant(ctx sdk.Context, prevCtx sdk.Context, proposalTable orm.AutoUInt64Table) (string, bool, error) {

	var msg string
	var broken bool

	prevIt, err := proposalTable.PrefixScan(prevCtx, 1, math.MaxUint64)
	if err != nil {
		return msg, broken, err
	}
	curIt, err := proposalTable.PrefixScan(ctx, 1, math.MaxUint64)
	if err != nil {
		return msg, broken, err
	}

	var curProposals []*group.Proposal
	_, err = orm.ReadAll(curIt, &curProposals)
	if err != nil {
		return msg, broken, err
	}

	var prevProposals []*group.Proposal
	_, err = orm.ReadAll(prevIt, &prevProposals)
	if err != nil {
		return msg, broken, err
	}

	for i := 0; i < len(prevProposals) && i < len(curProposals); i++ {
		prevYesCount, err := prevProposals[i].VoteState.GetYesCount()
		if err != nil {
			return msg, broken, err
		}
		curYesCount, err := curProposals[i].VoteState.GetYesCount()
		if err != nil {
			return msg, broken, err
		}
		prevNoCount, err := prevProposals[i].VoteState.GetNoCount()
		if err != nil {
			return msg, broken, err
		}
		curNoCount, err := curProposals[i].VoteState.GetNoCount()
		if err != nil {
			return msg, broken, err
		}
		prevAbstainCount, err := prevProposals[i].VoteState.GetAbstainCount()
		if err != nil {
			return msg, broken, err
		}
		curAbstainCount, err := curProposals[i].VoteState.GetAbstainCount()
		if err != nil {
			return msg, broken, err
		}
		prevVetoCount, err := prevProposals[i].VoteState.GetVetoCount()
		if err != nil {
			return msg, broken, err
		}
		curVetoCount, err := curProposals[i].VoteState.GetVetoCount()
		if err != nil {
			return msg, broken, err
		}
		if (curYesCount.Cmp(prevYesCount) == -1) || (curNoCount.Cmp(prevNoCount) == -1) || (curAbstainCount.Cmp(prevAbstainCount) == -1) || (curVetoCount.Cmp(prevVetoCount) == -1) {
			broken = true
			msg += "vote tally sums must never have less than the block before\n"
			return msg, broken, err
		}
	}
	return msg, broken, err
}
