package xrb

import "github.com/regen-network/regen-ledger/x/schema"
import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	prefixGeoAddress byte = iota
	prefixAccAddress
	prefixDataAddress
	prefixHashID

	prefixIDMin = prefixGeoAddress
	prefixIDMax = prefixHashID
)

const (
	prefixPropertyID byte = 0
)

type SchemaResolver interface {
	GetPropertyByURL(url string) Property
	GetPropertyByID(id schema.PropertyID) Property
}

type onChainSchemaResolver struct {
	keeper schema.Keeper
	ctx    sdk.Context
}

func (res onChainSchemaResolver) GetPropertyByURL(url string) Property {
	id := res.keeper.GetPropertyID(res.ctx, url)
	if id == 0 {
		return nil
	}
	return res.GetPropertyByID(id)
}

func (res onChainSchemaResolver) GetPropertyByID(id schema.PropertyID) Property {
	prop, found := res.keeper.GetProperty(res.ctx, id)
	if !found {
		return nil
	}
	return property{prop, id, prop.URI()}
}
