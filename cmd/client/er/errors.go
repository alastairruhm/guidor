package er

import "errors"

var (
	ErrNotSupported     = errors.New("database not supported")
	ErrDBAuthentication = errors.New("database authentication failed")
	// ErrBadRequest = errors.New("foo: bad request")
)
