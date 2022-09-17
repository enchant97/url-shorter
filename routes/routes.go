package routes

import (
	"github.com/enchant97/go-gincookieauth"
	"github.com/enchant97/go-gincookieauth/extras"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, secretKey []byte) {
	session := extras.MakeSession("AUTH", secretKey)
	cookieAuth := gincookieauth.CookieAuth(gincookieauth.CookieAuthConfig{
		AuthRequired: false,
	})

	r.GET("/:shortID", GetRedirect)

	frontendRoutes := r.Group("/")
	{
		frontendRoutes.Use(session)
		frontendRoutes.Use(cookieAuth)
		frontendRoutes.GET("/", GetIndex)
		frontendRoutes.GET("/checker", GetChecker)
		frontendRoutes.GET("/new", GetNew)
		frontendRoutes.POST("/new", PostNew)
	}
	userRoutes := r.Group("/users")
	{
		userRoutes.Use(session)
		userRoutes.Use(cookieAuth)
		userRoutes.GET("/new", GetNewUser)
		userRoutes.POST("/new", PostNewUser)
		userRoutes.GET("/login", GetLoginUser)
		userRoutes.POST("/login", PostLoginUser)
		userRoutes.GET("/logout", GetLogoutUser)
	}
	apiRoutes := r.Group("/api")
	{
		apiRoutes.POST("/short", PostApiNew)
		apiRoutes.GET("/short/:shortID", GetApiInfo)
	}
}
