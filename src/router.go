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

	// Router.Get("/api/debug", func(ctx *gear.Context) error {
	// 	user, err := user.Current()
	// 	if err != nil {
	// 		panic(fmt.Errorf("Unable to determine user's home directory: %s", err))
	// 	}
	// 	pick := func(x, y interface{}) interface{} {
	// 		return x
	// 	}

	// 	return ctx.JSON(200, map[string]interface{}{
	// 		"Uid":        user.Uid,
	// 		"Gid":        user.Gid,
	// 		"UserName":   user.Username,
	// 		"Name":       user.Name,
	// 		"HomeDir":    user.HomeDir,
	// 		"Env":        os.Environ(),
	// 		"Executable": pick(os.Executable()),
	// 		"Hostname":   pick(os.Hostname()),
	// 		"WD":         pick(os.Getwd()),
	// 	})
	// })
	return
}
