package main

import (
	"path/filepath"
	"text/template"

	"github.com/enchant97/url-shorter/core"
	"github.com/enchant97/url-shorter/core/db"
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
		funcMap := template.FuncMap{
			"timeToHumanOr": core.TimeToHumanOr,
		}
		r.AddFromFilesFuncs(filepath.Base(include), funcMap, files...)
	}
	return r
}

func main() {
	r := gin.Default()

	var appConfig core.AppConfig
	if err := appConfig.ParseConfig(); err != nil {
		panic(err)
	}

	db.InitDB(appConfig.SQLitePath)

	r.HTMLRender = loadTemplates("./templates")
	r.Static("/static", "./static")
	routes.InitRoutes(r, appConfig.SecretKey)
	r.Run("localhost:8080")
}
