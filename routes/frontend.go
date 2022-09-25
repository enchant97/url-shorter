package routes

import (
	"fmt"
	"net/http"

	"github.com/enchant97/go-gincookieauth"
	"github.com/enchant97/url-shorter/core"
	"github.com/enchant97/url-shorter/core/db"
	"github.com/enchant97/url-shorter/core/flash"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	core.HTMLTemplate(c, http.StatusOK, "index.html", gin.H{
		"pageTitle": "Home",
		"AppConfig": c.MustGet("AppConfig"),
	})
}

func GetChecker(c *gin.Context) {
	shortID := c.Query("short-id")
	// there is a short id given
	if shortID != "" {
		if decodedID, err := core.DecodeIDPadded(shortID); err == nil {
			shortRow := db.GetShortByID(uint(decodedID))
			core.HTMLTemplate(c, http.StatusOK, "checker.html", gin.H{
				"pageTitle": "Checker",
				"short":     shortRow,
				"shortID":   shortID,
			})
			return
		}
		flash.FlashWarning(c, fmt.Sprintf("Provided short id '%s' was not found.", shortID))
	}
	// no short id or invalid
	core.HTMLTemplate(c, http.StatusOK, "checker.html", gin.H{
		"pageTitle": "Checker",
	})
}

func GetNew(c *gin.Context) {
	core.HTMLTemplate(c, http.StatusOK, "new.html", gin.H{
		"pageTitle": "New",
	})
}

func GetNewAuthRequired(c *gin.Context) {
	core.HTMLTemplate(c, http.StatusUnauthorized, "error.html", gin.H{
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
	// is valid short id?
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	shortRow := db.GetShortByID(uint(decodedID))
	// does short id exist?
	if shortRow == nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"pageTitle":    "404",
			"errorTitle":   "Page Could Not Be Found",
			"errorMessage": fmt.Sprintf("The given short url with id '%s' could not be found, is your link valid?", shortID),
		})
		return
	}
	// is short id expired?
	if shortRow.IsExpired() || !shortRow.IsUsable() {
		go func() { db.DB.Delete(&shortRow) }()
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"pageTitle":    "404",
			"errorTitle":   "Page Could Not Be Found",
			"errorMessage": fmt.Sprintf("The given short url with id '%s' has expired and cannot be used.", shortID),
		})
		return
	}
	// short id OK, redirect to target URL
	go func() { shortRow.IncrVisitCount() }()
	c.Redirect(http.StatusTemporaryRedirect, shortRow.TargetURL)
}
