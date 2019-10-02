package test

import (
	"github.com/stretchr/testify/assert"
	"golang-starter/app/models"
	"golang-starter/config"
	"gopkg.in/thedevsaddam/gojsonq.v2"
	"net/http/httptest"
	"testing"
	"time"
)

/***
* check if user not send any required data
 */
func TestResetWithoutEmail(t *testing.T) {
	data := models.User{
	}
	w := post(data, "reset", true)
	assert.Equal(t, 400, w.Code)
}

/***
* check if user not send any required data
 */
func TestResetWithNotValidEmail(t *testing.T) {
	data := models.User{
		Email: "zizo199988",
	}
	w := post(data, "reset", true)
	assert.Equal(t, 400, w.Code)
}

/***
* first register new user
* send valid request to reset email
 */
func TestResetValidRequest(t *testing.T) {
	w, _ := registerNewUser()
	assert.Equal(t, 200, w.Code)
	r, _ := resetPassword()
	assert.Equal(t, 200, r.Code)
}

/***
* send recover request without token
 */
func TestRecoverWithoutToken(t *testing.T) {
	data := models.Recover{
		Password: "123456",
	}
	w := post(data, "recover", true)
	assert.Equal(t, 400, w.Code)
}

/***
* send recover request without password
 */
func TestRecoverWithoutPassword(t *testing.T) {
	data := models.Recover{
		Token: "77127127217219441273812371923",
	}
	w := post(data, "recover", true)
	assert.Equal(t, 400, w.Code)
}

/***
* first register new user
* send valid request to reset email
 */
func TestRecoverValidRequest(t *testing.T) {
	/**
	* first send register request to register new user
	 */
	w, data := registerNewUser()
	assert.Equal(t, 200, w.Code)
	/***
	* send reset password with the user email
	 */
	r, _ := resetPassword()
	assert.Equal(t, 200, r.Code)
	/***
	* send recover password then check if it not equal with the old data
	 */
	var user models.User
	config.DB.Find(&user, "email = ?", data.Email)
	recoverRequest := models.Recover{
		Token:    user.Token,
		Password: "123412312",
	}
	c := post(recoverRequest, "recover", false)
	time.Sleep(100000)
	responseData := responseData(c.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.NotEqual(t, user.Token, recoverResponse.Find("data.token"))
	assert.NotEqual(t, user.Password, recoverResponse.Find("data.password"))
	assert.Equal(t, 200, c.Code)
}

/**
* function make
*/
func resetPassword() (*httptest.ResponseRecorder, models.Reset) {
	resetRequest := models.Reset{
		Email: "zizo199988@gmail.com",
	}
	r := post(resetRequest, "reset", false)
	time.Sleep(100000)
	return r, resetRequest
}
