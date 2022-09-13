package routes

import (
	"net/http"

	"github.com/enchant97/url-shorter/core"
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"pageTitle": "Home",
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
	shortID := randstr.String(8)
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

func GetShortInfo(c *gin.Context) {
	shortID := c.Param("shortID")
	targetURL := core.FakeShortsDB[shortID]
	if targetURL == "" {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.HTML(http.StatusOK, "short_info.html", gin.H{
			"pageTitle": "Short Info",
			"short": createdShort{
				ShortID:   shortID,
				TargetURL: targetURL,
			},
		})
	}
}
