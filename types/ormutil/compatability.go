package ormutil

import (
	gogoproto "github.com/cosmos/gogoproto/proto"
	"google.golang.org/protobuf/proto"

	queryv1beta1 "cosmossdk.io/api/cosmos/base/query/v1beta1"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/regen-network/regen-ledger/orm/model/ormlist"
)

// PageReqToCosmosAPILegacy is a temporal adapter for ORM v-alpha-*
func PageReqToCosmosAPILegacy(from *query.PageRequest) *queryv1beta1.PageRequest {
	if from == nil {
		return &queryv1beta1.PageRequest{Limit: query.DefaultLimit}
	}
	return &queryv1beta1.PageRequest{
		Key: from.Key, Offset: from.Offset, Limit: from.Limit, CountTotal: from.CountTotal, Reverse: from.Reverse}
}

func PageReqToOrmPaginate(pg *query.PageRequest) ormlist.Option {
	return ormlist.Paginate(PageReqToCosmosAPILegacy(pg))
}

func PageResToCosmosTypes(from *queryv1beta1.PageResponse) *query.PageResponse {
	if from == nil {
		return nil
	}
	return &query.PageResponse{NextKey: from.NextKey, Total: from.Total}
}

// TODO: probably we can remove
func GogoPageReqToPulsarPageReq(from *query.PageRequest) (*queryv1beta1.PageRequest, error) {
	if from == nil {
		return &queryv1beta1.PageRequest{Limit: query.DefaultLimit}, nil
	}

	return &queryv1beta1.PageRequest{
		Key: from.Key, Offset: from.Offset, Limit: from.Limit, CountTotal: from.CountTotal, Reverse: from.Reverse}, nil
}

// TODO: probably we can remove
func PulsarPageResToGogoPageRes(from *queryv1beta1.PageResponse) (*query.PageResponse, error) {
	if from == nil {
		return nil, nil
	}

	to := &query.PageResponse{}
	err := PulsarToGogoSlow(from, to)
	return to, err
}

func PulsarToGogoSlow(from proto.Message, to gogoproto.Message) error {
	if from == nil {
		return nil
	}

	bz, err := proto.Marshal(from)
	if err != nil {
		return err
	}

	return gogoproto.Unmarshal(bz, to)
}

func GogoToPulsarSlow(from gogoproto.Message, to proto.Message) error {
	bz, err := gogoproto.Marshal(from)
	if err != nil {
		return err
	}

	return proto.Unmarshal(bz, to)
}
