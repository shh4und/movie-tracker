package config

import (
	"fmt"
	"os"

	"github.com/shh4und/movie-tracker/utils"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

type ENVvars struct {
	ApiKey               string
	JwtToken             string
	JwtExpirationSeconds int64
}

var Envs = GetEnvs()

func Init() error {
	var err error

	db, err = InitSQLite()
	if err != nil {
		return fmt.Errorf("Error at initialize sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB { return db }

func GetEnvs() ENVvars {
	// get the env file
	godotenv.Load("../.env")

	return ENVvars{
		ApiKey:               os.Getenv("API_KEY"),
		JwtToken:             os.Getenv("JWT_TK"),
		JwtExpirationSeconds: utils.ParseInt(os.Getenv("JWT_EXP")),
	}
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
