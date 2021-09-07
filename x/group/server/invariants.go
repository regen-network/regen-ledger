package server

import (
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm"
	regenmath "github.com/regen-network/regen-ledger/types/math"
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
		msg, broken := tallyVotesInvariant(ctx, prevCtx, s.proposalTable)
		return sdk.FormatInvariant(group.ModuleName, votesInvariant, msg), broken
	}
}

func (s serverImpl) groupTotalWeightInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		msg, broken := groupTotalWeightInvariant(ctx, s.groupTable, s.groupMemberByGroupIndex)
		return sdk.FormatInvariant(group.ModuleName, weightInvariant, msg), broken
	}
}

func (s serverImpl) tallyVotesSumInvariant() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		msg, broken := tallyVotesSumInvariant(ctx, s.groupTable, s.proposalTable, s.groupMemberTable, s.voteByProposalIndex, s.groupAccountTable)
		return sdk.FormatInvariant(group.ModuleName, votesSumInvariant, msg), broken
	}
}

func tallyVotesInvariant(ctx sdk.Context, prevCtx sdk.Context, proposalTable orm.AutoUInt64Table) (string, bool) {

	var msg string
	var broken bool

	prevIt, err := proposalTable.PrefixScan(prevCtx, 1, math.MaxUint64)
	if err != nil {
		msg += fmt.Sprintf("PrefixScan failure on proposal table at block height %d\n%v\n", prevCtx.BlockHeight(), err)
		return msg, broken
	}

	curIt, err := proposalTable.PrefixScan(ctx, 1, math.MaxUint64)
	if err != nil {
		msg += fmt.Sprintf("PrefixScan failure on proposal table at block height %d\n%v\n", ctx.BlockHeight(), err)
		return msg, broken
	}

	var curProposals []*group.Proposal
	_, err = orm.ReadAll(curIt, &curProposals)
	if err != nil {
		msg += fmt.Sprintf("error while getting all the proposals at block height %d\n%v\n", ctx.BlockHeight(), err)
		return msg, broken
	}

	var prevProposals []*group.Proposal
	_, err = orm.ReadAll(prevIt, &prevProposals)
	if err != nil {
		msg += fmt.Sprintf("error while getting all the proposals at block height %d\n%v\n", prevCtx.BlockHeight(), err)
		return msg, broken
	}

	for i := 0; i < len(prevProposals); i++ {
		if prevProposals[i].ProposalId == curProposals[i].ProposalId {
			prevYesCount, err := prevProposals[i].VoteState.GetYesCount()
			if err != nil {
				msg += fmt.Sprintf("error while getting yes votes weight of proposal at block height %d\n%v\n", prevCtx.BlockHeight(), err)
				return msg, broken
			}
			curYesCount, err := curProposals[i].VoteState.GetYesCount()
			if err != nil {
				msg += fmt.Sprintf("error while getting yes votes weight of proposal at block height %d\n%v\n", ctx.BlockHeight(), err)
				return msg, broken
			}
			prevNoCount, err := prevProposals[i].VoteState.GetNoCount()
			if err != nil {
				msg += fmt.Sprintf("error while getting no votes weight of proposal at block height %d\n%v\n", prevCtx.BlockHeight(), err)
				return msg, broken
			}
			curNoCount, err := curProposals[i].VoteState.GetNoCount()
			if err != nil {
				msg += fmt.Sprintf("error while getting no votes weight of proposal at block height %d\n%v\n", ctx.BlockHeight(), err)
				return msg, broken
			}
			prevAbstainCount, err := prevProposals[i].VoteState.GetAbstainCount()
			if err != nil {
				msg += fmt.Sprintf("error while getting abstain votes weight of proposal at block height %d\n%v\n", prevCtx.BlockHeight(), err)
				return msg, broken
			}
			curAbstainCount, err := curProposals[i].VoteState.GetAbstainCount()
			if err != nil {
				msg += fmt.Sprintf("error while getting abstain votes weight of proposal at block height %d\n%v\n", ctx.BlockHeight(), err)
				return msg, broken
			}
			prevVetoCount, err := prevProposals[i].VoteState.GetVetoCount()
			if err != nil {
				msg += fmt.Sprintf("error while getting veto votes weight of proposal at block height %d\n%v\n", prevCtx.BlockHeight(), err)
				return msg, broken
			}
			curVetoCount, err := curProposals[i].VoteState.GetVetoCount()
			if err != nil {
				msg += fmt.Sprintf("error while getting veto votes weight of proposal at block height %d\n%v\n", ctx.BlockHeight(), err)
				return msg, broken
			}
			if (curYesCount.Cmp(prevYesCount) == -1) || (curNoCount.Cmp(prevNoCount) == -1) || (curAbstainCount.Cmp(prevAbstainCount) == -1) || (curVetoCount.Cmp(prevVetoCount) == -1) {
				broken = true
				msg += "vote tally sums must never have less than the block before\n"
				return msg, broken
			}
		}
	}
	return msg, broken
}

