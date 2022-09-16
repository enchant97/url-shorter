package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	frontendRoutes := r.Group("/")
	{
		frontendRoutes.GET("/", GetIndex)
		frontendRoutes.GET("/checker", GetChecker)
		frontendRoutes.GET("/new", GetNew)
		frontendRoutes.POST("/new", PostNew)
		frontendRoutes.GET("/:shortID", GetRedirect)
	}
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/new", GetNewUser)
		userRoutes.POST("/new", PostNewUser)
	}
	apiRoutes := r.Group("/api")
	{
		apiRoutes.POST("/short", PostApiNew)
		apiRoutes.GET("/short/:shortID", GetApiInfo)
	}
}
