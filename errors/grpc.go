package errros

import (
	"cosmossdk.io/errors"
	"google.golang.org/grpc/codes"
)

const moduleName = "grpc"

var (
	ErrCanceled          = errors.RegisterWithGRPCCode(moduleName, 1, codes.Canceled, "canceled")
	ErrUnknown           = errors.RegisterWithGRPCCode(moduleName, 2, codes.Unknown, "unknown")
	ErrInvalidArgument   = errors.RegisterWithGRPCCode(moduleName, 3, codes.InvalidArgument, "invalid argument")
	ErrNotFound          = errors.RegisterWithGRPCCode(moduleName, 5, codes.NotFound, "not found")
	ErrAlreadyExists     = errors.RegisterWithGRPCCode(moduleName, 6, codes.AlreadyExists, "already exists")
	ErrPermissionDenied  = errors.RegisterWithGRPCCode(moduleName, 7, codes.PermissionDenied, "permission denied")
	ErrResourceExhausted = errors.RegisterWithGRPCCode(moduleName, 8, codes.ResourceExhausted, "resource exhausted")
	ErrUnimplemented     = errors.RegisterWithGRPCCode(moduleName, 12, codes.Unimplemented, "unimplemented")
	ErrInternal          = errors.RegisterWithGRPCCode(moduleName, 13, codes.Internal, "internal")
	ErrUnavailable       = errors.RegisterWithGRPCCode(moduleName, 14, codes.Unavailable, "unavailable")
	ErrUnauthenticated   = errors.RegisterWithGRPCCode(moduleName, 16, codes.Unauthenticated, "unauthenticated")
)
