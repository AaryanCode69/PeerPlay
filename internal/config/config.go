// Package config manages config
package config

import (
	"os"
)

type Config struct {
	Port string
}

func Load() (*Config, error) {
	cfg := &Config{
		Port: getEnv("Port", "8080"),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
