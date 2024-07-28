package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPQSQL() (*pgxpool.Pool, error) {
	logger := GetLogger("postgresql")

	dbpool, err := pgxpool.New(context.Background(), Envs.DBUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	logger.Info("postgresql connection established")
	return dbpool, nil
}
