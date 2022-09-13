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
	shortId := randstr.String(8)
	core.FakeShortsDB[shortId] = formValues.TargetUrl
	c.Redirect(http.StatusSeeOther, "/"+shortId+"/info")
}

func GetRedirect(c *gin.Context) {
	shortId := c.Param("short_id")
	targetUrl := core.FakeShortsDB[shortId]
	if targetUrl == "" {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Redirect(http.StatusTemporaryRedirect, targetUrl)
	}
}

func GetShortInfo(c *gin.Context) {
	shortId := c.Param("short_id")
	targetUrl := core.FakeShortsDB[shortId]
	if targetUrl == "" {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.HTML(http.StatusOK, "short_info.html", gin.H{
			"pageTitle": "Short Info",
			"shortId":   shortId,
			"targetUrl": targetUrl,
		})
	}
}
