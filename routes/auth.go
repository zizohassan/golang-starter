package routes

import (
	"github.com/gin-gonic/gin"
)

/***
* any route here will add after /auth
* and only login user will have access this routes
*/
func Auth(r *gin.RouterGroup) *gin.RouterGroup {


	return r
}
