package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		App   App
		Mongo Mongo
	}
	// App -.
	App struct {
		BotToken string `env:"BOT_TOKEN"`
	}
	Mongo struct {
		ConnectionString string `env:"MONGO_CONNECTION"`
	}
)

var config *Config

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
