package main

import (
	"github.com/joho/godotenv"
	"github.com/shh4und/movie-tracker/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	router.Init()
}
