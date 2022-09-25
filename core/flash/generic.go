package flash

import "github.com/gin-gonic/gin"

// Generic error flash
func FlashError(c *gin.Context, message string) {
	WriteFlash(c, Flash{
		Message: message,
		Type:    "error",
	})
}

// Generic warning flash
func FlashWarning(c *gin.Context, message string) {
	WriteFlash(c, Flash{
		Message: message,
		Type:    "warning",
	})
}

// Generic info flash
func FlashInfo(c *gin.Context, message string) {
	WriteFlash(c, Flash{
		Message: message,
		Type:    "info",
	})
}
