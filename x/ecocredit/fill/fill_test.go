package fill_test

import (
	"bytes"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/x/ecocredit/testutil"

	"github.com/regen-network/regen-ledger/x/ecocredit/fill"
	"github.com/rs/zerolog"
)

func TestFill(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := zerolog.New(buf)
	transferMgr := testutil.NewTestTransferManager(logger)
	db, err := ormdb.NewModuleDB(testutil.TestModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	fillMgr, err := fill.NewManager(db, transferMgr, logger)
	assert.NilError(t, err)
	ctx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

	market := &marketplacev1beta1.Market{
		Id:                1,
		CreditType:        "C",
		BankDenom:         "foo",
		PrecisionModifier: 3,
	}

	acct1 := sdk.AccAddress{0, 1, 2, 3, 4, 5}
	acct2 := sdk.AccAddress{5, 4, 3, 2, 1, 0}

	tests := []struct {
		name         string
		buyOrder     *marketplacev1beta1.BuyOrder
		sellOrder    *marketplacev1beta1.SellOrder
		expectStatus fill.Status
		expectErr    bool
	}{
		{
			"equal",
			&marketplacev1beta1.BuyOrder{
				Id:                 0,
				Buyer:              acct1,
				Quantity:           "10",
				MarketId:           1,
				BidPrice:           "",
				DisableAutoRetire:  false,
				DisablePartialFill: false,
				Expiration:         nil,
				Maker:              false,
			},
			&marketplacev1beta1.SellOrder{},
			fill.BothFilled,
			false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			status, err := fillMgr.Fill(ctx, market, test.buyOrder, test.sellOrder)
			if test.expectErr {
				assert.Assert(t, err != nil)
			} else {
				assert.NilError(t, err)
			}
			assert.Equal(t, test.expectStatus, status)
		})
	}
}
