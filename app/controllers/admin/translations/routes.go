package translations

import "github.com/gin-gonic/gin"

/**
* all admin modules route will store here
*/
func Routes(r *gin.RouterGroup) *gin.RouterGroup  {
	r.GET("translations" , Index)
	r.PUT("translations/:id" , Update)
	r.GET("translations/:id" , Show)

	return r
}
