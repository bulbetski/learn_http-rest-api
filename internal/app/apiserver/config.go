package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	Loglevel    string `toml:"log_level"`
	DatabaseUrl string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		Loglevel: "debug",
	}
}
