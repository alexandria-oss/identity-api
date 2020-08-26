package tracingutil

import (
	"errors"
	"github.com/alexandria-oss/common-go/exception"
	"go.opencensus.io/trace"
)

func CodeFromError(err error) int32 {
	switch {
	case errors.Is(err, exception.AlreadyExists):
		return trace.StatusCodeAlreadyExists
	case errors.Is(err, exception.NotFound):
		return trace.StatusCodeNotFound
	case errors.Is(err, exception.FieldRange):
		return trace.StatusCodeOutOfRange
	case errors.Is(err, exception.FieldFormat):
		return trace.StatusCodeInvalidArgument
	case errors.Is(err, exception.RequiredField):
		return trace.StatusCodeFailedPrecondition
	case errors.Is(err, exception.Invalid):
		return trace.StatusCodeInternal
	default:
		return trace.StatusCodeUnknown
	}
}
