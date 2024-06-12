package main

import (
	"github.com/enchant97/url-shorter/db/migrations"
	"github.com/enchant97/url-shorter/handlers"
	"github.com/go-fuego/fuego"
)

func main() {
	var appConfig AppConfig
	if err := appConfig.ParseConfig(); err != nil {
		panic(err)
	}

	if err := migrations.MigrateDB(appConfig.DbUri); err != nil {
		panic(err)
	}

	s := fuego.NewServer(fuego.WithPort(8080))

	web := fuego.Group(s, "/")
	web.Hide()

	fuego.Get(web, "/", handlers.GetIndex)

	s.Run()
}
