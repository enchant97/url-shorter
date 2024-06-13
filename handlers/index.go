package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/enchant97/url-shorter/components"
	"github.com/enchant97/url-shorter/core"
	"github.com/enchant97/url-shorter/db"
	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5/pgconn"
)

type UiHandler struct {
	appConfig core.AppConfig
	dao       *db.Queries
}

func (h UiHandler) New(appConfig core.AppConfig, dao *db.Queries) UiHandler {
	return UiHandler{
		appConfig: appConfig,
		dao:       dao,
	}
}

func (h *UiHandler) GetIndex(c fuego.ContextNoBody) (fuego.Templ, error) {
	return components.Index(), nil
}

func (h *UiHandler) GetLatestShorts(c fuego.ContextNoBody) (fuego.Templ, error) {
	if shorts, err := h.dao.LatestShorts(c.Context(), 6); err != nil {
		return nil, err
	} else {
		return components.LatestShorts(shorts), nil
	}
}

func (h *UiHandler) GetDashboard(c fuego.ContextNoBody) (fuego.Templ, error) {
	return components.DashboardPage(), nil
}

func (h *UiHandler) GetNewShort(c fuego.ContextNoBody) (fuego.Templ, error) {
	return components.CreateShortPage(), nil
}

type NewShortForm struct {
	Slug      string `form:"slug" validate:"omitempty,alphanum"`
	TargetUrl string `form:"targetUrl" validate:"required,http_url"`
}

func (h *UiHandler) PostNewShort(c *fuego.ContextWithBody[NewShortForm]) (fuego.Templ, error) {
	b := c.MustBody()
	if b.Slug == "" {
		b.Slug = core.GenerateRandomSlug(6)
	}
	if _, err := h.dao.CreateShort(c.Context(), db.CreateShortParams{
		Slug:      b.Slug,
		TargetUrl: b.TargetUrl,
	}); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			c.SetStatus(422)
			return components.FlashBox("shortened name already exists", components.FlashError), nil
		}
		return nil, err
	}
	shortenedLink := fmt.Sprintf("%s/@/%s", h.appConfig.PublicUrl, b.Slug)
	return components.CreateShortForm(&shortenedLink), nil
}

func (h *UiHandler) GetShortRedirect(c *fuego.ContextNoBody) (any, error) {
	slug := c.PathParam("slug")
	targetUrl, err := h.dao.GetShortTargetBySlug(c.Context(), slug)
	if err != nil {
		return nil, err
	}
	return c.Redirect(http.StatusTemporaryRedirect, targetUrl)
}
