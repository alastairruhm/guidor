package bll

import (
	"github.com/alastairruhm/guidor/src/model"
)

// Bll is Business Logic Layer with all models
type Bll struct {
	Models *model.All
}

// All ...
type All struct {
	Models   *model.All
	Database *Database
}

// Init ...
func (a *All) Init() *All {
	a.Models = new(model.All).Init()
	b := &Bll{a.Models}
	a.Database = &Database{b}
	return a
}
