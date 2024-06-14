package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/enchant97/url-shorter/components"
	"github.com/enchant97/url-shorter/core"
	"github.com/enchant97/url-shorter/db"
	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5"
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
	SlugType   string `form:"slugType" validate:"required"`
	CustomSlug string `form:"customSlug" validate:"omitempty,alphanum,max=32"`
	TargetUrl  string `form:"targetUrl" validate:"required,http_url,max=8000"`
}

func (h *UiHandler) PostNewShort(c *fuego.ContextWithBody[NewShortForm]) (fuego.Templ, error) {
	b, err := c.Body()
	if err != nil {
		return nil, err
	}
	var slug string
	if b.SlugType == "custom" && b.CustomSlug != "" {
		slug = b.CustomSlug
	} else if b.SlugType == "long" {
		slug = core.GenerateRandomSlug(h.appConfig.UIDLongLength)
	} else {
		slug = core.GenerateRandomSlug(h.appConfig.UIDShortLength)
	}
	if _, err := h.dao.CreateShort(c.Context(), db.CreateShortParams{
		Slug:      slug,
		TargetUrl: b.TargetUrl,
	}); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			c.SetStatus(422)
			return components.FlashBox("shortened name already exists", components.FlashError), nil
		}
		return nil, err
	}
	shortenedLink := fmt.Sprintf("%s/@/%s", h.appConfig.PublicUrl, slug)
	return components.CreateShortForm(&shortenedLink), nil
}

type UpdateShortForm struct {
	ID        int64  `form:"id" validate:"required"`
	TargetUrl string `form:"targetUrl" validate:"required,http_url,max=8000"`
}

func (h *UiHandler) GetUpdateShort(c *fuego.ContextNoBody) (any, error) {
	if id, err := strconv.ParseInt(c.PathParam("id"), 10, 64); err != nil {
		c.SetStatus(404)
		return "404", nil
	} else if short, err := h.dao.GetShortByID(c.Context(), id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.SetStatus(404)
			return "404", nil
		}
		return nil, err
	} else {
		return components.EditShortPage(short), nil
	}
}

func (h *UiHandler) PostUpdateShort(c *fuego.ContextWithBody[UpdateShortForm]) (fuego.Templ, error) {
	b, err := c.Body()
	if err != nil {
		return nil, err
	}
	if short, err := h.dao.UpdateShortByID(c.Context(), db.UpdateShortByIDParams{ID: b.ID, TargetUrl: b.TargetUrl}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.SetStatus(422)
			return components.FlashBox("short does not exist", components.FlashError), nil
		}
		return nil, err
	} else {
		return components.EditShortForm(short), nil
	}
}

func (h *UiHandler) GetShortRedirect(c *fuego.ContextNoBody) (any, error) {
	slug := c.PathParam("slug")
	targetUrl, err := h.dao.GetShortTargetBySlug(c.Context(), slug)
	if err != nil {
		return nil, err
	}
	return c.Redirect(http.StatusTemporaryRedirect, targetUrl)
}
