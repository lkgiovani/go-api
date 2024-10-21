package config

import "go-api/pkg/env"

type Config struct {
	Mysql struct {
		Url string
	}

	Server struct {
		Port string
		Ip   string
	}
}

func NewConfig() *Config {
	return &Config{
		Mysql: struct{ Url string }{Url: env.GetEnvOrDie("URL_MYSQL")},

		Server: struct {
			Port string
			Ip   string
		}{
			Port: env.GetEnvOrDie("SERVER_PORT"),
			Ip:   env.GetEnvOrDie("IP_HTTP"),
		},
	}
}
