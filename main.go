package main

import (
	"path/filepath"

	"github.com/enchant97/url-shorter/routes"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// Load templates allowing for inheritance
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func main() {
	r := gin.Default()
	r.HTMLRender = loadTemplates("./templates")
	r.Static("/static", "./static")
	r.GET("/", routes.GetIndex)
	r.GET("/new", routes.GetNew)
	r.POST("/new", routes.PostNew)
	r.GET("/:shortID", routes.GetRedirect)
	r.GET("/:shortID/info", routes.GetShortInfo)
	apiRoutes := r.Group("/api")
	{
		apiRoutes.POST("/short", routes.PostApiNew)
		apiRoutes.GET("/short/:shortID", routes.GetApiInfo)
	}
	r.Run("localhost:8080")
}
