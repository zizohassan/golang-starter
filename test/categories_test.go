package test

import (
	"github.com/stretchr/testify/assert"
	"golang-starter/app/models"
	"golang-starter/config"
	"golang-starter/helpers"
	"net/http/httptest"
	"testing"
)

var categoryUrl = "admin/categories"

///// show all case
func TestCategoriesShowAll(t *testing.T) {
	token := getTokenAsHeader(true)
	k := get(categoryUrl, false, token)
	assert.EqualValues(t, 0, returnResponseKey(k, "data.offset"))
	assert.Equal(t, 200, k.Code)
}

func TestCategoriesFilter(t *testing.T) {
	token := getTokenAsHeader(true)
	w := newCategory(t, false, token)
	assert.Equal(t, 200, w.Code)
	filter(t, categoryUrl, 1, "status", "equal", token)
	filter(t, categoryUrl, "Doctors", "name", "equal", token)
	filter(t, categoryUrl, 1, "status", "not-equal", token)
}

///// show function cases
func TestCategoriesShowWithValidId(t *testing.T) {
	token := getTokenAsHeader(true)
	w := newCategory(t, false, token)
	assert.Equal(t, 200, w.Code)
	k := get(categoryUrl+"/1", false, token)
	assert.Equal(t, "Doctors", returnResponseKey(k, "data.name"))
	assert.Equal(t, 200, k.Code)
}

func TestCategoriesShowWithNotValidId(t *testing.T) {
	token := getTokenAsHeader(true)
	k := get(categoryUrl+"/1000", false, token)
	assert.Equal(t, 404, k.Code)
}

///// delete case
func TestCategoriesDeleteWithValidId(t *testing.T) {
	token := getTokenAsHeader(true)
	w := newCategory(t, false, token)
	assert.Equal(t, 200, w.Code)
	k := deleter(categoryUrl+"/1", false, token)
	assert.Equal(t, 200, k.Code)
}

func TestCategoriesDeleteWithNotValidId(t *testing.T) {
	token := getTokenAsHeader(true)
	k := deleter(categoryUrl+"/1000", false, token)
	assert.Equal(t, 404, k.Code)
}

func TestCategoriesDeleteWithWrongRoute(t *testing.T) {
	token := getTokenAsHeader(true)
	k := deleter(categoryUrl, false, token)
	assert.Equal(t, 404, k.Code)
}

/// valid store update cases
func TestCategoriesStoreCategoryWithValidData(t *testing.T) {
	token := getTokenAsHeader(true)
	w := newCategory(t, false, token)
	assert.Equal(t, 200, w.Code)
}

func TestCategoriesUpdateCategoryWithValidData(t *testing.T) {
	token := getTokenAsHeader(true)
	_ = newCategory(t, false, token)
	var oldRow models.Category
	config.DB.First(&oldRow)
	data := models.Category{
		Name:   "New Data",
		Status: 1,
	}
	k := put(data, categoryUrl+"/1", false, token)
	assert.Equal(t, data.Name, returnResponseKey(k, "data.name"))
	assert.Equal(t, 200, k.Code)
}

/**
* Test Required inputs
 */
func TestCategoriesRequireInputs(t *testing.T) {
	token := getTokenAsHeader(true)
	///not send name
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(4),
	}, categoryUrl, false, token)
	///not send status
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Status: 1,
	}, categoryUrl, false, token)
	newCategory(t, false, token)
	url := categoryUrl + "/1"
	///not send name
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Name: helpers.RandomString(10),
	}, url, false, token)
	///not send status
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Status: 1,
	}, url, false, token)
}

/**
* Test inputs limitaion
 */
func TestCategoriesInputsLimitation(t *testing.T) {
	token := getTokenAsHeader(true)
	///min name fails
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Name:   helpers.RandomString(4),
		Status: 1,
	}, categoryUrl, false, token)
	///max name fails
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Name:   helpers.RandomString(80),
		Status: 1,
	}, categoryUrl, false, token)
	///max status fails
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Name:   helpers.RandomString(10),
		Status: 3,
	}, categoryUrl, false, token)
	///min status fails
	checkPostRequestWithHeadersDataIsValid(t, models.Category{
		Name:   helpers.RandomString(10),
		Status: 0,
	}, categoryUrl, false, token)
	///create new category
	newCategory(t, false, token)
	url := categoryUrl + "/1"
	///min name fails
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Name:   helpers.RandomString(4),
		Status: 1,
	}, url, false, token)
	///max name fails
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Name:   helpers.RandomString(80),
		Status: 1,
	}, url, false, token)
	///max status fails
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Name:   helpers.RandomString(10),
		Status: 3,
	}, url, false, token)
	///min status fails
	checkPutRequestWithHeadersDataIsValid(t, models.Category{
		Name:   helpers.RandomString(10),
		Status: 0,
	}, url, false, token)

}

func newCategory(t *testing.T, migrate bool, token map[string]string) *httptest.ResponseRecorder {
	data := models.Category{
		Name:   "Doctors",
		Status: 1,
	}
	w := post(data, categoryUrl, migrate, token)
	return w
}
