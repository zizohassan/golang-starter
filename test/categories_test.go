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
/// show all case
func TestShowAll(t *testing.T)  {
	k := get("admin/categories" , true)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, 0.0, recoverResponse.Find("data.offset"))
	assert.Equal(t, 200, k.Code)
}

func TestFilter(t *testing.T)  {
	filter(t  , "admin/categories?status=1" , 1 , "status", "equal")
	filter(t  , "admin/categories?name=doctor" , "Doctors" , "name", "equal")
	filter(t  , "admin/categories?status=4" , 1 , "status" , "not-equal")
}

func filter(t *testing.T , url string , value interface{} , key string , method string)  {
	w := register()
	assert.Equal(t, 200, w.Code)
	k := get(url , false)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	if method == "equal"{
		assert.EqualValues(t, value, recoverResponse.Find("data.records.[0]."+key))
	}else {
		assert.NotEqual(t, value, recoverResponse.Find("data.records.[0]."+key))
	}
	assert.Equal(t, 200, k.Code)
}

/// show function cases
func TestShowWithValidId(t *testing.T)  {
	w := register()
	assert.Equal(t, 200, w.Code)
	k := get("admin/categories/1" , false)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, "Doctors", recoverResponse.Find("data.name"))
	assert.Equal(t, 200, k.Code)
}

func TestShowWithNotValidId(t *testing.T)  {
	k := get("admin/categories/1000" , true)
	assert.Equal(t, 404, k.Code)
}

/// delete case
func TestDeleteWithValidId(t *testing.T)  {
	w := register()
	assert.Equal(t, 200, w.Code)
	k := deleter("admin/categories/1" , false)
	assert.Equal(t, 200, k.Code)
}

func TestDeleteWithNotValidId(t *testing.T)  {
	k := deleter("admin/categories/1000" , true)
	assert.Equal(t, 404, k.Code)
}

func TestDeleteWithWrongRoute(t *testing.T)  {
	k := deleter("admin/categories" , true)
	assert.Equal(t, 404, k.Code)
}

/// valid store update cases
func TestStoreCategoryWithValidData(t *testing.T)  {
	w := register()
	assert.Equal(t, 200, w.Code)
}

func TestUpdateCategoryWithValidData(t *testing.T)  {
	_ = register()
	var oldRow models.Category
	config.DB.Where("name = ?" , "Doctors").First(&oldRow)
	data := models.Category{
		Name:"New Data",
		Status:1,
	}
	k := put(data , "admin/categories/1" , false)
	responseData := responseData(k.Result().Body)
	recoverResponse := gojsonq.New().JSONString(responseData)
	assert.Equal(t, oldRow.Name, recoverResponse.Find("data.name"))
	assert.Equal(t, 200, k.Code)
}

/// validate update store requests
func TestValidateStoreUrl(t *testing.T)  {
	url := "admin/categories"
	CategoryWithBigNameData(t , url , "POST")
	CategoryWitOutName(t , url , "POST")
	CategoryWitOutStatus(t , url , "POST")
	CategoryNotValidStatus(t , url , "POST")
	CategoryEmptyName(t , url , "POST")
	CategoryEmptyStatus(t , url , "POST")
}

func TestValidateUpdateUrl(t *testing.T)  {
	url := "admin/categories/1"
	CategoryWithBigNameData(t , url , "PUT")
	CategoryWitOutName(t , url , "PUT")
	CategoryWitOutStatus(t , url , "PUT")
	CategoryNotValidStatus(t , url , "PUT")
	CategoryEmptyName(t , url , "PUT")
	CategoryEmptyStatus(t , url , "PUT")

}

func CategoryWithBigNameData(t *testing.T , url string , method string)  {
	data := models.Category{
		Name:helpers.RandomString(60),
		Status:1,
	}
	w := request(data , url, true , method)
	assert.Equal(t, 400, w.Code)
}

func CategoryWitOutName(t *testing.T , url string , method string)  {
	data := models.Category{
		Status:1,
	}
	w := request(data , url, true , method)
	assert.Equal(t, 400, w.Code)
}

func CategoryWitOutStatus(t *testing.T  , url string, method string)  {
	data := models.Category{
		Name:"Doctors",
	}
	w := request(data , url, true , method)
	assert.Equal(t, 400, w.Code)
}

func CategoryNotValidStatus(t *testing.T , url string, method string)  {
	data := models.Category{
		Name:"Doctors",
		Status:3,
	}
	w := request(data , url, true , method)
	assert.Equal(t, 400, w.Code)
}

func CategoryEmptyName(t *testing.T , url string, method string)  {
	data := models.Category{
		Name:"",
		Status:1,
	}
	w := request(data , url, true , method)
	assert.Equal(t, 400, w.Code)
}

func CategoryEmptyStatus(t *testing.T , url string, method string)  {
	data := models.Category{
		Name:"Doctors",
		Status:0,
	}
	w := request(data , url, true , method)
	assert.Equal(t, 400, w.Code)
}

func register() *httptest.ResponseRecorder {
	data := models.Category{
		Name:"Doctors",
		Status:1,
	}
	w := post(data , "admin/categories" , true)
	return w
}
