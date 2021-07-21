package testsuite

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	proto "github.com/gogo/protobuf/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s *IntegrationTestSuite) TestInitExportGenesis() {
	require := s.Require()
	ctx := s.genesisCtx
	cdc := s.fixture.Codec()

	now := time.Now()
	submittedAt, err := proto.TimestampProto(now)
	require.NoError(err)
	timeout, err := proto.TimestampProto(now.Add(time.Second * 1))
	require.NoError(err)

	groupAccount := &group.GroupAccountInfo{
		Address:       s.groupAccountAddr.String(),
		GroupId:       1,
		Admin:         s.addr1.String(),
		Version:       1,
		Metadata:      []byte("account metadata"),
		DerivationKey: []byte("account derivation key"),
	}
	err = groupAccount.SetDecisionPolicy(&group.ThresholdDecisionPolicy{
		Threshold: "1",
		Timeout:   proto.Duration{Seconds: 1},
	})
	require.NoError(err)

	proposal := &group.Proposal{
		ProposalId:          1,
		Address:             s.groupAccountAddr.String(),
		Metadata:            []byte("proposal metadata"),
		GroupVersion:        1,
		GroupAccountVersion: 1,
		Proposers: []string{
			s.addr1.String(),
		},
		SubmittedAt: *submittedAt,
		Status:      group.ProposalStatusClosed,
		Result:      group.ProposalResultAccepted,
		VoteState: group.Tally{
			YesCount:     "1",
			NoCount:      "0",
			AbstainCount: "0",
			VetoCount:    "0",
		},
		Timeout:        *timeout,
		ExecutorResult: group.ProposalExecutorResultSuccess,
	}
	err = proposal.SetMsgs([]sdk.Msg{&banktypes.MsgSend{
		FromAddress: s.groupAccountAddr.String(),
		ToAddress:   s.addr2.String(),
		Amount:      sdk.Coins{sdk.NewInt64Coin("test", 100)},
	}})
	require.NoError(err)

	genesisState := &group.GenesisState{
		GroupSeq:        2,
		Groups:          []*group.GroupInfo{{GroupId: 1, Admin: s.addr1.String(), Metadata: []byte("1"), Version: 1, TotalWeight: "1"}, {GroupId: 2, Admin: s.addr2.String(), Metadata: []byte("2"), Version: 2, TotalWeight: "2"}},
		GroupMembers:    []*group.GroupMember{{GroupId: 1, Member: &group.Member{Address: s.addr1.String(), Weight: "1", Metadata: []byte("member metadata")}}, {GroupId: 2, Member: &group.Member{Address: s.addr1.String(), Weight: "2", Metadata: []byte("member metadata")}}},
		GroupAccountSeq: 1,
		GroupAccounts:   []*group.GroupAccountInfo{groupAccount},
		ProposalSeq:     1,
		Proposals:       []*group.Proposal{proposal},
		Votes:           []*group.Vote{{ProposalId: proposal.ProposalId, Voter: s.addr1.String(), SubmittedAt: *submittedAt, Choice: group.Choice_CHOICE_YES}},
	}

	genesisBytes, err := cdc.MarshalJSON(genesisState)
	require.NoError(err)

	ecocreditGenesisState := ecocredit.DefaultGenesisState()
	ecocreditGenesisBytes, err := cdc.MarshalJSON(ecocreditGenesisState)
	require.NoError(err)

	genesisData := map[string]json.RawMessage{
		group.ModuleName:     genesisBytes,
		ecocredit.ModuleName: ecocreditGenesisBytes,
	}
	_, err = s.fixture.InitGenesis(ctx.Context, genesisData)
	require.NoError(err)

	for i, g := range genesisState.Groups {
		res, err := s.queryClient.GroupInfo(ctx, &group.QueryGroupInfoRequest{
			GroupId: g.GroupId,
		})
		require.NoError(err)
		require.Equal(g, res.Info)

		membersRes, err := s.queryClient.GroupMembers(ctx, &group.QueryGroupMembersRequest{
			GroupId: g.GroupId,
		})
		require.NoError(err)
		require.Equal(len(membersRes.Members), 1)
		require.Equal(membersRes.Members[0], genesisState.GroupMembers[i])
	}

	for _, g := range genesisState.GroupAccounts {
		res, err := s.queryClient.GroupAccountInfo(ctx, &group.QueryGroupAccountInfoRequest{
			Address: g.Address,
		})
		require.NoError(err)
		s.assertGroupAccountsEqual(g, res.Info)
	}

	for _, g := range genesisState.Proposals {
		res, err := s.queryClient.Proposal(ctx, &group.QueryProposalRequest{
			ProposalId: g.ProposalId,
		})
		require.NoError(err)
		s.assertProposalsEqual(g, res.Proposal)

		votesRes, err := s.queryClient.VotesByProposal(ctx, &group.QueryVotesByProposalRequest{
			ProposalId: g.ProposalId,
		})
		require.NoError(err)
		require.Equal(len(votesRes.Votes), 1)
		require.Equal(votesRes.Votes[0], genesisState.Votes[0])
	}

	exported, err := s.fixture.ExportGenesis(ctx.Context)
	require.NoError(err)

	var exportedGenesisState group.GenesisState
	err = cdc.UnmarshalJSON(exported[group.ModuleName], &exportedGenesisState)
	require.NoError(err)

	require.Equal(genesisState.Groups, exportedGenesisState.Groups)
	require.Equal(genesisState.GroupMembers, exportedGenesisState.GroupMembers)

	require.Equal(len(genesisState.GroupAccounts), len(exportedGenesisState.GroupAccounts))
	for i, g := range genesisState.GroupAccounts {
		res := exportedGenesisState.GroupAccounts[i]
		require.NoError(err)
		s.assertGroupAccountsEqual(g, res)
	}

	require.Equal(len(genesisState.Proposals), len(exportedGenesisState.Proposals))
	for i, g := range genesisState.Proposals {
		res := exportedGenesisState.Proposals[i]
		require.NoError(err)
		s.assertProposalsEqual(g, res)
	}
	require.Equal(genesisState.Votes, exportedGenesisState.Votes)

	require.Equal(genesisState.GroupSeq, exportedGenesisState.GroupSeq)
	require.Equal(genesisState.GroupAccountSeq, exportedGenesisState.GroupAccountSeq)
	require.Equal(genesisState.ProposalSeq, exportedGenesisState.ProposalSeq)

}

