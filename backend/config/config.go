package config

import (
	"fmt"
	"os"
	"strconv"

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

	// get the env file
	err = godotenv.Load("../.env")
	if err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}

	db, err = InitSQLite()
	if err != nil {
		return fmt.Errorf("Error at initialize sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB { return db }

func GetEnvs() ENVvars {
	exp, err := strconv.ParseInt(os.Getenv("JWT_EXP"), 10, 64)
	if err != nil {
		return ENVvars{}
	}
	return ENVvars{
		ApiKey:               os.Getenv("API_KEY"),
		JwtToken:             os.Getenv("JWT_TK"),
		JwtExpirationSeconds: exp,
	}
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
