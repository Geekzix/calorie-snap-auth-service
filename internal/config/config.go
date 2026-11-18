package config

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Config struct {
	HttpAddr        string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

func Load() (Config, error) {
	cfg := Config{
		HttpAddr:        getEnv("HTTP_ADDR", ":8080"),
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    10 * time.Second,
		IdleTimeout:     60 * time.Second,
		ShutdownTimeout: 10 * time.Second,
	}

	var err error

	if cfg.ReadTimeout, err = getDuration("HTTP_READ_TIMEOUT", cfg.ReadTimeout); err != nil {
		return Config{}, err
	}

	if cfg.WriteTimeout, err = getDuration("HTTP_WRITE_TIMEOUT", cfg.WriteTimeout); err != nil {
		return Config{}, err
	}

	if cfg.IdleTimeout, err = getDuration("HTTP_IDLE_TIMEOUT", cfg.IdleTimeout); err != nil {
		return Config{}, err
	}

	if cfg.ShutdownTimeout, err = getDuration("HTTP_SHUTDOWN_TIMEOUT", cfg.ShutdownTimeout); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func getDuration(key string, fallback time.Duration) (time.Duration, error) {
	value := os.Getenv(key)

	if value == "" {
		return fallback, nil
	}

	duration, err := time.ParseDuration(value)

	if err != nil {
		log.Fatalf("Invalid %s: %v\n", key, err)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("%s must be positive", key)
	}

	return duration, nil
}
