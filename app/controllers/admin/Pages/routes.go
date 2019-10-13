package pages

import "github.com/gin-gonic/gin"

/**
* all admin modules route will store here
*/
func Routes(r *gin.RouterGroup) *gin.RouterGroup  {
	r.GET("pages" , Index)
	r.PUT("pages/:id" , Update)
	r.GET("pages/:id" , Show)

	return r
}
