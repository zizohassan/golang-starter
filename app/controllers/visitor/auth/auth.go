package auth

import (
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
	// init user login struct to validate request
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
	user, valid := checkUserExistsNotBlocked(g, login.Email, "")
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
	 // update token then return with the new data
	token, _ := helpers.GenerateToken(user.Password + user.Email)
	config.DB.Model(&user).Update("token", token).First(&user)
	// now user is login we can return his info
	helpers.OkResponse(g, "you are login now", transformers.UserResponse(user))
}

/**
* Register new user on system
 */
func Register(g *gin.Context) {
	// init visitor login struct to validate request
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
		helpers.ReturnDuplicateData(g, "email")
		return
	}
	/**
	* set role and block
	* role 1 is user
	* block user (1 , 2) 2 is not block 1 is block
	*/
	user.Role =  1
	user.Status = models.ACTIVE
	/**
	* create new user based on register struct
	* token , role  , block will set with event
	*/
	config.DB.Create(&user)
	// now user is login we can return his info
	helpers.OkResponse(g, "Thank you for register in our system", transformers.UserResponse(*user))
}

/**
* recover password take request token
* select user that have this token
* if user token valid and user not block
* then user can  recover his password
 */
func Recover(g *gin.Context) {
	//init Reset struct to validate request
	recoverPassword := new(models.Recover)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := visitor.Recover(g.Request, recoverPassword)
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
	user, valid := checkUserExistsNotBlocked(g, "", recoverPassword.Token)
	if !valid {
		return
	}
	/**
	* now update token and update password
	* we update token to make it the old link not valid
	*/
	encPassword, _ := helpers.HashPassword(recoverPassword.Password)
	token, _ := helpers.GenerateToken(user.Password + user.Email)
	config.DB.Model(&user).Updates(map[string]interface{}{"password": encPassword, "token": token}).First(&user)
	// notice user that his password has been changes
	sendRecoverPasswordEmail(user)
	// return ok response
	helpers.OkResponse(g, "Your password has been set , and your token changes", transformers.UserResponse(user))
}

/***
* notice user that his password has been updated
 */
func sendRecoverPasswordEmail(user models.User) {
	msg := "Your Password has been updated to (" + user.Password + ")" + "\n"
	msg += "Do not worry your password is encrypted , this just note for your activity" + "\n"
	msg += os.Getenv("RESET_PASSWORD_URL") + user.Token
	helpers.SendMail(user.Email, "Your password has been updated", msg)
}

/**
* reset password
* with email you can send reset link
* to user email
 */
func Reset(g *gin.Context) {
	// init Reset struct to validate request
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
	user, valid := checkUserExistsNotBlocked(g, reset.Email, "")
	if !valid {
		return
	}
	sendRestLink(user)
	// return ok response
	var data map[string]interface{}
	helpers.OkResponse(g, "We send your reset password link on your email", data)
}

/**
* create reset password link
* send it to user email
*/
func sendRestLink(user models.User)  {
	msg := "Your Request To reset your password if you take this action click on this link to reset your password " + "\n"
	msg += os.Getenv("RESET_PASSWORD_URL") + user.Token
	helpers.SendMail(user.Email, "Reset Password Request", msg)
}
/**
* check if user exists
* check if user not blocked
 */
func checkUserExistsNotBlocked(g *gin.Context, email string, token string) (models.User, bool) {
	// init user struct binding data for user
	var user models.User
	/**
	* check if this email exists database
	* if this email will not found will return not found
	* will return 404 code
	* will select by email if token is empty
	* if token not empty select by token
	 */
	if token != "" {
		config.DB.Find(&user, "token = ? ", token)
	} else {
		config.DB.Find(&user, "email = ? ", email)
	}
	if user.ID == 0 {
		helpers.ReturnNotFound(g, "We not found this user on system")
		return user, false
	}
	// if user block
	if user.Status == models.BLOCK {
		helpers.ReturnForbidden(g, "You are blocked from the system")
		return user, false
	}
	return user, true
}
