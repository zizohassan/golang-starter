package visitor

import (
	"github.com/gin-gonic/gin"
	"starter/helpers"
	"starter/models"
	"starter/requests/visitor"
)

/**
* check if user have access to login in system
*/
func Login(g *gin.Context) {
	/**
	* init user login struct to validate and use later
	*/
	var login models.Login
	/**
	* get request and parse it to validation
	* if there any error will return with message
	*/
	err := visitor.Login(g.Request , login)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return
	*/
	if helpers.ReturnNotValidRequest(err, g) { return }
	///TODO:: here we will put login logic

	/**
	* now user is login we can return his info
	*/
	helpers.OkResponse(g , "you are login now" , models.UserResponse(models.User{}) , 200)
}
