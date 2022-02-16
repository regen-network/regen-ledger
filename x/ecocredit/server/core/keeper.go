package core

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
)

type Keeper struct {
	stateStore ecocreditv1beta1.StateStore
	bankKeeper ecocredit.BankKeeper
	params     ParamKeeper
}

func (k Keeper) Send(ctx context.Context, send *v1beta1.MsgSend) (*v1beta1.MsgSendResponse, error) {
	panic("implement me")
}

func (k Keeper) Retire(ctx context.Context, retire *v1beta1.MsgRetire) (*v1beta1.MsgRetireResponse, error) {
	panic("implement me")
}

func (k Keeper) Cancel(ctx context.Context, cancel *v1beta1.MsgCancel) (*v1beta1.MsgCancelResponse, error) {
	panic("implement me")
}

func (k Keeper) UpdateClassAdmin(ctx context.Context, admin *v1beta1.MsgUpdateClassAdmin) (*v1beta1.MsgUpdateClassAdminResponse, error) {
	panic("implement me")
}

func (k Keeper) UpdateClassIssuers(ctx context.Context, issuers *v1beta1.MsgUpdateClassIssuers) (*v1beta1.MsgUpdateClassIssuersResponse, error) {
	panic("implement me")
}

func (k Keeper) UpdateClassMetadata(ctx context.Context, metadata *v1beta1.MsgUpdateClassMetadata) (*v1beta1.MsgUpdateClassMetadataResponse, error) {
	panic("implement me")
}

type ParamKeeper interface {
	GetParamSet(ctx sdk.Context, params *ecocredit.Params)
}

var _ v1beta1.MsgServer = &Keeper{}
