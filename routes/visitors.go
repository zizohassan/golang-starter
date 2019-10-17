package routes

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/controllers/visitor/auth"
	"golang-starter/app/controllers/visitor/general"
	"golang-starter/app/controllers/visitor/pages"
)

/***
* any route here will add after /
* anyone will have access this routes
 */
func Visitor(r *gin.RouterGroup) *gin.RouterGroup {
	general.Routes(r)
	auth.Routes(r)
	pages.Routes(r)
	/// serve static files like images
	r.Static("/public" , "./public")

	return r
}
