package data

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the governance Querier
const (
	QueryData = "get"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryData:
			return queryData(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown data query endpoint")
		}
	}
}

func fromHex(str string) []byte {
	bz, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return bz
}

func queryData(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	id := path[0]

	value, err := keeper.GetData(ctx, fromHex(id))
	if err != nil {
		return nil, err
	}

	if len(value) == 0 {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve id")
	}

	return value, nil
}
