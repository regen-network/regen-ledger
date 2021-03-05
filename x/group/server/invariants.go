package server

import (
	"fmt"
	"math"

	"github.com/regen-network/regen-ledger/orm"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
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
		var curProposal []*group.Proposal
		ctx := types.Context{Context: sdkCtx}
		curIterator, err := s.proposalTable.PrefixScan(ctx, 1, math.MaxUint64)
		if err != nil {
			return "start value must be less than end value in iterator", false
		}
		curProposalRowID, err := orm.ReadAll(curIterator, &curProposal)
		if err != nil {
			return "cannot read all proposals in current block", false
		}
		_ = curProposalRowID
		var prevProposal []*group.Proposal
		if ctx.BlockHeight()-1 >= 0 {
			sdkCtx = sdkCtx.WithBlockHeight(ctx.BlockHeight() - 1)
		} else {
			return "Not enough blocks to perform TallyVotesInvariant", false
		}
		prevIterator, err := s.proposalTable.PrefixScan(sdkCtx, 1, math.MaxUint64)
		if err != nil {
			return "start value must be less than end value in iterator", false
		}
		prevProposalRowID, err := orm.ReadAll(prevIterator, &prevProposal)
		if err != nil {
			return "cannot read all proposals from previous block", false
		}
		_ = prevProposalRowID
		for i := 0; i < len(prevProposal) && i < len(curProposal); i++ {
			if int32(prevProposal[i].Status) == 1 && int32(curProposal[i].Status) == 1 {
				prevYesCount, err := prevProposal[i].VoteState.GetYesCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				curYesCount, err := curProposal[i].VoteState.GetYesCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				prevNoCount, err := prevProposal[i].VoteState.GetNoCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				curNoCount, err := curProposal[i].VoteState.GetNoCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				prevAbstainCount, err := prevProposal[i].VoteState.GetAbstainCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				curAbstainCount, err := curProposal[i].VoteState.GetAbstainCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				prevVetoCount, err := prevProposal[i].VoteState.GetVetoCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				curVetoCount, err := curProposal[i].VoteState.GetVetoCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				if (curYesCount.Cmp(prevYesCount) == -1) || (curNoCount.Cmp(prevNoCount) == -1) || (curAbstainCount.Cmp(prevAbstainCount) == -1) || (curVetoCount.Cmp(prevVetoCount) == -1) {
					broken = true
					msg += "vote tally sums must never have less than the block before\n"
				}
			}
		}
		return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", msg), broken
	}
}
