package core

import (
	"encoding/base64"

	"errors"

	"github.com/caarlos0/env/v6"
)

type Base64Decoded []byte

func (b *Base64Decoded) UnmarshalText(text []byte) error {
	decoded, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return errors.New("cannot decode base64 string")
	}
	*b = decoded
	return nil
}

type AppConfig struct {
	SecretKey        Base64Decoded `env:"SECRET_KEY,notEmpty"`
	SQLitePath       string        `env:"DB_SQLITE_PATH,notEmpty"`
	RequireLogin     bool          `env:"REQUIRE_LOGIN" envDefault:"false"`
	AllowNewAccounts bool          `env:"ALLOW_NEW_ACCOUNTS" envDefault:"true"`
}

// Load the config from OS
func (appConfig *AppConfig) ParseConfig() error {
	if err := env.Parse(appConfig); err != nil {
		return err
	}
	return nil
}
