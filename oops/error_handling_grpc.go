// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package oops

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/utils/grpckit"
)

// HandlingGRPC handle an error by defining HTTP response body and code
func HandlingGRPC(err error) error {
	if err == nil {
		return nil
	}

	var e *Error
	if !errors.As(err, &e) {
		return HandlingGRPC(Err(err))
	}

	if e.IsHandled() {
		return buildGRPCStatus(e)
	}

	e = handling(e.Err).(*Error)

	return buildGRPCStatus(e)
}

// buildGRPCStatus replicates the gRPC status construct
func buildGRPCStatus(e *Error) error {
	msg := e.Message

	if config.Get().Server.IsModeDevelopment() {
		msg += " @ "
		for ii := range e.Trace {
			msg += e.Trace[ii] + " || "
		}
		msg += e.Error()
	}

	rawError := e.Error()
	st, _ := status.New(codes.Aborted, msg).WithDetails(&grpckit.ErrorGRPC{
		Error:    &e.Message,
		Location: &msg,
		RawError: &rawError,
	})

	return st.Err()
}
