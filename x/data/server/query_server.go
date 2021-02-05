package server

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	"github.com/regen-network/regen-ledger/x/data/rdf"

	"github.com/regen-network/regen-ledger/x/data/rdf/compact"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.QueryServer = serverImpl{}

func (s serverImpl) ByHash(ctx types.Context, request *data.QueryByHashRequest) (*data.QueryByHashResponse, error) {
	return nil, fmt.Errorf("not implemented")
	//cid := request.Cid
	//
	//var timestamp gogotypes.Timestamp
	//store := ctx.KVStore(s.storeKey)
	//bz := store.Get(AnchorKey(cid))
	//if len(bz) == 0 {
	//	return nil, status.Error(codes.NotFound, "CID not found")
	//}
	//err := timestamp.Unmarshal(bz)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var signers []string
	//cidSignerPrefixKey := CIDSignerIndexPrefix(CIDBase64String(cid))
	//prefixStore := prefix.NewStore(store, cidSignerPrefixKey)
	//iterator := prefixStore.Iterator(nil, nil)
	//
	//for iterator.Valid() {
	//	signer := string(iterator.Key())
	//	signers = append(signers, signer)
	//	iterator.Next()
	//}
	//
	//content := store.Get(DataKey(cid))
	//
	//return &data.QueryByCidResponse{
	//	Timestamp: &timestamp,
	//	Signers:   signers,
	//	Content:   content,
	//}, err
}

func (s serverImpl) BySigner(ctx types.Context, request *data.QueryBySignerRequest) (*data.QueryBySignerResponse, error) {
	return nil, fmt.Errorf("not implemented")
	//store := prefix.NewStore(ctx.KVStore(s.storeKey), SignerCIDIndexPrefix(request.Signer))
	//
	//var cids [][]byte
	//pageRes, err := query.Paginate(store, request.Pagination, func(key []byte, value []byte) error {
	//	cids = append(cids, key)
	//	return nil
	//})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &data.QueryBySignerResponse{
	//	Cids:       cids,
	//	Pagination: pageRes,
	//}, nil
}

func (s serverImpl) ConvertToCompactDataset(ctx types.Context, request *data.ConvertToCompactDatasetRequest) (*data.ConvertToCompactDatasetResponse, error) {
	dataset, err := compact.Compact(request.Content, request.ContentType, serverIRIResolver{
		serverImpl: s,
		context:    ctx,
	})
	if err != nil {
		return nil, err
	}

	bz, err := proto.Marshal(dataset)
	if err != nil {
		return nil, err
	}

	return &data.ConvertToCompactDatasetResponse{CompactDataset: bz}, nil
}

type serverIRIResolver struct {
	serverImpl
	context types.Context
}

func (s serverIRIResolver) GetIDForIRI(iri rdf.IRI) []byte {
	panic("TODO")
}
