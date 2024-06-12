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

	ui := fuego.Group(s, "/")
	ui.Hide()
	fuego.Get(ui, "/", handlers.GetIndex)
	fuego.Get(ui, "/ui/", handlers.GetDashboard)
	fuego.Get(ui, "/ui/new", handlers.GetNewShort)
	fuego.Post(ui, "/ui/_post_new_short", handlers.PostNewShort)

	s.Run()
}
