package errros

import (
	"cosmossdk.io/errors"
	"google.golang.org/grpc/codes"
)

const moduleName = "grpc"

var (
	ErrCancelled        = errors.RegisterWithGRPCCode(moduleName, 1, codes.Canceled, "cancelled")
	ErrUnkown           = errors.RegisterWithGRPCCode(moduleName, 2, codes.Unknown, "unknown")
	ErrInvalidArgument  = errors.RegisterWithGRPCCode(moduleName, 3, codes.InvalidArgument, "invalid argument")
	ErrNotFound         = errors.RegisterWithGRPCCode(moduleName, 4, codes.NotFound, "not found")
	ErrAlreadyExists    = errors.RegisterWithGRPCCode(moduleName, 5, codes.AlreadyExists, "already exists")
	ErrPermissionDenied = errors.RegisterWithGRPCCode(moduleName, 6, codes.PermissionDenied, "permission denied")
	ErrInternal         = errors.RegisterWithGRPCCode(moduleName, 7, codes.Internal, "internal")
	ErrUnavailable      = errors.RegisterWithGRPCCode(moduleName, 8, codes.Unavailable, "unavailable")
	ErrUnauthenticated  = errors.RegisterWithGRPCCode(moduleName, 9, codes.Unauthenticated, "unauthenticated")
	ErrUnimplemented    = errors.RegisterWithGRPCCode(moduleName, 10, codes.Unimplemented, "unimplemented")
)
