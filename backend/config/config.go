package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// ConfigEnv holds the configuration values for the application.
type ConfigEnv struct {
	PublicHost           string
	Port                 string
	DBUser               string
	DBPassword           string
	DBUrl                string
	DBName               string
	ApiKey               string
	JwtToken             string
	JwtExpirationSeconds int64
}

// Envs is a global variable that holds the loaded environment configuration.
var Envs = GetEnvs()

// GetEnvs loads environment variables and returns a ConfigEnv struct.
func GetEnvs() ConfigEnv {
	godotenv.Load()

	return ConfigEnv{
		PublicHost:           getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                 getEnv("PORT", "8080"),
		DBUser:               getEnv("DB_USER", "root"),
		DBPassword:           getEnv("DB_PASS", "mypassword"),
		DBUrl:                getEnv("DB_URL", "url"),
		DBName:               getEnv("DB_NAME", "ecom"),
		JwtToken:             getEnv("JWT_TK", "not-so-secret-now-is-it?"),
		JwtExpirationSeconds: getEnvAsInt("JWT_EXP", 3600),
		ApiKey:               getEnv("API_KEY", ""),
	}
}

// getEnv retrieves the value of the environment variable named by the key.
// It returns the value, or the fallback if the variable is not present.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// getEnvAsInt retrieves the value of the environment variable named by the key
// and converts it to an int64. It returns the value, or the fallback if the
// variable is not present or cannot be converted.
func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
