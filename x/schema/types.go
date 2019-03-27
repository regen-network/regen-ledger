package schema

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PropertyType indicates the data type of a property
type PropertyType uint32

const (
	// TyObject is the PropertyType of objects or nodes
	TyObject PropertyType = iota
	// TyString is the PropertyType of strings
	TyString
	// TyInteger is the PropertyType of arbitrary-precision integers
	TyInteger
	// TyDouble is the PropertyType of double-precision floating point
	TyDouble
	// TyBool is the PropertyType of boolean values
	TyBool
)

// PropertyID is the integer ID of property starting from 1 with 0 indicating no property
type PropertyID uint64

// PropertyDefinition defines the schema for a property
type PropertyDefinition struct {
	// Creator is the entity that defined this property
	Creator sdk.AccAddress `json:"creator"`
	// Name is the human-readable name of the property within the creator's namespace of properties
	Name string `json:"name"`
	// Many indicates whether or not this property can be assigned more than once to a given node/object. If it is
	// false, then the property can only be assigned once for a given node/object
	Many bool `json:"many:omitempty"`
	// PropertyType indicates the data type of the property
	PropertyType PropertyType `json:"type,omitempty"`
}

func (prop PropertyDefinition) String() string {
	arity := "One"
	if prop.Many {
		arity = "Many"
	}
	return fmt.Sprintf(`Property
URL: %s,
Arity: %s
Type: %s
`, prop.URL(), arity, prop.PropertyType.String())
}

// URL returns the URL of the property
func (prop PropertyDefinition) URL() string {
	return fmt.Sprintf("%s/%s", prop.Creator.String(), prop.Name)
}

func (ty PropertyType) String() string {
	names := [...]string{
		"Object",
		"String",
		"Integer",
		"Double",
		"Bool",
	}
	if int(ty) > len(names) {
		return ""
	}
	return names[ty]
}
