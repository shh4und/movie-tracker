package handler

import (
	"github.com/shh4und/movie-tracker/config"
)

var (
	logger *config.Logger
	//db     *gorm.DB
	dbpg   *config.Postsql
	apiKEY string
)

func InitHandler() {
	logger = config.GetLogger("handler")
	//db = config.GetSQLite()
	dbpg = config.GetPSQL()
	apiKEY = config.Envs.ApiKey

}
