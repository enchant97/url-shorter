package routes

import (
	"net/http"

	"github.com/enchant97/go-gincookieauth"
	"github.com/enchant97/go-gincookieauth/extras"
	"github.com/enchant97/url-shorter/core"
	"github.com/gin-gonic/gin"
)

func ensureAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, exists := c.Get(gincookieauth.GlobalDataKey); exists {
			user := user.(gincookieauth.AuthData)
			if user.IsAuthenticated {
				c.Next()
				return
			}
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func setData(key string, value interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, value)
		c.Next()
	}
}

func InitRoutes(r *gin.Engine, appConfig core.AppConfig) {
	session := extras.MakeSession("AUTH", appConfig.SecretKey)
	cookieAuth := gincookieauth.CookieAuth(gincookieauth.CookieAuthConfig{
		AuthRequired: false,
	})
	ensureAuthHandler := ensureAuth()
	appConfigHandler := setData("AppConfig", appConfig)

	r.GET("/:shortID", GetRedirect)

	frontendRoutes := r.Group("/")
	{
		frontendRoutes.Use(session, cookieAuth, appConfigHandler)
		frontendRoutes.GET("/", GetIndex)
		frontendRoutes.GET("/checker", GetChecker)
	}
	userRoutes := r.Group("/users")
	{
		userRoutes.Use(session, cookieAuth, appConfigHandler)
		userRoutes.GET("/login", GetLoginUser)
		userRoutes.POST("/login", PostLoginUser)
		userRoutes.GET("/logout", GetLogoutUser, ensureAuthHandler)
	}
	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/short/:shortID", GetApiInfo)
	}
	// Routes with launch config
	if appConfig.RequireLogin {
		frontendRoutes.GET("/new", GetNewAuthRequired)
		frontendRoutes.POST("/new", PostNew, ensureAuthHandler)
		apiRoutes.POST("/short", PostApiNew, ensureAuthHandler)
	} else {
		frontendRoutes.GET("/new", GetNew)
		frontendRoutes.POST("/new", PostNew)
		apiRoutes.POST("/short", PostApiNew)
	}
	if appConfig.AllowNewAccounts {
		userRoutes.GET("/new", GetNewUser)
		userRoutes.POST("/new", PostNewUser)
	} else {
		userRoutes.GET("/new", GetNewUserAuthRequired)
		userRoutes.POST("/new", PostLoginUser, ensureAuthHandler)
	}
}
