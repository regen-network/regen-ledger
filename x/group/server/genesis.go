package server

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) InitGenesis(ctx types.Context, cdc codec.Codec, data json.RawMessage) ([]abci.ValidatorUpdate, error) {
	var genesisState group.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	if err := s.groupTable.Import(ctx, genesisState.Groups, genesisState.GroupSeq); err != nil {
		return nil, errors.Wrap(err, "groups")
	}

	if err := s.groupMemberTable.Import(ctx, genesisState.GroupMembers, 0); err != nil {
		return nil, errors.Wrap(err, "group members")
	}

	if err := s.groupAccountTable.Import(ctx, genesisState.GroupAccounts, 0); err != nil {
		return nil, errors.Wrap(err, "group accounts")
	}
	if err := s.groupAccountSeq.InitVal(ctx, genesisState.GroupAccountSeq); err != nil {
		return nil, errors.Wrap(err, "group account seq")
	}

	if err := s.proposalTable.Import(ctx, genesisState.Proposals, genesisState.ProposalSeq); err != nil {
		return nil, errors.Wrap(err, "proposals")
	}

	if err := s.voteTable.Import(ctx, genesisState.Votes, 0); err != nil {
		return nil, errors.Wrap(err, "votes")
	}

	return []abci.ValidatorUpdate{}, nil
}

func (s serverImpl) ExportGenesis(ctx types.Context, cdc codec.Codec) (json.RawMessage, error) {
	genesisState := group.NewGenesisState()

	var groups []*group.GroupInfo
	groupSeq, err := s.groupTable.Export(ctx, &groups)
	if err != nil {
		return nil, errors.Wrap(err, "groups")
	}
	genesisState.Groups = groups
	genesisState.GroupSeq = groupSeq

	var groupMembers []*group.GroupMember
	_, err = s.groupMemberTable.Export(ctx, &groupMembers)
	if err != nil {
		return nil, errors.Wrap(err, "group members")
	}
	genesisState.GroupMembers = groupMembers

	var groupAccounts []*group.GroupAccountInfo
	_, err = s.groupAccountTable.Export(ctx, &groupAccounts)
	if err != nil {
		return nil, errors.Wrap(err, "group accounts")
	}
	genesisState.GroupAccounts = groupAccounts
	genesisState.GroupAccountSeq = s.groupAccountSeq.CurVal(ctx)

	var proposals []*group.Proposal
	proposalSeq, err := s.proposalTable.Export(ctx, &proposals)
	if err != nil {
		return nil, errors.Wrap(err, "proposals")
	}
	genesisState.Proposals = proposals
	genesisState.ProposalSeq = proposalSeq

	var votes []*group.Vote
	_, err = s.voteTable.Export(ctx, &votes)
	if err != nil {
		return nil, errors.Wrap(err, "votes")
	}
	genesisState.Votes = votes

	genesisBytes := cdc.MustMarshalJSON(genesisState)
	return genesisBytes, nil
}
