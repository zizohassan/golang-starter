package auth

import (
	"github.com/gin-gonic/gin"
)

/**
* all admin modules route will store here
 */
func Routes(r *gin.RouterGroup) *gin.RouterGroup {
	/// start not auth routes
	r.POST("login", Login)
	r.POST("register", Register)
	r.POST("reset", Reset)
	r.POST("recover", Recover)

	return r
}
