package fixture

import (
	"context"
	"fmt"
	"reflect"

	"cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/gogo/protobuf/proto"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/grpc/encoding"
)

type router struct {
	cdc                encoding.Codec
	msgServiceRouter   *baseapp.MsgServiceRouter
	queryServiceRouter *baseapp.GRPCQueryRouter
}

func (rtr *router) invoker(methodName string, writeCondition func(context.Context, string, sdk.Msg) error) (Invoker, error) {
	return func(ctx context.Context, request interface{}, response interface{}, opts ...interface{}) error {
		req, ok := request.(proto.Message)
		if !ok {
			return fmt.Errorf("expected proto.Message, got %T for service method %s", request, methodName)
		}

		typeURL := TypeURL(req)

		msg, isMsg := request.(sdk.Msg)

		// cache wrap the multistore so that inter-module writes are atomic
		// see https://github.com/cosmos/cosmos-sdk/issues/8030
		regenCtx := sdk.UnwrapSDKContext(ctx)
		cacheMs := regenCtx.MultiStore().CacheMultiStore()
		ctx = sdk.WrapSDKContext(regenCtx.WithMultiStore(cacheMs))

		// msg handler
		if writeCondition != nil && isMsg {
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			err = writeCondition(ctx, methodName, msg)
			if err != nil {
				return err
			}

			// routing using baseapp.MsgServiceRouter
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			handler := rtr.msgServiceRouter.HandlerByTypeURL(typeURL)
			if handler == nil {
				return errors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized message route: %s;", typeURL)
			}

			res, err := handler(sdkCtx, msg)
			if err != nil {
				return err
			}
			// set events from response in the sdk context event manager.
			for _, e := range res.Events {
				sdkCtx.EventManager().EmitEvent(sdk.Event(e))
			}

			if len(res.MsgResponses) != 1 {
				panic(fmt.Sprintf("expected 1 msg response, got %d", len(res.MsgResponses)))
			}
			resValue := reflect.ValueOf(res.MsgResponses[0].GetCachedValue())
			reflect.ValueOf(response).Elem().Set(resValue.Elem())

			// only commit writes if there is no error so that calls are atomic
			cacheMs.Write()
		} else {
			// route query here
			handler := rtr.queryServiceRouter.Route(methodName)
			if handler == nil {
				panic(fmt.Sprintf("no handler found for %s", methodName))
			}
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			bz, err := rtr.cdc.Marshal(request)
			if err != nil {
				return err
			}
			queryResponse, err := handler(sdkCtx, abciTypes.RequestQuery{
				Data: bz,
			})
			if err != nil {
				return err
			}
			if err := rtr.cdc.Unmarshal(queryResponse.Value, response); err != nil {
				return err
			}
			cacheMs.Write()
		}
		return nil

	}, nil
}

func (rtr *router) testTxFactory(signers []sdk.AccAddress) InvokerFactory {
	signerMap := map[string]bool{}
	for _, signer := range signers {
		signerMap[signer.String()] = true
	}

	return func(callInfo CallInfo) (Invoker, error) {
		return rtr.invoker(callInfo.Method, func(_ context.Context, _ string, req sdk.Msg) error {
			for _, signer := range req.GetSigners() {
				if _, found := signerMap[signer.String()]; !found {
					return sdkerrors.ErrUnauthorized
				}
			}
			return nil
		})
	}
}

func (rtr *router) testQueryFactory() InvokerFactory {
	return func(callInfo CallInfo) (Invoker, error) {
		return rtr.invoker(callInfo.Method, nil)
	}
}

func TypeURL(req proto.Message) string {
	return "/" + proto.MessageName(req)
}
