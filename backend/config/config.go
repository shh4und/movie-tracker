package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
	apiKey string
)

func Init() error {
	var err error

	// get the env file
	err = godotenv.Load("../.env")
	if err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}

	apiKey = os.Getenv("API_KEY")

	db, err = InitSQLite()
	if err != nil {
		return fmt.Errorf("Error at initialize sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB { return db }

func GetApiKEY() string { return apiKey }

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
