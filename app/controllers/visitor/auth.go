package visitor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-starter/app/models"
	"golang-starter/app/requests/visitor"
	"golang-starter/app/transformers"
	"golang-starter/config"
	"golang-starter/helpers"
	"os"
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
	* check if user exists
	* check if user not blocked
	 */
	user, valid := checkUserExistsNotBlocked(g, login.Email)
	fmt.Println(user, valid)
	if !valid {
		return
	}
	/**
	* now check if password are valid
	* if user password is not valid we will return invalid email
	* or password
	 */
	check := helpers.CheckPasswordHash(login.Password, user.Password)
	if !check {
		helpers.ReturnNotFound(g, "your email or your password are not valid")
		return
	}
	/**
	* update token then return with the new data
	 */
	token, _ := helpers.GenerateToken(user.Password + user.Email)
	config.DB.Model(&user).Update("token", token).First(&user)
	/**
	* now user is login we can return his info
	 */
	helpers.OkResponse(g, "you are login now", transformers.UserResponse(user))
}

/**
* Register new user on system
 */
func Register(g *gin.Context) {
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
	config.DB.Find(&user, "email = ? ", user.Email)
	if user.ID != 0 {
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
	helpers.OkResponse(g, "Thank you for register in our system", transformers.UserResponse(*user))
}

/**
* reset password
* with email you can send reset link
* to user email
 */
func Reset(g *gin.Context) {
	/**
	* init Reset struct to validate request
	 */
	reset := new(models.Reset)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := visitor.Reset(g.Request, reset)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err, g) {
		return
	}
	/**
	* check if user exists
	* check if user not blocked
	 */
	user, valid := checkUserExistsNotBlocked(g, reset.Email)
	if !valid {
		return
	}
	/**
	* create reset password link
	 */
	msg := "Your Request To reset your password if you take this action click on this link to reset your password " + "\n"
	msg += os.Getenv("RESET_PASSWORD_URL") + user.Token
	helpers.SendMail(user.Email, "Reset Password Request", msg)
	/**
	* return ok response
	 */
	var data map[string]interface{}
	helpers.OkResponse(g, "We send your reset password link on your email", data)
}

/**
* check if user exists
* check if user not blocked
 */
func checkUserExistsNotBlocked(g *gin.Context, email string) (models.User, bool) {
	/**
	* init user struct binding data for user
	 */
	var user models.User
	/**
	* check if this email exists database
	* if this email will not found will return not found
	* will return 404 code
	 */
	config.DB.Find(&user, "email = ? ", email)
	if user.ID == 0 {
		helpers.ReturnNotFound(g, "We not found this user on system")
		return user, false
	}
	/***
	* if user block
	 */
	if user.Block == 1 {
		helpers.ReturnForbidden(g, "You are blocked from the system")
		return user, false
	}
	return user, true
}
