package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ApiKey    string `env:"API_KEY"`
	ApiSecret string `env:"API_SECRET"`
	BotToken  string `env:"BOT_TOKEN"`
	ChannelID string `env:"CHANNEL_ID"`
}

func ReadConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
