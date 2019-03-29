package binary

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"net/url"
)

type AccAddressID struct {
	sdk.AccAddress
}

func (id AccAddressID) URI() *url.URL {
	uri, err := url.Parse(id.String())
	if err != nil {
		panic(err)
	}
	return uri
}

type HashID struct {
	fragment string
}

func (id HashID) String() string {
	return fmt.Sprintf("_:_.%s", id.fragment)
}

func (id HashID) URI() *url.URL {
	// TODO this should really be a blank node or the actual hash of the graph plus the fragment
	uri, err := url.Parse(fmt.Sprintf("root:#%s", id.fragment))
	if err != nil {
		panic(err)
	}
	return uri
}

//func (n node) GetClasses() []string {
//	panic("implement me")
//}
