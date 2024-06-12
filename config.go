package main

import "github.com/caarlos0/env/v6"

type AppConfig struct {
	DbUri string `env:"DB_URI,notEmpty"`
}

func (appConfig *AppConfig) ParseConfig() error {
	if err := env.Parse(appConfig); err != nil {
		return err
	}
	return nil
}
