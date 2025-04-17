package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	LogLevel   string `env:"LOG_LEVEL" envDefault:"INFO"`
	Token      string `env:"TOKEN" envDefault:"7731167021:AAHb5ofDkyroDH0LUYznmRVeIycJCYe0c2M"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"123"`
	DBHost     string `env:"DB_HOST" envDefault:"db"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBName     string `env:"DB_NAME" envDefault:"postgres"`
}

func ReadConfig() (*Config, error) {
	config := Config{}

	err := env.Parse(&config)
	if err != nil {
		return nil, fmt.Errorf("read config error: %w", err)
	}

	return &config, err
}
