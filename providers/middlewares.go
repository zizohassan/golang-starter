package providers

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/middleware"
)

func middlewares(r *gin.Engine) *gin.Engine {
	/// run cors middleware
	r.Use(middleware.CORSMiddleware())

	return r
}
