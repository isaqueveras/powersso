package oops

import "net/http"

// Error defines the properties for a basic error response
type Error struct {
	Code       int      `json:"code"`
	Message    string   `json:"message"`
	Trace      []string `json:"-"`
	Err        error    `json:"-"`
	StatusCode int      `json:"-"`
}

// NewApiError creates and returns new normalized `Error` instance.
func NewApiError(code int, message string, statusCode int) *Error {
	return &Error{
		Code:       code,
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

// NotFoundError creates and returns 404 `Error`.
func NotFoundError(message string) *Error {
	if message == "" {
		message = "The requested resource wasn't found."
	}

	return NewApiError(1, message, http.StatusNotFound)
}
