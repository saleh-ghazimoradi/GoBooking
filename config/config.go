package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
	"time"
)

var Appconfig *Env

type Env struct {
	Server     Server
	Postgresql Postgresql
}

type Postgresql struct {
	DbHost       string        `env:"DB_HOST,required"`
	DbPort       string        `env:"DB_PORT,required"`
	DbUser       string        `env:"DB_USER,required"`
	DbPassword   string        `env:"DB_PASSWORD,required"`
	DbName       string        `env:"DB_NAME,required"`
	DbSslMode    string        `env:"DB_SSLMODE,required"`
	MaxOpenConns int           `env:"DB_MAX_OPEN_CONNECTIONS,required"`
	MaxIdleConns int           `env:"DB_MAX_IDLE_CONNECTIONS,required"`
	MaxIdleTime  time.Duration `env:"DB_MAX_IDLE_TIME,required"`
	Timeout      time.Duration `env:"DB_TIMEOUT,required"`
}

type Server struct {
	Port         string        `env:"SERVER_PORT,required"`
	Version      string        `env:"SERVER_VERSION,required"`
	IdleTimeout  time.Duration `env:"SERVER_IDLE_TIMEOUT,required"`
	ReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT,required"`
	WriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT,required"`
}

func EnvConfig() error {
	if err := godotenv.Load("app.env"); err != nil {
		log.Fatalf("Unable to load app.env file: %e", err)
	}

	envConfig := &Env{}

	if err := env.Parse(envConfig); err != nil {
		log.Fatalf("Unable to parse app.env file: %e", err)
	}

	postConfig := &Postgresql{}

	if err := env.Parse(postConfig); err != nil {
		log.Fatalf("Unable to parse app.env file: %e", err)
	}

	envConfig.Postgresql = *postConfig

	serverConfig := &Server{}

	if err := env.Parse(serverConfig); err != nil {
		log.Fatalf("Unable to parse app.env file: %e", err)
	}

	envConfig.Server = *serverConfig

	Appconfig = envConfig

	return nil
}
