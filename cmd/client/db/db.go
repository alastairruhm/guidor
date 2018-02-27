package db

import (
	"fmt"

	"github.com/alastairruhm/guidor/cmd/client/er"
	"github.com/teambition/gear/logging"
)

// Base ...
type Base interface {
	// Version show version of database
	Version() (string, error)
	// Type show type of database
	Type() string
	// Auth make login attemps
	Auth() error
	// Do backup action
	Backup() error
	// Close
	Close() error
}

// Config ...
type Config struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
	Type     string
}

// New - initialize database
func New(c Config) (Base, error) {
	var ctx Base
	var e error
	switch c.Type {
	case "mysql":
		ctx, e = NewMySQL(c)
		if e != nil {
			return nil, e
		}
	default:
		logging.Err(fmt.Errorf("databases config `type: %s`, but is not implement", c.Type))
		return nil, er.ErrNotSupported
	}
	return ctx, nil
}
