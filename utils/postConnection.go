package utils

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/saleh-ghazimoradi/GoBooking/config"
	"github.com/saleh-ghazimoradi/GoBooking/logger"
)

func PostURI() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Appconfig.Postgresql.DbHost, config.Appconfig.Postgresql.DbPort, config.Appconfig.Postgresql.DbUser, config.Appconfig.Postgresql.DbPassword, config.Appconfig.Postgresql.DbName, config.Appconfig.Postgresql.DbSslMode)
}

func PostConnection() (*sql.DB, error) {
	postURI := PostURI()

	logger.Logger.Info("Connecting to Postgres with options: " + postURI)

	db, err := sql.Open("postgres", postURI)
	if err != nil {
		return nil, fmt.Errorf("error opening Postgres connection: %w", err)
	}

	db.SetMaxOpenConns(config.Appconfig.Postgresql.MaxOpenConns)
	db.SetMaxIdleConns(config.Appconfig.Postgresql.MaxIdleConns)
	db.SetConnMaxLifetime(config.Appconfig.Postgresql.MaxIdleTime)
	ctx, cancel := context.WithTimeout(context.Background(), config.Appconfig.Postgresql.Timeout)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error pinging Postgres database: %w", err)
	}

	return db, nil
}
