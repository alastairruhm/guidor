package src

import (
	"github.com/teambition/gear"
)

// Version is app version
const Version = "v1.0.0"

// New returns a app instance
func New() *gear.App {
	app := gear.New()
	// logger.Init()

	// app.Use(cors.New())

	// app.UseHandler(logger.Default())

	// app.UseHandler(newRouter())
	return app
}
