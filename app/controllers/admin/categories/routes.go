package categories

import "github.com/gin-gonic/gin"

/**
* all admin modules route will store here
*/
func Routes(r *gin.RouterGroup) *gin.RouterGroup  {
	r.GET("categories" , Index)
	r.POST("categories" , Store)
	r.PUT("categories/:id" , Update)
	r.GET("categories/:id" , Show)
	r.DELETE("categories/:id" , Delete)

	return r
}
