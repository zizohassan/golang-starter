package test

import (
	"github.com/stretchr/testify/assert"
	"golang-starter/app/models"
	"golang-starter/config"
	"golang-starter/helpers"
	"net/http/httptest"
	"testing"
)

var usersUrl = "admin/users"

///// show all case
func TestUsersShowAll(t *testing.T) {
	k := get(usersUrl, false, getTokenAsHeader(true))
	assert.Equal(t, 0.0, returnResponseKey(k, "data.offset"))
	assert.Equal(t, 200, k.Code)
}

func TestUsersFilter(t *testing.T) {
	token := getTokenAsHeader(true)
	w := newUser(t, false, token)
	assert.Equal(t, 200, w.Code)
	filter(t, usersUrl, 1, "role", "equal", token)
	filter(t, usersUrl, "Abdel Aziz", "name", "equal", token)
	filter(t, usersUrl, 3, "block", "not-equal", token)
	filter(t, usersUrl, 2, "block", "equal", token)
	filter(t, usersUrl, "zizo1999988@gmail.com", "email", "equal", token)
}

///// show function cases
func TestUsersShowWithValidId(t *testing.T) {
	token := getTokenAsHeader(true)
	w := newUser(t, false, token)
	assert.Equal(t, 200, w.Code)
	k := get(usersUrl+"/2", false, token)
	assert.Equal(t, "Abdel Aziz", returnResponseKey(k, "data.name"))
	assert.Equal(t, 200, k.Code)
}

func TestUsersShowWithNotValidId(t *testing.T) {
	token := getTokenAsHeader(true)
	k := get(usersUrl+"/1000", false, token)
	assert.Equal(t, 404, k.Code)
}

///// delete case
func TestUsersDeleteWithValidId(t *testing.T) {
	token := getTokenAsHeader(true)
	k := deleter(usersUrl+"/1", false, token)
	assert.Equal(t, 200, k.Code)
}

func TestUsersDeleteWithNotValidId(t *testing.T) {
	token := getTokenAsHeader(true)
	k := deleter(usersUrl+"/1000", false, token)
	assert.Equal(t, 404, k.Code)
}

func TestUsersDeleteWithWrongRoute(t *testing.T) {
	token := getTokenAsHeader(true)
	k := deleter(usersUrl, false, token)
	assert.Equal(t, 404, k.Code)
}

/// valid store update cases
func TestStoreUsersWithValidData(t *testing.T) {
	token := getTokenAsHeader(true)
	w := newUser(t, false, token)
	assert.Equal(t, "Abdel Aziz", returnResponseKey(w, "data.name"))
	assert.Equal(t, 200, w.Code)
}

/**
* check if user has register with email before
 */
func TestAddUserWithExistEmail(t *testing.T)  {
	token := getTokenAsHeader(true)
	k := post(existsEmailData() , usersUrl , false , token)
	assert.Equal(t, 409, k.Code)
}

func TestUpdateUserWithExistEmail(t *testing.T)  {
	token := getTokenAsHeader(true)
	k := post(existsEmailData() , usersUrl , false , token)
	assert.Equal(t, 409, k.Code)
}

/**
* Test not valid inputs
 */
func TestAddUserNotValidInputs(t *testing.T) {
	token := getTokenAsHeader(true)
	///not valid email
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:  "Abdel Aziz hassan Abdel Aziz",
		Role:  2,
		Block: 1,
		Email: "zizo199988",
		Password: "12121221",
	}, usersUrl, false , token)
	usersUrl = usersUrl + "/1"
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:  "Abdel Aziz hassan Abdel Aziz",
		Role:  2,
		Block: 1,
		Email: "zizo199988",
		Password: "12121221",
	}, usersUrl, false , token)
}


func TestUpdateUsersWithValidWithOutPasswordData(t *testing.T) {
	token := getTokenAsHeader(true)
	_ = newUser(t, false, token)
	data := models.User{
		Name:  "Abdel Aziz hassan Abdel Aziz",
		Role:  2,
		Block: 1,
		Email: "zizo@gmail.com",
	}
	k := put(data, usersUrl+"/2", false, token)
	var row models.User
	config.DB.Where("id = 2").Find(&row)
	d := getDataMap(k)
	assert.EqualValues(t, row.Name, d["name"])
	assert.EqualValues(t, row.Role, d["role"])
	assert.EqualValues(t, row.Block, d["block"])
	assert.EqualValues(t, row.Email, d["email"])
	assert.Equal(t, 200, k.Code)
}

