package erros

import (
	"cosmossdk.io/errors"
	"google.golang.org/grpc/codes"
)

const codeSpace = "grpc"

var (
	ErrCanceled           = errors.RegisterWithGRPCCode(codeSpace, 1, codes.Canceled, "canceled")
	ErrUnknown            = errors.RegisterWithGRPCCode(codeSpace, 2, codes.Unknown, "unknown")
	ErrInvalidArgument    = errors.RegisterWithGRPCCode(codeSpace, 3, codes.InvalidArgument, "invalid argument")
	ErrDeadlineExceeded   = errors.RegisterWithGRPCCode(codeSpace, 4, codes.DeadlineExceeded, "deadline exceeded")
	ErrNotFound           = errors.RegisterWithGRPCCode(codeSpace, 5, codes.NotFound, "not found")
	ErrAlreadyExists      = errors.RegisterWithGRPCCode(codeSpace, 6, codes.AlreadyExists, "already exists")
	ErrPermissionDenied   = errors.RegisterWithGRPCCode(codeSpace, 7, codes.PermissionDenied, "permission denied")
	ErrResourceExhausted  = errors.RegisterWithGRPCCode(codeSpace, 8, codes.ResourceExhausted, "resource exhausted")
	ErrFailedPrecondition = errors.RegisterWithGRPCCode(codeSpace, 9, codes.FailedPrecondition, "failed precondition")
	ErrAborted            = errors.RegisterWithGRPCCode(codeSpace, 10, codes.Aborted, "aborted")
	ErrOutOfRange         = errors.RegisterWithGRPCCode(codeSpace, 11, codes.OutOfRange, "out of range")
	ErrUnimplemented      = errors.RegisterWithGRPCCode(codeSpace, 12, codes.Unimplemented, "unimplemented")
	ErrInternal           = errors.RegisterWithGRPCCode(codeSpace, 13, codes.Internal, "internal")
	ErrUnavailable        = errors.RegisterWithGRPCCode(codeSpace, 14, codes.Unavailable, "unavailable")
	ErrDataLoss           = errors.RegisterWithGRPCCode(codeSpace, 15, codes.DataLoss, "data loss")
	ErrUnauthenticated    = errors.RegisterWithGRPCCode(codeSpace, 16, codes.Unauthenticated, "unauthenticated")
)
