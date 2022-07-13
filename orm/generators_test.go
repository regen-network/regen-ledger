package orm_test

import (
	"pgregory.net/rapid"

	"github.com/regen-network/regen-ledger/orm/testdata"
)

// genGroupMember generates a new group member. At the moment it doesn't
// generate empty strings for Group or Member.
var genGroupMember = rapid.Custom(func(t *rapid.T) *testdata.GroupMember {
	return &testdata.GroupMember{
		Group:  []byte(rapid.StringN(1, 100, 150).Draw(t, "group").(string)),
		Member: []byte(rapid.StringN(1, 100, 150).Draw(t, "member").(string)),
		Weight: rapid.Uint64().Draw(t, "weight").(uint64),
	}
})
