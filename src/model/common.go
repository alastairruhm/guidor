package model

import (
	"github.com/teambition/gear"
	"gopkg.in/mgo.v2"
)

func dbError(err error) error {
	if err == nil {
		return nil
	}
	if err == mgo.ErrNotFound {
		return gear.ErrNotFound.WithMsg(err.Error())
	}
	return gear.ErrInternalServerError.From(err)
}

// All ....
type All struct {
	Database *Database
}

// Init ...
func (a *All) Init() *All {
	a.Database = new(Database)
	return a
}
