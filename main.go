package main

import (
	"context"

	"github.com/enchant97/url-shorter/core"
	"github.com/enchant97/url-shorter/db"
	"github.com/enchant97/url-shorter/db/migrations"
	"github.com/enchant97/url-shorter/handlers"
	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5"
)

func main() {
	var appConfig core.AppConfig
	if err := appConfig.ParseConfig(); err != nil {
		panic(err)
	}

	if err := migrations.MigrateDB(appConfig.DbUri); err != nil {
		panic(err)
	}
	ctx := context.Background()
	dbConn, err := pgx.Connect(ctx, appConfig.DbUri)
	if err != nil {
		panic(err)
	}
	dao := db.New(dbConn)

	s := fuego.NewServer(fuego.WithPort(8080))

	uiHandler := handlers.UiHandler{}.New(appConfig, dao)

	ui := fuego.Group(s, "/")
	ui.Hide()
	fuego.Get(ui, "/@/{slug}", uiHandler.GetShortRedirect)
	fuego.Get(ui, "/", uiHandler.GetIndex)
	fuego.Get(ui, "/ui/", uiHandler.GetDashboard)
	fuego.Get(ui, "/ui/new", uiHandler.GetNewShort)
	fuego.Post(ui, "/ui/_post_new_short", uiHandler.PostNewShort)

	s.Run()
}
