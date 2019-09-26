package routes

import (
	"github.com/gin-gonic/gin"
	"starter/controllers/visitor"
)

/***
* any route here will add after /
* anyone will have access this routes
 */
func Visitor(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST("login" , visitor.Login)

	return r
}
