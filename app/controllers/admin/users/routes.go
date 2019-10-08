package users

import "github.com/gin-gonic/gin"

/**
* all admin modules route will store here
*/
func Routes(r *gin.RouterGroup) *gin.RouterGroup  {
	r.GET("users" , Index)
	r.POST("users" , Store)
	r.PUT("users/:id" , Update)
	r.GET("users/:id" , Show)
	r.DELETE("users/:id" , Delete)

	return r
}