func groupTotalWeightInvariant(ctx sdk.Context, groupTable orm.AutoUInt64Table, groupMemberByGroupIndex orm.UInt64Index) (string, bool) {

	var msg string
	var broken bool

	var groupInfo group.GroupInfo
	var groupMember group.GroupMember

	groupIt, err := groupTable.PrefixScan(ctx, 1, math.MaxUint64)
	if err != nil {
		msg += fmt.Sprintf("PrefixScan failure on group table\n%v\n", err)
		return msg, broken
	}
	defer groupIt.Close()

	for {
		membersWeight := regenmath.NewDecFromInt64(0)
		_, err := groupIt.LoadNext(&groupInfo)
		if orm.ErrIteratorDone.Is(err) {
			break
		}
		memIt, err := groupMemberByGroupIndex.Get(ctx, groupInfo.GroupId)
		if err != nil {
			msg += fmt.Sprintf("error while returning group member iterator for group with ID %d\n%v\n", groupInfo.GroupId, err)
			return msg, broken
		}
		defer memIt.Close()

		for {
			_, err = memIt.LoadNext(&groupMember)
			if orm.ErrIteratorDone.Is(err) {
				break
			}
			curMemWeight, err := regenmath.NewNonNegativeDecFromString(groupMember.GetMember().GetWeight())
			if err != nil {
				msg += fmt.Sprintf("error while parsing non-nengative decimal for group member %s\n%v\n", groupMember.Member.Address, err)
				return msg, broken
			}
			membersWeight, err = membersWeight.Add(curMemWeight)
			if err != nil {
				msg += fmt.Sprintf("decimal addition error while adding group member voting weight to total voting weight\n%v\n", err)
				return msg, broken
			}
		}
		groupWeight, err := regenmath.NewNonNegativeDecFromString(groupInfo.GetTotalWeight())
		if err != nil {
			msg += fmt.Sprintf("error while parsing non-nengative decimal for group with ID %d\n%v\n", groupInfo.GroupId, err)
			return msg, broken
		}

		if groupWeight.Cmp(membersWeight) != 0 {
			broken = true
			msg += fmt.Sprintf("group's TotalWeight must be equal to the sum of its members' weights\ngroup weight: %s\nSum of group members weights: %s\n", groupWeight.String(), membersWeight.String())
			break
		}
	}
	return msg, broken
}

