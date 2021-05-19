package server

import (
	"fmt"
	"math"

	"github.com/cockroachdb/apd/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm"
	regenMath "github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/group"
)

const (
	votesInvariant    = "Tally-Votes"
	weightInvariant   = "Group-TotalWeight"
	votesSumInvariant = "Tally-Votes-Sum"
)

func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(group.ModuleName, votesInvariant, s.tallyVotesInvariant())
	ir.RegisterRoute(group.ModuleName, weightInvariant, s.groupTotalWeightInvariant())
	ir.RegisterRoute(group.ModuleName, votesSumInvariant, s.tallyVotesSumInvariant())
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

func (s serverImpl) tallyVotesSumInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		msg, broken, err := tallyVotesSumInvariant(ctx, s.proposalTable, s.groupMemberTable, s.voteByProposalIndex, s.groupAccountTable)
		if err != nil {
			panic(err)
		}
		return sdk.FormatInvariant(group.ModuleName, votesSumInvariant, msg), broken
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

	for {
		membersWeight := apd.New(0, 0)
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

func tallyVotesSumInvariant(ctx sdk.Context, proposalTable orm.AutoUInt64Table, groupMemberTable orm.PrimaryKeyTable, voteByProposalIndex orm.UInt64Index, groupAccountTable orm.PrimaryKeyTable) (string, bool, error) {
	var msg string
	var broken bool

	var proposal group.Proposal
	var groupAcc group.GroupAccountInfo
	var groupMem group.GroupMember
	var vote group.Vote

	proposalIt, err := proposalTable.PrefixScan(ctx, 1, math.MaxUint64)
	if err != nil {
		return msg, broken, err
	}
	defer proposalIt.Close()

	for {

		totalVotingWeight := apd.New(0, 0)
		yesVoteWeight := apd.New(0, 0)
		noVoteWeight := apd.New(0, 0)
		abstainVoteWeight := apd.New(0, 0)
		vetoVoteWeight := apd.New(0, 0)

		_, err := proposalIt.LoadNext(&proposal)
		if orm.ErrIteratorDone.Is(err) {
			break
		}

		address, err := sdk.AccAddressFromBech32(proposal.Address)
		if err != nil {
			return msg, broken, err
		}

		err = groupAccountTable.GetOne(ctx, address.Bytes(), &groupAcc)
		if err != nil {
			break
		}

		voteIt, err := voteByProposalIndex.Get(ctx, proposal.ProposalId)
		if err != nil {
			return msg, broken, err
		}
		defer voteIt.Close()

		for {
			_, err := voteIt.LoadNext(&vote)
			if orm.ErrIteratorDone.Is(err) {
				break
			}

			groupMem = group.GroupMember{GroupId: groupAcc.GroupId, Member: &group.Member{Address: vote.Voter}}

			err = groupMemberTable.GetOne(ctx, groupMem.PrimaryKey(), &groupMem)
			if err != nil {
				return msg, broken, err
			}

			curMemVotingWeight, err := regenMath.ParseNonNegativeDecimal(groupMem.Member.Weight)
			if err != nil {
				return msg, broken, err
			}
			err = regenMath.Add(totalVotingWeight, totalVotingWeight, curMemVotingWeight)
			if err != nil {
				return msg, broken, err
			}

			switch vote.Choice {
			case group.Choice_CHOICE_YES:
				err = regenMath.Add(yesVoteWeight, yesVoteWeight, curMemVotingWeight)
				if err != nil {
					return msg, broken, err
				}
			case group.Choice_CHOICE_NO:
				err = regenMath.Add(noVoteWeight, noVoteWeight, curMemVotingWeight)
				if err != nil {
					return msg, broken, err
				}
			case group.Choice_CHOICE_ABSTAIN:
				err = regenMath.Add(abstainVoteWeight, abstainVoteWeight, curMemVotingWeight)
				if err != nil {
					return msg, broken, err
				}
			case group.Choice_CHOICE_VETO:
				err = regenMath.Add(vetoVoteWeight, vetoVoteWeight, curMemVotingWeight)
				if err != nil {
					return msg, broken, err
				}
			}
		}

		totalProposalVotes, err := proposal.VoteState.TotalCounts()
		if err != nil {
			return msg, broken, err
		}
		proposalYesCount, err := proposal.VoteState.GetYesCount()
		if err != nil {
			return msg, broken, err
		}
		proposalNoCount, err := proposal.VoteState.GetNoCount()
		if err != nil {
			return msg, broken, err
		}
		proposalAbstainCount, err := proposal.VoteState.GetAbstainCount()
		if err != nil {
			return msg, broken, err
		}
		proposalVetoCount, err := proposal.VoteState.GetVetoCount()
		if err != nil {
			return msg, broken, err
		}

		if totalProposalVotes.Cmp(totalVotingWeight) != 0 {
			broken = true
			msg += "proposal VoteState must correspond to the sum of votes weights\n"
			fmt.Printf("Proposal with ID %d has total proposal votes %s, but got sum of votes weights %s\n", proposal.ProposalId, totalProposalVotes.String(), totalVotingWeight.String())
			break
		}

		if (yesVoteWeight.Cmp(proposalYesCount) != 0) || (noVoteWeight.Cmp(proposalNoCount) != 0) || (abstainVoteWeight.Cmp(proposalAbstainCount) != 0) || (vetoVoteWeight.Cmp(proposalVetoCount) != 0) {
			broken = true
			msg += "proposal VoteState must correspond to the vote choice\n"
			fmt.Printf("Proposal with ID %d and voter address %s must correspond to the vote choice\n", proposal.ProposalId, vote.Voter)
			break
		}
	}
	return msg, broken, err
}
