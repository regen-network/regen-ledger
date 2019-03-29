package schema

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/graph"
)

type onChainSchemaResolver struct {
	keeper Keeper
	ctx    sdk.Context
}

func NewOnChainSchemaResolver(keeper Keeper, ctx sdk.Context) graph.SchemaResolver {
	return &onChainSchemaResolver{keeper: keeper, ctx: ctx}
}

func (res onChainSchemaResolver) GetPropertyID(p graph.Property) graph.PropertyID {
	return res.keeper.GetPropertyID(res.ctx, p.URI().String())
}

func (res onChainSchemaResolver) GetPropertyByURL(url string) graph.Property {
	id := res.keeper.GetPropertyID(res.ctx, url)
	if id == 0 {
		return nil
	}
	return res.GetPropertyByID(id)
}

func (res onChainSchemaResolver) GetPropertyByID(id graph.PropertyID) graph.Property {
	prop, found := res.keeper.GetPropertyDefinition(res.ctx, id)
	if !found {
		return nil
	}
	return NewProperty(prop, prop.URI())
}

func (res onChainSchemaResolver) MinPropertyID() graph.PropertyID {
	return 0
}

func (res onChainSchemaResolver) MaxPropertyID() graph.PropertyID {
	return res.keeper.GetLastPropertyID(res.ctx)
}
