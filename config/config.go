package config

import (
	"errors"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	// get the env file
	err := godotenv.Load(".env")
	if err != nil {
		return errors.New("Error loading .env file")
	}
	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
