package routes

import (
	"net/http"

	"github.com/enchant97/url-shorter/core"
	"github.com/enchant97/url-shorter/core/db"
	"github.com/gin-gonic/gin"
)

func PostApiNew(c *gin.Context) {
	var formValues core.CreateShort
	if err := c.ShouldBindJSON(&formValues); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"detail": "missing required field(s)",
		})
		return
	}
	shortRow := formValues.GenerateShort()
	if err := shortRow.Create(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"detail": "unable to create short",
		})
		return
	}
	c.JSON(http.StatusOK, core.ShortToAPIShort(shortRow))
}

func GetApiInfo(c *gin.Context) {
	shortID := c.Param("shortID")
	decodedID, err := core.DecodeIDPadded(shortID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	shortRow := db.GetShortByID(uint(decodedID))
	if shortRow == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"detail": "short id not found",
		})
	} else {
		c.JSON(
			http.StatusOK,
			core.ShortToAPIShort(*shortRow),
		)
	}
}
