package main

import (
	"github.com/joho/godotenv"
	"github.com/shh4und/movie-tracker/config"
	"github.com/shh4und/movie-tracker/router"
)

func main() {
	// get the env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// initialize configs
	err = config.Init()
	if err != nil {
		panic("Error loading configs")
	}

	// initialize router
	router.Init()

	//1:44:38
}
