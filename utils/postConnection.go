package utils

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/saleh-ghazimoradi/GoBooking/config"
	"github.com/saleh-ghazimoradi/GoBooking/logger"
)

func PostURI(cfg *config.Env) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Postgresql.DbHost, cfg.Postgresql.DbPort, cfg.Postgresql.DbUser, cfg.Postgresql.DbPassword, cfg.Postgresql.DbName, cfg.Postgresql.DbSslMode)
}

func PostConnection(cfg *config.Env) (*sql.DB, error) {
	postURI := PostURI(cfg)

	logger.Logger.Info("Connecting to Postgres with options: " + postURI)

	db, err := sql.Open("postgres", postURI)
	if err != nil {
		return nil, fmt.Errorf("error opening Postgres connection: %w", err)
	}

	db.SetMaxOpenConns(cfg.Postgresql.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Postgresql.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.Postgresql.MaxIdleTime)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Postgresql.Timeout)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error pinging Postgres database: %w", err)
	}

	return db, nil
}
