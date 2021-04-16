package apiserver

import "github.com/bulbetski/learn_http-rest-api/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	Loglevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Loglevel: "debug",
		Store:    store.NewConfig(),
	}
}
