package services

import "github.com/shh4und/movie-tracker/config"

var (
	logger *config.Logger
	apiKEY string
)

func InitServices() {
	logger = config.GetLogger("services")
	apiKEY = config.Envs.ApiKey

}
