package test

import (
	"github.com/stretchr/testify/assert"
	"golang-starter/app/models"
	"golang-starter/config"
	"golang-starter/helpers"
	"gopkg.in/thedevsaddam/gojsonq.v2"
	"testing"
)

var settingUrl = "/admin/settings"

///// show all case
func TestSettingsShowAll(t *testing.T) {
	k := get(settingUrl, false, getTokenAsHeader(t, true))
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, 0.0, recoverResponse.Find("data.offset"))
	assert.Equal(t, 200, k.Code)
}

func TestSettingsFilter(t *testing.T) {
	token := getTokenAsHeader(t, true)
	newSetting()
	filter(t, settingUrl, "Facebook", "name", "equal", token)
	filter(t, settingUrl, "facebook", "slug", "equal", token)
	filter(t, settingUrl, "http://facebook.com", "value", "equal", token)
	filter(t, settingUrl, "text", "setting_type", "equal", token)
}

///// show function cases
func TestSettingsShowWithValidId(t *testing.T) {
	token := getTokenAsHeader(t, true)
	_ = newSetting()
	k := get(settingUrl+"/1", false, token)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, "facebook", recoverResponse.Find("data.slug"))
	assert.Equal(t, 200, k.Code)
}

func TestSettingsShowWithNotValidId(t *testing.T) {
	token := getTokenAsHeader(t, true)
	k := get(settingUrl+"/1000", false, token)
	assert.Equal(t, 404, k.Code)
}

/// valid store update cases
func TestSettingsUpdateCategoryWithValidData(t *testing.T) {
	token := getTokenAsHeader(t, true)
	_ = newSetting()
	var oldRow models.Setting
	config.DB.First(&oldRow)
	data := models.Setting{
		Name:        "FacebookFacebook",
		Slug:        "facebook",
		Value:       "http://facebook.com",
		SettingType: "text",
	}
	k := put(data, settingUrl+"/1", false, token)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, data.Name, recoverResponse.Find("data.name"))
	assert.Equal(t, 200, k.Code)
}

/**
* Test Required inputs
 */
func TestSettingsRequireInputs(t *testing.T) {
	token := getTokenAsHeader(t, true)
	newTranslation()
	settingUrl := settingUrl + "/1"
	///not send SettingType
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(5),
		Name:  helpers.RandomString(8),
	}, settingUrl, false, token)
	///not send slug
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Value:       helpers.RandomString(5),
		Name:        helpers.RandomString(8),
		SettingType: helpers.RandomString(4),
	}, settingUrl, false, token)
	///not send Value
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Slug:        helpers.RandomString(10),
		Name:        helpers.RandomString(8),
		SettingType: helpers.RandomString(4),
	}, settingUrl, false, token)
	///not send Name
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Slug:        helpers.RandomString(10),
		Value:       helpers.RandomString(5),
		SettingType: helpers.RandomString(4),
	}, settingUrl, false, token)
}

/**
* Test inputs Limitaion
 */
func TestSettingsInputsLimitation(t *testing.T) {
	token := getTokenAsHeader(t, true)
	///min name fails
	newCategory(t, false, token)
	url := settingUrl + "/1"
	///min Name fails
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Name:  helpers.RandomString(1),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(10),
		SettingType:  helpers.RandomString(10),
	}, url, false, token)
	///max Name fails
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Name:  helpers.RandomString(60),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(10),
		SettingType:  helpers.RandomString(10),
	}, url, false, token)
	///min Slug fails
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Name:  helpers.RandomString(10),
		Slug:  helpers.RandomString(1),
		Value: helpers.RandomString(10),
		SettingType:  helpers.RandomString(10),
	}, url, false, token)
	///max Slug fails
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Name:  helpers.RandomString(10),
		Slug:  helpers.RandomString(60),
		Value: helpers.RandomString(10),
		SettingType:  helpers.RandomString(10),
	}, url, false, token)
	///min Value fails
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Name:  helpers.RandomString(10),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(1),
		SettingType:  helpers.RandomString(10),
	}, url, false, token)
	///max Value fails
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Name:  helpers.RandomString(10),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(300),
		SettingType:  helpers.RandomString(10),
	}, url, false, token)
	///min Slug fails
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Name:  helpers.RandomString(10),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(10),
		SettingType:  helpers.RandomString(1),
	}, url, false, token)
	///max Slug fails
	checkPutRequestWithHeadersDataIsValid(t, models.Setting{
		Name:  helpers.RandomString(10),
		Slug:  helpers.RandomString(10),
		Value: helpers.RandomString(10),
		SettingType:  helpers.RandomString(60),
	}, url, false, token)

}

func newSetting() models.Setting {
	data := models.Setting{
		Name:        "Facebook",
		Slug:        "facebook",
		Value:       "http://facebook.com",
		SettingType: "text",
	}
	config.DB.Create(&data)
	return data
}