func tallyVotesSumInvariant(ctx sdk.Context, groupTable orm.AutoUInt64Table, proposalTable orm.AutoUInt64Table, groupMemberTable orm.PrimaryKeyTable, voteByProposalIndex orm.UInt64Index, groupAccountTable orm.PrimaryKeyTable) (string, bool) {
	var msg string
	var broken bool

	var groupInfo group.GroupInfo
	var proposal group.Proposal
	var groupAcc group.GroupAccountInfo
	var groupMem group.GroupMember
	var vote group.Vote

	proposalIt, err := proposalTable.PrefixScan(ctx, 1, math.MaxUint64)
	if err != nil {
		msg += fmt.Sprintf("PrefixScan failure on proposal table\n%v\n", err)
		return msg, broken
	}
	defer proposalIt.Close()

	for {

		totalVotingWeight := regenmath.NewDecFromInt64(0)
		yesVoteWeight := regenmath.NewDecFromInt64(0)
		noVoteWeight := regenmath.NewDecFromInt64(0)
		abstainVoteWeight := regenmath.NewDecFromInt64(0)
		vetoVoteWeight := regenmath.NewDecFromInt64(0)

		_, err := proposalIt.LoadNext(&proposal)
		if orm.ErrIteratorDone.Is(err) {
			break
		}

		address, err := sdk.AccAddressFromBech32(proposal.Address)
		if err != nil {
			msg += fmt.Sprintf("error while converting proposal address of type string to type AccAddress\n%v\n", err)
			return msg, broken
		}

		err = groupAccountTable.GetOne(ctx, orm.AddLengthPrefix(address.Bytes()), &groupAcc)
		if err != nil {
			msg += fmt.Sprintf("group account not found for address: %s\n%v\n", proposal.Address, err)
			return msg, broken
		}

		if proposal.GroupAccountVersion != groupAcc.Version {
			msg += fmt.Sprintf("group account with address %s was modified\n", groupAcc.Address)
			return msg, broken
		}

		_, err = groupTable.GetOne(ctx, groupAcc.GroupId, &groupInfo)
		if err != nil {
			msg += fmt.Sprintf("group info not found for group id %d\n%v\n", groupAcc.GroupId, err)
			return msg, broken
		}

		if groupInfo.Version != proposal.GroupVersion {
			msg += fmt.Sprintf("group with id %d was modified\n", groupInfo.GroupId)
			return msg, broken
		}

		voteIt, err := voteByProposalIndex.Get(ctx, proposal.ProposalId)
		if err != nil {
			msg += fmt.Sprintf("error while returning vote iterator for proposal with ID %d\n%v\n", proposal.ProposalId, err)
			return msg, broken
		}
		defer voteIt.Close()

		for {
			_, err := voteIt.LoadNext(&vote)
			if orm.ErrIteratorDone.Is(err) {
				break
			}

			groupMem = group.GroupMember{GroupId: groupAcc.GroupId, Member: &group.Member{Address: vote.Voter}}

			err = groupMemberTable.GetOne(ctx, orm.PrimaryKey(&groupMem), &groupMem)
			if err != nil {
				msg += fmt.Sprintf("group member not found with group ID %d and group member %s\n%v\n", groupAcc.GroupId, vote.Voter, err)
				return msg, broken
			}

			curMemVotingWeight, err := regenmath.NewNonNegativeDecFromString(groupMem.Member.Weight)
			if err != nil {
				msg += fmt.Sprintf("error while parsing non-negative decimal for group member %s\n%v\n", groupMem.Member.Address, err)
				return msg, broken
			}
			totalVotingWeight, err = totalVotingWeight.Add(curMemVotingWeight)
			if err != nil {
				msg += fmt.Sprintf("decimal addition error while adding current member voting weight to total voting weight\n%v\n", err)
				return msg, broken
			}

			switch vote.Choice {
			case group.Choice_CHOICE_YES:
				yesVoteWeight, err = yesVoteWeight.Add(curMemVotingWeight)
				if err != nil {
					msg += fmt.Sprintf("decimal addition error\n%v\n", err)
					return msg, broken
				}
			case group.Choice_CHOICE_NO:
				noVoteWeight, err = noVoteWeight.Add(curMemVotingWeight)
				if err != nil {
					msg += fmt.Sprintf("decimal addition error\n%v\n", err)
					return msg, broken
				}
			case group.Choice_CHOICE_ABSTAIN:
				abstainVoteWeight, err = abstainVoteWeight.Add(curMemVotingWeight)
				if err != nil {
					msg += fmt.Sprintf("decimal addition error\n%v\n", err)
					return msg, broken
				}
			case group.Choice_CHOICE_VETO:
				vetoVoteWeight, err = vetoVoteWeight.Add(curMemVotingWeight)
				if err != nil {
					msg += fmt.Sprintf("decimal addition error\n%v\n", err)
					return msg, broken
				}
			}
		}

		totalProposalVotes, err := proposal.VoteState.TotalCounts()
		if err != nil {
			msg += fmt.Sprintf("error while getting total weighted votes of proposal with ID %d\n%v\n", proposal.ProposalId, err)
			return msg, broken
		}
		proposalYesCount, err := proposal.VoteState.GetYesCount()
		if err != nil {
			msg += fmt.Sprintf("error while getting the weighted sum of yes votes for proposal with ID %d\n%v\n", proposal.ProposalId, err)
			return msg, broken
		}
		proposalNoCount, err := proposal.VoteState.GetNoCount()
		if err != nil {
			msg += fmt.Sprintf("error while getting the weighted sum of no votes for proposal with ID %d\n%v\n", proposal.ProposalId, err)
			return msg, broken
		}
		proposalAbstainCount, err := proposal.VoteState.GetAbstainCount()
		if err != nil {
			msg += fmt.Sprintf("error while getting the weighted sum of abstain votes for proposal with ID %d\n%v\n", proposal.ProposalId, err)
			return msg, broken
		}
		proposalVetoCount, err := proposal.VoteState.GetVetoCount()
		if err != nil {
			msg += fmt.Sprintf("error while getting the weighted sum of veto votes for proposal with ID %d\n%v\n", proposal.ProposalId, err)
			return msg, broken
		}

		if totalProposalVotes.Cmp(totalVotingWeight) != 0 {
			broken = true
			msg += fmt.Sprintf("proposal VoteState must correspond to the sum of votes weights\nProposal with ID %d has total proposal votes %s, but got sum of votes weights %s\n", proposal.ProposalId, totalProposalVotes.String(), totalVotingWeight.String())
			break
		}

		if (yesVoteWeight.Cmp(proposalYesCount) != 0) || (noVoteWeight.Cmp(proposalNoCount) != 0) || (abstainVoteWeight.Cmp(proposalAbstainCount) != 0) || (vetoVoteWeight.Cmp(proposalVetoCount) != 0) {
			broken = true
			msg += fmt.Sprintf("proposal VoteState must correspond to the vote choice\nProposal with ID %d and voter address %s must correspond to the vote choice\n", proposal.ProposalId, vote.Voter)
			break
		}
	}
	return msg, broken
}
