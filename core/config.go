package core

import "github.com/caarlos0/env/v6"

type AppConfig struct {
	DbUri          string `env:"DB_URI,notEmpty"`
	PublicUrl      string `env:"PUBLIC_URL,notEmpty"`
	UIDShortLength uint   `env:"UID_SHORT_LENGTH" envDefault:"6"`
	UIDLongLength  uint   `env:"UID_LONG_LENGTH" envDefault:"128"`
}

func (appConfig *AppConfig) ParseConfig() error {
	if err := env.Parse(appConfig); err != nil {
		return err
	}
	return nil
}
