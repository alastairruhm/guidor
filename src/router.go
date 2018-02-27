package src

import (
	"github.com/alastairruhm/guidor/src/api"
	"github.com/alastairruhm/guidor/src/bll"
	"github.com/teambition/gear"
)

func noOp(ctx *gear.Context) error {
	return gear.ErrNotFound.WithMsg("noOp")
}

func newRouter() (Router *gear.Router) {
	Router = gear.NewRouter()

	blls := new(bll.All).Init()

	databaseAPI := new(api.Database).Init(blls)

	// Register a database instance
	Router.Post("/api/databases/", databaseAPI.Register)
	return
}
