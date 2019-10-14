package test

import (
	"github.com/stretchr/testify/assert"
	"golang-starter/app/models"
	"golang-starter/config"
	"golang-starter/helpers"
	"gopkg.in/thedevsaddam/gojsonq.v2"
	"testing"
)

///// show all case
func TestTranslationsShowAll(t *testing.T) {
	k := get("admin/translations", false, getTokenAsHeader(true))
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, 0.0, recoverResponse.Find("data.offset"))
	assert.Equal(t, 200, k.Code)
}

func TestTranslationsFilter(t *testing.T) {
	translationUrl := "admin/translations"
	token := getTokenAsHeader(true)
	newTranslation()
	filter(t, translationUrl, "Home", "value", "equal", token)
	filter(t, translationUrl, 1, "page_id", "equal", token)
	filter(t, translationUrl, "home_page_title", "slug", "equal", token)
	filter(t, translationUrl, "ar", "lang", "equal", token)
}

///// show function cases
func TestTranslationsShowWithValidId(t *testing.T) {
	translationUrl := "admin/translations/1"
	token := getTokenAsHeader(true)
	_ = newTranslation()
	k := get(translationUrl, false, token)
	assert.EqualValues(t, 1, returnResponseKey(k, "data.page_id"))
	assert.Equal(t, 200, k.Code)
}

func TestTranslationsShowWithNotValidId(t *testing.T) {
	translationUrl := "admin/translations"
	token := getTokenAsHeader(true)
	k := get(translationUrl+"/1000", false, token)
	assert.Equal(t, 404, k.Code)
}

/// valid store update cases
func TestTranslationsUpdateWithValidData(t *testing.T) {
	translationUrl := "admin/translations"
	token := getTokenAsHeader(true)
	_ = newPage()
	_ = newTranslation()
	var oldRow models.Translation
	config.DB.First(&oldRow)
	data := models.Translation{
		PageId: 1,
		Slug:   "home_page_title",
		Value:  "Homed",
		Lang:   "ar",
	}
	k := put(data, translationUrl+"/1", false, token)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.EqualValues(t, data.PageId, recoverResponse.Find("data.page_id"))
	assert.Equal(t, 200, k.Code)
}

func TestTranslationsUpdateWithValidDataAndPageIdZero(t *testing.T) {
	translationUrl := "admin/translations"
	token := getTokenAsHeader(true)
	_ = newTranslation()
	var oldRow models.Translation
	config.DB.First(&oldRow)
	data := models.Translation{
		PageId: 0,
		Slug:   "home_page_title",
		Value:  "Homed",
		Lang:   "ar",
	}
	k := put(data, translationUrl+"/1", false, token)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.EqualValues(t, data.PageId, recoverResponse.Find("data.page_id"))
	assert.Equal(t, 200, k.Code)
}

func TestTranslationsUpdateWithValidDataAndNotSendPagId(t *testing.T) {
	translationUrl := "admin/translations"
	token := getTokenAsHeader(true)
	_ = newTranslation()
	var oldRow models.Translation
	config.DB.First(&oldRow)
	data := models.Translation{
		Slug:   "home_page_title",
		Value:  "Homed",
		Lang:   "ar",
	}
	k := put(data, translationUrl+"/1", false, token)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.EqualValues(t, data.PageId, recoverResponse.Find("data.page_id"))
	assert.Equal(t, 200, k.Code)
}

/**
* Test Required inputs
 */
func TestTranslationsRequireInputs(t *testing.T) {
	token := getTokenAsHeader(true)
	newTranslation()
	translationUrl := "admin/translations/1"
	///not send slug
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		PageId: 1,
		Value:  helpers.RandomString(5),
		Lang:   helpers.RandomString(2),
	}, translationUrl, false, token)
	///not send Value
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		PageId: 1,
		Slug:   helpers.RandomString(10),
		Lang:   helpers.RandomString(2),
	}, translationUrl, false, token)
	///not send Lang
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		PageId: 1,
		Slug:   helpers.RandomString(10),
		Value:  helpers.RandomString(5),
	}, translationUrl, false, token)
}


/**
* Test not valid inputs
 */
func TestTranslationsNotValidInputs(t *testing.T) {
	token := getTokenAsHeader(true)
	translationUrl :=  "admin/translations/1"
	m := make(map[string]interface{})
	m["page_id"] = "asdasdasd"
	m["slug"] =  helpers.RandomString(10)
	m["value"] = helpers.RandomString(10)
	m["lang"] = helpers.RandomString(2)

	checkPutRequestWithHeadersDataIsValid(t,m , translationUrl, false , token)
}

func TestTranslationsWithNotValidPageId(t *testing.T) {
	token := getTokenAsHeader(true)
	translationUrl :=  "admin/translations/1"
	m := make(map[string]interface{})
	m["page_id"] = 100
	m["slug"] =  helpers.RandomString(10)
	m["value"] = helpers.RandomString(5)
	m["lang"] = helpers.RandomString(2)
	k := put(m, translationUrl, false, token)
	assert.Equal(t, 404, k.Code)
}

/**
* Test inputs limitaion
 */
func TestTranslationsInputsLimitation(t *testing.T) {
	token := getTokenAsHeader(true)
	///min name fails
	newTranslation()
	url := "admin/translations/1"
	///min Value fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		PageId: 1,
		Slug:   helpers.RandomString(10),
		Value:  helpers.RandomString(1),
		Lang:   helpers.RandomString(2),
	}, url, false, token)
	///max Value fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		PageId: 1,
		Slug:   helpers.RandomString(10),
		Value:  helpers.RandomString(300),
		Lang:   helpers.RandomString(2),
	}, url, false, token)
	///min Lang fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		PageId: 1,
		Slug:   helpers.RandomString(10),
		Value:  helpers.RandomString(10),
		Lang:   helpers.RandomString(1),
	}, url, false, token)
	///max Lang fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		PageId: 1,
		Slug:   helpers.RandomString(10),
		Value:  helpers.RandomString(10),
		Lang:   helpers.RandomString(20),
	}, url, false, token)
	///min Slug fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		PageId: 1,
		Slug:   helpers.RandomString(1),
		Value:  helpers.RandomString(10),
		Lang:   helpers.RandomString(2),
	}, url, false, token)
	///max Slug fails
	checkPutRequestWithHeadersDataIsValid(t, models.Translation{
		PageId: 1,
		Slug:   helpers.RandomString(60),
		Value:  helpers.RandomString(10),
		Lang:   helpers.RandomString(2),
	}, url, false, token)

}

func newTranslation() models.Translation {
	data := models.Translation{
		PageId: 1,
		Slug:   "home_page_title",
		Value:  "Home",
		Lang:   "ar",
	}
	config.DB.Create(&data)
	return data
}
