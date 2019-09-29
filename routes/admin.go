package routes

import (
	"github.com/gin-gonic/gin"
)

/***
* any route here will add after /admin
* admin only  will have access this routes
 */
func Admin(r *gin.RouterGroup) *gin.RouterGroup {
	return r
}



