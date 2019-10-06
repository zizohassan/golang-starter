package test

import (
	"github.com/stretchr/testify/assert"
	"golang-starter/app/models"
	"golang-starter/config"
	"golang-starter/helpers"
	"gopkg.in/thedevsaddam/gojsonq.v2"
	"net/http/httptest"
	"testing"
)

///// show all case
func TestShowAll(t *testing.T) {
	k := get("admin/categories", false, getTokenAsHeader(t , true))
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, 0.0, recoverResponse.Find("data.offset"))
	assert.Equal(t, 200, k.Code)
}

func TestFilter(t *testing.T) {
	token := getTokenAsHeader(t , true)
	w := newCategory(t , false , token)
	assert.Equal(t, 200, w.Code)
	filter(t, "admin/categories?status=1", 1, "status", "equal" , token)
	filter(t, "admin/categories?name=doctor", "Doctors", "name", "equal" , token)
	filter(t, "admin/categories?status=4", 1, "status", "not-equal" , token)
}

func filter(t *testing.T, url string, value interface{}, key string, method string , token map[string]string ) {
	k := get(url, false , token)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	if method == "equal" {
		assert.EqualValues(t, value, recoverResponse.Find("data.records.[0]."+key))
	} else {
		assert.NotEqual(t, value, recoverResponse.Find("data.records.[0]."+key))
	}
	assert.Equal(t, 200, k.Code)
}

///// show function cases
func TestShowWithValidId(t *testing.T) {
	token := getTokenAsHeader(t , true)
	w := newCategory(t , false , token)
	assert.Equal(t, 200, w.Code)
	k := get("admin/categories/1", false, token)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, "Doctors", recoverResponse.Find("data.name"))
	assert.Equal(t, 200, k.Code)
}

func TestShowWithNotValidId(t *testing.T) {
	token := getTokenAsHeader(t , true)
	k := get("admin/categories/1000", false, token)
	assert.Equal(t, 404, k.Code)
}

///// delete case
func TestDeleteWithValidId(t *testing.T) {
	token := getTokenAsHeader(t , true)
	w := newCategory(t , false , token)
	assert.Equal(t, 200, w.Code)
	k := deleter("admin/categories/1", false, token)
	assert.Equal(t, 200, k.Code)
}

func TestDeleteWithNotValidId(t *testing.T) {
	token := getTokenAsHeader(t , true)
	k := deleter("admin/categories/1000", false , token)
	assert.Equal(t, 404, k.Code)
}

func TestDeleteWithWrongRoute(t *testing.T) {
	token := getTokenAsHeader(t , true)
	k := deleter("admin/categories", false , token)
	assert.Equal(t, 404, k.Code)
}


/// valid store update cases
func TestStoreCategoryWithValidData(t *testing.T) {
	token := getTokenAsHeader(t , true)
	w := newCategory(t , false , token)
	assert.Equal(t, 200, w.Code)
}

func TestUpdateCategoryWithValidData(t *testing.T) {
	token := getTokenAsHeader(t , true)
	_ = newCategory(t , false , token)
	var oldRow models.Category
	config.DB.First(&oldRow)
	data := models.Category{
		Name:   "New Data",
		Status: 1,
	}
	k := put(data, "admin/categories/1", false, token)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, oldRow.Name, recoverResponse.Find("data.name"))
	assert.Equal(t, 200, k.Code)
}

/**
* Test Required inputs
 */
func TestCategoriesRequireInputs(t *testing.T) {
	url := "admin/categories"
	token := getTokenAsHeader(t , true)
	///not send name
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(4),
	}, url, false , token)
	///not send status
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Status: 1,
	}, url, false , token)
	newCategory(t , false  , token)
	url = "admin/categories/1"
	///not send name
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(10),
	}, url, false , token)
	///not send status
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Status: 1,
	}, url, false , token)
}

/**
* Test inputs limitaion
 */
func TestCategoriesInputsLimitation(t *testing.T) {
	url := "admin/categories"
	token := getTokenAsHeader(t , true)
	///min name fails
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(4),
		Status: 1,
	}, url, false , token)
	///max name fails
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(80),
		Status: 1,
	}, url, false , token)
	///max status fails
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(10),
		Status: 3,
	}, url, false , token)
	///min status fails
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(10),
		Status: 0,
	}, url, false , token)
	///create new category
	newCategory(t , false  , token)
	url = "admin/categories/1"
	///min name fails
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(4),
		Status: 1,
	}, url, false , token)
	///max name fails
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(80),
		Status: 1,
	}, url, false , token)
	///max status fails
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(10),
		Status: 3,
	}, url, false , token)
	///min status fails
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(10),
		Status: 0,
	}, url, false , token)

}

func newCategory(t *testing.T , migrate bool , token map[string]string) *httptest.ResponseRecorder {
	data := models.Category{
		Name:   "Doctors",
		Status: 1,
	}
	w := post(data, "admin/categories", migrate, token)
	return w
}
