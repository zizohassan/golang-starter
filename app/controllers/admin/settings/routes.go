package settings

import "github.com/gin-gonic/gin"

/**
* all admin modules route will store here
*/
func Routes(r *gin.RouterGroup) *gin.RouterGroup  {
	r.GET("settings" , Index)
	r.PUT("settings/:id" , Update)
	r.GET("settings/:id" , Show)

	return r
}
