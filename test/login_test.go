package test

import (
	"github.com/stretchr/testify/assert"
	"golang-starter/app/models"
	"golang-starter/config"
	"golang-starter/helpers"
	"testing"
)

/**
* register with valid data and
* login with this data
*/
func TestLoginWithValidData(t *testing.T)  {
	registerNewUser(true)
	loginData := models.User{
		Email:    "zizo199988@gmail.com",
		Password: "123457",
	}
	l := postWitOutHeader(loginData , "login" , false)
	assert.Equal(t, 200, l.Code)
}

/**
* register with valid data
* login with not valid email
*/
func TestLoginWithNotValidEmail(t *testing.T)  {
	registerNewUser(true)
	loginData := models.Login{
		Email:"zizo@gmail.com",
		Password:"123457",
	}
	l := postWitOutHeader(loginData , "login" , false)
	assert.Equal(t, 404, l.Code)
}
/**
* register with valid data
* login with not valid password
*/
func TestLoginWithNotValidPassword(t *testing.T)  {
	registerNewUser(true)
	loginData := models.Login{
		Email:"zizo199988@gmail.com",
		Password:"12345745",
	}
	l := postWitOutHeader(loginData , "login" , false)
	assert.Equal(t, 404, l.Code)
}

/**
* register with valid data
* login with block user
 */
func TestLoginWithBlockUser(t *testing.T) {
	registerNewUser(true)
	var user models.User
	config.DB.First(&user).Update("block", 1)
	loginData := models.Login{
		Email:    user.Email,
		Password: "123457",
	}
	l := postWitOutHeader(loginData, "login", false)
	assert.Equal(t, 403, l.Code)
}

/**
* Test Required inputs
 */
func TestLoginRequireInputs(t *testing.T) {
	///not send email
	checkPostRequestWithOutHeadersDataIsValid(t, models.Login{
		Password: "123457",
	}, "login", true)
	///not send password
	checkPostRequestWithOutHeadersDataIsValid(t, models.Login{
		Email: "zizo199988@gmail.com",
	}, "login", true)
}

/**
* Test not valid inputs
 */
func TestLoginNotValidInputs(t *testing.T) {
	///not valid email
	checkPostRequestWithOutHeadersDataIsValid(t, models.Login{
		Email:    "sasdasd",
		Password: "123457",
	}, "login", true)
}

/**
* Test inputs limitaion
 */
func TestLoginInputsLimitation(t *testing.T) {
	///min email fails
	checkPostRequestWithOutHeadersDataIsValid(t, models.Login{
		Email:    "s@s.i",
		Password: "123457",
	}, "login", true)
	///max email fail
	checkPostRequestWithOutHeadersDataIsValid(t, models.Login{
		Email:    helpers.RandomString(50) + "@gmail.com",
		Password: "123457",
	}, "login", true)
	///max password fails
	checkPostRequestWithOutHeadersDataIsValid(t, models.Login{
		Email:    "zizohassan@gmail.com",
		Password: helpers.RandomString(52),
	}, "login", true)
	///min password fails
	checkPostRequestWithOutHeadersDataIsValid(t, models.Login{
		Email:    "zizohassan@gmail.com",
		Password: helpers.RandomString(5),
	}, "login", true)
}
