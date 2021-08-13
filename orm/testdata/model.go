package testdata

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrTest = errors.Register("orm_testdata", 9999, "test")
)

func (g GroupMember) PrimaryKeyFields() []interface{} {
	return []interface{}{[]byte(g.Group), []byte(g.Member)}
}

func (g GroupInfo) PrimaryKeyFields() []interface{} {
	return []interface{}{g.GroupId}
}

func (g GroupInfo) ValidateBasic() error {
	if g.Description == "invalid" {
		return errors.Wrap(ErrTest, "description")
	}
	return nil
}

func (g GroupMember) ValidateBasic() error {
	return nil
}
