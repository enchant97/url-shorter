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
	c.Redirect(http.StatusSeeOther, "/users/login")
}

func GetLoginUser(c *gin.Context) {
	if core.GetAuthenticatedUserID(c) != nil {
		// user already logged in
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	c.HTML(http.StatusOK, "login-user.html", gin.H{
		"pageTitle": "Login",
	})
}

func PostLoginUser(c *gin.Context) {
	var userLogin core.LoginUser
	if err := c.ShouldBind(&userLogin); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if user := db.GetUserByUsername(userLogin.Username); user != nil {
		if user.IsPasswordMatch(userLogin.Password) {
			core.SetAuthenticatedUserID(c, user.ID)
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
	}
	c.Redirect(http.StatusSeeOther, "/users/login")
}

func GetLogoutUser(c *gin.Context) {
	core.RemoveAuthenticatedUser(c)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
