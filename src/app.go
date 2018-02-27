package src

import (
	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware/cors"
)

// Version is app version
const Version = "v1.0.0"

// New returns a app instance
func New() *gear.App {
	app := gear.New()
	app.Use(cors.New())
	app.UseHandler(newRouter())
	return app
}
