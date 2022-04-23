package core

import "github.com/regen-network/regen-ledger/types/testutil"

var (
	batchDenom     = "A00-00000000-00000000-000"
	batchIssuance1 = BatchIssuance{Recipient: testutil.GenAddress(), TradableAmount: "12"}
	batchIssuance2 = BatchIssuance{Recipient: testutil.GenAddress(), TradableAmount: "12", RetiredAmount: "20", RetirementJurisdiction: "CH"}
	batchIssuances = []*BatchIssuance{&batchIssuance1, &batchIssuance2}
	batchOrigTx    = OriginTx{Typ: "Polygon", Id: "0x1234"}
)
