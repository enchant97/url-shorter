package main

import (
	"github.com/enchant97/url-shorter/handlers"
	"github.com/go-fuego/fuego"
)

func main() {
	s := fuego.NewServer(fuego.WithPort(8080))

	web := fuego.Group(s, "/")
	web.Hide()

	fuego.Get(web, "/", handlers.GetIndex)

	s.Run()
}
