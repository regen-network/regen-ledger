package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/gogo/protobuf/types"
	"github.com/regen-network/regen-ledger/x/group"
)

const (
	GroupInfo        = "group-info"
	GroupMembers     = "group-members"
	GroupAccountInfo = "group-accout-info"
	GroupProposals   = "group-proposals"
)

func GetGroups(r *rand.Rand, accounts []simtypes.Account) []*group.GroupInfo {
	groups := make([]*group.GroupInfo, 3)
	for i := 0; i < 3; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		groups[i] = &group.GroupInfo{
			GroupId:     1,
			Admin:       acc.Address.String(),
			Metadata:    []byte(simtypes.RandStringOfLength(r, 10)),
			Version:     1,
			TotalWeight: "30",
		}
	}
	return groups
}

func GetGroupMembers(r *rand.Rand, accounts []simtypes.Account) []*group.GroupMember {
	groupMembers := make([]*group.GroupMember, 3)
	for i := 0; i < 3; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		groupMembers[i] = &group.GroupMember{
			GroupId: 1,
			Member: &group.Member{
				Address:  acc.Address.String(),
				Weight:   "10",
				Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
			},
		}
	}
	return groupMembers
}

func GetGroupAccounts(r *rand.Rand, accounts []simtypes.Account) []*group.GroupAccountInfo {
	groupMembers := make([]*group.GroupAccountInfo, 3)
	for i := 0; i < 3; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		groupMembers[i] = &group.GroupAccountInfo{
			GroupId:  1,
			Admin:    acc.Address.String(),
			Metadata: []byte(simtypes.RandStringOfLength(r, 10)),
		}
	}
	return groupMembers
}

func GetProposals(r *rand.Rand, simState *module.SimulationState) []*group.Proposal {
	groupMembers := make([]*group.Proposal, 3)
	for i := 0; i < 3; i++ {
		acc, _ := simtypes.RandomAcc(r, simState.Accounts)
		groupMembers[i] = &group.Proposal{
			ProposalId:          1,
			Proposers:           []string{simState.Accounts[0].Address.String(), simState.Accounts[1].Address.String()},
			Address:             acc.Address.String(),
			GroupVersion:        1,
			GroupAccountVersion: 1,
			Status:              group.ProposalStatusClosed,
			Result:              group.ProposalResultAccepted,
			VoteState:           group.Tally{},
			ExecutorResult:      group.ProposalExecutorResultNotRun,
			Metadata:            []byte(simtypes.RandStringOfLength(r, 50)),
			SubmittedAt:         types.Timestamp{},
			Timeout:             types.Timestamp{},
		}
	}
	return groupMembers
}

// RandomizedGenState generates a random GenesisState for the group module.
func RandomizedGenState(simState *module.SimulationState) {

	// groups
	var groups []*group.GroupInfo
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GroupInfo, &groups, simState.Rand,
		func(r *rand.Rand) { groups = GetGroups(r, simState.Accounts) },
	)

	// group members
	var members []*group.GroupMember
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GroupMembers, &members, simState.Rand,
		func(r *rand.Rand) { members = GetGroupMembers(r, simState.Accounts) },
	)

	// group accounts
	var groupAccounts []*group.GroupAccountInfo
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GroupAccountInfo, &groupAccounts, simState.Rand,
		func(r *rand.Rand) { groupAccounts = GetGroupAccounts(r, simState.Accounts) },
	)

	// proposals
	var proposals []*group.Proposal
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GroupProposals, &proposals, simState.Rand,
		func(r *rand.Rand) { proposals = GetProposals(r, simState) },
	)

	groupGenesis := group.GenesisState{
		GroupSeq:        1,
		Groups:          groups,
		GroupMembers:    members,
		GroupAccountSeq: 1,
		GroupAccounts:   groupAccounts,
		ProposalSeq:     1,
		Proposals:       proposals,
	}

	bz, err := json.MarshalIndent(&groupGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", group.ModuleName, bz)

	simState.GenState[group.ModuleName] = simState.Cdc.MustMarshalJSON(&groupGenesis)

}
