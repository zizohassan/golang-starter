package faqs

import "github.com/gin-gonic/gin"

/**
* all admin modules route will store here
*/
func Routes(r *gin.RouterGroup) *gin.RouterGroup  {
	r.GET("faqs" , Index)
	r.POST("faqs" , Store)
	r.PUT("faqs/:id" , Update)
	r.GET("faqs/:id" , Show)
	r.DELETE("faqs/:id" , Delete)

	return r
}
