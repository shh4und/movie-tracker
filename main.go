package main

import (
	"github.com/shh4und/movie-tracker/config"
	"github.com/shh4und/movie-tracker/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	// initialize router
	router.Init()

	//1:44:38
}
