package config

import (
	"time"

	"github.com/Geekzix/calorie-snap-commons/utils"
)

type Config struct {
	HttpAddr        string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
	DatabaseURL     string
}

func Load() (Config, error) {
	cfg := Config{
		HttpAddr:        utils.GetEnvString("HTTP_ADDR", ":8080"),
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    10 * time.Second,
		IdleTimeout:     60 * time.Second,
		ShutdownTimeout: 10 * time.Second,
		DatabaseURL:     utils.GetEnvString("DATABASE_URL", "postgres://auth_user:password@localhost:5432/auth_db?sslmode=disable"),
	}

	var err error

	if cfg.ReadTimeout, err = utils.GetDuration("HTTP_READ_TIMEOUT", cfg.ReadTimeout); err != nil {
		return Config{}, err
	}

	if cfg.WriteTimeout, err = utils.GetDuration("HTTP_WRITE_TIMEOUT", cfg.WriteTimeout); err != nil {
		return Config{}, err
	}

	if cfg.IdleTimeout, err = utils.GetDuration("HTTP_IDLE_TIMEOUT", cfg.IdleTimeout); err != nil {
		return Config{}, err
	}

	if cfg.ShutdownTimeout, err = utils.GetDuration("HTTP_SHUTDOWN_TIMEOUT", cfg.ShutdownTimeout); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
