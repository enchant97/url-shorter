package core

import (
	"github.com/enchant97/go-gincookieauth/extras"
	"github.com/enchant97/url-shorter/core/flash"
	"github.com/gin-gonic/gin"
)

// Replaces gin.Context.HTML() to add specific data into them
func HTMLTemplate(c *gin.Context, code int, name string, obj gin.H) {
	obj[flash.FlashesKey] = flash.ReadFlashes(c)
	extras.TemplateWithAuth(c, code, name, obj)
}
