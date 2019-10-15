package pages

import "github.com/gin-gonic/gin"

/**
* all admin modules route will store here
 */
func Routes(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET("pages", Index)
	r.PUT("pages/:id", Update)
	r.GET("pages/:id", Show)
	r.POST("pages/image/:id", UploadImage)
	r.DELETE("pages/image/:id", DeleteImage)
	r.DELETE("pages/images/:id", DeletePageImages)

	return r
}
