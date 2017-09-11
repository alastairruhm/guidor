package mongodb

import (
	"github.com/alastairruhm/guidor/src/logger"
)

// MaxPool max pool size
var MaxPool int

func init() {
	MaxPool = 300

	// init method to start db
	logger.Info("check mongodb connection")
	checkAndInitServiceConnection()
}

func checkAndInitServiceConnection() {
	if service.baseSession == nil {
		service.URL = "127.0.0.1:27017"
		err := service.New()
		if err != nil {
			// logger.Err(err)
			logger.Fatal(err)
		}
	}
}
