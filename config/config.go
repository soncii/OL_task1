package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port     string `env:"PORT" envDefault:":3000"`
	DB       string `env:"DB" envDefault:"in-memory"`
	Capacity int    `env:"CAPACITY" envDefault:"100"`
}

func NewConfig() (config *Config, err error) {
	cfg := Config{}
	error := env.Parse(&cfg)
	if error != nil {
		return nil, nil
	}
	return &cfg, error
}
