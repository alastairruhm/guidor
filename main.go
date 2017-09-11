package main

import (
	"flag"
	"os"
	"regexp"

	"github.com/alastairruhm/guidor/src"
	"github.com/teambition/gear/logging"
)

var (
	portReg = regexp.MustCompile(`^\d+$`)
	port    = flag.String("port", "8080", `Server port.`)
)

func main() {
	flag.Parse()
	if portReg.MatchString(*port) {
		*port = ":" + *port
	}
	if *port == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	app := src.New()
	// start app
	logging.Info("IP Service start " + *port)
	app.Listen(*port)
}
