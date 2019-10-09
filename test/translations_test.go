package test

import (
	"github.com/stretchr/testify/assert"
	"golang-starter/app/models"
	"golang-starter/config"
	"golang-starter/helpers"
	"gopkg.in/thedevsaddam/gojsonq.v2"
	"testing"
)

var translationUrl = "/admin/translations"

///// show all case
func TestTranslationsShowAll(t *testing.T) {
	k := get(translationUrl, false, getTokenAsHeader(true))
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, 0.0, recoverResponse.Find("data.offset"))
	assert.Equal(t, 200, k.Code)
}

func TestTranslationsFilter(t *testing.T) {
	token := getTokenAsHeader(true)
	newTranslation()
	filter(t, translationUrl, "Home", "value", "equal", token)
	filter(t, translationUrl, "home", "page", "equal", token)
	filter(t, translationUrl, "home_page_title", "slug", "equal", token)
	filter(t, translationUrl, "ar", "lang", "equal", token)
}

///// show function cases
func TestTranslationsShowWithValidId(t *testing.T) {
	token := getTokenAsHeader(true)
	_ = newTranslation()
	k := get(translationUrl+"/1", false, token)
	assert.Equal(t, "home", returnResponseKey(k, "data.page"))
	assert.Equal(t, 200, k.Code)
}

func TestTranslationsShowWithNotValidId(t *testing.T) {
	token := getTokenAsHeader(true)
	k := get(translationUrl+"/1000", false, token)
	assert.Equal(t, 404, k.Code)
}

/// valid store update cases
func TestTranslationsUpdateCategoryWithValidData(t *testing.T) {
	token := getTokenAsHeader(true)
	_ = newTranslation()
	var oldRow models.Translation
	config.DB.First(&oldRow)
	data := models.Translation{
		Page:  "homeEE",
		Slug:  "home_page_title",
		Value: "Homed",
		Lang:  "ar",
	}
	k := put(data, translationUrl+"/1", false, token)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, data.Page, recoverResponse.Find("data.page"))
	assert.Equal(t, 200, k.Code)
}

/**
* Test Required inputs
 */
func TestTranslationsRequireInputs(t *testing.T) {
	token := getTokenAsHeader(true)
	newTranslation()
	translationUrl := translationUrl + "/1"
	///not send page
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(5),
		Lang:  helpers.RandomString(2),
	}, translationUrl, false, token)
	///not send slug
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Page:  helpers.RandomString(7),
		Value: helpers.RandomString(5),
		Lang:  helpers.RandomString(2),
	}, translationUrl, false, token)
	///not send Value
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Page: helpers.RandomString(7),
		Slug: helpers.RandomString(10),
		Lang: helpers.RandomString(2),
	}, translationUrl, false, token)
	///not send Lang
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Page:  helpers.RandomString(7),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(5),
	}, translationUrl, false, token)
}

/**
* Test inputs limitaion
 */
func TestTranslationsInputsLimitation(t *testing.T) {
	token := getTokenAsHeader(true)
	///min name fails
	newCategory(t, false, token)
	url := translationUrl + "/1"
	///min Value fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Page:  helpers.RandomString(7),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(1),
		Lang:  helpers.RandomString(2),
	}, url, false, token)
	///max Value fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Page:  helpers.RandomString(7),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(300),
		Lang:  helpers.RandomString(2),
	}, url, false, token)
	///min Page fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Page:  helpers.RandomString(1),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(10),
		Lang:  helpers.RandomString(2),
	}, url, false, token)
	///max Page fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Page:  helpers.RandomString(50),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(10),
		Lang:  helpers.RandomString(2),
	}, url, false, token)
	///min Lang fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Page:  helpers.RandomString(4),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(10),
		Lang:  helpers.RandomString(1),
	}, url, false, token)
	///max Lang fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Page:  helpers.RandomString(4),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(10),
		Lang:  helpers.RandomString(20),
	}, url, false, token)
	///min Slug fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Page:  helpers.RandomString(4),
		Slug:  helpers.RandomString(1),
		Value: helpers.RandomString(10),
		Lang:  helpers.RandomString(2),
	}, url, false, token)
	///max Slug fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		Page:  helpers.RandomString(4),
		Slug:  helpers.RandomString(60),
		Value: helpers.RandomString(10),
		Lang:  helpers.RandomString(2),
	}, url, false, token)

}

func newTranslation() models.Translation {
	data := models.Translation{
		Page:  "home",
		Slug:  "home_page_title",
		Value: "Home",
		Lang:  "ar",
	}
	config.DB.Create(&data)
	return data
}
