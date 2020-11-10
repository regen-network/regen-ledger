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
	handlers map[string]handler
}

type registrar struct {
	*router
	baseServer   gogogrpc.Server
	commitWrites bool
}

var _ gogogrpc.Server = registrar{}

func (t registrar) RegisterService(sd *grpc.ServiceDesc, ss interface{}) {
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

func (t router) invoker(methodName string, addr sdk.AccAddress) (Invoker, error) {
	handler, found := t.handlers[methodName]
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("cannot find method named %s", methodName))
	}

	i := func(ctx context.Context, request interface{}, response interface{}, opts ...interface{}) error {
		msgReq, ok := request.(sdk.MsgRequest)
		if !ok {
			return fmt.Errorf("expected %T, got %T", (*sdk.MsgRequest)(nil), request)
		}

		err := msgReq.ValidateBasic()
		if err != nil {
			return err
		}

		signers := msgReq.GetSigners()
		if len(signers) != 1 {
			return fmt.Errorf("expected a signle signer %s, got %+v", addr, signers)
		}

		if !bytes.Equal(addr, signers[0]) {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
				fmt.Sprintf("expected %s, got %s", signers[0], addr))
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

		// only commit writes if there are no errors and commitWrites is true
		if handler.commitWrites {
			cacheMs.Write()
		}

		return nil
	}

	return i, nil
}

func (t router) invokerFactory(moduleName string) InvokerFactory {
	return func(callInfo CallInfo) (Invoker, error) {
		moduleId := callInfo.Caller
		if moduleName != moduleId.ModuleName {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
				fmt.Sprintf("expected a call from module %s, but module %s is calling", moduleName, moduleId.ModuleName))
		}

		return t.invoker(callInfo.Method, moduleId.Address())
	}
}
