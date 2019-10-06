package providers

import (
	"github.com/gin-gonic/gin"
	"golang-starter/app/middleware"
	"os"
	"golang-starter/routes"
)

/**
* sets routing group you can edit any group
* slugs by edit env file
*/
func Routing(r *gin.Engine) *gin.Engine {
	admin := r.Group(os.Getenv("ADMIN_SLUG"))
	admin.Use(middleware.Admin())
	{
		routes.Admin(admin)
	}
	/// Auth users only can access these routes
	auth := r.Group(os.Getenv("AUTH_SLUG"))
	admin.Use(middleware.Auth())
	{
		routes.Auth(auth)
	}
	/// any one can access this routes
	visitor := r.Group(os.Getenv("VISTORS_SLUG"))
	{
		routes.Visitor(visitor)
	}

	return r
}
