package core

import (
	"github.com/enchant97/go-gincookieauth/extras"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const FlashesKey = "Flashes"

// A flashed message
type Flash struct {
	Message string
	Type    string
}

// Add a new flashed message to session
func WriteFlash(c *gin.Context, flash Flash) {
	session := sessions.Default(c)
	session.AddFlash(flash)
	if err := session.Save(); err != nil {
		panic(err)
	}
}

// Read all flashed messages for session
func ReadFlashes(c *gin.Context) []Flash {
	session := sessions.Default(c)
	flashes := session.Flashes()
	if len(flashes) != 0 {
		if err := session.Save(); err != nil {
			panic(err)
		}
	}
	var flashesMessages []Flash
	for _, flash := range flashes {
		flashesMessages = append(flashesMessages, flash.(Flash))
	}
	return flashesMessages
}

// Replaces gin.Context.HTML() to add specific data into them
func HTMLTemplate(c *gin.Context, code int, name string, obj gin.H) {
	obj[FlashesKey] = ReadFlashes(c)
	extras.TemplateWithAuth(c, code, name, obj)
}
