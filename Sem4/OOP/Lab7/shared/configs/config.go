package configs

import (
	"github.com/joho/godotenv"
	"time"
)

type Config struct {
	TokenKey      string        `mapstructure:"TOKEN_KEY"`
	TokenDuration time.Duration `mapstructure:"TOKEN_DURATION"`
	Address       string        `mapstructure:"ADDRESS"`
}

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	var config Config

	return config, nil
}
