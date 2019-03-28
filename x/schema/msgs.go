package schema

import (
	"encoding/json"
	"fmt"
	"regexp"
)
import sdk "github.com/cosmos/cosmos-sdk/types"

// Route returns the route of the message
func (PropertyDefinition) Route() string { return "schema" }

// Type returns the type of the message
func (PropertyDefinition) Type() string { return "schema.define-property" }

// PropertyNameRegex defines the valid characters for property names. Only lowercase ascii letters are allowed, property
// names must start with a letter and can otherwise also contain numbers and underscores. Snake case is preferred
// as a number of sources point suggest it has enhanced readability as opposed to camel case
const PropertyNameRegex = "^[a-z][a-z0-9_]+$"

var propertyNameRegexCompiled = regexp.MustCompile(PropertyNameRegex)

// ValidateBasic ensures that Creator and Name are non-empty and the PropertyType is valid
func (prop PropertyDefinition) ValidateBasic() sdk.Error {
	if len(prop.Creator) == 0 {
		return sdk.ErrUnknownRequest("property creator cannot be empty")
	}
	if len(prop.Name) == 0 {
		return sdk.ErrUnknownRequest("property name cannot be empty")
	}
	if !propertyNameRegexCompiled.MatchString(prop.Name) {
		return sdk.ErrUnknownRequest(fmt.Sprintf("property name %s doesn't conform to regex %s", prop.Name, PropertyNameRegex))
	}
	if prop.PropertyType.String() == "" {
		return sdk.ErrUnknownRequest(fmt.Sprintf("unknown PropertyType %d", prop.PropertyType))
	}
	if prop.Arity.String() == "" {
		return sdk.ErrUnknownRequest(fmt.Sprintf("unknown Arity %d", prop.PropertyType))
	}
	return nil
}

// GetSignBytes returns the bytes over which to sign
func (prop PropertyDefinition) GetSignBytes() []byte {
	b, err := json.Marshal(prop)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners returns the addresses which must sign the message
func (prop PropertyDefinition) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{prop.Creator}
}
