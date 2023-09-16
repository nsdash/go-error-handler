package error

import (
	"errors"
	customError "github.com/nsdash/go-error-lib"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcErrorHandler struct {
}

func NewGrpcErrorHandler() GrpcErrorHandler {
	return GrpcErrorHandler{}
}

func (h GrpcErrorHandler) Handle(err error) customError.Error {
	if statusErr, ok := status.FromError(err); ok {
		return h.parseError(statusErr)
	}

	applicationError := customError.NewApplicationError(errors.New("unknown error"), customError.CodeUnknown)

	return &applicationError
}

func (h GrpcErrorHandler) parseError(statusErr *status.Status) customError.Error {
	errorCode := statusErr.Code()
	errorMessage := statusErr.Message()

	applicationError := customError.NewApplicationError(errors.New(errorMessage), h.grpcErrorToCustom(errorCode))
	return &applicationError
}

func (h GrpcErrorHandler) grpcErrorToCustom(code codes.Code) int {
	if code == codes.InvalidArgument {
		return customError.CodeFieldIsInvalid
	}

	if code == codes.NotFound {
		return customError.CodeFieldEntityNotFound
	}

	if code == codes.AlreadyExists {
		return customError.CodeFieldIsUnique
	}

	return customError.CodeFieldUnauthorized
}
