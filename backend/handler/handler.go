package handler

import (
	"github.com/shh4und/movie-tracker/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
	apiKEY string
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
	apiKEY = config.Envs.ApiKey

}
