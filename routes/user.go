package routes

import (
	"net/http"

	"github.com/enchant97/url-shorter/core"
	"github.com/enchant97/url-shorter/core/db"
	"github.com/gin-gonic/gin"
)

func GetNewUser(c *gin.Context) {
	c.HTML(http.StatusOK, "new-user.html", gin.H{
		"pageTitle": "New User",
	})
}

func PostNewUser(c *gin.Context) {
	var newUser core.CreateUser
	if err := c.ShouldBind(&newUser); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userRow := db.User{
		Username: newUser.Username,
	}
	userRow.SetPassword(newUser.Password)
	if err := db.DB.Create(&userRow).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// TODO navigate to login page
	c.Redirect(http.StatusSeeOther, "/")
}
