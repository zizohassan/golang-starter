package routes

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/controllers/admin/categories"
)

/***
* any route here will add after /admin
* admin only  will have access this routes
 */
func Admin(r *gin.RouterGroup) *gin.RouterGroup {
	/// category routes
	categories.Routes(r)

	return r
}



