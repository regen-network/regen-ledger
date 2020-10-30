package server

import (
	"context"
	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.QueryServer = serverImpl{}

func (s serverImpl) Data(ctx context.Context, request *data.QueryDataRequest) (*data.QueryDataResponse, error) {
	panic("implement me")
}
