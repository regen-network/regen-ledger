package keeper

import (
	"context"
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

type EventReducer struct {
	reducerMap map[reflect.Type]func(context.Context, proto.Message) error
}

func NewEventReducer(reducers ...any) EventReducer {
	reducerMap := map[reflect.Type]func(context.Context, proto.Message) error{}
	for _, reducer := range reducers {
		typ := reflect.TypeOf(reducer)
		numMethods := typ.NumMethod()
		for i := 0; i < numMethods; i++ {
			method := typ.Method(i)
			if !method.IsExported() {
				continue
			}

			if method.Type.NumIn() != 2 || method.Type.NumOut() != 1 {
				continue
			}

			if method.Type.In(0) != reflect.TypeOf((*context.Context)(nil)).Elem() {
				continue
			}

			if method.Type.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
				continue
			}

			evtType := method.Type.In(1)

			if !evtType.Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
				fmt.Printf("reduder method %s does not take a proto.Message as its second argument\n", method.Name)
				continue
			}

			fmt.Printf("registering reducer for event type %s\n", evtType)
			reducerMap[evtType] = func(ctx context.Context, evt proto.Message) error {
				return method.Func.Call([]reflect.Value{reflect.ValueOf(reducer), reflect.ValueOf(ctx), reflect.ValueOf(evt)})[0].Interface().(error)
			}
		}
	}
	return EventReducer{
		reducerMap: reducerMap,
	}
}

func (er EventReducer) Reduce(ctx context.Context, evt proto.Message) error {
	reducer, ok := er.reducerMap[reflect.TypeOf(evt)]
	if !ok {
		fmt.Printf("no reducer found for event type %T\n", evt)
		return nil
	}
	return reducer(ctx, evt)
}

type EventEmitter struct {
	redux EventReducer
}

func (er EventReducer) Emitter() EventEmitter {
	return EventEmitter{
		redux: er,
	}
}

func (ee EventEmitter) Emit(ctx context.Context, evt proto.Message) error {
	err := ee.redux.Reduce(ctx, evt)
	if err != nil {
		return err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.EventManager().EmitTypedEvent(evt)
}
