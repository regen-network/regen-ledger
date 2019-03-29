package binary

import (
	"github.com/regen-network/regen-ledger/graph"
	"github.com/regen-network/regen-ledger/graph/impl"
	"github.com/regen-network/regen-ledger/x/schema"
)
import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	prefixGeoAddress byte = iota
	prefixAccAddress
	prefixHashID

	prefixIDMin = prefixGeoAddress
	prefixIDMax = prefixHashID
)

const (
	prefixPropertyID byte = 0
)

// SchemaResolver resolves properties against a schema
type SchemaResolver interface {
	GetPropertyByURL(url string) graph.Property
	GetPropertyByID(id schema.PropertyID) graph.Property
	GetPropertyID(p graph.Property) schema.PropertyID
}

type onChainSchemaResolver struct {
	keeper schema.Keeper
	ctx    sdk.Context
}

func (res onChainSchemaResolver) GetPropertyID(p graph.Property) schema.PropertyID {
	return res.keeper.GetPropertyID(res.ctx, p.URI().String())
}

func (res onChainSchemaResolver) GetPropertyByURL(url string) graph.Property {
	id := res.keeper.GetPropertyID(res.ctx, url)
	if id == 0 {
		return nil
	}
	return res.GetPropertyByID(id)
}

func (res onChainSchemaResolver) GetPropertyByID(id schema.PropertyID) graph.Property {
	prop, found := res.keeper.GetPropertyDefinition(res.ctx, id)
	if !found {
		return nil
	}
	return impl.NewProperty(prop, id, prop.URI())
}
