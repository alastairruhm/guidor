package api

import (
	"github.com/alastairruhm/guidor/src/bll"
	"github.com/alastairruhm/guidor/src/model"
	"github.com/alastairruhm/guidor/src/schema"
	"github.com/teambition/gear"
)

// Database is API object for teams
//
// @Name Database
// @Description Database API
// @Accepts json
// @Produces json
type Database struct {
	models      *model.All
	databaseBll *bll.Database
}

// Init ...
func (a *Database) Init(blls *bll.All) *Database {
	a.models = blls.Models
	a.databaseBll = blls.Database
	return a
}

type tplDatabaseCreate struct {
	IP        string `json:"ip,omitempty" swaggo:"false,,"`
	Hostname  string `json:"hostname" swaggo:"true"`
	DbType    string `json:"db_type" swaggo:"true"`
	DbVersion string `json:"db_version" swaggo:"true"`
	DbName    string `json:"db_name" swaggo:"false"`
}

func (t *tplDatabaseCreate) Validate() error {
	if len(t.Hostname) == 0 {
		return gear.ErrBadRequest.WithMsg("hostname required")
	}
	return nil
}

// Register ...
//
// @Title Register
// @Summary Register a database instance
// @Description Register a database
// @Param body body tplTeamCreate true "database body"
// @Success 200 schema.DatabaseResult
// @Failure 400 string
// @Failure 401 string
// @Router POST /api/databases
func (a *Database) Register(ctx *gear.Context) error {
	b := new(tplDatabaseCreate)
	if err := ctx.ParseBody(b); err != nil {
		return gear.ErrBadRequest.From(err)
	}

	res, err := a.databaseBll.Register(&schema.Database{
		IP:        b.IP,
		Hostname:  b.Hostname,
		DbVersion: b.DbVersion,
		DbType:    b.DbType,
		DbName:    b.DbName,
	})

	if err != nil {
		return gear.ErrInternalServerError.From(err)
	}
	return ctx.JSON(200, map[string]schema.DatabaseResult{"database": *res})
}
