package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/config"
	"golang-starter/helpers"
)

/**
* This middleware will Allow only auth user
* Will block not auth user
* if user allow to access then this middleware will add
* one header with user information you can use later (AUTH_DATA)
* in function you call
 */
func Auth() gin.HandlerFunc {
	return func(g *gin.Context) {
		var user models.User
		/// get Authorization header to check if user send it first
		adminToken := g.GetHeader("Authorization")
		if adminToken == "" {
			helpers.ReturnYouAreNotAuthorize(g)
			g.Abort()
			return
		}
		/// check if token exits in database
		config.DB.Where("token = ? " , adminToken).First(&user)
		if user.ID == 0 {
			helpers.ReturnYouAreNotAuthorize(g)
			g.Abort()
			return
		}
		/// check if user block or not
		if user.Status != models.BLOCK {
			helpers.ReturnYouAreNotAuthorize(g)
			g.Abort()
			return
		}
		/// not set header with user information
		userJson, _ := json.Marshal(&user)
		g.Request.Header.Set("AUTH_DATA", string(userJson))
		g.Next()
	}
}
