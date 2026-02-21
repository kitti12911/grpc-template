package interceptor

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"grpc-template/internal/apperror"
)

var codeMap = map[apperror.Code]codes.Code{
	apperror.CodeInternal:     codes.Internal,
	apperror.CodeNotFound:     codes.NotFound,
	apperror.CodeAlreadyExist: codes.AlreadyExists,
	apperror.CodeInvalidInput: codes.InvalidArgument,
	apperror.CodeUnauthorized: codes.Unauthenticated,
	apperror.CodeForbidden:    codes.PermissionDenied,
}

func ErrorHandler() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		resp, err := handler(ctx, req)

		if err == nil {
			return resp, nil
		}

		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		if appErr, ok := errors.AsType[*apperror.Error](err); ok {
			slog.ErrorContext(ctx, "request failed",
				"method", info.FullMethod,
				"error", err.Error(),
			)

			grpcCode, exists := codeMap[appErr.Code()]

			if !exists {
				grpcCode = codes.Internal
			}

			return nil, status.Error(grpcCode, appErr.Message())
		}

		slog.ErrorContext(ctx, "unexpected error",
			"method", info.FullMethod,
			"error", err.Error(),
		)

		return nil, status.Error(codes.Internal, "internal server error")
	}
}
