package oops

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	pgxCode        int = 1000
	jsonCode       int = 2000
	internalCode   int = 3000
	defaultCode    int = 4000
	timeParseError int = 5000
)

// Error defines the properties for a basic error response
type Error struct {
	Code       int      `json:"code"`
	Message    string   `json:"message"`
	Trace      []string `json:"-"`
	Err        error    `json:"-"`
	StatusCode int      `json:"-"`
}

// NewError creates and returns new normalized `Error` instance.
func NewError(message string, statusCode int) *Error {
	return &Error{
		Code:       defaultCode,
		Message:    message,
		Trace:      []string{},
		Err:        nil,
		StatusCode: statusCode,
	}
}

// Error makes it compatible with the `error` interface
func (e *Error) Error() string {
	return e.Message
}

// Wrap wraps an error adding an information message
func Wrap(err error, message string) error {
	return errors.Wrap(Err(err), message)
}

// NotFoundError creates and returns 404 `Error`.
func NotFoundError(message string) *Error {
	if message == "" {
		message = "The requested resource wasn't found."
	}

	return NewError(message, http.StatusNotFound)
}

// Handling handles an error by setting a message and a response status code
func Handling(err error, ctx *gin.Context) {
	var e *Error

	if !errors.As(err, &e) {
		Handling(Err(err), ctx)
		return
	}

	e.Message = err.Error()
	e.Trace, _ = reconstructStackTrace(err, e)

	if reqIDVal, hasRID := ctx.Get("RID"); hasRID {
		if reqID, ok := reqIDVal.(string); ok && len(reqID) > 0 {
			e.Message = fmt.Sprintf("%s [%s]", e.Message, reqID[:6])
		}
	}

	ctx.AbortWithStatusJSON(e.StatusCode, e)
	ctx.Set("error", err)
}

// Err builds annotated error instance from any error value
func Err(err error) error {
	var e *Error
	if !errors.As(err, &e) {
		err = toError(err)
	} else if err == e {
		err = toError(err)
	}

	return errors.WithStack(err)
}
