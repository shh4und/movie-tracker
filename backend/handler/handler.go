package handler

import (
	"github.com/shh4und/movie-tracker/config"
)

var (
	logger *config.Logger
	dbpg   *config.Postsql
	apiKEY string
)

func InitHandler() {
	logger = config.GetLogger("handler")
	dbpg = config.GetPSQL()
	apiKEY = config.Envs.ApiKey

}
