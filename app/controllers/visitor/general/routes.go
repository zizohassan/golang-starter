package general

import "github.com/gin-gonic/gin"

/**
* all admin modules route will store here
 */
func Routes(r *gin.RouterGroup) *gin.RouterGroup {
	/// init project
	r.GET("init", Init)

	return r
}
