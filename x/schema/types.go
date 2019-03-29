package schema

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/graph"
	"net/url"
)

// PropertyID is the integer ID of property starting from 1 with 0 indicating no property
type PropertyID int64

// PropertyDefinition defines the schema for a property
type PropertyDefinition struct {
	// Creator is the entity that defined this property
	Creator sdk.AccAddress `json:"creator"`
	// Name is the human-readable name of the property within the creator's namespace of properties
	Name string `json:"name"`
	// Many indicates whether or not this property can be assigned more than once to a given node/object. If it is
	// false, then the property can only be assigned once for a given node/object
	Arity graph.Arity `json:"arity:omitempty"`
	// PropertyType indicates the data type of the property
	PropertyType graph.PropertyType `json:"type,omitempty"`
}

func (prop PropertyDefinition) String() string {
	return fmt.Sprintf(`Property
URI: %s,
Arity: %s
Type: %s
`, prop.URI(), prop.Arity.String(), prop.PropertyType.String())
}

// URL returns the URL of the property
func (prop PropertyDefinition) URI() *url.URL {
	uri, err := url.Parse(fmt.Sprintf("%s/%s", prop.Creator.String(), prop.Name))
	if err != nil {
		panic(err)
	}
	return uri
}
