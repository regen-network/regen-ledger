package server

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) InitGenesis(ctx types.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState group.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	if err := orm.ImportTableData(ctx, s.groupTable, genesisState.Groups, 0); err != nil {
		panic(errors.Wrap(err, "groups"))
	}
	if err := s.groupSeq.InitVal(ctx, genesisState.GroupSeq); err != nil {
		panic(errors.Wrap(err, "group seq"))
	}

	if err := orm.ImportTableData(ctx, s.groupMemberTable, genesisState.GroupMembers, 0); err != nil {
		panic(errors.Wrap(err, "group members"))
	}

	if err := orm.ImportTableData(ctx, s.groupAccountTable, genesisState.GroupAccounts, 0); err != nil {
		panic(errors.Wrap(err, "group accounts"))
	}
	if err := s.groupAccountSeq.InitVal(ctx, genesisState.GroupAccountSeq); err != nil {
		panic(errors.Wrap(err, "group account seq"))
	}

	if err := orm.ImportTableData(ctx, s.proposalTable, genesisState.Proposals, genesisState.ProposalSeq); err != nil {
		panic(errors.Wrap(err, "proposals"))
	}

	if err := orm.ImportTableData(ctx, s.voteTable, genesisState.Votes, 0); err != nil {
		panic(errors.Wrap(err, "votes"))
	}

	return []abci.ValidatorUpdate{}
}

func (s serverImpl) ExportGenesis(ctx types.Context, cdc codec.JSONMarshaler) json.RawMessage {
	genesisState := group.NewGenesisState()

	var groups []*group.GroupInfo
	groupSeq, err := orm.ExportTableData(ctx, s.groupTable, &groups)
	if err != nil {
		panic(errors.Wrap(err, "groups"))
	}
	genesisState.Groups = groups
	genesisState.GroupSeq = groupSeq

	var groupMembers []*group.GroupMember
	_, err = orm.ExportTableData(ctx, s.groupMemberTable, &groupMembers)
	if err != nil {
		panic(errors.Wrap(err, "group members"))
	}
	genesisState.GroupMembers = groupMembers

	var groupAccounts []*group.GroupAccountInfo
	_, err = orm.ExportTableData(ctx, s.groupAccountTable, &groupAccounts)
	if err != nil {
		panic(errors.Wrap(err, "group accounts"))
	}
	genesisState.GroupAccounts = groupAccounts

	var proposals []*group.Proposal
	proposalSeq, err := orm.ExportTableData(ctx, s.proposalTable, &proposals)
	if err != nil {
		panic(errors.Wrap(err, "proposals"))
	}
	genesisState.Proposals = proposals
	genesisState.ProposalSeq = proposalSeq

	var votes []*group.Vote
	_, err = orm.ExportTableData(ctx, s.voteTable, &votes)
	if err != nil {
		panic(errors.Wrap(err, "votes"))
	}
	genesisState.Votes = votes

	genesisBytes := cdc.MustMarshalJSON(genesisState)
	return genesisBytes
}
