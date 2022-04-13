package testutil

import "github.com/cosmos/cosmos-sdk/testutil/testdata"

// GenAddress generates a valid bech32 address
func GenAddress() string {
	_, _, addr := testdata.KeyTestPubAddr()
	return addr.String()
}
