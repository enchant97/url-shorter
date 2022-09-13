package routes

import (
	"net/http"

	"github.com/enchant97/url-shorter/core"
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
)

func PostApiNew(c *gin.Context) {
	var formValues newShortFormValues
	if err := c.ShouldBindJSON(&formValues); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "missing required field(s)",
		})
		return
	}
	shortID := randstr.String(8)
	core.FakeShortsDB[shortID] = formValues.TargetUrl
	c.JSON(
		http.StatusOK,
		createdShort{
			TargetURL: formValues.TargetUrl,
			ShortID:   shortID,
		})
}

func GetApiInfo(c *gin.Context) {
	shortID := c.Param("shortID")
	targetURL := core.FakeShortsDB[shortID]
	if targetURL == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"detail": "short id not found",
		})
	} else {
		c.JSON(
			http.StatusOK,
			createdShort{
				ShortID:   shortID,
				TargetURL: targetURL,
			})
	}
}
