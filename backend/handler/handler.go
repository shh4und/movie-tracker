package handler

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shh4und/movie-tracker/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
	dbpq   *pgxpool.Pool
	apiKEY string
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
	dbpq = config.GetPQSQL()
	apiKEY = config.Envs.ApiKey

}
