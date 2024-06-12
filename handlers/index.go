package handlers

import (
	"github.com/enchant97/url-shorter/components"
	"github.com/go-fuego/fuego"
)

func GetIndex(c fuego.ContextNoBody) (fuego.Templ, error) {
	return components.Index(), nil
}

func GetDashboard(c fuego.ContextNoBody) (fuego.Templ, error) {
	return components.DashboardPage(), nil
}

func GetNewShort(c fuego.ContextNoBody) (fuego.Templ, error) {
	return components.CreateShortPage(""), nil
}

type NewShortForm struct {
	Slug      string `form:"slug" validate:"required"`
	TargetUrl string `form:"targetUrl" validate:"required"`
}

func PostNewShort(c *fuego.ContextWithBody[NewShortForm]) (fuego.Templ, error) {
	b := c.MustBody()
	return components.CreateShortForm("", "https://example.com/"+b.Slug), nil
}
