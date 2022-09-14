package routes

import (
	"net/http"

	"github.com/enchant97/url-shorter/core"
	"github.com/enchant97/url-shorter/core/db"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"pageTitle": "Home",
	})
}

func GetChecker(c *gin.Context) {
	shortID := c.Query("short-id")
	if shortID != "" {
		shortRow := db.GetShortByShortID(shortID)
		c.HTML(http.StatusOK, "checker.html", gin.H{
			"pageTitle": "Checker",
			"short":     shortRow,
		})
		return
	}
	c.HTML(http.StatusOK, "checker.html", gin.H{
		"pageTitle": "Checker",
	})
}

func GetNew(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{
		"pageTitle": "New",
	})
}

func PostNew(c *gin.Context) {
	var formValues core.CreateShort
	if err := c.ShouldBind(&formValues); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	shortID := core.MakeShortID()
	db.CreateNewShort(shortID, formValues.TargetURL)
	c.Redirect(http.StatusSeeOther, "/checker?short-id="+shortID+"")
}

func GetRedirect(c *gin.Context) {
	shortID := c.Param("shortID")
	shortRow := db.GetShortByShortID(shortID)
	if shortRow == (db.Short{}) {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		db.ShortVisitCountIncr(shortID)
		c.Redirect(http.StatusTemporaryRedirect, shortRow.TargetURL)
	}
}
