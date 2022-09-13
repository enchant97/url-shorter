package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
)

// until real database is implemented
var fakeShortsDB = map[string]string{}

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"pageTitle": "Home",
	})
}

func getNew(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{
		"pageTitle": "New",
	})
}

type NewShortFormValues struct {
	TargetUrl string `form:"target-url"`
}

func postNew(c *gin.Context) {
	var formValues NewShortFormValues
	if err := c.ShouldBind(&formValues); err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}
	shortId := randstr.String(8)
	fakeShortsDB[shortId] = formValues.TargetUrl
	c.Redirect(http.StatusSeeOther, "/"+shortId+"/info")
}

func getRedirect(c *gin.Context) {
	shortId := c.Param("short_id")
	targetUrl := fakeShortsDB[shortId]
	if targetUrl == "" {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Redirect(http.StatusTemporaryRedirect, targetUrl)
	}
}

func getShortInfo(c *gin.Context) {
	shortId := c.Param("short_id")
	targetUrl := fakeShortsDB[shortId]
	if targetUrl == "" {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.HTML(http.StatusOK, "short_info.html", gin.H{
			"pageTitle": "Short Info",
			"shortId":   shortId,
			"targetUrl": targetUrl,
		})
	}
}

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
	r.GET("/", getIndex)
	r.GET("/new", getNew)
	r.POST("/new", postNew)
	r.GET("/:short_id", getRedirect)
	r.GET("/:short_id/info", getShortInfo)
	r.Run("localhost:8080")
}
