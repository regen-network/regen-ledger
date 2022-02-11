package basket_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	"github.com/regen-network/regen-ledger/x/ecocredit"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/golang/mock/gomock"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket/mocks"
	"github.com/stretchr/testify/require"
)

// this is an example of how we will unit test the basket functionality with mocks
func TestKeeperExample(t *testing.T) {
	ctrl := gomock.NewController(t)
	db, err := ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)

	bankKeeper := mocks.NewMockBankKeeper(ctrl)
	ecocreditKeeper := mocks.NewMockEcocreditKeeper(ctrl)
	sk := sdk.NewKVStoreKey("test")
	k := basket.NewKeeper(db, ecocreditKeeper, bankKeeper, sk, ecocredit.ModuleName)
	require.NotNil(t, k)
}
