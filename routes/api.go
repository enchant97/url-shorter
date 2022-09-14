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
	shortID := core.MakeShortID()
	db.CreateNewShort(shortID, formValues.TargetURL)
	c.JSON(
		http.StatusOK,
		core.Short{
			TargetURL: formValues.TargetURL,
			ShortID:   shortID,
		})
}

func GetApiInfo(c *gin.Context) {
	shortID := c.Param("shortID")
	shortRow := db.GetShortByShortID(shortID)
	if shortRow == (db.Short{}) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"detail": "short id not found",
		})
	} else {
		c.JSON(
			http.StatusOK,
			core.Short{
				ShortID:   shortID,
				TargetURL: shortRow.TargetURL,
			})
	}
}
