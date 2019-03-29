package binary

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"net/url"
)

// AccAddressID wraps sdk.AccAddress as a Node ID type
type AccAddressID struct {
	sdk.AccAddress
}

// URI returns the URI representation of the AccAddressID
func (id AccAddressID) URI() *url.URL {
	uri, err := url.Parse(id.String())
	if err != nil {
		panic(err)
	}
	return uri
}

// HashID represents a node ID that is a URI hash fragment of the unlabelled root node
type HashID struct {
	fragment string
}

func (id HashID) String() string {
	return fmt.Sprintf("_:_.%s", id.fragment)
}

// URI returns the URI representation of the HashID
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
