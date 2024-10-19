package config

import "go-api/pkg/env"

type Config struct {
	mysql struct {
		url string
	}

	Server struct {
		Port string
		Ip   string
	}
}

func NewConfig() *Config {
	return &Config{
		mysql: struct{ url string }{url: env.GetEnvOrDie("URL_MYSQL")},

		Server: struct {
			Port string
			Ip   string
		}{
			Port: env.GetEnvOrDie("SERVER_PORT"),
			Ip:   env.GetEnvOrDie("IP_HTTP"),
		},
	}
}
