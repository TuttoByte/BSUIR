package configs

import (
	"os"
	"time"
)

type Config struct {
	TokenKey      string        `mapstructure:"TOKEN_KEY"`
	TokenDuration time.Duration `mapstructure:"TOKEN_DURATION"`
	Address       string        `mapstructure:"ADDRESS"`
}

func Load() (Config, error) {
	var config Config

	config.TokenKey = os.Getenv("TOKEN_KEY")
	config.TokenDuration, _ = time.ParseDuration(os.Getenv("TOKEN_DURATION"))
	config.Address = os.Getenv("ADDRESS")

	return config, nil
}
