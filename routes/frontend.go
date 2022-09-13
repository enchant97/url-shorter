package routes

import (
	"net/http"

	"github.com/enchant97/url-shorter/core"
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
		targetURL := core.FakeShortsDB[shortID]
		short := createdShort{
			ShortID:   shortID,
			TargetURL: targetURL,
		}
		c.HTML(http.StatusOK, "checker.html", gin.H{
			"pageTitle": "Checker",
			"short":     short,
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
	var formValues newShortFormValues
	if err := c.ShouldBind(&formValues); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	shortID := core.MakeShortID()
	core.FakeShortsDB[shortID] = formValues.TargetURL
	c.Redirect(http.StatusSeeOther, "/"+shortID+"/info")
}

func GetRedirect(c *gin.Context) {
	shortID := c.Param("shortID")
	targetURL := core.FakeShortsDB[shortID]
	if targetURL == "" {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Redirect(http.StatusTemporaryRedirect, targetURL)
	}
}
