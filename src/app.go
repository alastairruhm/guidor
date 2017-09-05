package src

import (
	"github.com/alastairruhm/guidor/src/logger"
	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware/cors"
)

// Version is app version
const Version = "v1.0.0"

// New returns a app instance
func New(bindHost string) *gear.App {
	app := gear.New()
	logger.Init()

	app.Use(cors.New())
	// app.Use(secure.Default)
	app.UseHandler(logger.Default())

	return app
}
