package test

import (
	"github.com/stretchr/testify/assert"
	"golang-starter/app/models"
	"golang-starter/config"
	"golang-starter/helpers"
	"net/http/httptest"
	"testing"
	"time"
)

/***
* first register new user
* send valid request to reset email
*/
func TestResetValidRequest(t *testing.T) {
	w, _ := registerNewUser(true)
	assert.Equal(t, 200, w.Code)
	r, _ := resetPassword(false)
	assert.Equal(t, 200, r.Code)
}

/***
* first register new user
* send valid request to reset email
 */
func TestRecoverValidRequest(t *testing.T) {
	/**
	* first send register request to register new user
	 */
	_, data := registerNewUser(true)
	/***
	* send reset password with the user email
	 */
	r, _ := resetPassword(false)
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
	c := postWitOutHeader(recoverRequest, "recover", false)
	time.Sleep(100000)
	d := getDataMap(c)
	assert.NotEqual(t, user.Token, d["token"])
	assert.NotEqual(t, user.Password, d["password"])
	assert.Equal(t, 200, c.Code)
}

/**
* Test Required inputs
 */
func TestResetRequireInputs(t *testing.T) {
	///not send email
	checkPostRequestWithOutHeadersDataIsValid(t, models.Reset{

	}, "reset", true)
	/// do not send token
	checkPostRequestWithOutHeadersDataIsValid(t, models.Recover{
		Password: "123232323",
	}, "recover", true)
	/// do not send password
	checkPostRequestWithOutHeadersDataIsValid(t, models.Recover{
		Token: "123232323",
	}, "recover", true)
}

/**
* Test not valid inputs
 */
func TestResetNotValidInputs(t *testing.T) {
	///not valid email
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email: "sasdasd",
	}, "reset", true)
}

/**
* Test inputs limitaion
 */
func TestResetInputsLimitation(t *testing.T) {
	///min email fails
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email: "s@s.i",
	}, "reset", true)
	///max email fail
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Email: helpers.RandomString(50) + "@gmail.com",
	}, "reset", true)
	///max password fail
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Password:helpers.RandomString(50),
		Token:"21323123123123",
	}, "reset", true)
	///min password fail
	checkPostRequestWithOutHeadersDataIsValid(t, models.User{
		Password:helpers.RandomString(5),
		Token:"21323123123123",
	}, "reset", true)
}

func resetPassword(migrate bool) (*httptest.ResponseRecorder, models.Reset) {
	resetRequest := models.Reset{
		Email: "zizo199988@gmail.com",
	}
	r := postWitOutHeader(resetRequest, "reset", migrate)
	time.Sleep(100000)
	return r, resetRequest
}
