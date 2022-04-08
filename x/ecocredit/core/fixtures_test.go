package core

import "github.com/cosmos/cosmos-sdk/testutil/testdata"

func genAddress() string {
	_, _, addr :=
		testdata.KeyTestPubAddr()
	return addr.String()
}

var (
	batchDenom     = "A00-00000000-00000000-000"
	batchIssuance1 = BatchIssuance{Recipient: genAddress(), TradableAmount: "12"}
	batchIssuance2 = BatchIssuance{Recipient: genAddress(), TradableAmount: "12", RetiredAmount: "20", RetirementLocation: "CH"}
	batchIssuances = []*BatchIssuance{&batchIssuance1, &batchIssuance2}
	batchOrigTx    = OriginTx{Typ: "Polygon", Id: "0x1234"}
)
