// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ecocredit

import (
	context "context"
	types "github.com/regen-network/regen-ledger/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// ClassInfo queries for information on a credit class.
	ClassInfo(ctx context.Context, in *QueryClassInfoRequest, opts ...grpc.CallOption) (*QueryClassInfoResponse, error)
	// BatchInfo queries for information on a credit batch.
	BatchInfo(ctx context.Context, in *QueryBatchInfoRequest, opts ...grpc.CallOption) (*QueryBatchInfoResponse, error)
	// Balance queries the balance (both tradable and retired) of a given credit
	// batch for a given account.
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
	// Supply queries the tradable and retired supply of a credit batch.
	Supply(ctx context.Context, in *QuerySupplyRequest, opts ...grpc.CallOption) (*QuerySupplyResponse, error)
	// Precision queries the number of decimal places that can be used to
	// represent credits in a batch. See Tx/SetPrecision for more details.
	Precision(ctx context.Context, in *QueryPrecisionRequest, opts ...grpc.CallOption) (*QueryPrecisionResponse, error)
}

type queryClient struct {
	cc         grpc.ClientConnInterface
	_ClassInfo types.Invoker
	_BatchInfo types.Invoker
	_Balance   types.Invoker
	_Supply    types.Invoker
	_Precision types.Invoker
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc: cc}
}

func (c *queryClient) ClassInfo(ctx context.Context, in *QueryClassInfoRequest, opts ...grpc.CallOption) (*QueryClassInfoResponse, error) {
	if invoker := c._ClassInfo; invoker != nil {
		var out QueryClassInfoResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._ClassInfo, err = invokerConn.Invoker("/regen.ecocredit.v1alpha1.Query/ClassInfo")
		if err != nil {
			var out QueryClassInfoResponse
			err = c._ClassInfo(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryClassInfoResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha1.Query/ClassInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BatchInfo(ctx context.Context, in *QueryBatchInfoRequest, opts ...grpc.CallOption) (*QueryBatchInfoResponse, error) {
	if invoker := c._BatchInfo; invoker != nil {
		var out QueryBatchInfoResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._BatchInfo, err = invokerConn.Invoker("/regen.ecocredit.v1alpha1.Query/BatchInfo")
		if err != nil {
			var out QueryBatchInfoResponse
			err = c._BatchInfo(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryBatchInfoResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha1.Query/BatchInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	if invoker := c._Balance; invoker != nil {
		var out QueryBalanceResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Balance, err = invokerConn.Invoker("/regen.ecocredit.v1alpha1.Query/Balance")
		if err != nil {
			var out QueryBalanceResponse
			err = c._Balance(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha1.Query/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Supply(ctx context.Context, in *QuerySupplyRequest, opts ...grpc.CallOption) (*QuerySupplyResponse, error) {
	if invoker := c._Supply; invoker != nil {
		var out QuerySupplyResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Supply, err = invokerConn.Invoker("/regen.ecocredit.v1alpha1.Query/Supply")
		if err != nil {
			var out QuerySupplyResponse
			err = c._Supply(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QuerySupplyResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha1.Query/Supply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Precision(ctx context.Context, in *QueryPrecisionRequest, opts ...grpc.CallOption) (*QueryPrecisionResponse, error) {
	if invoker := c._Precision; invoker != nil {
		var out QueryPrecisionResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._Precision, err = invokerConn.Invoker("/regen.ecocredit.v1alpha1.Query/Precision")
		if err != nil {
			var out QueryPrecisionResponse
			err = c._Precision(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryPrecisionResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha1.Query/Precision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ClassInfo queries for information on a credit class.
	ClassInfo(types.Context, *QueryClassInfoRequest) (*QueryClassInfoResponse, error)
	// BatchInfo queries for information on a credit batch.
	BatchInfo(types.Context, *QueryBatchInfoRequest) (*QueryBatchInfoResponse, error)
	// Balance queries the balance (both tradable and retired) of a given credit
	// batch for a given account.
	Balance(types.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
	// Supply queries the tradable and retired supply of a credit batch.
	Supply(types.Context, *QuerySupplyRequest) (*QuerySupplyResponse, error)
	// Precision queries the number of decimal places that can be used to
	// represent credits in a batch. See Tx/SetPrecision for more details.
	Precision(types.Context, *QueryPrecisionRequest) (*QueryPrecisionResponse, error)
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_ClassInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClassInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ClassInfo(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha1.Query/ClassInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ClassInfo(types.UnwrapSDKContext(ctx), req.(*QueryClassInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BatchInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBatchInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BatchInfo(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha1.Query/BatchInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BatchInfo(types.UnwrapSDKContext(ctx), req.(*QueryBatchInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha1.Query/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(types.UnwrapSDKContext(ctx), req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Supply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Supply(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha1.Query/Supply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Supply(types.UnwrapSDKContext(ctx), req.(*QuerySupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Precision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPrecisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Precision(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha1.Query/Precision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Precision(types.UnwrapSDKContext(ctx), req.(*QueryPrecisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.ecocredit.v1alpha1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClassInfo",
			Handler:    _Query_ClassInfo_Handler,
		},
		{
			MethodName: "BatchInfo",
			Handler:    _Query_BatchInfo_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
		{
			MethodName: "Supply",
			Handler:    _Query_Supply_Handler,
		},
		{
			MethodName: "Precision",
			Handler:    _Query_Precision_Handler,
		},
	},
	Metadata: "regen/ecocredit/v1alpha1/query.proto",
}

const (
	QueryClassInfoMethod = "/regen.ecocredit.v1alpha1.Query/ClassInfo"
	QueryBatchInfoMethod = "/regen.ecocredit.v1alpha1.Query/BatchInfo"
	QueryBalanceMethod   = "/regen.ecocredit.v1alpha1.Query/Balance"
	QuerySupplyMethod    = "/regen.ecocredit.v1alpha1.Query/Supply"
	QueryPrecisionMethod = "/regen.ecocredit.v1alpha1.Query/Precision"
)
