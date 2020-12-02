package server

import (
	"bytes"
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogogrpc "github.com/gogo/protobuf/grpc"
	"google.golang.org/grpc"
	"reflect"
)

type handler struct {
	f            func(ctx context.Context, args, reply interface{}) error
	commitWrites bool
}

type router struct {
	handlers         map[string]handler
	providedServices map[reflect.Type]bool
}

type registrar struct {
	*router
	baseServer   gogogrpc.Server
	commitWrites bool
}

var _ gogogrpc.Server = registrar{}

func (t registrar) RegisterService(sd *grpc.ServiceDesc, ss interface{}) {
	t.providedServices[reflect.TypeOf(sd.HandlerType)] = true

	t.baseServer.RegisterService(sd, ss)

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
			if !resValue.IsZero() {
				reflect.ValueOf(reply).Elem().Set(resValue.Elem())
			}
			return nil
		}
		t.handlers[fqName] = handler{
			f:            f,
			commitWrites: t.commitWrites,
		}
	}
}

func (t router) invoker(methodName string, writeCondition func(sdk.MsgRequest) error) (Invoker, error) {
	handler, found := t.handlers[methodName]
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("cannot find method named %s", methodName))
	}

	if writeCondition != nil && handler.commitWrites {
		// msg handler
		return func(ctx context.Context, request interface{}, response interface{}, opts ...interface{}) error {
			msgReq, ok := request.(sdk.MsgRequest)
			if !ok {
				return fmt.Errorf("expected %T, got %T", (*sdk.MsgRequest)(nil), request)
			}

			err := msgReq.ValidateBasic()
			if err != nil {
				return err
			}

			err = writeCondition(msgReq)
			if err != nil {
				return err
			}

			// cache wrap the multistore so that writes are batched
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			ms := sdkCtx.MultiStore()
			cacheMs := ms.CacheMultiStore()
			sdkCtx = sdkCtx.WithMultiStore(cacheMs)
			ctx = sdk.WrapSDKContext(sdkCtx)

			err = handler.f(ctx, request, response)
			if err != nil {
				return err
			}

			cacheMs.Write()

			return nil
		}, nil
	} else {
		// query handler
		return func(ctx context.Context, request interface{}, response interface{}, opts ...interface{}) error {
			// cache wrap the multistore so that writes are batched
			sdkCtx := sdk.UnwrapSDKContext(ctx)
			ms := sdkCtx.MultiStore()
			cacheMs := ms.CacheMultiStore()
			sdkCtx = sdkCtx.WithMultiStore(cacheMs)
			ctx = sdk.WrapSDKContext(sdkCtx)

			err := handler.f(ctx, request, response)
			if err != nil {
				return err
			}

			cacheMs.Write()

			return nil
		}, nil
	}
}

func (t router) invokerFactory(moduleName string) InvokerFactory {
	return func(callInfo CallInfo) (Invoker, error) {
		moduleId := callInfo.Caller
		if moduleName != moduleId.ModuleName {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
				fmt.Sprintf("expected a call from module %s, but module %s is calling", moduleName, moduleId.ModuleName))
		}

		addr := moduleId.Address()

		writeCondition := func(msgReq sdk.MsgRequest) error {
			signers := msgReq.GetSigners()
			if len(signers) != 1 {
				return fmt.Errorf("expected a signle signer %s, got %+v", addr, signers)
			}

			if !bytes.Equal(addr, signers[0]) {
				return sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
					fmt.Sprintf("expected %s, got %s", signers[0], addr))
			}

			return nil
		}

		return t.invoker(callInfo.Method, writeCondition)
	}
}

func (t router) testTxFactory(signers []sdk.AccAddress) InvokerFactory {
	signerMap := map[string]bool{}
	for _, signer := range signers {
		signerMap[signer.String()] = true
	}

	return func(callInfo CallInfo) (Invoker, error) {
		return t.invoker(callInfo.Method, func(req sdk.MsgRequest) error {
			for _, signer := range req.GetSigners() {
				if _, found := signerMap[signer.String()]; !found {
					return sdkerrors.ErrUnauthorized
				}
			}
			return nil
		})
	}
}

func (t router) testQueryFactory() InvokerFactory {
	return func(callInfo CallInfo) (Invoker, error) {
		return t.invoker(callInfo.Method, nil)
	}
}
