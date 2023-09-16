package error

import (
	"errors"
	customError "github.com/nsdash/go-error-lib"
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

	applicationError := customError.NewApplicationError(errors.New(errorMessage), int(errorCode))
	return &applicationError
}
