package server

import (
	"bytes"
	"context"
	"fmt"
	"reflect"

	"github.com/regen-network/regen-ledger/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogogrpc "github.com/gogo/protobuf/grpc"
	"google.golang.org/grpc"
)

type handler struct {
	f            func(ctx context.Context, args, reply interface{}) error
	commitWrites bool
	moduleName   string
}

type router struct {
	handlers         map[string]handler
	providedServices map[reflect.Type]bool
	antiReentryMap   map[string]bool
	authzMiddleware  AuthorizationMiddleware
	legacyRouter     sdk.Router
}

type registrar struct {
	*router
	baseServer   gogogrpc.Server
	commitWrites bool
	moduleName   string
}

var _ gogogrpc.Server = registrar{}

func (r registrar) RegisterService(sd *grpc.ServiceDesc, ss interface{}) {
	r.providedServices[reflect.TypeOf(sd.HandlerType)] = true

	r.baseServer.RegisterService(sd, ss)

	for _, method := range sd.Methods {
		fqName := fmt.Sprintf("/%s/%s", sd.ServiceName, method.MethodName)
		methodHandler := method.Handler
		f := func(ctx context.Context, args, reply interface{}) error {
			res, err := methodHandler(ss, ctx, func(i interface{}) error { return nil },
				func(ctx context.Context, _ interface{}, _ *grpc.UnaryServerInfo, unaryHandler grpc.UnaryHandler) (resp interface{}, err error) {
					return unaryHandler(ctx, args)
				})
			if err != nil {
				return err
			}

			resValue := reflect.ValueOf(res)
			if !resValue.IsZero() && reply != nil {
				reflect.ValueOf(reply).Elem().Set(resValue.Elem())
			}
			return nil
		}
		r.handlers[fqName] = handler{
			f:            f,
			commitWrites: r.commitWrites,
			moduleName:   r.moduleName,
		}
	}
}

func (rtr *router) invoker(methodName string, writeCondition func(context.Context, string, sdk.MsgRequest) error) (types.Invoker, error) {
	var handler handler
	// In case of ServiceMsg, we can use ADR 033 router handler
	if isServiceMsg(methodName) {
		h, found := rtr.handlers[methodName]
		handler = h
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("cannot find method named %s", methodName))
		}
	}

	moduleName := handler.moduleName

	return func(ctx context.Context, request interface{}, response interface{}, opts ...interface{}) error {
		msg, isMsg := request.(sdk.Msg)
		// msg handler
		if writeCondition != nil && (handler.commitWrites || isMsg) {
			if rtr.antiReentryMap[moduleName] {
				return fmt.Errorf("re-entrant module calls not allowed for security reasons! module %s is already on the call stack", moduleName)
			}

			rtr.antiReentryMap[moduleName] = true
			defer delete(rtr.antiReentryMap, moduleName)

			msgReq, ok := request.(sdk.MsgRequest)
			if !ok {
				return fmt.Errorf("expected %T, got %T", (*sdk.MsgRequest)(nil), request)
			}

			err := msgReq.ValidateBasic()
			if err != nil {
				return err
			}

			err = writeCondition(ctx, methodName, msgReq)
			if err != nil {
				return err
			}

			// cache wrap the multistore so that inter-module writes are atomic
			// see https://github.com/cosmos/cosmos-sdk/issues/8030
			regenCtx := types.UnwrapSDKContext(ctx)
			cacheMs := regenCtx.MultiStore().CacheMultiStore()
			ctx = sdk.WrapSDKContext(regenCtx.WithMultiStore(cacheMs))

			if isServiceMsg(methodName) {
				err = handler.f(ctx, request, response)
				if err != nil {
					return err
				}
			} else {
				// legacy sdk.Msg routing using sdk.Router
				// for routing non ServiceMsg
				msgRoute := msg.Route()
				sdkCtx := sdk.UnwrapSDKContext(ctx)
				handler := rtr.legacyRouter.Route(sdkCtx, msgRoute)
				if handler == nil {
					return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized message route: %s;", msgRoute)
				}

				_, err = handler(sdkCtx, msg)
				if err != nil {
					return err
				}
			}

			// only commit writes if there is no error so that calls are atomic
			cacheMs.Write()
		} else {
			// query handler

			// cache wrap the multistore so that writes are batched
			sdkCtx := types.UnwrapSDKContext(ctx)
			cacheMs := sdkCtx.MultiStore().CacheMultiStore()
			ctx = types.Context{Context: sdkCtx.WithMultiStore(cacheMs)}

			err := handler.f(ctx, request, response)
			if err != nil {
				return err
			}

			cacheMs.Write()
		}
		return nil

	}, nil

}

func (rtr *router) invokerFactory(moduleName string) InvokerFactory {
	return func(callInfo CallInfo) (types.Invoker, error) {
		moduleID := callInfo.Caller
		if moduleName != moduleID.ModuleName {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
				fmt.Sprintf("expected a call from module %s, but module %s is calling", moduleName, moduleID.ModuleName))
		}

		moduleAddr := moduleID.Address()

		writeCondition := func(ctx context.Context, methodName string, msgReq sdk.MsgRequest) error {
			signers := msgReq.GetSigners()
			if len(signers) != 1 {
				return fmt.Errorf("inter module Msg invocation requires a single expected signer (%s), but %s expects multiple signers (%+v),  ", moduleAddr, methodName, signers)
			}

			signer := signers[0]

			if bytes.Equal(moduleAddr, signer) {
				return nil
			}

			if rtr.authzMiddleware != nil && rtr.authzMiddleware(sdk.UnwrapSDKContext(ctx), methodName, msgReq, moduleAddr) {
				return nil
			}

			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
				fmt.Sprintf("expected %s, got %s", signers[0], moduleAddr))
		}

		return rtr.invoker(callInfo.Method, writeCondition)
	}
}

func (rtr *router) testTxFactory(signers []sdk.AccAddress) InvokerFactory {
	signerMap := map[string]bool{}
	for _, signer := range signers {
		signerMap[signer.String()] = true
	}

	return func(callInfo CallInfo) (types.Invoker, error) {
		return rtr.invoker(callInfo.Method, func(_ context.Context, _ string, req sdk.MsgRequest) error {
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
	return func(callInfo CallInfo) (types.Invoker, error) {
		return rtr.invoker(callInfo.Method, nil)
	}
}
