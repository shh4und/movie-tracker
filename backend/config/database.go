package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// pgInstance is a global variable that holds the PostgreSQL connection pool instance.
// singleton instance
var pgInstance *Postsql

// logger is a global logger instance for database operations.
var logger = NewLogger("postgresql")

// Postsql is a struct that holds the PostgreSQL connection pool.
type Postsql struct {
	DB *pgxpool.Pool
}

// InitPSQL initializes the PostgreSQL connection pool and returns a Postsql instance.
func InitPSQL() (*Postsql, error) {
	db, err := pgxpool.New(context.Background(), Envs.DBUrl)
	if err != nil {
		logger.Errorf("unable to create connection pool: %v", err)
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	pgInstance = &Postsql{db}
	logger.Info("postgresql connection established")
	return pgInstance, nil
}

// Ping checks the connection to the PostgreSQL database.
func (pg *Postsql) Ping(ctx context.Context) error {
	return pg.DB.Ping(ctx)
}

// Close closes the PostgreSQL connection pool.
func (pg *Postsql) Close() {
	pg.DB.Close()
	logger.Info("postgresql connection closed")
}

// Init initializes the PostgreSQL connection pool and assigns it to the global pgInstance variable.
func Init() error {
	var err error

	pgInstance, err = InitPSQL()
	if err != nil {
		return fmt.Errorf("Error at initialize PostgreSQL: %v", err)
	}

	return nil
}

// GetPSQL returns the global PostgreSQL connection pool instance.
func GetPSQL() *Postsql { return pgInstance }
