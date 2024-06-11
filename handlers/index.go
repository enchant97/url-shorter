package handlers

import (
	"github.com/enchant97/url-shorter/components"
	"github.com/go-fuego/fuego"
)

func GetIndex(c fuego.ContextNoBody) (fuego.Templ, error) {
	return components.Index(), nil
}
