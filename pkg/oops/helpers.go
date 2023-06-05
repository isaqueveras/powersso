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

	"github.com/isaqueveras/powersso/pkg/i18n"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type wrappedError interface {
	Unwrap() error
}

// handling wraps errors to provide user readable messages
func handling(rawError error) error {
	message, code, responseStatus := i18n.Value("errors.default"), 0, http.StatusBadRequest

	switch err := rawError.(type) {
	case pgx.PgError:
		message, code = handlePgxError(&err)
		rawError = errors.Errorf("%s: %s", err.Error(), err.Hint)

	case *json.UnmarshalTypeError:
		message, code = i18n.Value("errors.handling.json_unmarshal_type_error", err.Value, err.Field, err.Type.String()), jsonCode+1

	case *reflect.ValueError:
		message, code = i18n.Value("errors.handling.reflect_value_error", err.Kind.String()), internalCode+1

	case *strconv.NumError:
		message, code = i18n.Value("errors.handling.strconv_num_error", err.Num), internalCode+2

	case *time.ParseError:
		message, code = i18n.Value("errors.handling.time_parse", err.Value), timeParseError+1

	case *Error:
		rawError, message, code, responseStatus = err, err.Message, err.Code, err.StatusCode

	case error:
		switch err {
		case sql.ErrNoRows:
			message, code, responseStatus = i18n.Value("errors.handling.error.sql_no_rows"), defaultCode+1, http.StatusNotFound

		case io.EOF:
			message, code = i18n.Value("errors.handling.error.io_eof"), defaultCode+2

		case strconv.ErrSyntax:
			message, code = i18n.Value("errors.handling.error.strconv_err_syntax"), defaultCode+3
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
		return i18n.Value("errors.handling.pgx.23505"), pgxCode + 1
	case "23502":
		return i18n.Value("errors.handling.pgx.23502"), pgxCode + 2
	case "23503":
		return i18n.Value("errors.handling.pgx.23503"), pgxCode + 3
	case "42P01", "42703":
		return i18n.Value("errors.handling.pgx.42P01_42703"), pgxCode + 4
	case "42601", "42803", "42883":
		return i18n.Value("errors.handling.pgx.42601_42803_42883"), pgxCode + 5
	case "22001":
		return i18n.Value("errors.handling.pgx.22001"), pgxCode + 6
	case "42702":
		return i18n.Value("errors.handling.pgx.42702"), pgxCode + 7
	case "55P03":
		return i18n.Value("errors.handling.pgx.55P03"), pgxCode + 8
	case "22P02":
		return i18n.Value("errors.handling.pgx.22P02"), pgxCode + 9
	case "25006":
		return i18n.Value("errors.handling.pgx.25006"), pgxCode + 10
	}

	return i18n.Value("errors.handling.pgx.default"), pgxCode
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
