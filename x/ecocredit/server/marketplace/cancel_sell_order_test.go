package marketplace

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func TestSell_CancelOrder(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	expir := time.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).Times(1)

	balBefore, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	supBefore, err := s.coreStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, Expiration: &expir},
		},
	})
	assert.NilError(t, err)

	_, err = s.k.CancelSellOrder(s.ctx, &marketplace.MsgCancelSellOrder{SellOrderId: res.SellOrderIds[0], Seller: s.addr.String()})
	assert.NilError(t, err)

	balAfter, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	supAfter, err := s.coreStore.BatchSupplyTable().Get(s.ctx, 1)
	assert.NilError(t, err)

	assert.DeepEqual(t, balBefore, balAfter, cmpopts.IgnoreUnexported(api.BatchBalance{}))
	assert.DeepEqual(t, supBefore, supAfter, cmpopts.IgnoreUnexported(api.BatchSupply{}))
}

func TestSell_CancelOrderInvalid(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	batchDenom := "C01-20200101-20200201-001"
	start, end := timestamppb.Now(), timestamppb.Now()
	expir := time.Now()
	ask := sdk.NewInt64Coin("ufoo", 10)
	creditType := ecocredit.CreditType{
		Name:         "carbon",
		Abbreviation: "C",
		Unit:         "tonnes",
		Precision:    6,
	}
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], "C01", start, end, creditType)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any interface{}, p *ecocredit.Params) {
		p.CreditTypes = []*ecocredit.CreditType{&creditType}
	}).Times(1)

	_, _, otherAddr := testdata.KeyTestPubAddr()

	res, err := s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.addr.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &ask, Expiration: &expir},
		},
	})
	assert.NilError(t, err)

	_, err = s.k.CancelSellOrder(s.ctx, &marketplace.MsgCancelSellOrder{
		Seller:      otherAddr.String(),
		SellOrderId: res.SellOrderIds[0],
	})
	assert.ErrorContains(t, err, sdkerrors.ErrUnauthorized.Error())
}
