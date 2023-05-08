package errors

import (
	"errors"
)

var (
	ErrInvalidRequest          = errors.New("invalid_request")
	ErrUnsupportedResponseType = errors.New("unsupported_response_type")
	ErrInvalidScope            = errors.New("invalid_scope")
	ErrUnauthorizedClient      = errors.New("unauthorized_client")
	ErrServerError             = errors.New("server_error")
)

var descriptions = map[error]string{
	ErrInvalidRequest:          "request missing one or more required parameters",
	ErrUnsupportedResponseType: "passed response type is unsupported",
	ErrInvalidScope:            "one or more passed scopes are unsupported",
	ErrUnauthorizedClient:      "client unauthorized to perform action",
	ErrServerError:             "unexpected condition encountered",
}

var errorsMap = map[error]struct{}{
	ErrInvalidScope:            {},
	ErrUnsupportedResponseType: {},
	ErrInvalidScope:            {},
	ErrUnauthorizedClient:      {},
	ErrServerError:             {},
}

func GetDescription(err error) string {
	v, _ := descriptions[err]
	return v
}

func IsHandledError(err error) bool {
	_, ok := errorsMap[err]
	return ok
}
