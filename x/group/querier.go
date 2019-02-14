package agent

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"strconv"
)

// query endpoints supported by the governance Querier
const (
	QueryAgent = "get"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAgent:
			return queryAgent(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown data query endpoint")
		}
	}
}


func queryAgent(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	idStr := path[0]

	id, parseErr := strconv.ParseUint(idStr, 10, 64)
	if parseErr != nil {
		return []byte{}, sdk.ErrUnknownRequest("Can't parse id")
	}

	info, err := keeper.GetAgentInfo(ctx, AgentID(id))

	if err != nil {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve id")
	}

	res, jsonErr := codec.MarshalJSONIndent(keeper.cdc, info)
	if jsonErr != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", jsonErr.Error()))
	}
	return res, nil
}
