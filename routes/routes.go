package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, secretKey []byte) {
	store := cookie.NewStore(secretKey)
	session := sessions.Sessions("APP-SESSION", store)

	r.GET("/:shortID", GetRedirect)

	frontendRoutes := r.Group("/")
	{
		frontendRoutes.Use(session)
		frontendRoutes.GET("/", GetIndex)
		frontendRoutes.GET("/checker", GetChecker)
		frontendRoutes.GET("/new", GetNew)
		frontendRoutes.POST("/new", PostNew)
	}
	userRoutes := r.Group("/users")
	{
		userRoutes.Use(session)
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
