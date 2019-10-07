package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
)

/**
* here we set the default language
* if we have Accept-Language in header we will work on it
* else we will pass the default language
*/
func Language() gin.HandlerFunc {
	return func(g *gin.Context) {
		lang := g.GetHeader("Accept-Language")
		if lang == "" {
			g.Request.Header.Set("Accept-Language", os.Getenv("DEFAULT_LANG"))
		}
		g.Next()
	}
}
