package migrations

import (
	"embed"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed *.sql
var migrationsFS embed.FS

func MigrateDB(connectionURI string) error {
	d, err := iofs.New(migrationsFS, ".")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, connectionURI)
	if err != nil {
		return err
	}
	return m.Up()
}
