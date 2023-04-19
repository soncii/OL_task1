package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port            string `env:"PORT" envDefault:":3000"`
	TransactionsURL string `env:"TR_URL" envDefault:""`
	DB              string `env:"DB" envDefault:"pg"`
	DbHost          string `env:"DB_HOST"`
	DbName          string `env:"DB_NAME"`
	DbPort          string `env:"DB_PORT"`
	DbUser          string `env:"DB_USER"`
	DbPassword      string `env:"DB_PASSWORD"`
	DbSSL           string `env:"DB_SSL"`
	HashCost        int    `env:"HASH_COST"`
	JWTSecret       string `env:"JWT_SECRET"`
}

func NewConfig() (*Config, error) {
	//err := godotenv.Load()
	//if err != nil {
	//	return nil, err
	//}
	cfg := Config{}
	err1 := env.Parse(&cfg)
	if err1 != nil {
		return nil, err1
	}
	return &cfg, nil
}
