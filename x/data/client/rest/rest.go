package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gorilla/mux"
)

const (
//restId = "id"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {
	//r.HandleFunc(fmt.Sprintf("/%s/data/{%s}", storeName, restId), getDataHandler(cdc, cliCtx, storeName)).Methods("GET")
}

//func getDataHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		vars := mux.Vars(r)
//		paramType := vars[restId]
//
//		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/get/%s", storeName, paramType), nil)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
//			return
//		}
//
//		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
//	}
//}
