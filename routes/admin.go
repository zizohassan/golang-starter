package routes

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/controllers/admin/categories"
	"golang-starter/app/controllers/admin/settings"
	"golang-starter/app/controllers/admin/translations"
)

/***
* any route here will add after /admin
* admin only  will have access this routes
 */
func Admin(r *gin.RouterGroup) *gin.RouterGroup {
	/// category routes
	categories.Routes(r)
	translations.Routes(r)
	settings.Routes(r)

	return r
}



