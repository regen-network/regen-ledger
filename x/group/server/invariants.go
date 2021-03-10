package server

import (
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) RegisterInvariantsHandler(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(group.ModuleName, "Tally-Votes", s.TallyVotesInvariant())
}

func (s serverImpl) AllInvariants() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return s.TallyVotesInvariant()(ctx)
	}
}

func (s serverImpl) TallyVotesInvariant() sdk.Invariant {
	return func(sdkCtx sdk.Context) (string, bool) {
		var msg string
		var broken bool
		ctx := types.Context{Context: sdkCtx}
		if ctx.BlockHeight()-1 < 0 {
			return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", "Not enough blocks to perform TallyVotesInvariant"), false
		}
		sdkCtx = sdkCtx.WithBlockHeight(ctx.BlockHeight() - 1)
		it1, err := s.proposalTable.PrefixScan(ctx, 1, math.MaxUint64)
		if err != nil {
			panic(err)
		}
		it2, err := s.proposalTable.PrefixScan(sdkCtx, 1, math.MaxUint64)
		if err != nil {
			panic(err)
		}
		var t require.TestingT
		var curProposals []*group.Proposal
		curProposalRowID, err := orm.ReadAll(it1, &curProposals)
		require.NoError(t, err, &curProposals)

		_ = curProposalRowID
		var prevProposals []*group.Proposal
		prevProposalRowID, err := orm.ReadAll(it2, &prevProposals)
		require.NoError(t, err, &curProposals)

		_ = prevProposalRowID

		for i := 0; i < len(prevProposals) && i < len(curProposals); i++ {
			prevYesCount, err := prevProposals[i].VoteState.GetYesCount()
			if err != nil {
				return fmt.Sprint(err), false
			}
			curYesCount, err := curProposals[i].VoteState.GetYesCount()
			if err != nil {
				return fmt.Sprint(err), false
			}
			prevNoCount, err := prevProposals[i].VoteState.GetNoCount()
			if err != nil {
				return fmt.Sprint(err), false
			}
			curNoCount, err := curProposals[i].VoteState.GetNoCount()
			if err != nil {
				return fmt.Sprint(err), false
			}
			prevAbstainCount, err := prevProposals[i].VoteState.GetAbstainCount()
			if err != nil {
				return fmt.Sprint(err), false
			}
			curAbstainCount, err := curProposals[i].VoteState.GetAbstainCount()
			if err != nil {
				return fmt.Sprint(err), false
			}
			prevVetoCount, err := prevProposals[i].VoteState.GetVetoCount()
			if err != nil {
				return fmt.Sprint(err), false
			}
			curVetoCount, err := curProposals[i].VoteState.GetVetoCount()
			if err != nil {
				return fmt.Sprint(err), false
			}
			if (curYesCount.Cmp(prevYesCount) == -1) || (curNoCount.Cmp(prevNoCount) == -1) || (curAbstainCount.Cmp(prevAbstainCount) == -1) || (curVetoCount.Cmp(prevVetoCount) == -1) {
				broken = true
				msg += "vote tally sums must never have less than the block before\n"
			}
		}
		return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", msg), broken
	}
}
