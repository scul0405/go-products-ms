package grpc_errors

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrNoCtxMetaData    = errors.New("no ctx metadata")
	ErrInvalidSessionId = errors.New("invalid session id")
	ErrEmailExists      = errors.New("email already exists")
)

func ParseGRPCErrStatusCode(err error) codes.Code {
	switch {
	case errors.Is(err, ErrNotFound):
		return codes.NotFound
	case errors.Is(err, ErrNoCtxMetaData):
		return codes.Unauthenticated
	case errors.Is(err, ErrInvalidSessionId):
		return codes.PermissionDenied
	case errors.Is(err, ErrEmailExists):
		return codes.AlreadyExists
	case errors.Is(err, sql.ErrNoRows):
		return codes.NotFound
	case errors.Is(err, context.Canceled):
		return codes.Canceled
	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded
	case strings.Contains(err.Error(), "Validate"):
		return codes.InvalidArgument
	case strings.Contains(err.Error(), "redis"):
		return codes.NotFound
	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded
	}

	return codes.Internal
}

// MapGRPCErrCodeToHttpStatus Map GRPC errors codes to http status
func MapGRPCErrCodeToHttpStatus(code codes.Code) int {
	switch code {
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.AlreadyExists:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.InvalidArgument:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

// ErrorResponse GRPC Error response
func ErrorResponse(err error, msg string) error {
	return status.Errorf(ParseGRPCErrStatusCode(err), fmt.Sprintf("%s: %v", msg, err))
}
