package core

func genAddress() string {
	_, _, addr := testdata.KeyTestPubAddr()
	return addr.String()
}

var (
	batchDenom = "A00-00000000-00000000-000"
)
