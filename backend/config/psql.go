package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPSQL() (*Postsql, error) {
	logger := GetLogger("postgresql")
	db, err := pgxpool.New(context.Background(), Envs.DBUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	pgInstance = &Postsql{db}
	logger.Info("postgresql connection established")
	return pgInstance, nil
}

func (pg *Postsql) Ping(ctx context.Context) error {
	return pg.DB.Ping(ctx)
}

func (pg *Postsql) Close() {
	pg.DB.Close()
}
