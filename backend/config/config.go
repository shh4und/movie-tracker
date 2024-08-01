package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	dbsqlite   *gorm.DB
	logger     *Logger
	pgInstance *Postsql
)

type Postsql struct {
	DB *pgxpool.Pool
}

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

var Envs = GetEnvs()

func Init() error {
	var err error
	dbsqlite, err = InitSQLite()
	if err != nil {
		return fmt.Errorf("Error at initialize sqlite: %v", err)
	}

	pgInstance, err = InitPSQL()
	if err != nil {
		return fmt.Errorf("Error at initialize PostgreSQL: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB { return dbsqlite }

func GetPSQL() *Postsql { return pgInstance }

func GetEnvs() ConfigEnv {
	// get the env file
	godotenv.Load()

	return ConfigEnv{
		PublicHost:           getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                 getEnv("PORT", "8080"),
		DBUser:               getEnv("DB_USER", "root"),
		DBPassword:           getEnv("DB_PASS", "mypassword"),
		DBUrl:                getEnv("DB_URL", "url"),
		DBName:               getEnv("DB_NAME", "ecom"),
		JwtToken:             getEnv("JWT_TK", "not-so-secret-now-is-it?"),
		JwtExpirationSeconds: getEnvAsInt("JWT_EXP", 60),
	}
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

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
