package handler

import (
	"github.com/shh4und/movie-tracker/config"
)

var (
	logger *config.Logger
	dbpg   *config.Postsql
)

func InitHandler() {
	logger = config.NewLogger("handler")
	dbpg = config.GetPSQL()
}
