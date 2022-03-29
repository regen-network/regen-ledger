package ormutil

import (
	queryv1beta1 "github.com/cosmos/cosmos-sdk/api/cosmos/base/query/v1beta1"
	"github.com/cosmos/cosmos-sdk/types/query"
	gogoproto "github.com/gogo/protobuf/proto"
	"google.golang.org/protobuf/proto"
)

func GogoPageReqToPulsarPageReq(from *query.PageRequest) (*queryv1beta1.PageRequest, error) {
	if from == nil {
		return &queryv1beta1.PageRequest{}, nil
	}

	to := &queryv1beta1.PageRequest{}
	err := GogoToPulsarSlow(from, to)
	return to, err
}

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
