package routes

import (
	"fmt"
	"net/http"

	"github.com/enchant97/go-gincookieauth"
	"github.com/enchant97/go-gincookieauth/extras"
	"github.com/enchant97/url-shorter/core"
	"github.com/enchant97/url-shorter/core/db"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	extras.TemplateWithAuth(c, http.StatusOK, "index.html", gin.H{
		"pageTitle": "Home",
		"AppConfig": c.MustGet("AppConfig"),
	})
}

func GetChecker(c *gin.Context) {
	shortID := c.Query("short-id")
	if shortID != "" {
		if decodedID, err := core.DecodeIDPadded(shortID); err == nil {
			shortRow := db.GetShortByID(uint(decodedID))
			extras.TemplateWithAuth(c, http.StatusOK, "checker.html", gin.H{
				"pageTitle": "Checker",
				"short":     shortRow,
				"shortID":   shortID,
			})
			return
		}
	}
	extras.TemplateWithAuth(c, http.StatusOK, "checker.html", gin.H{
		"pageTitle": "Checker",
	})
}

func GetNew(c *gin.Context) {
	extras.TemplateWithAuth(c, http.StatusOK, "new.html", gin.H{
		"pageTitle": "New",
	})
}

func GetNewAuthRequired(c *gin.Context) {
	extras.TemplateWithAuth(c, http.StatusUnauthorized, "error.html", gin.H{
		"pageTitle":    "New",
		"errorTitle":   "Login Required",
		"errorMessage": "This page is restricted to logged in users only.",
	})
}

func PostNew(c *gin.Context) {
	var formValues core.CreateShort
	if err := c.ShouldBind(&formValues); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	short := formValues.GenerateShort()
	if userID := gincookieauth.GetUserID(c); userID != nil {
		userID := (*userID).(uint)
		short.OwnerID = &userID
	}
	if err := short.Create(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/checker?short-id="+core.EncodeIDPadded(short.ID))
}

func GetRedirect(c *gin.Context) {
	shortID := c.Param("shortID")
	decodedID, err := core.DecodePossibleShortID(shortID)
	if err != nil {
		// not valid short id
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	shortRow := db.GetShortByID(uint(decodedID))
	if shortRow == nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"pageTitle":    "404",
			"errorTitle":   "Page Could Not Be Found",
			"errorMessage": fmt.Sprintf("The given short url with id '%s' could not be found, is your link valid?", shortID),
		})
	} else {
		if shortRow.IsExpired() || !shortRow.IsUsable() {
			defer func() { go func() { db.DB.Delete(&shortRow) }() }()
			c.HTML(http.StatusNotFound, "error.html", gin.H{
				"pageTitle":    "404",
				"errorTitle":   "Page Could Not Be Found",
				"errorMessage": fmt.Sprintf("The given short url with id '%s' has expired and cannot be used.", shortID),
			})
		} else {
			// execute update in goroutine, as client doesn't need to wait for this
			defer func() { go func() { shortRow.IncrVisitCount() }() }()
			c.Redirect(http.StatusTemporaryRedirect, shortRow.TargetURL)
		}
	}
}
