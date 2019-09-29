package routes

import (
	"github.com/gin-gonic/gin"
	"golang-starter/controllers/visitor"
)

/***
* any route here will add after /
* anyone will have access this routes
 */
func Visitor(r *gin.RouterGroup) *gin.RouterGroup {
	/// start auth apis
	r.POST("login" , visitor.Login)
	r.POST("register" , visitor.Register)
	/// end auth apis

	return r
}
