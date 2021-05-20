package simulation

import (
	"math/rand"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/regen-network/regen-ledger/x/group"
)

const (
	GroupInfo        = "group-info"
	GroupMembers     = "group-members"
	GroupAccountInfo = "group-accout-info"
	GroupProposals   = "group-proposals"
	GroupVote        = "group-vote"
)

func getGroups(r *rand.Rand, accounts []simtypes.Account) []*group.GroupInfo {
	groups := make([]*group.GroupInfo, 3)
	for i := 0; i < 3; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		groups[i] = &group.GroupInfo{
			GroupId:     uint64(i + 1),
			Admin:       acc.Address.String(),
			Metadata:    []byte(simtypes.RandStringOfLength(r, 10)),
			Version:     1,
			TotalWeight: "10",
		}
	}
	return groups
}

func getGroupMembers(r *rand.Rand, accounts []simtypes.Account) []*group.GroupMember {
	groupMembers := make([]*group.GroupMember, 3)
	for i := 0; i < 3; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		groupMembers[i] = &group.GroupMember{
			GroupId: uint64(i + 1),
			Member: &group.Member{
				Address:  acc.Address.String(),
				Weight:   "10",
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
		}
	}
	return groupMembers
}

func getGroupAccounts(r *rand.Rand, simState *module.SimulationState) []*group.GroupAccountInfo {
	groupMembers := make([]*group.GroupAccountInfo, 3)
	for i := 0; i < 3; i++ {
		acc, _ := simtypes.RandomAcc(r, simState.Accounts)
		any, err := codectypes.NewAnyWithValue(group.NewThresholdDecisionPolicy("10", gogotypes.Duration{Seconds: 1}))
		if err != nil {
			panic(err)
		}
		groupMembers[i] = &group.GroupAccountInfo{
			GroupId:        uint64(i + 1),
			Admin:          acc.Address.String(),
			Address:        acc.Address.String(),
			Version:        1,
			DecisionPolicy: any,
			Metadata:       []byte(simtypes.RandStringOfLength(r, 10)),
		}
	}
	return groupMembers
}

func getProposals(r *rand.Rand, simState *module.SimulationState) []*group.Proposal {
	proposals := make([]*group.Proposal, 3)
	proposers := []string{simState.Accounts[0].Address.String(), simState.Accounts[1].Address.String()}
	for i := 0; i < 3; i++ {
		from, _ := simtypes.RandomAcc(r, simState.Accounts)
		to, _ := simtypes.RandomAcc(r, simState.Accounts)
		fromAddr := from.Address.String()

		proposal := &group.Proposal{
			ProposalId:          uint64(i + 1),
			Proposers:           proposers,
			Address:             fromAddr,
			GroupVersion:        uint64(i + 1),
			GroupAccountVersion: uint64(i + 1),
			Status:              group.ProposalStatusSubmitted,
			Result:              group.ProposalResultAccepted,
			VoteState: group.Tally{
				YesCount:     "1",
				NoCount:      "1",
				AbstainCount: "1",
				VetoCount:    "0",
			},
			ExecutorResult: group.ProposalExecutorResultNotRun,
			Metadata:       []byte(simtypes.RandStringOfLength(r, 50)),
			SubmittedAt:    gogotypes.Timestamp{Seconds: 1},
			Timeout:        gogotypes.Timestamp{Seconds: 1000},
		}
		err := proposal.SetMsgs([]sdk.Msg{&banktypes.MsgSend{
			FromAddress: fromAddr,
			ToAddress:   to.Address.String(),
			Amount:      sdk.NewCoins(sdk.NewInt64Coin("test", 10)),
		}})
		if err != nil {
			panic(err)
		}

		proposals[i] = proposal
	}

	return proposals
}

func getVotes(r *rand.Rand, simState *module.SimulationState) []*group.Vote {
	votes := make([]*group.Vote, 3)

	for i := 0; i < 3; i++ {
		votes[i] = &group.Vote{
			ProposalId:  uint64(i + 1),
			Voter:       simState.Accounts[i].Address.String(),
			Choice:      getVoteChoice(i),
			Metadata:    []byte(simtypes.RandStringOfLength(r, 50)),
			SubmittedAt: gogotypes.Timestamp{Seconds: 10},
		}
	}

	return votes
}

func getVoteChoice(index int) group.Choice {
	switch index {
	case 0:
		return group.Choice_CHOICE_YES
	case 1:
		return group.Choice_CHOICE_NO
	case 2:
		return group.Choice_CHOICE_ABSTAIN
	default:
		return group.Choice_CHOICE_VETO
	}
}

// RandomizedGenState generates a random GenesisState for the group module.
func RandomizedGenState(simState *module.SimulationState) {

	// groups
	var groups []*group.GroupInfo
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GroupInfo, &groups, simState.Rand,
		func(r *rand.Rand) { groups = getGroups(r, simState.Accounts) },
	)

	// group members
	var members []*group.GroupMember
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GroupMembers, &members, simState.Rand,
		func(r *rand.Rand) { members = getGroupMembers(r, simState.Accounts) },
	)

	// group accounts
	var groupAccounts []*group.GroupAccountInfo
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GroupAccountInfo, &groupAccounts, simState.Rand,
		func(r *rand.Rand) { groupAccounts = getGroupAccounts(r, simState) },
	)

	// proposals
	var proposals []*group.Proposal
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GroupProposals, &proposals, simState.Rand,
		func(r *rand.Rand) { proposals = getProposals(r, simState) },
	)

	// votes
	var votes []*group.Vote
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GroupVote, &votes, simState.Rand,
		func(r *rand.Rand) { votes = getVotes(r, simState) },
	)

	groupGenesis := group.GenesisState{
		GroupSeq:        3,
		Groups:          groups,
		GroupMembers:    members,
		GroupAccountSeq: 3,
		GroupAccounts:   groupAccounts,
		ProposalSeq:     3,
		Proposals:       proposals,
		Votes:           votes,
	}

	simState.GenState[group.ModuleName] = simState.Cdc.MustMarshalJSON(&groupGenesis)
}
