package middleware

import (
	"github.com/gin-gonic/gin"
)

/**
* This middle ware will Allow only
* Will block not admin role admin role is (2)
* if user allow to access then this middleware will add
* one header with user information you can use later (ADMIN_DATA)
* in function you call
 */
func Admin() gin.HandlerFunc {
	return func(g *gin.Context) {
		//var user models.User
		///// get Authorization header to check if user send it first
		//adminToken := g.GetHeader("Authorization")
		//if adminToken == "" {
		//	helpers.ReturnYouAreNotAuthorize(g)
		//	g.Abort()
		//	return
		//}
		///// check if token exits in database
		//config.DB.Where("token = ? and role = ?", adminToken, 2).First(&user)
		//if user.ID == 0 {
		//	helpers.ReturnYouAreNotAuthorize(g)
		//	g.Abort()
		//	return
		//}
		///// check if user block or not
		//if user.Block != 2 {
		//	helpers.ReturnYouAreNotAuthorize(g)
		//	g.Abort()
		//	return
		//}
		///// not set header with user information
		//userJson, _ := json.Marshal(&user)
		//g.Request.Header.Set("ADMIN_DATA", string(userJson))
		g.Next()
	}
}
