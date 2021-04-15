package server

import (
	"math"

	"github.com/cockroachdb/apd/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"

	regenMath "github.com/regen-network/regen-ledger/math"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/x/group"
)

const (
	votesInvariant  = "Tally-Votes"
	weightInvariant = "Group-TotalWeight"
)

func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(group.ModuleName, votesInvariant, s.tallyVotesInvariant())
	ir.RegisterRoute(group.ModuleName, weightInvariant, s.groupTotalWeightInvariant())
}

func (s serverImpl) tallyVotesInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if ctx.BlockHeight()-1 < 0 {
			return sdk.FormatInvariant(group.ModuleName, votesInvariant, "Not enough blocks to perform TallyVotesInvariant"), false
		}
		prevCtx, _ := ctx.CacheContext()
		prevCtx = prevCtx.WithBlockHeight(ctx.BlockHeight() - 1)
		msg, broken, err := tallyVotesInvariant(ctx, prevCtx, s.proposalTable)
		if err != nil {
			panic(err)
		}
		return sdk.FormatInvariant(group.ModuleName, votesInvariant, msg), broken
	}
}

func (s serverImpl) groupTotalWeightInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		msg, broken, err := groupTotalWeightInvariant(ctx, s.groupTable, s.groupMemberByGroupIndex)
		if err != nil {
			panic(err)
		}
		return sdk.FormatInvariant(group.ModuleName, weightInvariant, msg), broken
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

	for i := 0; i < len(prevProposals); i++ {
		if prevProposals[i].ProposalId == curProposals[i].ProposalId {
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
	}
	return msg, broken, err
}

func groupTotalWeightInvariant(ctx sdk.Context, groupTable orm.Table, groupMemberByGroupIndex orm.UInt64Index) (string, bool, error) {

	var msg string
	var broken bool

	var groupInfo group.GroupInfo
	var groupMember group.GroupMember

	groupIt, err := groupTable.PrefixScan(ctx, nil, nil)
	if err != nil {
		return msg, broken, err
	}
	defer groupIt.Close()

	membersWeight := apd.New(0, 0)

	for {
		_, err := groupIt.LoadNext(&groupInfo)
		if orm.ErrIteratorDone.Is(err) {
			break
		}
		memIt, err := groupMemberByGroupIndex.Get(ctx, groupInfo.GroupId)
		if err != nil {
			return msg, broken, err
		}
		defer memIt.Close()

		for {
			_, err = memIt.LoadNext(&groupMember)
			if orm.ErrIteratorDone.Is(err) {
				break
			}
			curMemWeight, err := regenMath.ParseNonNegativeDecimal(groupMember.GetMember().GetWeight())
			if err != nil {
				return msg, broken, err
			}
			err = regenMath.Add(membersWeight, membersWeight, curMemWeight)
			if err != nil {
				return msg, broken, err
			}
		}
		groupWeight, err := regenMath.ParseNonNegativeDecimal(groupInfo.GetTotalWeight())
		if err != nil {
			return msg, broken, err
		}
		if groupWeight.Cmp(membersWeight) != 0 {
			broken = true
			msg += "group's TotalWeight must be equal to the sum of its members' weights\n"
			break
		}
	}
	return msg, broken, err
}
