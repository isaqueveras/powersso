// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package oops

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	defaultCode    int = 1000
	jsonCode       int = 2000
	internalCode   int = 3000
	pgxCode        int = 4000
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
func Handling(ctx *gin.Context, err error) {
	var e *Error

	if !errors.As(err, &e) {
		Handling(ctx, Err(err))
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
		err = handling(err)
	} else if err == e {
		err = handling(err)
	}

	return errors.WithStack(err)
}
