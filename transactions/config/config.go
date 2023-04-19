package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port       string `env:"PORT" envDefault:":3000"`
	DB         string `env:"DB" envDefault:"pg"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
	DbPort     string `env:"DB_PORT"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbSSL      string `env:"DB_SSL"`
}

func NewConfig() (config *Config, err error) {
	//err = godotenv.Load()
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	cfg := Config{}
	err1 := env.Parse(&cfg)
	if err1 != nil {
		fmt.Println(err1)
		return nil, err1
	}
	fmt.Printf("%v", cfg)
	return &cfg, nil
}
