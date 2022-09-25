package routes

import (
	"net/http"

	"github.com/enchant97/go-gincookieauth"
	"github.com/enchant97/url-shorter/core"
	"github.com/enchant97/url-shorter/core/db"
	"github.com/enchant97/url-shorter/core/flash"
	"github.com/gin-gonic/gin"
)

func GetNewUser(c *gin.Context) {
	core.HTMLTemplate(c, http.StatusOK, "new-user.html", gin.H{
		"pageTitle": "New User",
	})
}

func GetNewUserAuthRequired(c *gin.Context) {
	core.HTMLTemplate(c, http.StatusUnauthorized, "error.html", gin.H{
		"pageTitle":    "New User",
		"errorTitle":   "Account Creation Disabled",
		"errorMessage": "No new user accounts can be created, as it has been disabled.",
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
	if gincookieauth.GetUserID(c) != nil {
		// user already logged in
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	core.HTMLTemplate(c, http.StatusOK, "login-user.html", gin.H{
		"pageTitle": "Login",
	})
}

func PostLoginUser(c *gin.Context) {
	var userLogin core.LoginUser
	// validate login form fields
	if err := c.ShouldBind(&userLogin); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// check user from db, if match is found do login stuff
	if user := db.GetUserByUsername(userLogin.Username); user != nil {
		if user.IsPasswordMatch(userLogin.Password) {
			gincookieauth.LoginUser(c, user.ID)
			c.Redirect(http.StatusSeeOther, "/")
			return
		}
	}
	// failed login check
	flash.FlashError(c, "Login details invalid")
	c.Redirect(http.StatusSeeOther, "/users/login")
}

func GetLogoutUser(c *gin.Context) {
	gincookieauth.LogoutUser(c, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
