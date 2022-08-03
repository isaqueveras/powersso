// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package oops

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type wrappedError interface {
	Unwrap() error
}

// toError wraps errors to provide user readable messages
func toError(rawError error) error {
	message, code, responseStatus := "Unknown error", 0, http.StatusBadRequest

	switch err := rawError.(type) {
	case pgx.PgError:
		message, code = handlePgxError(&err)
		rawError = errors.Errorf("%s: %s", err.Error(), err.Hint)

	case *json.UnmarshalTypeError:
		message, code = fmt.Sprintf("Value type %v not supported in field %v. expected type %v", err.Value, err.Field, err.Type.String()), jsonCode+1

	case *reflect.ValueError:
		message, code = fmt.Sprintf("Cannot access value of type %v", err.Kind.String()), internalCode+1

	case *strconv.NumError:
		message, code = fmt.Sprintf("Unable to convert value %v", err.Num), internalCode+2

	case *time.ParseError:
		message, code = fmt.Sprintf("Impossible converter %v", err.Value), timeParseError+1

	case *Error:
		rawError, message, code, responseStatus = err, err.Message, err.Code, err.StatusCode

	case error:
		switch err {
		case sql.ErrNoRows:
			message, code, responseStatus = "Register not found", defaultCode+1, http.StatusNotFound

		case io.EOF:
			message, code = "No data available for reading", defaultCode+2

		case strconv.ErrSyntax:
			message, code = "Invalid format for string conversion", defaultCode+3
		}
	case nil:
		return nil
	}

	return &Error{
		Code:       code,
		Message:    message,
		Err:        rawError,
		StatusCode: responseStatus,
	}
}

func handlePgxError(err *pgx.PgError) (string, int) {
	switch err.Code {
	case "23505":
		return "Duplicate record", pgxCode + 1

	case "23502":
		return "Required data not specified", pgxCode + 2

	case "23503":
		return "Data indicated is not a valid reference", pgxCode + 3

	case "42P01", "42703":
		return "Incorrect access of elements in data records: Syntax error", pgxCode + 4

	case "42601", "42803", "42883":
		return "Incorrect use of function or operator when accessing data records: Syntax error", pgxCode + 5

	case "22001":
		return "Data exceeds record capacity in database", pgxCode + 6

	case "42702":
		return "Ambiguous reference: Syntax error", pgxCode + 7

	case "55P03":
		return "Required data is isolated and cannot be accessed", pgxCode + 8

	case "22P02":
		return "Specified data does not represent a valid type", pgxCode + 9
	}

	return "Unknown data error", pgxCode
}

// reconstructStackTrace reconstructs the stack trace
func reconstructStackTrace(err error, bound error) (output []string, traced bool) {
	var (
		wrapped wrappedError
		tracer  stackTracer
	)

	if errors.As(err, &wrapped) {
		internal := wrapped.Unwrap()

		// stop looking as we found our error instance
		if internal != bound {
			output, traced = reconstructStackTrace(internal, bound)
		}

		if !traced && errors.As(err, &tracer) {
			var stack errors.StackTrace = tracer.StackTrace()
			for _, frame := range stack {
				output = append(output, fmt.Sprintf("%+v", frame))
			}
			traced = true
		}
	}

	return
}
