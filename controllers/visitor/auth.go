package visitor

import (
	"github.com/gin-gonic/gin"
	"golang-starter/config"
	"golang-starter/controllers/functions"
	"golang-starter/helpers"
	"golang-starter/models"
	"golang-starter/requests/visitor"
)

/**
* check if user have access to login in system
 */
func Login(g *gin.Context) {
	/**
	* init user login struct to validate request
	*/
	login := new(models.Login)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	*/
	err := visitor.Login(g.Request, login)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	*/
	if helpers.ReturnNotValidRequest(err, g) {
		return
	}
	/**
	* init user struct binding data for user
	 */
	var user models.User
	/**
	* check if this email exists database
	* if this email will not found will return not found
	* will return 404 code
	*/
	config.DB.Find(&user, "email = ? ", login.Email)
	if user.ID == 0 {
		helpers.ReturnNotFound(g, "We not found this user on system")
		return
	}
	/**
	* now check if password are valid
	* if user password is not valid we will return invalid email
	* or password
	*/
	check := functions.CheckPasswordHash(login.Password, user.Password)
	if !check {
		 helpers.ReturnNotFound(g, "your email or your password are not valid")
		return
	}
	/***
	* if user block
	*/
	if user.Block == 1{
		helpers.ReturnForbidden(g , "You are blocked from the system")
		return
	}
	/**
	* update token then return with the new data
	*/
	token, _ := functions.GenerateToken(user)
	config.DB.Model(&user).Update("token" , token).First(&user)
	/**
	* now user is login we can return his info
	*/
	helpers.OkResponse(g, "you are login now", models.UserResponse(user))
}

/**
* Register new user on system
*/
func Register(g *gin.Context)  {
	/**
	* init visitor login struct to validate request
	 */
	user := new(models.User)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := visitor.Register(g.Request, user)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err, g) {
		return
	}
	/**
	* check if this email exists database
	* if this email found will return
	 */
	checkUser := new(models.User)
	config.DB.Find(&checkUser, "email = ? ", user.Email)
	if checkUser.ID != 0 {
		helpers.ReturnFoundRow(g, "We found this email in our system")
		return
	}
	/**
	* create new user based on register struct
	* token , role  , block will set with event
	*/
	config.DB.Create(&user)
	/**
	* now user is login we can return his info
	 */
	helpers.OkResponse(g, "Thank you for register in our system", models.UserResponse(*user))
}
