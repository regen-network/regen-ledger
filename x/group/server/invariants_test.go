package server

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm"
)

func TestTallyVotesInvariant(t *testing.T) {
	// ff := server.NewFixtureFactory(t, 6)
	var ctx sdk.Context //ff.Setup().Context()
	// c := ff.Setup().Context()
	blockHeight := ctx.BlockHeight()
	var proposalTable orm.AutoUInt64Table

	invar := Invar{
		sdkCtx:        ctx,
		proposalTable: proposalTable,
	}
	invarRes := tallyVotesInvariant(invar)

}
