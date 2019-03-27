package schema

import (
	"encoding/json"
	"fmt"
)
import sdk "github.com/cosmos/cosmos-sdk/types"

// Route returns the route of the message
func (PropertyDefinition) Route() string { return "schema" }

// Type returns the type of the message
func (PropertyDefinition) Type() string { return "schema.define-property" }

// ValidateBasic ensures that Creator and Name are non-empty and the PropertyType is valid
func (prop PropertyDefinition) ValidateBasic() sdk.Error {
	if len(prop.Creator) == 0 {
		return sdk.ErrUnknownRequest("property creator cannot be empty")
	}
	if len(prop.Name) == 0 {
		return sdk.ErrUnknownRequest("property name cannot be empty")
	}
	if prop.PropertyType.String() == "" {
		return sdk.ErrUnknownRequest(fmt.Sprintf("unknown PropertyType %d", prop.PropertyType))
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