func TestUpdateUsersWithValidWithPasswordData(t *testing.T) {
	token := getTokenAsHeader(true)
	_ = newUser(t, false, token)
	var oldRow models.User
	config.DB.Where("id = 2").Find(&oldRow)
	data := models.User{
		Name:     "Abdel Aziz hassan Abdel Aziz",
		Role:     2,
		Block:    1,
		Email:    "zizo@gmail.com",
		Password: "1234567",
	}
	k := put(data, usersUrl+"/2", false, token)
	var row models.User
	config.DB.Where("id = 2").Find(&row)
	d := getDataMap(k)
	assert.EqualValues(t, row.Name, d["name"])
	assert.EqualValues(t, row.Role, d["role"])
	assert.EqualValues(t, row.Block, d["block"])
	assert.EqualValues(t, row.Email, d["email"])
	assert.NotEqual(t, row.Password, oldRow.Password)
	assert.Equal(t, 200, k.Code)
}

/**
* Test Required inputs
 */
func TestUsersRequireInputs(t *testing.T) {
	token := getTokenAsHeader(true)
	///not send name
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Role:  2,
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///not send role
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:  "Abdel Aziz hassan Abdel Aziz",
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///not send block
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:  "Abdel Aziz hassan Abdel Aziz",
		Role:  2,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///not send email
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:  "Abdel Aziz hassan Abdel Aziz",
		Role:  2,
		Block: 1,
		Password: "12121221",
	}, usersUrl, false, token)
	///not send password
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:  "Abdel Aziz hassan Abdel Aziz",
		Role:  2,
		Block: 1,
		Email: "zizo199988@gmail.com",
	}, usersUrl, false, token)
	usersUrl = usersUrl + "/1"
	///not send name
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Role:  2,
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///not send role
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:  "Abdel Aziz hassan Abdel Aziz",
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///not send block
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:  "Abdel Aziz hassan Abdel Aziz",
		Role:  2,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///not send email
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:  "Abdel Aziz hassan Abdel Aziz",
		Role:  2,
		Block: 1,
		Password: "12121221",
	}, usersUrl, false, token)
}

/**
* Test inputs limitaion
 */
func TestUsersInputsLimitation(t *testing.T) {
	token := getTokenAsHeader(true)
	///min send name
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(5),
		Role:  2,
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///max send name
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(80),
		Role:  2,
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///min send Role
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  0,
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///max send Role
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  100,
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///min send Block
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 0,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///max send Block
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 3,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///min send Email
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 0,
		Email: "z@2.c",
		Password: "12121221",
	}, usersUrl, false, token)
	///max send Email
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 3,
		Email: helpers.RandomString(100)+"@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///min send Password
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 0,
		Email: "zizo199988@gmail.com",
		Password: helpers.RandomString(2),
	}, usersUrl, false, token)
	///max send Password
	checkPostRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 3,
		Email: "zizo199988@gmail.com",
		Password: helpers.RandomString(1000),
	}, usersUrl, false, token)
	usersUrl = usersUrl + "/1"
	///min send name
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(5),
		Role:  2,
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///max send name
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(80),
		Role:  2,
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///min send Role
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  0,
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///max send Role
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  100,
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///min send Block
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 0,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///max send Block
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 3,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///min send Email
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 0,
		Email: "z@2.c",
		Password: "12121221",
	}, usersUrl, false, token)
	///max send Email
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 3,
		Email: helpers.RandomString(100)+"@gmail.com",
		Password: "12121221",
	}, usersUrl, false, token)
	///min send Password
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 0,
		Email: "zizo199988@gmail.com",
		Password: helpers.RandomString(2),
	}, usersUrl, false, token)
	///max send Password
	checkPutRequestWithHeadersDataIsValid(t, models.User{
		Name:helpers.RandomString(10),
		Role:  1,
		Block: 3,
		Email: "zizo199988@gmail.com",
		Password: helpers.RandomString(1000),
	}, usersUrl, false, token)

}

func newUser(t *testing.T, migrate bool, token map[string]string) *httptest.ResponseRecorder {
	data := models.User{
		Name:     "Abdel Aziz",
		Role:     1,
		Block:    2,
		Email:    "zizo1999988@gmail.com",
		Password: "1234567",
	}
	w := post(data, usersUrl, migrate, token)
	return w
}

func existsEmailData() models.User {
	return models.User{
		Name:  "Abdel Aziz hassan Abdel Aziz",
		Role:  2,
		Block: 1,
		Email: "zizo199988@gmail.com",
		Password: "12121221",
	}
}