func (s *IntegrationTestSuite) assertGroupAccountsEqual(g *group.GroupAccountInfo, other *group.GroupAccountInfo) {
	require := s.Require()
	require.Equal(g.Address, other.Address)
	require.Equal(g.GroupId, other.GroupId)
	require.Equal(g.Admin, other.Admin)
	require.Equal(g.Metadata, other.Metadata)
	require.Equal(g.Version, other.Version)
	require.Equal(g.GetDecisionPolicy(), other.GetDecisionPolicy())
}

func (s *IntegrationTestSuite) assertProposalsEqual(g *group.Proposal, other *group.Proposal) {
	require := s.Require()
	require.Equal(g.ProposalId, other.ProposalId)
	require.Equal(g.Address, other.Address)
	require.Equal(g.Metadata, other.Metadata)
	require.Equal(g.Proposers, other.Proposers)
	require.Equal(g.SubmittedAt, other.SubmittedAt)
	require.Equal(g.GroupVersion, other.GroupVersion)
	require.Equal(g.GroupAccountVersion, other.GroupAccountVersion)
	require.Equal(g.Status, other.Status)
	require.Equal(g.Result, other.Result)
	require.Equal(g.VoteState, other.VoteState)
	require.Equal(g.Timeout, other.Timeout)
	require.Equal(g.ExecutorResult, other.ExecutorResult)
	require.Equal(g.GetMsgs(), other.GetMsgs())
}
