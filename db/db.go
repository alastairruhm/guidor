package db

// Base ...
type Base interface {
	// Version show version of database
	Version() (string, error)
	// Type show type of database
	Type() string
	// Auth make login attemps
	Auth() error
}

// Auth ...
func Auth(ctx Base) error {
	return ctx.Auth()
}
